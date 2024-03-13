package document_tracker

import (
	"context"
	"database/sql"
	"time"

	"github.com/backent/fra-golang/models"
)

type RepositoryDocumentTrackerInterface interface {
	Create(ctx context.Context, tx *sql.Tx, documentUUID string, documentCreatedAt time.Time) error
	Increment(ctx context.Context, tx *sql.Tx, documentUUID string, viewIncrementBy int, searchIncrementBy int) error
	GetByUUId(ctx context.Context, tx *sql.Tx, documentUUID string) (models.DocumentTracker, error)
}
