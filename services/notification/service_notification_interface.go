package notification

import (
	"context"

	"github.com/backent/fra-golang/web/notification"
)

type ServiceNotificationInterface interface {
	ReadAll(ctx context.Context, request notification.NotificationRequestReadAll)
	FindAll(ctx context.Context, request notification.NotificationRequestFindAll) []notification.NotificationResponse
}
