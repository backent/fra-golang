package document_tracker

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/backent/fra-golang/models"
)

type RepositoryDocumentTrackerImpl struct {
}

func NewRepositoryDocumentTrackerImpl() RepositoryDocumentTrackerInterface {
	return &RepositoryDocumentTrackerImpl{}
}

func (implementation *RepositoryDocumentTrackerImpl) Create(ctx context.Context, tx *sql.Tx, documentUUID string, documentCreatedAt time.Time) error {
	query := fmt.Sprintf(`INSERT INTO %s (document_uuid, document_created_at) VALUES (?, ?)`, models.DocumentTrackerTable)

	_, err := tx.ExecContext(ctx, query, documentUUID, documentCreatedAt)
	return err
}

func (implementation *RepositoryDocumentTrackerImpl) Increment(ctx context.Context, tx *sql.Tx, documentUUID string, viewIncrementBy int, searchIncrementBy int) error {
	query := fmt.Sprintf("UPDATE %s SET `viewed_count` = `viewed_count` + ?, `searched_count` = `searched_count` + ?  WHERE document_uuid = ?", models.DocumentTrackerTable)

	_, err := tx.ExecContext(ctx, query, viewIncrementBy, searchIncrementBy, documentUUID)
	return err
}

func (implementation *RepositoryDocumentTrackerImpl) GetByUUId(ctx context.Context, tx *sql.Tx, documentUUID string) (models.DocumentTracker, error) {
	var documentTracker models.DocumentTracker

	query := fmt.Sprintf("SELECT id, document_uuid, viewed_count, searched_count, document_created_at FROM %s WHERE document_uuid = ? LIMIT 1", models.DocumentTrackerTable)

	rows, err := tx.QueryContext(ctx, query, documentUUID)
	if err != nil {
		return documentTracker, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&documentTracker.Id, &documentTracker.DocumentUUID, &documentTracker.ViewedCount, &documentTracker.SearchedCount, &documentTracker.DocumentCreatedAt)
		if err != nil {
			return documentTracker, err
		}
	} else {
		return documentTracker, errors.New("document tracker not found")
	}

	return documentTracker, nil

}
