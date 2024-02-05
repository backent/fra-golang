package notification

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryNotificationImpl struct {
}

func NewRepositoryNotificationImpl() RepositoryNotificationInterface {
	return &RepositoryNotificationImpl{}
}
func (implementation *RepositoryNotificationImpl) ReadAll(ctx context.Context, tx *sql.Tx, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET `read` = 1 WHERE user_id = ?", models.NotificationTable)

	_, err := tx.ExecContext(ctx, query, userId)
	return err
}
func (implementation *RepositoryNotificationImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]models.Notification, error) {

	query := fmt.Sprintf("SELECT id, user_id, document_id, title, subtitle, action, `read`, created_at FROM %s WHERE user_id = ? ORDER BY id DESC", models.NotificationTable)

	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err = rows.Scan(
			&notification.Id,
			&notification.UserId,
			&notification.DocumentId,
			&notification.Title,
			&notification.Subtitle,
			&notification.Action,
			&notification.Read,
			&notification.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	return notifications, nil
}

func (implementation *RepositoryNotificationImpl) Create(ctx context.Context, tx *sql.Tx, notification models.Notification) (models.Notification, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, document_id, title, subtitle, action) VALUES (?, ?, ?, ?, ?)
	`, models.NotificationTable)

	result, err := tx.ExecContext(ctx, query, notification.UserId, notification.DocumentId, notification.Title, notification.Subtitle, notification.Action)
	if err != nil {
		return notification, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return notification, err
	}

	notification.Id = int(id)

	return notification, err
}
