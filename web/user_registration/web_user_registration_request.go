package user_registration

import "strings"

type UserRegistrationRequestApply struct {
	Nik string `json:"nik" validate:"required,max=20"`
}

type UserRegistrationRequestFindAll struct {
	WithRisk       bool
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
		return "DESC"
	}
	return implementation.orderDirection
}

type UserRegistrationRequestCurrentUserRegistration struct {
	UserRegistrationId int
}
