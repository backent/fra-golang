package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	webDocument "github.com/backent/fra-golang/web/document"
	"github.com/go-playground/validator/v10"
)

type DocumentMiddleware struct {
	Validate *validator.Validate
	repositoriesDocument.RepositoryDocumentInterface
}

func NewDocumentMiddleware(validator *validator.Validate, repositoriesDocument repositoriesDocument.RepositoryDocumentInterface) *DocumentMiddleware {
	return &DocumentMiddleware{
		Validate:                    validator,
		RepositoryDocumentInterface: repositoriesDocument,
	}
}

func (implementation *DocumentMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestCreate) {
	err := implementation.Validate.Struct(request)
	helpers.PanifIfError(err)
}

func (implementation *DocumentMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestUpdate) {
	err := implementation.Validate.Struct(request)
	helpers.PanifIfError(err)

	document, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.Id = document.Id
	request.DocumentId = document.DocumentId
	request.UserId = document.UserId
}

func (implementation *DocumentMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestDelete) {

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindById) {

	_, err := implementation.RepositoryDocumentInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *DocumentMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webDocument.DocumentRequestFindAll) {
}
