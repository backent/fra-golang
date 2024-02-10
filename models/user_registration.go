package models

import (
	"database/sql"
	"time"
)

type UserRegistration struct {
	Id        int
	Nik       string
	Status    string
	RejectBy  sql.NullInt32
	ApproveBy sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NullAbleUserRegistration struct {
	Id        sql.NullInt32
	Nik       sql.NullString
	Status    sql.NullString
	RejectBy  sql.NullInt32
	ApproveBy sql.NullInt32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

var UserRegistrationTable string = "user_registrations"

func NullAbleUserRegistrationToUserRegistration(nullAbleUserRegistration NullAbleUserRegistration) UserRegistration {
	return UserRegistration{
		Id:        int(nullAbleUserRegistration.Id.Int32),
		Nik:       nullAbleUserRegistration.Nik.String,
		Status:    nullAbleUserRegistration.Status.String,
		RejectBy:  nullAbleUserRegistration.RejectBy,
		ApproveBy: nullAbleUserRegistration.ApproveBy,
		CreatedAt: nullAbleUserRegistration.CreatedAt.Time,
		UpdatedAt: nullAbleUserRegistration.UpdatedAt.Time,
	}
}
