package user_registration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryUserRegistrationImpl struct {
}

func NewRepositoryUserRegistrationImpl() RepositoryUserRegistrationInterface {
	return &RepositoryUserRegistrationImpl{}
}

func (implementation *RepositoryUserRegistrationImpl) Create(ctx context.Context, tx *sql.Tx, user_registration models.UserRegistration) (models.UserRegistration, error) {
	query := fmt.Sprintf("INSERT INTO %s (nik, status) VALUES (?, ?)", models.UserRegistrationTable)
	result, err := tx.ExecContext(ctx, query, user_registration.Nik, user_registration.Status)
	if err != nil {
		return user_registration, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user_registration, err
	}

	user_registration.Id = int(id)

	return user_registration, nil

}

func (implementation *RepositoryUserRegistrationImpl) FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.UserRegistration, error) {
	var userRegistration models.UserRegistration
	query := fmt.Sprintf("SELECT id, nik, status, reject_by, approve_by, created_at, updated_at FROM %s WHERE nik = ? LIMIT 1", models.UserRegistrationTable)

	rows, err := tx.QueryContext(ctx, query, nik)
	if err != nil {
		return userRegistration, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&userRegistration.Id, &userRegistration.Nik, &userRegistration.Status, &userRegistration.RejectBy, &userRegistration.ApproveBy, &userRegistration.CreatedAt, &userRegistration.UpdatedAt)
		if err != nil {
			return userRegistration, err
		}
	} else {
		return userRegistration, errors.New("not found")
	}

	return userRegistration, nil
}
