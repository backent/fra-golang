package document

import (
	"time"

	"github.com/backent/fra-golang/models"
)

type DocumentResponse struct {
	Id          int       `json:"id"`           // id
	Uuid        string    `json:"uuid"`         // uuid
	CreatedBy   int       `json:"created_by"`   // created_by
	ActionBy    int       `json:"action_by"`    // action_by
	Action      string    `json:"action"`       // action
	ProductName string    `json:"product_name"` // product_name
	CreatedAt   time.Time `json:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at"`   // updated_at
}

func DocumentModelToDocumentResponse(document models.Document) DocumentResponse {
	return DocumentResponse{
		Id:          document.Id,
		Uuid:        document.Uuid,
		CreatedBy:   document.CreatedBy,
		ActionBy:    document.ActionBy,
		Action:      document.Action,
		ProductName: document.ProductName,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}
}

func BulkDocumentModelToDocumentResponse(documents []models.Document) []DocumentResponse {
	var bulk []DocumentResponse
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponse(document))
	}
	return bulk
}

type DocumentResponseWithDetail struct {
	Id          int       `json:"id"`           // id
	Uuid        string    `json:"uuid"`         // uuid
	CreatedBy   int       `json:"created_by"`   // created_by
	ActionBy    int       `json:"action_by"`    // action_by
	Action      string    `json:"action"`       // action
	ProductName string    `json:"product_name"` // product_name
	CreatedAt   time.Time `json:"created_at"`   // created_at
	UpdatedAt   time.Time `json:"updated_at"`   // updated_at

	UserDetail userResponse `json:"user_detail"` // user detail
}

type userResponse struct {
	Id   int    `json:"id"`
	Nik  string `json:"nik"`
	Name string `json:"name"`
}

func DocumentModelToDocumentResponseWithDetail(document models.Document) DocumentResponseWithDetail {
	return DocumentResponseWithDetail{
		Id:          document.Id,
		Uuid:        document.Uuid,
		CreatedBy:   document.CreatedBy,
		ActionBy:    document.ActionBy,
		Action:      document.Action,
		ProductName: document.ProductName,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
		UserDetail:  userModelToUserResponse(document.UserDetail),
	}
}

func BulkDocumentModelToDocumentResponseWithDetail(documents []models.Document) []DocumentResponseWithDetail {
	var bulk []DocumentResponseWithDetail
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponseWithDetail(document))
	}
	return bulk
}

func userModelToUserResponse(user models.User) userResponse {
	return userResponse{
		Id:   user.Id,
		Name: user.Name,
		Nik:  user.Nik,
	}
}
