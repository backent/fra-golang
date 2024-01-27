package rejectnote

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryRejectNoteInterface interface {
	Create(ctx context.Context, tx *sql.Tx, rejectNote models.RejectNote) (models.RejectNote, error)
}
