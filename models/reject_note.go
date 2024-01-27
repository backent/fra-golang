package models

import "time"

type RejectNote struct {
	Id                     int       // id
	DocumentId             int       // document_id
	RiskId                 int       // risk_id
	Fraud                  string    // fraud
	RiskSource             string    // risk_source
	RootCause              string    // root_cause
	BisproControlProcedure string    // bispro_control_procedure
	QualitativeImpact      string    // qualitative_impact
	Assessment             string    // assessment
	Justification          string    // justification
	Strategy               string    // strategy
	CreatedAt              time.Time // created_at
	UpdatedAt              time.Time // updated_at
}

var RejectNoteTable = "reject_notes"
