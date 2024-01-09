package user

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryUserInterface interface {
	Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	Update(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int) ([]models.User, error)
	FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.User, error)
	FindAllWithDocumentsDetail(ctx context.Context, tx *sql.Tx, take int, skip int) ([]models.User, error)
}
