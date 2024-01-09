package document

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryDocumentInterface interface {
	Create(ctx context.Context, tx *sql.Tx, document models.Document) (models.Document, error)
	Update(ctx context.Context, tx *sql.Tx, document models.Document) (models.Document, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.Document, error)
	FindByDocumentId(ctx context.Context, tx *sql.Tx, documentId string) (models.Document, error)
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int) ([]models.Document, error)
	FindAllWithUserDetail(ctx context.Context, tx *sql.Tx, take int, skip int) ([]models.Document, error)
}
