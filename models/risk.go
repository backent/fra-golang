package models

import (
	"database/sql"
	"time"
)

type Risk struct {
	Id                     int       // id
	DocumentId             int       // document_id
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
	AssessmentLikehood     string    // assessment_likehood
	AssessmentImpact       string    // assessment_impact
	AssessmentRiskLevel    string    // assessment_risk_level
	CreatedAt              time.Time // created_at
	UpdatedAt              time.Time // updated_at

	UserDetail User
}

type NullAbleRisk struct {
	Id                     sql.NullInt32
	DocumentId             sql.NullInt32
	RiskName               sql.NullString
	FraudSchema            sql.NullString
	FraudMotive            sql.NullString
	FraudTechnique         sql.NullString
	RiskSource             sql.NullString
	RootCause              sql.NullString
	BisproControlProcedure sql.NullString
	QualitativeImpact      sql.NullString
	LikehoodJustification  sql.NullString
	ImpactJustification    sql.NullString
	StartegyAgreement      sql.NullString
	StrategyRecomendation  sql.NullString
	AssessmentLikehood     sql.NullString
	AssessmentImpact       sql.NullString
	AssessmentRiskLevel    sql.NullString
	CreatedAt              sql.NullTime
	UpdatedAt              sql.NullTime
}

var RiskTable string = "risks"

func NullAbleRiskToRisk(nullAbleRisk NullAbleRisk) Risk {
	return Risk{
		Id:                     int(nullAbleRisk.Id.Int32),
		DocumentId:             int(nullAbleRisk.DocumentId.Int32),
		RiskName:               nullAbleRisk.RiskName.String,
		FraudSchema:            nullAbleRisk.FraudSchema.String,
		FraudMotive:            nullAbleRisk.FraudMotive.String,
		FraudTechnique:         nullAbleRisk.FraudTechnique.String,
		RiskSource:             nullAbleRisk.RiskSource.String,
		RootCause:              nullAbleRisk.RootCause.String,
		BisproControlProcedure: nullAbleRisk.BisproControlProcedure.String,
		QualitativeImpact:      nullAbleRisk.QualitativeImpact.String,
		LikehoodJustification:  nullAbleRisk.LikehoodJustification.String,
		ImpactJustification:    nullAbleRisk.ImpactJustification.String,
		StartegyAgreement:      nullAbleRisk.StartegyAgreement.String,
		StrategyRecomendation:  nullAbleRisk.StrategyRecomendation.String,
		AssessmentLikehood:     nullAbleRisk.AssessmentLikehood.String,
		AssessmentImpact:       nullAbleRisk.AssessmentImpact.String,
		AssessmentRiskLevel:    nullAbleRisk.AssessmentRiskLevel.String,
		CreatedAt:              nullAbleRisk.CreatedAt.Time,
		UpdatedAt:              nullAbleRisk.UpdatedAt.Time,
	}
}
