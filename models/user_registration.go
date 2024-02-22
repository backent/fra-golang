package models

import (
	"database/sql"
	"time"
)

type UserRegistration struct {
	Id        int
	Nik       string
	Name      string
	Email     string
	Status    string
	RejectBy  sql.NullInt32
	ApproveBy sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NullAbleUserRegistration struct {
	Id        sql.NullInt32
	Nik       sql.NullString
	Name      sql.NullString
	Email     sql.NullString
	Status    sql.NullString
	RejectBy  sql.NullInt32
	ApproveBy sql.NullInt32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type UserLdap struct {
	Nik   string
	Name  string
	Email string
}

var UserRegistrationTable string = "user_registrations"

func NullAbleUserRegistrationToUserRegistration(nullAbleUserRegistration NullAbleUserRegistration) UserRegistration {
	return UserRegistration{
		Id:        int(nullAbleUserRegistration.Id.Int32),
		Nik:       nullAbleUserRegistration.Nik.String,
		Name:      nullAbleUserRegistration.Name.String,
		Email:     nullAbleUserRegistration.Email.String,
		Status:    nullAbleUserRegistration.Status.String,
		RejectBy:  nullAbleUserRegistration.RejectBy,
		ApproveBy: nullAbleUserRegistration.ApproveBy,
		CreatedAt: nullAbleUserRegistration.CreatedAt.Time,
		UpdatedAt: nullAbleUserRegistration.UpdatedAt.Time,
	}
}
