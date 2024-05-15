package document

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/backent/fra-golang/config"
	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/models/elastic"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesDocumentTracker "github.com/backent/fra-golang/repositories/document_tracker"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
)

type ServiceDocumentImpl struct {
	DB       *sql.DB
	EsClient *elasticsearch.Client
	repositoriesDocument.RepositoryDocumentInterface
	*middlewares.DocumentMiddleware
	repositoriesRisk.RepositoryRiskInterface
	repositoriesRejectNote.RepositoryRejectNoteInterface
	repositoriesUser.RepositoryUserInterface
	repositoriesNotification.RepositoryNotificationInterface
	repositoriesDocument.RepositoryDocumentSearchInterface
	repositoriesDocumentTracker.RepositoryDocumentTrackerInterface
}

func NewServiceDocumentImpl(
	db *sql.DB,
	esClient *elasticsearch.Client,
	repositoriesDocument repositoriesDocument.RepositoryDocumentInterface,
	documentMiddleware *middlewares.DocumentMiddleware,
	repositoriesRisk repositoriesRisk.RepositoryRiskInterface,
	repositoriesRejectNote repositoriesRejectNote.RepositoryRejectNoteInterface,
	repositoriesUser repositoriesUser.RepositoryUserInterface,
	repositoriesNotification repositoriesNotification.RepositoryNotificationInterface,
	repositoriesDocumentSearch repositoriesDocument.RepositoryDocumentSearchInterface,
	repositoriesDocumentTracker repositoriesDocumentTracker.RepositoryDocumentTrackerInterface,
) ServiceDocumentInterface {
	return &ServiceDocumentImpl{
		DB:                                 db,
		EsClient:                           esClient,
		RepositoryDocumentInterface:        repositoriesDocument,
		DocumentMiddleware:                 documentMiddleware,
		RepositoryRiskInterface:            repositoriesRisk,
		RepositoryRejectNoteInterface:      repositoriesRejectNote,
		RepositoryUserInterface:            repositoriesUser,
		RepositoryNotificationInterface:    repositoriesNotification,
		RepositoryDocumentSearchInterface:  repositoriesDocumentSearch,
		RepositoryDocumentTrackerInterface: repositoriesDocumentTracker,
	}
}

func (implementation *ServiceDocumentImpl) Create(ctx context.Context, request webDocument.DocumentRequestCreate) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	// check first if it is new document
	isNewDocument := request.Uuid == ""

	implementation.DocumentMiddleware.Create(ctx, tx, &request)

	trackerChan := make(chan error)
	documentChan := make(chan error)

	if isNewDocument {
		go func() {
			defer close(trackerChan)
			err = implementation.RepositoryDocumentTrackerInterface.Create(ctx, tx, request.Uuid, time.Now())
			trackerChan <- err
		}()
	} else {
		close(trackerChan)
	}

	var document models.Document

	go func() {
		defer close(documentChan)
		document, err = createDocumentAndRisk(ctx, tx, implementation, request)
		documentChan <- err
	}()
	helpers.PanicIfError(<-trackerChan)
	helpers.PanicIfError(<-documentChan)

	blastNotification(ctx, tx, document, implementation.RepositoryUserInterface, implementation.RepositoryNotificationInterface)

	return webDocument.DocumentModelToDocumentResponse(document)
}

func createDocumentAndRisk(ctx context.Context, tx *sql.Tx, implementation *ServiceDocumentImpl, request webDocument.DocumentRequestCreate) (models.Document, error) {
	var document models.Document
	document = models.Document{
		Uuid:        request.Uuid,
		CreatedBy:   request.CreatedBy,
		ActionBy:    request.ActionBy,
		Action:      request.Action,
		ProductName: request.ProductName,
		Category:    request.Category,
	}

	document, err := implementation.RepositoryDocumentInterface.Create(ctx, tx, document)
	if err != nil {
		return document, err
	}

	for _, riskRequest := range request.Risks {
		risk := models.Risk{
			DocumentId:             document.Id,
			RiskName:               riskRequest.RiskName,
			FraudSchema:            riskRequest.FraudSchema,
			FraudMotive:            riskRequest.FraudMotive,
			FraudTechnique:         riskRequest.FraudTechnique,
			RiskSource:             riskRequest.RiskSource,
			RootCause:              riskRequest.RootCause,
			BisproControlProcedure: riskRequest.BisproControlProcedure,
			QualitativeImpact:      riskRequest.QualitativeImpact,
			LikehoodJustification:  riskRequest.LikehoodJustification,
			ImpactJustification:    riskRequest.ImpactJustification,
			StartegyAgreement:      riskRequest.StartegyAgreement,
			StrategyRecomendation:  riskRequest.StrategyRecomendation,
			AssessmentLikehood:     riskRequest.AssessmentLikehood,
			AssessmentImpact:       riskRequest.AssessmentImpact,
			AssessmentRiskLevel:    riskRequest.AssessmentRiskLevel,
		}

		_, err = implementation.RepositoryRiskInterface.Create(ctx, tx, risk)
		if err != nil {
			return document, err
		}
		document.RiskDetail = append(document.RiskDetail, risk)
	}

	err = implementation.RepositoryDocumentSearchInterface.IndexProduct(implementation.EsClient, document)
	if err != nil {
		log.Println(err)
	}

	return document, nil
}

