package risk

import (
	"strings"

	"github.com/backent/fra-golang/web"
)

type RiskRequestCreate struct {
	RiskName               string `json:"risk_name" validate:"required,max=50"`                  // risk_name
	FraudSchema            string `json:"fraud_schema" validate:"required,max=1000"`             // fraud_schema
	FraudMotive            string `json:"fraud_motive" validate:"required,max=1000"`             // fraud_motive
	FraudTechnique         string `json:"fraud_technique" validate:"required,max=1000"`          // fraud_technique
	RiskSource             string `json:"risk_source" validate:"required,max=1000"`              // risk_source
	RootCause              string `json:"root_cause" validate:"required,max=1000"`               // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure" validate:"required,max=1000"` // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact" validate:"required,max=1000"`       // qualitative_impact
	LikehoodJustification  string `json:"likehood_justification" validate:"required,max=1000"`   // likehood_justification
	ImpactJustification    string `json:"impact_justification" validate:"required,max=1000"`     // impact_justification
	StartegyAgreement      string `json:"strategy_agreement" validate:"required,max=1000"`       // strategy_agreement
	StrategyRecomendation  string `json:"strategy_recomendation" validate:"required,max=1000"`   // strategy_recomendation
	AssessmentLikehood     string `json:"assessment_likehood" validate:"required"`               // assessment_likehood
	AssessmentImpact       string `json:"assessment_impact" validate:"required"`                 // assessment_impact
	AssessmentRiskLevel    string `json:"assessment_risk_level" validate:"required"`             // assessment_risk_level
	Action                 string `json:"action" validate:"required,max=50"`                     // action

	DocumentId string
	UserId     int
	ActionBy   int
}

type RiskRequestUpdate struct {
	RiskName               string `json:"risk_name" validate:"required"`                // risk_name
	FraudSchema            string `json:"fraud_schema" validate:"required"`             // fraud_schema
	FraudMotive            string `json:"fraud_motive" validate:"required"`             // fraud_motive
	FraudTechnique         string `json:"fraud_technique" validate:"required"`          // fraud_technique
	RiskSource             string `json:"risk_source" validate:"required"`              // risk_source
	RootCause              string `json:"root_cause" validate:"required"`               // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure" validate:"required"` // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact" validate:"required"`       // qualitative_impact
	LikehoodJustification  string `json:"likehood_justification" validate:"required"`   // likehood_justification
	ImpactJustification    string `json:"impact_justification" validate:"required"`     // impact_justification
	StartegyAgreement      string `json:"strategy_agreement" validate:"required"`       // strategy_agreement
	StrategyRecomendation  string `json:"strategy_recomendation" validate:"required"`   // strategy_recomendation
	Action                 string `json:"action" validate:"required"`                   // action
	AssessmentLikehood     string `json:"assessment_likehood" validate:"required"`      // assessment_likehood
	AssessmentImpact       string `json:"assessment_impact" validate:"required"`        // assessment_impact
	AssessmentRiskLevel    string `json:"assessment_risk_level" validate:"required"`    // assessment_risk_level

	Id         int
	DocumentId string
	UserId     int
	ActionBy   int
}
type RiskRequestDelete struct {
	Id int `json:"id"`
}

type RiskRequestFindById struct {
	Id int `json:"id"`
}

type RiskRequestFindAll struct {
	WithUser       bool
	take           int
	skip           int
	orderBy        string
	orderDirection string
}

func NewRiskRequestFindAll() web.RequestPagination {
	return &RiskRequestFindAll{}
}

func (implementation *RiskRequestFindAll) SetSkip(skip int) {
	implementation.skip = skip
}

func (implementation *RiskRequestFindAll) SetTake(take int) {
	implementation.take = take
}

func (implementation *RiskRequestFindAll) GetTake() int {
	return implementation.take
}

func (implementation *RiskRequestFindAll) GetSkip() int {
	return implementation.skip
}

func (implementation *RiskRequestFindAll) SetOrderBy(orderBy string) {
	implementation.orderBy = orderBy
}

func (implementation *RiskRequestFindAll) SetOrderDirection(orderDirection string) {
	implementation.orderDirection = strings.ToUpper(orderDirection)
}

func (implementation *RiskRequestFindAll) GetOrderBy() string {
	// set default order by
	if implementation.orderBy == "" {
		return "created_at"
	}
	return implementation.orderBy
}

func (implementation *RiskRequestFindAll) GetOrderDirection() string {
	// set default order direction
	if implementation.orderDirection == "" {
		return "DESC"
	}
	return implementation.orderDirection
}
