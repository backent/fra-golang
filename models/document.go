package models

import "time"

type Document struct {
	Id          int       // id
	Uuid        string    // uuid
	CreatedBy   int       // created_by
	ActionBy    int       // action_by
	Action      string    // action
	ProductName string    // product_name
	CreatedAt   time.Time // created_at
	UpdatedAt   time.Time // updated_at

	RiskDetail Risk
	UserDetail User
}

var DocumentTable string = "documents"
