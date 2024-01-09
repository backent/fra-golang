package user

import "strings"

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
	WithDocument   bool
	take           int
	skip           int
	orderBy        string
	orderDirection string
}

func (implementation *UserRequestFindAll) SetSkip(skip int) {
	implementation.skip = skip
}
func (implementation *UserRequestFindAll) SetTake(take int) {
	implementation.take = take
}
func (implementation *UserRequestFindAll) GetSkip() int {
	return implementation.skip
}
func (implementation *UserRequestFindAll) GetTake() int {
	return implementation.take
}

func (implementation *UserRequestFindAll) SetOrderBy(orderBy string) {
	implementation.orderBy = orderBy
}

func (implementation *UserRequestFindAll) SetOrderDirection(orderDirection string) {
	implementation.orderDirection = strings.ToUpper(orderDirection)
}

func (implementation *UserRequestFindAll) GetOrderBy() string {
	// set default order by
	if implementation.orderBy == "" {
		return "created_at"
	}
	return implementation.orderBy
}

func (implementation *UserRequestFindAll) GetOrderDirection() string {
	// set default order direction
	if implementation.orderDirection == "" {
		return "DESC"
	}
	return implementation.orderDirection
}
