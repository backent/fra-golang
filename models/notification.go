package models

import "time"

type Notification struct {
	Id         int
	UserId     int
	DocumentId int
	Title      string
	Subtitle   string
	Read       int
	Action     string
	CreatedAt  time.Time
}

var NotificationTable = "notifications"
