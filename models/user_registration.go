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
	RejectBy  int
	ApproveBy int
	Unit      string
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
	Unit      sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type UserLdap struct {
	Nik   string
	Name  string
	Email string
}

var UserRegistrationTable string = "users"

func NullAbleUserRegistrationToUserRegistration(nullAbleUserRegistration NullAbleUserRegistration) UserRegistration {
	return UserRegistration{
		Id:        int(nullAbleUserRegistration.Id.Int32),
		Nik:       nullAbleUserRegistration.Nik.String,
		Name:      nullAbleUserRegistration.Name.String,
		Email:     nullAbleUserRegistration.Email.String,
		Status:    nullAbleUserRegistration.Status.String,
		RejectBy:  int(nullAbleUserRegistration.RejectBy.Int32),
		ApproveBy: int(nullAbleUserRegistration.ApproveBy.Int32),
		CreatedAt: nullAbleUserRegistration.CreatedAt.Time,
		UpdatedAt: nullAbleUserRegistration.UpdatedAt.Time,
	}
}
