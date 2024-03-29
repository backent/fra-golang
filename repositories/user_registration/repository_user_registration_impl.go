package user_registration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/models"
)

type RepositoryUserRegistrationImpl struct {
}

func NewRepositoryUserRegistrationImpl() RepositoryUserRegistrationInterface {
	return &RepositoryUserRegistrationImpl{}
}

func (implementation *RepositoryUserRegistrationImpl) Create(ctx context.Context, tx *sql.Tx, user_registration models.UserRegistration) (models.UserRegistration, error) {
	query := fmt.Sprintf("INSERT INTO %s (nik, name, email, apply_status, password) VALUES (?, ?, ?, ?, ?)", models.UserRegistrationTable)
	result, err := tx.ExecContext(ctx, query, user_registration.Nik, user_registration.Name, user_registration.Email, user_registration.Status, user_registration.Password)
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

func (implementation *RepositoryUserRegistrationImpl) Update(ctx context.Context, tx *sql.Tx, user_registration models.UserRegistration) (models.UserRegistration, error) {
	query := fmt.Sprintf(`UPDATE %s SET 
		nik = ?,
		name = ?,
		email = ?,
		apply_status = ?,
		apply_reject_by = ?,
		apply_approved_by = ?,
		unit = ? WHERE id = ?`, models.UserRegistrationTable)
	_, err := tx.ExecContext(ctx, query, user_registration.Nik, user_registration.Name, user_registration.Email, user_registration.Status, user_registration.RejectBy, user_registration.ApproveBy, user_registration.Unit, user_registration.Id)
	if err != nil {
		return user_registration, err
	}

	return user_registration, nil

}

func (implementation *RepositoryUserRegistrationImpl) FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.UserRegistration, error) {
	var userRegistration models.UserRegistration
	query := fmt.Sprintf("SELECT id, nik, name, email, apply_status, apply_reject_by, apply_approved_by, created_at, updated_at FROM %s WHERE nik = ? LIMIT 1", models.UserRegistrationTable)

	rows, err := tx.QueryContext(ctx, query, nik)
	if err != nil {
		return userRegistration, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&userRegistration.Id, &userRegistration.Nik, &userRegistration.Name, &userRegistration.Email, &userRegistration.Status, &userRegistration.RejectBy, &userRegistration.ApproveBy, &userRegistration.CreatedAt, &userRegistration.UpdatedAt)
		if err != nil {
			return userRegistration, err
		}
	} else {
		return userRegistration, errors.New("not found")
	}

	return userRegistration, nil
}

func (implementation *RepositoryUserRegistrationImpl) FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string, queryStatus string) ([]models.UserRegistration, int, error) {
	var conditionalQueryStatus string
	var conditionalQueryStatusValue []interface{}
	if queryStatus == "" {
		queryStatus = "1"
		conditionalQueryStatus = "AND 1 = ?"
		conditionalQueryStatusValue = append(conditionalQueryStatusValue, "1")
	} else {
		for _, val := range strings.Split(queryStatus, ",") {
			conditionalQueryStatusValue = append(conditionalQueryStatusValue, val)
		}
		helpers.Placeholders(len(conditionalQueryStatusValue))
		conditionalQueryStatus = fmt.Sprintf("AND apply_status IN (%s)", helpers.Placeholders(len(conditionalQueryStatusValue)))
	}

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s WHERE 1 = 1 %s
		)
		SELECT
		a.id,
		a.nik,
		a.name,
		a.email,
		a.apply_status,
		a.created_at,
		a.updated_at,
		b.count
		FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ?) a LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true
	`, models.UserRegistrationTable, conditionalQueryStatus, orderBy, orderDirection)

	var args []interface{}
	args = append(args, conditionalQueryStatusValue...)
	args = append(args, skip, take)

	var userRegistrations []models.UserRegistration
	var total int

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, total, err
	}
	defer rows.Close()

	for rows.Next() {
		var userRegistration models.NullAbleUserRegistration
		err = rows.Scan(
			&userRegistration.Id,
			&userRegistration.Nik,
			&userRegistration.Name,
			&userRegistration.Email,
			&userRegistration.Status,
			&userRegistration.CreatedAt,
			&userRegistration.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, total, err
		}

		userRegistrations = append(userRegistrations, models.NullAbleUserRegistrationToUserRegistration(userRegistration))
	}

	return userRegistrations, total, nil
}
