package document

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	webDocument "github.com/backent/fra-golang/web/document"
)

type ServiceDocumentImpl struct {
	DB *sql.DB
	repositoriesDocument.RepositoryDocumentInterface
	*middlewares.DocumentMiddleware
	repositoriesRisk.RepositoryRiskInterface
	repositoriesRejectNote.RepositoryRejectNoteInterface
}

func NewServiceDocumentImpl(
	db *sql.DB,
	repositoriesDocument repositoriesDocument.RepositoryDocumentInterface,
	documentMiddleware *middlewares.DocumentMiddleware,
	repositoriesRisk repositoriesRisk.RepositoryRiskInterface,
	repositoriesRejectNote repositoriesRejectNote.RepositoryRejectNoteInterface,
) ServiceDocumentInterface {
	return &ServiceDocumentImpl{
		DB:                            db,
		RepositoryDocumentInterface:   repositoriesDocument,
		DocumentMiddleware:            documentMiddleware,
		RepositoryRiskInterface:       repositoriesRisk,
		RepositoryRejectNoteInterface: repositoriesRejectNote,
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
	}

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

	documents, total, err := implementation.RepositoryDocumentInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.CreatedBy, request.QueryAction)
	helpers.PanicIfError(err)

	if documents == nil {
		return []webDocument.DocumentResponse{}, total
	}

	return webDocument.BulkDocumentModelToDocumentResponse(documents), total
}
func (implementation *ServiceDocumentImpl) FindAllWithDetail(ctx context.Context, request webDocument.DocumentRequestFindAll) ([]webDocument.DocumentResponseWithDetail, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.FindAll(ctx, tx, &request)

	documents, total, err := implementation.RepositoryDocumentInterface.FindAllWithDetail(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection())
	helpers.PanicIfError(err)
	if documents != nil {
		return webDocument.BulkDocumentModelToDocumentResponseWithDetail(documents), total
	} else {
		return []webDocument.DocumentResponseWithDetail{}, total
	}

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
		fmt.Println(found)
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
}
