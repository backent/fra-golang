package user

import (
	"github.com/backent/fra-golang/models"
)

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

type UserResponseWithDocumentsDetail struct {
	Id              int                `json:"id"`
	Nik             string             `json:"nik"`
	Name            string             `json:"name"`
	DocumentsDetail []documentResponse `json:"documents_detail"`
}

type documentResponse struct {
	Id         int    `json:"id"`          // id
	DocumentId string `json:"document_id"` // document_id
	UserId     int    `json:"user_id"`     // user_id
	RiskName   string `json:"risk_name"`   // risk_name
}

func UserModelToUserResponseWithDocumentsDetail(user models.User) UserResponseWithDocumentsDetail {
	return UserResponseWithDocumentsDetail{
		Id:              user.Id,
		Name:            user.Name,
		Nik:             user.Nik,
		DocumentsDetail: bulkDocumentModelToDocumentResponse(user.DocumentsDetail),
	}
}

func BulkUserModelToUserResponseWithDocumentsDetail(users []models.User) []UserResponseWithDocumentsDetail {
	var bulk []UserResponseWithDocumentsDetail
	for _, user := range users {
		bulk = append(bulk, UserModelToUserResponseWithDocumentsDetail(user))
	}
	return bulk
}

func documentModelToDocumentResponse(document models.Document) documentResponse {
	return documentResponse{
		Id:         document.Id,
		DocumentId: document.DocumentId,
		UserId:     document.UserId,
		RiskName:   document.RiskName,
	}
}

func bulkDocumentModelToDocumentResponse(documents []models.Document) []documentResponse {
	var bulk []documentResponse
	for _, document := range documents {
		bulk = append(bulk, documentModelToDocumentResponse(document))
	}
	return bulk
}
