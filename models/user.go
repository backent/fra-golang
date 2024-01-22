package models

type User struct {
	Id          int
	Nik         string
	Name        string
	Password    string
	Role        string
	RisksDetail []Risk
}

var UserTable string = "users"
