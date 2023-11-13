package user

import "github.com/backent/fra-golang/models"

type UserResponse struct {
	Id   int    `json:"id"`
	Nik  string `json:"nik"`
	Name string `json:"name"`
}

func UserModelToUserResponse(user models.User) UserResponse {
	return UserResponse{
		Id:   user.Id,
		Name: user.Name,
		Nik:  user.Nik,
	}
}

func BulkUserModelToUserResponse(users []models.User) []UserResponse {
	var bulk []UserResponse
	for _, user := range users {
		bulk = append(bulk, UserModelToUserResponse(user))
	}
	return bulk
}
