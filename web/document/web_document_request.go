package document

import (
	"strings"

	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/web/risk"
)

type DocumentRequestCreate struct {
	Uuid        string                   `json:"uuid"`
	Action      string                   `json:"action" validate:"required,max=40"`        // action
	ProductName string                   `json:"product_name" validate:"required,max=100"` // product_name
	Risks       []risk.RiskRequestCreate `json:"risks" validate:"required,gt=0,dive"`      // risks

	CreatedBy int
	ActionBy  int
}

type DocumentRequestUpdate struct {
	Action      string `json:"action" validate:"required,max=40"`        // action
	ProductName string `json:"product_name" validate:"required,max=100"` // product_name

	Id        int
	Uuid      string
	CreatedBy int
	ActionBy  int
}
type DocumentRequestDelete struct {
	Id int `json:"id"`
}

type DocumentRequestFindById struct {
	Id int `json:"id"`
}

type DocumentRequestFindAll struct {
	WithDetail     bool
	take           int
	skip           int
	orderBy        string
	orderDirection string
	CreatedBy      int
	QueryAction    string
}

func (implementation *DocumentRequestFindAll) SetSkip(skip int) {
	implementation.skip = skip
}

func (implementation *DocumentRequestFindAll) SetTake(take int) {
	implementation.take = take
}

func (implementation *DocumentRequestFindAll) GetTake() int {
	return implementation.take
}

func (implementation *DocumentRequestFindAll) GetSkip() int {
	return implementation.skip
}

func (implementation *DocumentRequestFindAll) SetOrderBy(orderBy string) {
	implementation.orderBy = orderBy
}

func (implementation *DocumentRequestFindAll) SetOrderDirection(orderDirection string) {
	implementation.orderDirection = strings.ToUpper(orderDirection)
}

func (implementation *DocumentRequestFindAll) GetOrderBy() string {
	// set default order by
	if implementation.orderBy == "" {
		return "created_at"
	}
	return implementation.orderBy
}

func (implementation *DocumentRequestFindAll) GetOrderDirection() string {
	// set default order direction
	if implementation.orderDirection == "" {
		return "DESC"
	}
	return implementation.orderDirection
}

type DocumentRequestGetProductDistinct struct {
}

type DocumentRequestApprove struct {
	Id int `json:"id" validate:"required"`

	Document models.Document
}
