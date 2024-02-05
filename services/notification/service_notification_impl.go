package notification

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRejectNote "github.com/backent/fra-golang/repositories/rejectnote"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	"github.com/backent/fra-golang/web/notification"
)

type ServiceNotificationImpl struct {
	DB *sql.DB
	repositoriesNotification.RepositoryNotificationInterface
	*middlewares.NotificationMiddleware
	repositoriesRisk.RepositoryRiskInterface
	repositoriesRejectNote.RepositoryRejectNoteInterface
}

func NewServiceNotificationImpl(
	db *sql.DB,
	repositoriesNotification repositoriesNotification.RepositoryNotificationInterface,
	notificationMiddleware *middlewares.NotificationMiddleware,
	repositoriesRisk repositoriesRisk.RepositoryRiskInterface,
	repositoriesRejectNote repositoriesRejectNote.RepositoryRejectNoteInterface,
) ServiceNotificationInterface {
	return &ServiceNotificationImpl{
		DB:                              db,
		RepositoryNotificationInterface: repositoriesNotification,
		NotificationMiddleware:          notificationMiddleware,
		RepositoryRiskInterface:         repositoriesRisk,
		RepositoryRejectNoteInterface:   repositoriesRejectNote,
	}
}
func (implementation *ServiceNotificationImpl) ReadAll(ctx context.Context, request notification.NotificationRequestReadAll) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.NotificationMiddleware.ReadAll(ctx, tx, &request)

	err = implementation.RepositoryNotificationInterface.ReadAll(ctx, tx, request.UserId)
	helpers.PanicIfError(err)

}
func (implementation *ServiceNotificationImpl) FindAll(ctx context.Context, request notification.NotificationRequestFindAll) []notification.NotificationResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.NotificationMiddleware.FindAll(ctx, tx, &request)

	notifications, err := implementation.RepositoryNotificationInterface.FindAll(ctx, tx, request.UserId)
	helpers.PanicIfError(err)

	if len(notifications) > 0 {
		return notification.BulkNotificationModelToNotificationResponse(notifications)
	}
	return []notification.NotificationResponse{}
}
