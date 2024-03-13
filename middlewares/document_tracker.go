package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesDocument "github.com/backent/fra-golang/repositories/document"
	webDocumentTracker "github.com/backent/fra-golang/web/document_tracker"
	"github.com/go-playground/validator/v10"
)

type DocumentTrackerMiddleware struct {
	Validate *validator.Validate
	repositoriesDocument.RepositoryDocumentInterface
	repositoriesAuth.RepositoryAuthInterface
}

func NewDocumentTrackerMiddleware(
	validator *validator.Validate,
	repositoriesDocument repositoriesDocument.RepositoryDocumentInterface,
	repositoriesAuth repositoriesAuth.RepositoryAuthInterface,
) *DocumentTrackerMiddleware {
	return &DocumentTrackerMiddleware{
		Validate:                    validator,
		RepositoryDocumentInterface: repositoriesDocument,
		RepositoryAuthInterface:     repositoriesAuth,
	}
}

func (implementation *DocumentTrackerMiddleware) TrackCount(ctx context.Context, tx *sql.Tx, request *webDocumentTracker.DocumentTrackerRequestTrackCount) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryDocumentInterface.FindByUUID(ctx, tx, request.Uuid)
	helpers.PanicIfError(err)
}