func (implementation *ServiceDocumentImpl) Update(ctx context.Context, request webDocument.DocumentRequestUpdate) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Update(ctx, tx, &request)

	document := models.Document{
		Id:          request.Id,
		Uuid:        request.Uuid,
		CreatedBy:   request.CreatedBy,
		ActionBy:    request.ActionBy,
		Action:      request.Action,
		ProductName: request.ProductName,
	}

	document, err = implementation.RepositoryDocumentInterface.Update(ctx, tx, document)
	helpers.PanicIfError(err)

	return webDocument.DocumentModelToDocumentResponse(document)
}
func (implementation *ServiceDocumentImpl) Delete(ctx context.Context, request webDocument.DocumentRequestDelete) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Delete(ctx, tx, &request)

	err = implementation.RepositoryDocumentInterface.Delete(ctx, tx, request.Id)
	helpers.PanicIfError(err)

}
func (implementation *ServiceDocumentImpl) FindById(ctx context.Context, request webDocument.DocumentRequestFindById) webDocument.DocumentResponseWithDetail {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.FindById(ctx, tx, &request)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	return webDocument.DocumentModelToDocumentResponseWithDetail(document)
}
func (implementation *ServiceDocumentImpl) FindAll(ctx context.Context, request webDocument.DocumentRequestFindAll) ([]webDocument.DocumentResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.FindAll(ctx, tx, &request)

	documents, total, err := implementation.RepositoryDocumentInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.QueryAction, request.QueryCategory)
	helpers.PanicIfError(err)

	if documents == nil {
		return []webDocument.DocumentResponse{}, total
	}

	return webDocument.BulkDocumentModelToDocumentResponse(documents), total
}
func (implementation *ServiceDocumentImpl) FindAllWithDetail(ctx context.Context, request webDocument.DocumentRequestFindAll) ([]webDocument.DocumentResponseWithDetail, int) {
	panic("no longer need this func")

}

func (implementation *ServiceDocumentImpl) GetProductDistinct(ctx context.Context, request webDocument.DocumentRequestGetProductDistinct) []webDocument.DocumentResponseGetProductDistinct {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.GetProductDistinct(ctx, tx, &request)

	documents, err := implementation.RepositoryDocumentInterface.GetProductDistinct(ctx, tx)
	helpers.PanicIfError(err)

	return webDocument.BulkDocumentModelToBulkDocumentResponseGetProductDistinct(documents)
}

func (implementation *ServiceDocumentImpl) Approve(ctx context.Context, request webDocument.DocumentRequestApprove) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Approve(ctx, tx, &request)

	document, err := implementation.RepositoryDocumentInterface.Create(ctx, tx, request.Document)
	helpers.PanicIfError(err)

	for _, riskRequest := range request.Document.RiskDetail {
		risk := models.Risk{
			DocumentId:             document.Id,
			RiskName:               riskRequest.RiskName,
			FraudSchema:            riskRequest.FraudSchema,
			FraudMotive:            riskRequest.FraudMotive,
			FraudTechnique:         riskRequest.FraudTechnique,
			RiskSource:             riskRequest.RiskSource,
			RootCause:              riskRequest.RootCause,
			BisproControlProcedure: riskRequest.BisproControlProcedure,
			QualitativeImpact:      riskRequest.QualitativeImpact,
			LikehoodJustification:  riskRequest.LikehoodJustification,
			ImpactJustification:    riskRequest.ImpactJustification,
			StartegyAgreement:      riskRequest.StartegyAgreement,
			StrategyRecomendation:  riskRequest.StrategyRecomendation,
			AssessmentLikehood:     riskRequest.AssessmentLikehood,
			AssessmentImpact:       riskRequest.AssessmentImpact,
			AssessmentRiskLevel:    riskRequest.AssessmentRiskLevel,
		}

		_, err = implementation.RepositoryRiskInterface.Create(ctx, tx, risk)
		helpers.PanicIfError(err)
	}

	blastNotification(ctx, tx, document, implementation.RepositoryUserInterface, implementation.RepositoryNotificationInterface)
	err = implementation.RepositoryDocumentSearchInterface.IndexProduct(implementation.EsClient, document)
	if err != nil {
		log.Println(err)
	}

}

