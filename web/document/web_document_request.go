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

type DocumentRequestReject struct {
	Id int `json:"id" validate:"required"`

	RejectNote []RejectNoteRequest `json:"reject_note" validate:"required,gt=0,dive"`
	Document   models.Document
}

type RejectNoteRequest struct {
	RiskId                 int    `json:"risk_id" validate:"required"` // risk_id
	Fraud                  string `json:"fraud"`                       // fraud
	RiskSource             string `json:"risk_source"`                 // risk_source
	RootCause              string `json:"root_cause"`                  // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure"`    // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact"`          // qualitative_impact
	Assessment             string `json:"assessment"`                  // assessment
	Justification          string `json:"justification"`               // justification
	Strategy               string `json:"strategy"`                    // strategy

}

type DocumentRequestMonitoringList struct {
	WithDetail     bool
	take           int
	skip           int
	orderBy        string
	orderDirection string
	QueryAction    string
	QueryPeriod    int
	QueryName      string
}

func (implementation *DocumentRequestMonitoringList) SetSkip(skip int) {
	implementation.skip = skip
}

func (implementation *DocumentRequestMonitoringList) SetTake(take int) {
	implementation.take = take
}

func (implementation *DocumentRequestMonitoringList) GetTake() int {
	return implementation.take
}

func (implementation *DocumentRequestMonitoringList) GetSkip() int {
	return implementation.skip
}

func (implementation *DocumentRequestMonitoringList) SetOrderBy(orderBy string) {
	implementation.orderBy = orderBy
}

func (implementation *DocumentRequestMonitoringList) SetOrderDirection(orderDirection string) {
	implementation.orderDirection = strings.ToUpper(orderDirection)
}

func (implementation *DocumentRequestMonitoringList) GetOrderBy() string {
	// set default order by
	if implementation.orderBy == "" {
		return "created_at"
	}
	return implementation.orderBy
}

func (implementation *DocumentRequestMonitoringList) GetOrderDirection() string {
	// set default order direction
	if implementation.orderDirection == "" {
		return "DESC"
	}
	return implementation.orderDirection
}

type DocumentRequestTrackerProduct struct {
	QuerySearch string
}
