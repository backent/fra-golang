package models

type User struct {
	Id              int
	Nik             string
	Name            string
	Password        string
	DocumentsDetail []Document
}

var UserTable string = "users"
