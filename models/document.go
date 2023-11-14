package models

type Document struct {
	Id                     int    // id
	DocumentId             string // document_id
	UserId                 int    // user_id
	RiskName               string // risk_name
	FraudSchema            string // fraud_schema
	FraudMotive            string // fraud_motive
	FraudTechnique         string // fraud_technique
	RiskSource             string // risk_source
	RootCause              string // root_cause
	BisproControlProcedure string // bispro_control_procedure
	QualitativeImpact      string // qualitative_impact
	LikehoodJustification  string // likehood_justification
	ImpactJustification    string // impact_justification
	StartegyAgreement      string // strategy_agreement
	StrategyRecomendation  string // strategy_recomendation
}

var DocumentTable string = "documents"
