package risk

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryRiskInterface interface {
	Create(ctx context.Context, tx *sql.Tx, risk models.Risk) (models.Risk, error)
	Update(ctx context.Context, tx *sql.Tx, risk models.Risk) (models.Risk, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.Risk, error)
	FindByDocumentId(ctx context.Context, tx *sql.Tx, riskId string) (models.Risk, error)
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Risk, int, error)
	FindAllWithUserDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.Risk, int, error)
}
