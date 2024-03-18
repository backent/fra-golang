package models

import (
	"time"
)

type UserHistoryLogin struct {
	Id         int
	UserId     int
	CreatedAt  time.Time
	UserDetail User
}

var UserHistoryLoginTable string = "users_history_login"
