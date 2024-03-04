package user

import (
	"github.com/backent/fra-golang/models"
)

type UserResponse struct {
	Id     int    `json:"id"`
	Nik    string `json:"nik"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Role   string `json:"role"`
}

func UserModelToUserResponse(user models.User) UserResponse {
	return UserResponse{
		Id:     user.Id,
		Name:   user.Name,
		Nik:    user.Nik,
		Email:  user.Email,
		Status: user.ApplyStatus,
		Role:   user.Role,
	}
}

func BulkUserModelToUserResponse(users []models.User) []UserResponse {
	var bulk []UserResponse
	for _, user := range users {
		bulk = append(bulk, UserModelToUserResponse(user))
	}
	return bulk
}

type UserResponseWithRisksDetail struct {
	Id          int            `json:"id"`
	Nik         string         `json:"nik"`
	Name        string         `json:"name"`
	RisksDetail []riskResponse `json:"risks_detail"`
}

type riskResponse struct {
	Id         int    `json:"id"`          // id
	DocumentId string `json:"document_id"` // document_id
	UserId     int    `json:"user_id"`     // user_id
	RiskName   string `json:"risk_name"`   // risk_name
}

func UserModelToUserResponseWithRisksDetail(user models.User) UserResponseWithRisksDetail {
	return UserResponseWithRisksDetail{
		Id:          user.Id,
		Name:        user.Name,
		Nik:         user.Nik,
		RisksDetail: bulkRiskModelToRiskResponse(user.RisksDetail),
	}
}

func BulkUserModelToUserResponseWithRisksDetail(users []models.User) []UserResponseWithRisksDetail {
	var bulk []UserResponseWithRisksDetail
	for _, user := range users {
		bulk = append(bulk, UserModelToUserResponseWithRisksDetail(user))
	}
	return bulk
}

func riskModelToRiskResponse(risk models.Risk) riskResponse {
	return riskResponse{
		Id:         risk.Id,
		DocumentId: "a", // temp handle with static value to remove error
		RiskName:   risk.RiskName,
	}
}

func bulkRiskModelToRiskResponse(risks []models.Risk) []riskResponse {
	var bulk []riskResponse
	for _, risk := range risks {
		bulk = append(bulk, riskModelToRiskResponse(risk))
	}
	return bulk
}
