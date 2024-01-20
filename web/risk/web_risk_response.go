package risk

import (
	"time"

	"github.com/backent/fra-golang/models"
)

type RiskResponse struct {
	Id                     int       `json:"id"`                       // id
	DocumentId             string    `json:"document_id"`              // document_id
	UserId                 int       `json:"user_id"`                  // user_id
	RiskName               string    `json:"risk_name"`                // risk_name
	FraudSchema            string    `json:"fraud_schema"`             // fraud_schema
	FraudMotive            string    `json:"fraud_motive"`             // fraud_motive
	FraudTechnique         string    `json:"fraud_technique"`          // fraud_technique
	RiskSource             string    `json:"risk_source"`              // risk_source
	RootCause              string    `json:"root_cause"`               // root_cause
	BisproControlProcedure string    `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string    `json:"qualitative_impact"`       // qualitative_impact
	LikehoodJustification  string    `json:"likehood_justification"`   // likehood_justification
	ImpactJustification    string    `json:"impact_justification"`     // impact_justification
	StartegyAgreement      string    `json:"strategy_agreement"`       // strategy_agreement
	StrategyRecomendation  string    `json:"strategy_recomendation"`   // strategy_recomendation
	AssessmentLikehood     string    `json:"assessment_likehood"`      // assessment_likehood
	AssessmentImpact       string    `json:"assessment_impact"`        // assessment_impact
	AssessmentRiskLevel    string    `json:"assessment_risk_level"`    // assessment_risk_level
	Action                 string    `json:"action"`                   // action
	ActionBy               int       `json:"action_by"`                // action_by
	CreatedAt              time.Time `json:"created_at"`               // created_at
	UpdatedAt              time.Time `json:"updated_at"`               // updated_at
}

func RiskModelToRiskResponse(risk models.Risk) RiskResponse {
	return RiskResponse{
		Id:                     risk.Id,
		DocumentId:             "a", // temp handle with static value to remove error
		RiskName:               risk.RiskName,
		FraudSchema:            risk.FraudSchema,
		FraudMotive:            risk.FraudMotive,
		FraudTechnique:         risk.FraudTechnique,
		RiskSource:             risk.RiskSource,
		RootCause:              risk.RootCause,
		BisproControlProcedure: risk.BisproControlProcedure,
		QualitativeImpact:      risk.QualitativeImpact,
		LikehoodJustification:  risk.LikehoodJustification,
		ImpactJustification:    risk.ImpactJustification,
		StartegyAgreement:      risk.StartegyAgreement,
		StrategyRecomendation:  risk.StrategyRecomendation,
		AssessmentLikehood:     risk.AssessmentLikehood,
		AssessmentImpact:       risk.AssessmentImpact,
		AssessmentRiskLevel:    risk.AssessmentRiskLevel,
		CreatedAt:              risk.CreatedAt,
		UpdatedAt:              risk.UpdatedAt,
	}
}

func BulkRiskModelToRiskResponse(risks []models.Risk) []RiskResponse {
	var bulk []RiskResponse
	for _, risk := range risks {
		bulk = append(bulk, RiskModelToRiskResponse(risk))
	}
	return bulk
}

type RiskResponseWithUserDetail struct {
	Id                     int       `json:"id"`                       // id
	DocumentId             string    `json:"document_id"`              // document_id
	UserId                 int       `json:"user_id"`                  // user_id
	RiskName               string    `json:"risk_name"`                // risk_name
	FraudSchema            string    `json:"fraud_schema"`             // fraud_schema
	FraudMotive            string    `json:"fraud_motive"`             // fraud_motive
	FraudTechnique         string    `json:"fraud_technique"`          // fraud_technique
	RiskSource             string    `json:"risk_source"`              // risk_source
	RootCause              string    `json:"root_cause"`               // root_cause
	BisproControlProcedure string    `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string    `json:"qualitative_impact"`       // qualitative_impact
	LikehoodJustification  string    `json:"likehood_justification"`   // likehood_justification
	ImpactJustification    string    `json:"impact_justification"`     // impact_justification
	StartegyAgreement      string    `json:"strategy_agreement"`       // strategy_agreement
	StrategyRecomendation  string    `json:"strategy_recomendation"`   // strategy_recomendation
	AssessmentLikehood     string    `json:"assessment_likehood"`      // assessment_likehood
	AssessmentImpact       string    `json:"assessment_impact"`        // assessment_impact
	AssessmentRiskLevel    string    `json:"assessment_risk_level"`    // assessment_risk_level
	Action                 string    `json:"action"`                   // action
	ActionBy               int       `json:"action_by"`                // action_by
	CreatedAt              time.Time `json:"created_at"`               // created_at
	UpdatedAt              time.Time `json:"updated_at"`               // updated_at

	UserDetail userResponse `json:"user_detail"` // user detail
}

type userResponse struct {
	Id   int    `json:"id"`
	Nik  string `json:"nik"`
	Name string `json:"name"`
}

func RiskModelToRiskResponseWithUserDetail(risk models.Risk) RiskResponseWithUserDetail {
	return RiskResponseWithUserDetail{
		Id:                     risk.Id,
		DocumentId:             "a", // temp handle with static value to remove error
		RiskName:               risk.RiskName,
		FraudSchema:            risk.FraudSchema,
		FraudMotive:            risk.FraudMotive,
		FraudTechnique:         risk.FraudTechnique,
		RiskSource:             risk.RiskSource,
		RootCause:              risk.RootCause,
		BisproControlProcedure: risk.BisproControlProcedure,
		QualitativeImpact:      risk.QualitativeImpact,
		LikehoodJustification:  risk.LikehoodJustification,
		ImpactJustification:    risk.ImpactJustification,
		StartegyAgreement:      risk.StartegyAgreement,
		StrategyRecomendation:  risk.StrategyRecomendation,
		AssessmentLikehood:     risk.AssessmentLikehood,
		AssessmentImpact:       risk.AssessmentImpact,
		AssessmentRiskLevel:    risk.AssessmentRiskLevel,
		CreatedAt:              risk.CreatedAt,
		UpdatedAt:              risk.UpdatedAt,
		UserDetail:             userModelToUserResponse(risk.UserDetail),
	}
}

func BulkRiskModelToRiskResponseWithUserDetail(risks []models.Risk) []RiskResponseWithUserDetail {
	var bulk []RiskResponseWithUserDetail
	for _, risk := range risks {
		bulk = append(bulk, RiskModelToRiskResponseWithUserDetail(risk))
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
