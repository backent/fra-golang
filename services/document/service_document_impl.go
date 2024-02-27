package document

import (
	"context"
	"database/sql"
	"log"

	"github.com/backent/fra-golang/config"
	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/elastic/go-elasticsearch/v8"
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
) ServiceDocumentInterface {
	return &ServiceDocumentImpl{
		DB:                                db,
		EsClient:                          esClient,
		RepositoryDocumentInterface:       repositoriesDocument,
		DocumentMiddleware:                documentMiddleware,
		RepositoryRiskInterface:           repositoriesRisk,
		RepositoryRejectNoteInterface:     repositoriesRejectNote,
		RepositoryUserInterface:           repositoriesUser,
		RepositoryNotificationInterface:   repositoriesNotification,
		RepositoryDocumentSearchInterface: repositoriesDocumentSearch,
	}
}

func (implementation *ServiceDocumentImpl) Create(ctx context.Context, request webDocument.DocumentRequestCreate) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Create(ctx, tx, &request)

	document := models.Document{
		Uuid:        request.Uuid,
		CreatedBy:   request.CreatedBy,
		ActionBy:    request.ActionBy,
		Action:      request.Action,
		ProductName: request.ProductName,
		Category:    request.Category,
	}

	document, err = implementation.RepositoryDocumentInterface.Create(ctx, tx, document)
	helpers.PanicIfError(err)

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
		helpers.PanicIfError(err)
		document.RiskDetail = append(document.RiskDetail, risk)
	}

	err = implementation.RepositoryDocumentSearchInterface.IndexProduct(implementation.EsClient, document)
	if err != nil {
		log.Println(err)
	}

	blastNotification(ctx, tx, document, implementation.RepositoryUserInterface, implementation.RepositoryNotificationInterface)

	return webDocument.DocumentModelToDocumentResponse(document)
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
		}
	}
}

func (implementation *ServiceDocumentImpl) SummaryDashboard(ctx context.Context, request webDocument.DocumentRequestSummaryDashboard) webDocument.DocumentResponseSummaryDashboard {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.SummaryDashboard(ctx, tx, &request)
	helpers.PanicIfError(err)

	documents, err := implementation.RepositoryDocumentInterface.GetProductCurrentYear(ctx, tx)
	helpers.PanicIfError(err)

	var summary webDocument.DocumentResponseSummaryDashboard
	for _, data := range documents {
		switch data.Action {
		case "approve":
			summary.SummaryAssessment.Release++
		case "reject":
			summary.SummaryAssessment.Return++
		case "submit":
			summary.SummaryAssessment.Received++
		case "update":
			summary.SummaryAssessment.Received++
		}
		summary.SummaryAssessment.Total++
	}

	return summary
}
