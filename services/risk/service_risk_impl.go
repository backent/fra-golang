package risk

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	webRisk "github.com/backent/fra-golang/web/risk"
)

type ServiceRiskImpl struct {
	DB *sql.DB
	repositoriesRisk.RepositoryRiskInterface
	*middlewares.RiskMiddleware
}

func NewServiceRiskImpl(db *sql.DB, repositoriesRisk repositoriesRisk.RepositoryRiskInterface, riskMiddleware *middlewares.RiskMiddleware) ServiceRiskInterface {
	return &ServiceRiskImpl{
		DB:                      db,
		RepositoryRiskInterface: repositoriesRisk,
		RiskMiddleware:          riskMiddleware,
	}
}

func (implementation *ServiceRiskImpl) Create(ctx context.Context, request webRisk.RiskRequestCreate) webRisk.RiskResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.RiskMiddleware.Create(ctx, tx, &request)

	risk := models.Risk{
		DocumentId:             1, // temp handle with static just to remove error
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
		AssessmentLikehood:     request.AssessmentLikehood,
		AssessmentImpact:       request.AssessmentImpact,
		AssessmentRiskLevel:    request.AssessmentRiskLevel,
	}

	risk, err = implementation.RepositoryRiskInterface.Create(ctx, tx, risk)
	helpers.PanicIfError(err)

	return webRisk.RiskModelToRiskResponse(risk)
}
func (implementation *ServiceRiskImpl) Update(ctx context.Context, request webRisk.RiskRequestUpdate) webRisk.RiskResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.RiskMiddleware.Update(ctx, tx, &request)

	risk := models.Risk{
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

	risk, err = implementation.RepositoryRiskInterface.Update(ctx, tx, risk)
	risk.Id = request.Id
	risk.DocumentId = 1 // temp handle with static value to remove error
	helpers.PanicIfError(err)

	return webRisk.RiskModelToRiskResponse(risk)
}
func (implementation *ServiceRiskImpl) Delete(ctx context.Context, request webRisk.RiskRequestDelete) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.RiskMiddleware.Delete(ctx, tx, &request)

	err = implementation.RepositoryRiskInterface.Delete(ctx, tx, request.Id)
	helpers.PanicIfError(err)

}
func (implementation *ServiceRiskImpl) FindById(ctx context.Context, request webRisk.RiskRequestFindById) webRisk.RiskResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.RiskMiddleware.FindById(ctx, tx, &request)

	risk, err := implementation.RepositoryRiskInterface.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	return webRisk.RiskModelToRiskResponse(risk)
}
func (implementation *ServiceRiskImpl) FindAll(ctx context.Context, request webRisk.RiskRequestFindAll) ([]webRisk.RiskResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.RiskMiddleware.FindAll(ctx, tx, &request)

	risks, total, err := implementation.RepositoryRiskInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection())
	helpers.PanicIfError(err)

	return webRisk.BulkRiskModelToRiskResponse(risks), total
}
