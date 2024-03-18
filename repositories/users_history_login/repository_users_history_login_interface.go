package users_history_login

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryUserHistoryLoginInterface interface {
	Create(ctx context.Context, tx *sql.Tx, userHistoryLogin models.UserHistoryLogin) error
	FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, year string, month string) ([]models.UserHistoryLogin, error)
}
