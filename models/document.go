package models

import (
	"database/sql"
	"time"
)

type Document struct {
	Id          int       // id
	Uuid        string    // uuid
	CreatedBy   int       // created_by
	ActionBy    int       // action_by
	Action      string    // action
	ProductName string    // product_name
	CreatedAt   time.Time // created_at
	UpdatedAt   time.Time // updated_at

	RiskDetail            []Risk
	UserDetail            User
	RelatedDocumentDetail []RelatedDocument
}

type RelatedDocument struct {
	Id          int       // id
	ProductName string    // product_name
	Uuid        string    // uuid
	Action      string    // action
	CreatedAt   time.Time // created_at
	UserDetail  User
}

type NullAbleRelatedDocument struct {
	Id          sql.NullInt32  // id
	ProductName sql.NullString // product_name
	Uuid        sql.NullString // uuid
	Action      sql.NullString // action
	CreatedAt   sql.NullTime   // created_at
	User        NullAbleUser
}

func NullAbleRelatedDocumentToRelatedDocument(nullAbleRelatedDocument NullAbleRelatedDocument) RelatedDocument {
	return RelatedDocument{
		Id:          int(nullAbleRelatedDocument.Id.Int32),
		ProductName: nullAbleRelatedDocument.ProductName.String,
		Uuid:        nullAbleRelatedDocument.Uuid.String,
		Action:      nullAbleRelatedDocument.Action.String,
		CreatedAt:   nullAbleRelatedDocument.CreatedAt.Time,
	}
}

var DocumentTable string = "documents"
