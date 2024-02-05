package notification

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryNotificationInterface interface {
	ReadAll(ctx context.Context, tx *sql.Tx, userId int) error
	FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]models.Notification, error)
	Create(ctx context.Context, tx *sql.Tx, notification models.Notification) (models.Notification, error)
}
