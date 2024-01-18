package middlewares

import (
	"context"
	"database/sql"
	"strings"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	webRisk "github.com/backent/fra-golang/web/risk"
	"github.com/go-playground/validator/v10"
)

type RiskMiddleware struct {
	Validate *validator.Validate
	repositoriesRisk.RepositoryRiskInterface
	repositoriesAuth.RepositoryAuthInterface
}

func NewRiskMiddleware(validator *validator.Validate, repositoriesRisk repositoriesRisk.RepositoryRiskInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface) *RiskMiddleware {
	return &RiskMiddleware{
		Validate:                validator,
		RepositoryRiskInterface: repositoriesRisk,
		RepositoryAuthInterface: repositoriesAuth,
	}
}

func (implementation *RiskMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webRisk.RiskRequestCreate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	request.AssessmentImpact = strings.ToLower(request.AssessmentImpact)
	request.AssessmentLikehood = strings.ToLower(request.AssessmentLikehood)
	request.AssessmentRiskLevel = strings.ToLower(request.AssessmentRiskLevel)

}

func (implementation *RiskMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webRisk.RiskRequestUpdate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	risk, err := implementation.RepositoryRiskInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.AssessmentImpact = strings.ToLower(request.AssessmentImpact)
	request.AssessmentLikehood = strings.ToLower(request.AssessmentLikehood)
	request.AssessmentRiskLevel = strings.ToLower(request.AssessmentRiskLevel)

	request.Id = risk.Id
	request.DocumentId = "a" // temp handle with static value to remove error
}

func (implementation *RiskMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webRisk.RiskRequestDelete) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryRiskInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *RiskMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webRisk.RiskRequestFindById) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryRiskInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *RiskMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webRisk.RiskRequestFindAll) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}
