package user

type UserRequestCreate struct {
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRequestUpdate struct {
	Id       int    `json:"id"`
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRequestDelete struct {
	Id int `json:"id"`
}

type UserRequestFindById struct {
	Id int `json:"id"`
}

type UserRequestFindAll struct {
	Id int `json:"id"`
}
