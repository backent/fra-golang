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
	FindByUUID(ctx context.Context, tx *sql.Tx, documentUuid string) ([]models.Document, error)
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string, documentAction string, documentCategory string) ([]models.Document, int, error)
	FindAllWithDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Document, int, error)
	GetProductDistinct(ctx context.Context, tx *sql.Tx) ([]models.Document, error)
	FindAllNoGroup(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string, documentAction string, month int, name string) ([]models.Document, int, error)
	GetNonDraftProductByUUID(ctx context.Context, tx *sql.Tx, uuid string) ([]models.Document, error)
	TrackerProductByName(ctx context.Context, tx *sql.Tx, name string) ([]models.Document, error)
}
