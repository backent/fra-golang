package document

import (
	"time"

	"github.com/backent/fra-golang/models"
)

type DocumentResponse struct {
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

func DocumentModelToDocumentResponse(document models.Document) DocumentResponse {
	return DocumentResponse{
		Id:                     document.Id,
		DocumentId:             document.DocumentId,
		UserId:                 document.UserId,
		RiskName:               document.RiskName,
		FraudSchema:            document.FraudSchema,
		FraudMotive:            document.FraudMotive,
		FraudTechnique:         document.FraudTechnique,
		RiskSource:             document.RiskSource,
		RootCause:              document.RootCause,
		BisproControlProcedure: document.BisproControlProcedure,
		QualitativeImpact:      document.QualitativeImpact,
		LikehoodJustification:  document.LikehoodJustification,
		ImpactJustification:    document.ImpactJustification,
		StartegyAgreement:      document.StartegyAgreement,
		StrategyRecomendation:  document.StrategyRecomendation,
		AssessmentLikehood:     document.AssessmentLikehood,
		AssessmentImpact:       document.AssessmentImpact,
		AssessmentRiskLevel:    document.AssessmentRiskLevel,
		Action:                 document.Action,
		ActionBy:               document.ActionBy,
		CreatedAt:              document.CreatedAt,
		UpdatedAt:              document.UpdatedAt,
	}
}

func BulkDocumentModelToDocumentResponse(documents []models.Document) []DocumentResponse {
	var bulk []DocumentResponse
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponse(document))
	}
	return bulk
}

type DocumentResponseWithUserDetail struct {
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

func DocumentModelToDocumentResponseWithUserDetail(document models.Document) DocumentResponseWithUserDetail {
	return DocumentResponseWithUserDetail{
		Id:                     document.Id,
		DocumentId:             document.DocumentId,
		UserId:                 document.UserId,
		RiskName:               document.RiskName,
		FraudSchema:            document.FraudSchema,
		FraudMotive:            document.FraudMotive,
		FraudTechnique:         document.FraudTechnique,
		RiskSource:             document.RiskSource,
		RootCause:              document.RootCause,
		BisproControlProcedure: document.BisproControlProcedure,
		QualitativeImpact:      document.QualitativeImpact,
		LikehoodJustification:  document.LikehoodJustification,
		ImpactJustification:    document.ImpactJustification,
		StartegyAgreement:      document.StartegyAgreement,
		StrategyRecomendation:  document.StrategyRecomendation,
		AssessmentLikehood:     document.AssessmentLikehood,
		AssessmentImpact:       document.AssessmentImpact,
		AssessmentRiskLevel:    document.AssessmentRiskLevel,
		Action:                 document.Action,
		ActionBy:               document.ActionBy,
		CreatedAt:              document.CreatedAt,
		UpdatedAt:              document.UpdatedAt,
		UserDetail:             userModelToUserResponse(document.UserDetail),
	}
}

func BulkDocumentModelToDocumentResponseWithUserDetail(documents []models.Document) []DocumentResponseWithUserDetail {
	var bulk []DocumentResponseWithUserDetail
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponseWithUserDetail(document))
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
