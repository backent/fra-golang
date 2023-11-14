package document

import "github.com/backent/fra-golang/models"

type DocumentResponse struct {
	Id                     int    `json:"id"`                       // id
	DocumentId             string `json:"document_id"`              // document_id
	UserId                 int    `json:"user_id"`                  // user_id
	RiskName               string `json:"risk_name"`                // risk_name
	FraudSchema            string `json:"fraud_schema"`             // fraud_schema
	FraudMotive            string `json:"fraud_motive"`             // fraud_motive
	FraudTechnique         string `json:"fraud_technique"`          // fraud_technique
	RiskSource             string `json:"risk_source"`              // risk_source
	RootCause              string `json:"root_cause"`               // root_cause
	BisproControlProcedure string `json:"bispro_control_procedure"` // bispro_control_procedure
	QualitativeImpact      string `json:"qualitative_impact"`       // qualitative_impact
	LikehoodJustification  string `json:"likehood_justification"`   // likehood_justification
	ImpactJustification    string `json:"impact_justification"`     // impact_justification
	StartegyAgreement      string `json:"strategy_agreement"`       // strategy_agreement
	StrategyRecomendation  string `json:"strategy_recomendation"`   // strategy_recomendation
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
	}
}

func BulkDocumentModelToDocumentResponse(documents []models.Document) []DocumentResponse {
	var bulk []DocumentResponse
	for _, document := range documents {
		bulk = append(bulk, DocumentModelToDocumentResponse(document))
	}
	return bulk
}
