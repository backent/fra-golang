package models

import (
	"database/sql"
	"time"
)

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

type NullAbleRejectNote struct {
	Id                     sql.NullInt32  // id
	DocumentId             sql.NullInt32  // document_id
	RiskId                 sql.NullInt32  // risk_id
	Fraud                  sql.NullString // fraud
	RiskSource             sql.NullString // risk_source
	RootCause              sql.NullString // root_cause
	BisproControlProcedure sql.NullString // bispro_control_procedure
	QualitativeImpact      sql.NullString // qualitative_impact
	Assessment             sql.NullString // assessment
	Justification          sql.NullString // justification
	Strategy               sql.NullString // strategy
	CreatedAt              sql.NullTime   // created_at
	UpdatedAt              sql.NullTime   // updated_at
}

var RejectNoteTable = "reject_notes"
