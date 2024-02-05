package notification

import (
	"time"

	"github.com/backent/fra-golang/models"
)

type NotificationResponse struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	DocumentId int       `json:"document_id"`
	Title      string    `json:"title"`
	Subtitle   string    `json:"subtitle"`
	Read       int       `json:"read"`
	Action     string    `json:"action"`
	CreatedAt  time.Time `json:"created_at"`
}

func NotificationModelToNotificationResponse(notification models.Notification) NotificationResponse {
	return NotificationResponse{
		Id:         notification.Id,
		UserId:     notification.UserId,
		DocumentId: notification.DocumentId,
		Title:      notification.Title,
		Subtitle:   notification.Subtitle,
		Read:       notification.Read,
		Action:     notification.Action,
		CreatedAt:  notification.CreatedAt,
	}
}

func BulkNotificationModelToNotificationResponse(notifications []models.Notification) []NotificationResponse {
	var bulk []NotificationResponse
	for _, notification := range notifications {
		bulk = append(bulk, NotificationModelToNotificationResponse(notification))
	}

	return bulk
}
