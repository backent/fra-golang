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
	Id        int       // id
	CreatedAt time.Time // created_at
}

type NullAbleRelatedDocument struct {
	Id        sql.NullInt32 // id
	CreatedAt sql.NullTime  // created_at
}

func NullAbleRelatedDocumentToRelatedDocument(nullAbleRelatedDocument NullAbleRelatedDocument) RelatedDocument {
	return RelatedDocument{
		Id:        int(nullAbleRelatedDocument.Id.Int32),
		CreatedAt: nullAbleRelatedDocument.CreatedAt.Time,
	}
}

var DocumentTable string = "documents"
