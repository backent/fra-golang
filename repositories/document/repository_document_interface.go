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
	FindByUUID(ctx context.Context, tx *sql.Tx, documentUuid string) (models.Document, error)
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error)
	FindAllWithDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error)
}
