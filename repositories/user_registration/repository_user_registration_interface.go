package user_registration

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/models"
)

type RepositoryUserRegistrationInterface interface {
	Create(ctx context.Context, tx *sql.Tx, user_registration models.UserRegistration) (models.UserRegistration, error)
	FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.UserRegistration, error)
}
