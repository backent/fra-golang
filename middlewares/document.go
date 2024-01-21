package middlewares

import (
	"context"
	"database/sql"
	"strings"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/go-playground/validator/v10"
)

type DocumentMiddleware struct {
	Validate *validator.Validate
	repositoriesDocument.RepositoryDocumentInterface
	repositoriesAuth.RepositoryAuthInterface
}

func NewDocumentMiddleware(validator *validator.Validate, repositoriesDocument repositoriesDocument.RepositoryDocumentInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface) *DocumentMiddleware {
	return &DocumentMiddleware{
		Validate:                    validator,
		RepositoryDocumentInterface: repositoriesDocument,
		RepositoryAuthInterface:     repositoriesAuth,
	}
}

func (implementation *DocumentMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestCreate) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	for index := range request.Risks {
		request.Risks[index].AssessmentImpact = strings.ToLower(request.Risks[index].AssessmentImpact)
		request.Risks[index].AssessmentLikehood = strings.ToLower(request.Risks[index].AssessmentLikehood)
		request.Risks[index].AssessmentRiskLevel = strings.ToLower(request.Risks[index].AssessmentRiskLevel)
	}

	request.CreatedBy = userId
	request.ActionBy = userId
}

func (implementation *DocumentMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestUpdate) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.Id = document.Id
	request.Uuid = document.Uuid
	request.CreatedBy = document.CreatedBy
	request.ActionBy = userId
}

func (implementation *DocumentMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestDelete) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindById) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindAll) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *DocumentMiddleware) GetProductDistinct(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestGetProductDistinct) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}
