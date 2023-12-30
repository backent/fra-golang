package user

type UserRequestCreate struct {
	Nik      string `json:"nik" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRequestUpdate struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password"`

	Nik             string `json:"nik"`
	CurrentPassword string
}

type UserRequestDelete struct {
	Id int `json:"id"`
}

type UserRequestFindById struct {
	Id int `json:"id"`
}

type UserRequestFindAll struct {
	WithDocument bool
}
