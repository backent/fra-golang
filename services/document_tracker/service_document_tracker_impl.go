package document_tracker

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	repositoriesDocumentTracker "github.com/backent/fra-golang/repositories/document_tracker"
	webDocumentTracker "github.com/backent/fra-golang/web/document_tracker"
)

type ServiceDocumentTrackerImpl struct {
	DB *sql.DB
	repositoriesDocumentTracker.RepositoryDocumentTrackerInterface
	*middlewares.DocumentTrackerMiddleware
}

func NewServiceDocumentTrackerImpl(db *sql.DB, repositoriesDocumentTracker repositoriesDocumentTracker.RepositoryDocumentTrackerInterface, userMiddleware *middlewares.DocumentTrackerMiddleware) ServiceDocumentTrackerInterface {
	return &ServiceDocumentTrackerImpl{
		DB:                                 db,
		RepositoryDocumentTrackerInterface: repositoriesDocumentTracker,
		DocumentTrackerMiddleware:          userMiddleware,
	}
}

func (implementation *ServiceDocumentTrackerImpl) TrackCount(ctx context.Context, request webDocumentTracker.DocumentTrackerRequestTrackCount) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.DocumentTrackerMiddleware.TrackCount(ctx, tx, &request)

	switch request.Type {
	case "view":
		err = implementation.RepositoryDocumentTrackerInterface.Increment(ctx, tx, request.Uuid, 1, 0)
		helpers.PanicIfError(err)
	case "search":
		err = implementation.RepositoryDocumentTrackerInterface.Increment(ctx, tx, request.Uuid, 0, 1)
		helpers.PanicIfError(err)
	}

}
