package models

import "time"

type Document struct {
	Id                     int       // id
	Status                 bool      // status
	DocumentId             string    // document_id
	UserId                 int       // user_id
	RiskName               string    // risk_name
	FraudSchema            string    // fraud_schema
	FraudMotive            string    // fraud_motive
	FraudTechnique         string    // fraud_technique
	RiskSource             string    // risk_source
	RootCause              string    // root_cause
	BisproControlProcedure string    // bispro_control_procedure
	QualitativeImpact      string    // qualitative_impact
	LikehoodJustification  string    // likehood_justification
	ImpactJustification    string    // impact_justification
	StartegyAgreement      string    // strategy_agreement
	StrategyRecomendation  string    // strategy_recomendation
	Action                 string    // action
	ActionBy               int       // action_by
	CreatedAt              time.Time // created_at
	UpdatedAt              time.Time // updated_at

	UserDetail User
}

var DocumentTable string = "documents"