func (implementation *ServiceDocumentImpl) Reject(ctx context.Context, request webDocument.DocumentRequestReject) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Reject(ctx, tx, &request)

	document, err := implementation.RepositoryDocumentInterface.Create(ctx, tx, request.Document)
	helpers.PanicIfError(err)

	rejectNoteMap := make(map[int]webDocument.RejectNoteRequest)
	for _, rejectNote := range request.RejectNote {
		rejectNoteMap[rejectNote.RiskId] = rejectNote
	}

	for _, risk := range document.RiskDetail {
		// store risk id to use when find original id from request
		requestRiskId := risk.Id

		// create new risk
		risk.DocumentId = document.Id
		risk, err = implementation.RepositoryRiskInterface.Create(ctx, tx, risk)
		helpers.PanicIfError(err)

		// use previous risk id to find the correct reject note
		rejectNote, found := rejectNoteMap[requestRiskId]
		if found {
			rejectNoteModel := models.RejectNote{
				DocumentId:             document.Id,
				RiskId:                 risk.Id,
				Fraud:                  rejectNote.Fraud,
				RiskSource:             rejectNote.RiskSource,
				RootCause:              rejectNote.RootCause,
				BisproControlProcedure: rejectNote.BisproControlProcedure,
				QualitativeImpact:      rejectNote.QualitativeImpact,
				Assessment:             rejectNote.Assessment,
				Justification:          rejectNote.Justification,
				Strategy:               rejectNote.Strategy,
			}

			_, err = implementation.RepositoryRejectNoteInterface.Create(ctx, tx, rejectNoteModel)
			helpers.PanicIfError(err)

		}
	}

	blastNotification(ctx, tx, document, implementation.RepositoryUserInterface, implementation.RepositoryNotificationInterface)
	err = implementation.RepositoryDocumentSearchInterface.IndexProduct(implementation.EsClient, document)
	if err != nil {
		log.Println(err)
	}
}

func (implementation *ServiceDocumentImpl) MonitoringList(ctx context.Context, request webDocument.DocumentRequestMonitoringList) ([]webDocument.DocumentResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.MonitoringList(ctx, tx, &request)

	documents, total, err := implementation.RepositoryDocumentInterface.FindAllNoGroup(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.QueryAction, request.QueryPeriod, request.QueryName)
	helpers.PanicIfError(err)

	if documents == nil {
		return []webDocument.DocumentResponse{}, total
	}

	return webDocument.BulkDocumentModelToDocumentResponse(documents), total
}

func (implementation *ServiceDocumentImpl) TrackerProduct(ctx context.Context, request webDocument.DocumentRequestTrackerProduct) []webDocument.DocumentTrackerProduct {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.TrackerProduct(ctx, tx, &request)

	documents, err := implementation.RepositoryDocumentInterface.TrackerProductByName(ctx, tx, request.QuerySearch)
	helpers.PanicIfError(err)

	if documents == nil {
		return []webDocument.DocumentTrackerProduct{}
	}
	return webDocument.BulkDocumetModelToDocumentTrackerProduct(documents)
}

