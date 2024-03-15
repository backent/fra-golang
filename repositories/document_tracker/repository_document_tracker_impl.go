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

func (implementation *RepositoryDocumentTrackerImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
	take int,
	skip int,
	orderBy string,
	orderDirection string,
	year string,
) ([]models.DocumentTracker, int, error) {

	query := fmt.Sprintf(`
	WITH main_table AS (
		SELECT * FROM %s WHERE YEAR(document_created_at) = ?
	),
	main_table_document AS (
		SELECT * FROM %s WHERE action != 'approve'
	), group_by_uuid AS (
		SELECT d1.*
		FROM main_table_document d1
		JOIN (
				SELECT uuid, MAX(id) AS max_id
				FROM main_table_document
				GROUP BY uuid
		) d2 ON d1.uuid = d2.uuid AND d1.id = d2.max_id
	), main_table_document_after_grouped AS (
		SELECT * FROM group_by_uuid
	)

	SELECT 
	 a.id,
	 a.document_uuid,
	 a.viewed_count,
	 a.searched_count,
	 c.product_name,
	 c.category,
	 b.count

	FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ? ) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true
	LEFT JOIN main_table_document_after_grouped c ON a.document_uuid = c.uuid
	`, models.DocumentTrackerTable, models.DocumentTable, orderBy, orderDirection)

	rows, err := tx.QueryContext(ctx, query, year, skip, take)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var documentTrackers []models.DocumentTracker
	var total int

	for rows.Next() {
		var documentTracker models.DocumentTracker
		var document models.Document

		err := rows.Scan(
			&documentTracker.Id,
			&documentTracker.DocumentUUID,
			&documentTracker.ViewedCount,
			&documentTracker.SearchedCount,
			&document.ProductName,
			&document.Category,
			&total,
		)
		if err != nil {
			return nil, 0, err
		}

		documentTracker.DocumentDetail = document
		documentTrackers = append(documentTrackers, documentTracker)
	}
	return documentTrackers, total, nil
}
