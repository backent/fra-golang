package models

import "database/sql"

type User struct {
	Id          int
	Nik         string
	Name        string
	Password    string
	Role        string
	RisksDetail []Risk
}

type NullAbleUser struct {
	Id       sql.NullInt32
	Nik      sql.NullString
	Name     sql.NullString
	Password sql.NullString
	Role     sql.NullString
}

var UserTable string = "users"

func NullAbleUserToUser(nullAbleUser NullAbleUser) User {
	return User{
		Id:       int(nullAbleUser.Id.Int32),
		Nik:      nullAbleUser.Nik.String,
		Name:     nullAbleUser.Name.String,
		Password: nullAbleUser.Password.String,
		Role:     nullAbleUser.Role.String,
	}
}
