package user_registration

import (
	"strings"

	"github.com/backent/fra-golang/models"
)

type UserRegistrationRequestApply struct {
	Nik      string `json:"nik" validate:"required,max=20"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type UserRegistrationRequestCheckUserLDAP struct {
	Nik      string `json:"nik" validate:"required,max=20"`
	NikValid bool
}

type UserRegistrationRequestApprove struct {
	Id   int    `json:"id" validate:"required"`
	Unit string `json:"unit" validate:"required,oneof=communication datacomm wireless internet"`

	User models.UserRegistration
}

type UserRegistrationRequestReject struct {
	Id int `json:"id" validate:"required"`

	User models.UserRegistration
}

type UserRegistrationRequestFindAll struct {
	QueryStatus    string
	take           int
	skip           int
	orderBy        string
	orderDirection string
}

func (implementation *UserRegistrationRequestFindAll) SetSkip(skip int) {
	implementation.skip = skip
}
func (implementation *UserRegistrationRequestFindAll) SetTake(take int) {
	implementation.take = take
}
func (implementation *UserRegistrationRequestFindAll) GetSkip() int {
	return implementation.skip
}
func (implementation *UserRegistrationRequestFindAll) GetTake() int {
	return implementation.take
}

func (implementation *UserRegistrationRequestFindAll) SetOrderBy(orderBy string) {
	implementation.orderBy = orderBy
}

func (implementation *UserRegistrationRequestFindAll) SetOrderDirection(orderDirection string) {
	implementation.orderDirection = strings.ToUpper(orderDirection)
}

func (implementation *UserRegistrationRequestFindAll) GetOrderBy() string {
	// set default order by
	if implementation.orderBy == "" {
		return "created_at"
	}
	return implementation.orderBy
}

func (implementation *UserRegistrationRequestFindAll) GetOrderDirection() string {
	// set default order direction
	if implementation.orderDirection == "" {
		return "ASC"
	}
	return implementation.orderDirection
}

type UserRegistrationRequestCurrentUserRegistration struct {
	UserRegistrationId int
}
