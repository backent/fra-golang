package rejectnote

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryRejectNoteImpl struct {
}

func NewRepositoryRejectNote() RepositoryRejectNoteInterface {
	return &RepositoryRejectNoteImpl{}
}

func (implementation *RepositoryRejectNoteImpl) Create(ctx context.Context, tx *sql.Tx, rejectNote models.RejectNote) (models.RejectNote, error) {

	query := fmt.Sprintf(`INSERT INTO %s (
		document_id,
		risk_id,
		fraud,
		risk_source,
		root_cause,
		bispro_control_procedure,
		qualitative_impact,
		assessment,
		justification,
		strategy
		) VALUES (?,?,?,?,?,?,?,?,?,?)`, models.RejectNoteTable)

	res, err := tx.ExecContext(ctx, query,
		rejectNote.DocumentId,
		rejectNote.RiskId,
		rejectNote.Fraud,
		rejectNote.RiskSource,
		rejectNote.RootCause,
		rejectNote.BisproControlProcedure,
		rejectNote.QualitativeImpact,
		rejectNote.Assessment,
		rejectNote.Justification,
		rejectNote.Strategy,
	)

	if err != nil {
		return rejectNote, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return rejectNote, err
	}

	rejectNote.Id = int(id)

	return rejectNote, nil
}
