package user

import (
	"strings"

	"github.com/backent/fra-golang/models"
)

type UserRequestCreate struct {
	Nik      string `json:"nik" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRequestUpdate struct {
	Id              int    `json:"id"`
	Name            string `json:"name" validate:"required"`
	Password        string `json:"password"`
	IsEmptyPassword bool   `json:"is_empty_password"`
	Unit            string `json:"unit" validate:"required,oneof=communication datacomm wireless internet"`
	Role            string `json:"role" validate:"required,oneof=superadmin reviewer author guest"`

	Nik             string `json:"nik"`
	CurrentPassword string
	User            models.User
}

type UserRequestDelete struct {
	Id int `json:"id"`
}

type UserRequestFindById struct {
	Id int `json:"id"`
}

type UserRequestFindAll struct {
	WithRisk       bool
	take           int
	skip           int
	orderBy        string
	orderDirection string
	QueryStatus    string
	QuerySearch    string
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

type UserRequestCurrentUser struct {
	UserId int
}
