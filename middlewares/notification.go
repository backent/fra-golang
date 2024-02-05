package middlewares

import (
	"context"
	"database/sql"

	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesNotification "github.com/backent/fra-golang/repositories/notification"
	repositoriesRisk "github.com/backent/fra-golang/repositories/risk"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webNotification "github.com/backent/fra-golang/web/notification"
	"github.com/go-playground/validator/v10"
)

type NotificationMiddleware struct {
	Validate *validator.Validate
	repositoriesNotification.RepositoryNotificationInterface
	repositoriesAuth.RepositoryAuthInterface
	repositoriesUser.RepositoryUserInterface
	repositoriesRisk.RepositoryRiskInterface
}

func NewNotificationMiddleware(
	validator *validator.Validate,
	repositoriesNotification repositoriesNotification.RepositoryNotificationInterface,
	repositoriesAuth repositoriesAuth.RepositoryAuthInterface,
	repositoriesUser repositoriesUser.RepositoryUserInterface,
	repositoriesRisk repositoriesRisk.RepositoryRiskInterface,
) *NotificationMiddleware {
	return &NotificationMiddleware{
		Validate:                        validator,
		RepositoryNotificationInterface: repositoriesNotification,
		RepositoryAuthInterface:         repositoriesAuth,
		RepositoryUserInterface:         repositoriesUser,
		RepositoryRiskInterface:         repositoriesRisk,
	}
}

func (implementation *NotificationMiddleware) ReadAll(ctx context.Context, tx *sql.Tx, request *webNotification.NotificationRequestReadAll) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	request.UserId = userId
}

func (implementation *NotificationMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webNotification.NotificationRequestFindAll) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	request.UserId = userId

}
