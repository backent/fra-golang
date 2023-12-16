package document

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/google/uuid"
)

type ServiceDocumentImpl struct {
	DB *sql.DB
	repositoriesDocument.RepositoryDocumentInterface
	*middlewares.DocumentMiddleware
}

func NewServiceDocumentImpl(db *sql.DB, repositoriesDocument repositoriesDocument.RepositoryDocumentInterface, documentMiddleware *middlewares.DocumentMiddleware) ServiceDocumentInterface {
	return &ServiceDocumentImpl{
		DB:                          db,
		RepositoryDocumentInterface: repositoriesDocument,
		DocumentMiddleware:          documentMiddleware,
	}
}

func (implementation *ServiceDocumentImpl) Create(ctx context.Context, request webDocument.DocumentRequestCreate) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Create(ctx, tx, &request)

	document := models.Document{
		DocumentId:             uuid.New().String(),
		UserId:                 2,
		RiskName:               request.RiskName,
		FraudSchema:            request.FraudSchema,
		FraudMotive:            request.FraudMotive,
		FraudTechnique:         request.FraudTechnique,
		RiskSource:             request.RiskSource,
		RootCause:              request.RootCause,
		BisproControlProcedure: request.BisproControlProcedure,
		QualitativeImpact:      request.QualitativeImpact,
		LikehoodJustification:  request.LikehoodJustification,
		ImpactJustification:    request.ImpactJustification,
		StartegyAgreement:      request.StartegyAgreement,
		StrategyRecomendation:  request.StrategyRecomendation,
	}

	document, err = implementation.RepositoryDocumentInterface.Create(ctx, tx, document)
	helpers.PanicIfError(err)

	return webDocument.DocumentModelToDocumentResponse(document)
}
func (implementation *ServiceDocumentImpl) Update(ctx context.Context, request webDocument.DocumentRequestUpdate) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.Update(ctx, tx, &request)

	document := models.Document{
		RiskName:               request.RiskName,
		FraudSchema:            request.FraudSchema,
		FraudMotive:            request.FraudMotive,
		FraudTechnique:         request.FraudTechnique,
		RiskSource:             request.RiskSource,
		RootCause:              request.RootCause,
		BisproControlProcedure: request.BisproControlProcedure,
		QualitativeImpact:      request.QualitativeImpact,
		LikehoodJustification:  request.LikehoodJustification,
		ImpactJustification:    request.ImpactJustification,
		StartegyAgreement:      request.StartegyAgreement,
		StrategyRecomendation:  request.StrategyRecomendation,
	}

	document, err = implementation.RepositoryDocumentInterface.Update(ctx, tx, document)
	document.Id = request.Id
	document.DocumentId = request.DocumentId
	document.UserId = request.UserId
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
func (implementation *ServiceDocumentImpl) FindById(ctx context.Context, request webDocument.DocumentRequestFindById) webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.FindById(ctx, tx, &request)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	return webDocument.DocumentModelToDocumentResponse(document)
}
func (implementation *ServiceDocumentImpl) FindAll(ctx context.Context, request webDocument.DocumentRequestFindAll) []webDocument.DocumentResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentMiddleware.FindAll(ctx, tx, &request)

	documents, err := implementation.RepositoryDocumentInterface.FindAll(ctx, tx)
	helpers.PanicIfError(err)

	return webDocument.BulkDocumentModelToDocumentResponse(documents)
}
