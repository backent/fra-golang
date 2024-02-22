package user_registration

import (
	"github.com/backent/fra-golang/models"
)

type UserRegistrationResponse struct {
	Id     int    `json:"id"`
	Nik    string `json:"nik"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func UserRegistrationModelToUserRegistrationResponse(user_registration models.UserRegistration) UserRegistrationResponse {
	return UserRegistrationResponse{
		Id:     user_registration.Id,
		Nik:    user_registration.Nik,
		Name:   user_registration.Name,
		Email:  user_registration.Email,
		Status: user_registration.Status,
	}
}

func BulkUserRegistrationModelToUserRegistrationResponse(user_registrations []models.UserRegistration) []UserRegistrationResponse {
	var bulk []UserRegistrationResponse
	for _, user_registration := range user_registrations {
		bulk = append(bulk, UserRegistrationModelToUserRegistrationResponse(user_registration))
	}
	return bulk
}
