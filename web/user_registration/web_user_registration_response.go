package user_registration

import (
	"github.com/backent/fra-golang/models"
)

type UserRegistrationResponse struct {
	Id     int    `json:"id"`
	Nik    string `json:"nik"`
	Status string `json:"name"`
}

func UserRegistrationModelToUserRegistrationResponse(user_registration models.UserRegistration) UserRegistrationResponse {
	return UserRegistrationResponse{
		Id:     user_registration.Id,
		Nik:    user_registration.Nik,
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