func blastNotification(ctx context.Context, tx *sql.Tx, document models.Document, repositoryUserInterface repositoriesUser.RepositoryUserInterface, repositoryNotificationInterface repositoriesNotification.RepositoryNotificationInterface) {
	// get all users

	users, _, err := repositoryUserInterface.FindAll(ctx, tx, 99999, 0, "id", "asc", "approve", "")
	helpers.PanicIfError(err)

	// blast notifications
	for _, user := range users {
		title, subtitle, err := config.NotificationGenerator(user.Role, document.Action, document.ProductName)
		if err == nil {
			notification := models.Notification{
				UserId:     user.Id,
				DocumentId: document.Id,
				Title:      title,
				Subtitle:   subtitle,
				Action:     document.Action,
			}
			repositoryNotificationInterface.Create(ctx, tx, notification)

			recipient := helpers.RecipientDocumentNotification{
				Title:    title,
				Subtitle: subtitle,
				Email:    user.Email,
				Subject:  title,
			}
			go helpers.SendMailWithoutAuth(recipient)
		}
	}
}

func (implementation *ServiceDocumentImpl) SearchGlobal(ctx context.Context, request webDocument.DocumentRequestSearchGlobal) ([]elastic.DocumentSearchGlobal, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.SearchGlobal(ctx, tx, &request)

	docs, total, err := implementation.RepositoryDocumentSearchInterface.SearchByProductName(implementation.EsClient, request.QuerySearch, request.GetTake(), request.GetSkip())
	helpers.PanicIfError(err)

	if total == 0 {
		return []elastic.DocumentSearchGlobal{}, total
	}

	return docs, total
}

func (implementation *ServiceDocumentImpl) UploadFinal(ctx context.Context, request webDocument.DocumentRequestUploadFinal) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.UploadFinal(ctx, tx, &request)

	request.Document.FileOriginalName = request.FileHandler.Filename

	var separator = "."
	var sliceSplittedFileName = strings.Split(request.FileHandler.Filename, separator)
	request.Document.FileName = uuid.New().String() + separator + sliceSplittedFileName[len(sliceSplittedFileName)-1]

	document, err := implementation.RepositoryDocumentInterface.Create(ctx, tx, request.Document)
	helpers.PanicIfError(err)

	for _, riskRequest := range request.Document.RiskDetail {
		risk := models.Risk{
			DocumentId:             document.Id,
			RiskName:               riskRequest.RiskName,
			FraudSchema:            riskRequest.FraudSchema,
			FraudMotive:            riskRequest.FraudMotive,
			FraudTechnique:         riskRequest.FraudTechnique,
			RiskSource:             riskRequest.RiskSource,
			RootCause:              riskRequest.RootCause,
			BisproControlProcedure: riskRequest.BisproControlProcedure,
			QualitativeImpact:      riskRequest.QualitativeImpact,
			LikehoodJustification:  riskRequest.LikehoodJustification,
			ImpactJustification:    riskRequest.ImpactJustification,
			StartegyAgreement:      riskRequest.StartegyAgreement,
			StrategyRecomendation:  riskRequest.StrategyRecomendation,
			AssessmentLikehood:     riskRequest.AssessmentLikehood,
			AssessmentImpact:       riskRequest.AssessmentImpact,
			AssessmentRiskLevel:    riskRequest.AssessmentRiskLevel,
		}

		_, err = implementation.RepositoryRiskInterface.Create(ctx, tx, risk)
		helpers.PanicIfError(err)
	}

	err = helpers.SaveFile(request.File, request.Document.FileName, os.Getenv("DOCUMENT_FINAL_STORAGE_PATH"))
	helpers.PanicIfError(err)

	blastNotification(ctx, tx, document, implementation.RepositoryUserInterface, implementation.RepositoryNotificationInterface)
	err = implementation.RepositoryDocumentSearchInterface.IndexProduct(implementation.EsClient, document)
	if err != nil {
		log.Println(err)
	}

}

func (implementation *ServiceDocumentImpl) ServeFile(ctx context.Context, request webDocument.DocumentRequestServeFile) *webDocument.DocumentResponseServeFile {
	// Extract the requested file path
	staticDir := os.Getenv("DOCUMENT_FINAL_STORAGE_PATH")
	filePath := filepath.Join(staticDir, request.FileName)
	// Check if the file exists
	fileInfo, err := os.Stat(filePath)
	if err != nil || fileInfo.IsDir() {
		panic(exceptions.NewNotFoundError("file not found"))
	}

	// Add your custom logic here if needed

	file, err := os.Open(filePath)
	helpers.PanicIfError(err)

	return &webDocument.DocumentResponseServeFile{
		File:     file,
		FileInfo: fileInfo,
	}
}
