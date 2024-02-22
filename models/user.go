package models

import "database/sql"

type User struct {
	Id             int
	Nik            string
	Name           string
	Email          string
	ApplyStatus    string
	ApplyRejectBy  int
	ApplyApproveBy int
	Password       string
	Role           string
	RisksDetail    []Risk
}

type NullAbleUser struct {
	Id             sql.NullInt32
	Nik            sql.NullString
	Name           sql.NullString
	Email          sql.NullString
	ApplyStatus    sql.NullString
	ApplyRejectBy  sql.NullInt32
	ApplyApproveBy sql.NullInt32
	Password       sql.NullString
	Role           sql.NullString
}

var UserTable string = "users"

func NullAbleUserToUser(nullAbleUser NullAbleUser) User {
	return User{
		Id:             int(nullAbleUser.Id.Int32),
		Nik:            nullAbleUser.Nik.String,
		Name:           nullAbleUser.Name.String,
		Email:          nullAbleUser.Email.String,
		ApplyStatus:    nullAbleUser.ApplyStatus.String,
		ApplyApproveBy: int(nullAbleUser.ApplyApproveBy.Int32),
		ApplyRejectBy:  int(nullAbleUser.ApplyRejectBy.Int32),
		Password:       nullAbleUser.Password.String,
		Role:           nullAbleUser.Role.String,
	}
}
