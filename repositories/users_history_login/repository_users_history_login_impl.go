package users_history_login

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/models"
)

type RepositoryUserHistoryLoginImpl struct {
}

func NewRepositoryUserHistoryLoginImpl() RepositoryUserHistoryLoginInterface {
	return &RepositoryUserHistoryLoginImpl{}
}

func (implementation *RepositoryUserHistoryLoginImpl) Create(ctx context.Context, tx *sql.Tx, userHistoryLogin models.UserHistoryLogin) error {
	query := fmt.Sprintf(`INSERT INTO %s (user_id) VALUES (?)`, models.UserHistoryLoginTable)

	_, err := tx.ExecContext(ctx, query, userHistoryLogin.UserId)
	return err
}

func (implementation *RepositoryUserHistoryLoginImpl) FindAll(
	ctx context.Context,
	tx *sql.Tx,
	take int,
	skip int,
	year string,
	month string,
	userException string,
) ([]models.UserHistoryLogin, error) {

	var userExceptionQuery string
	var userExceptionQueryValue []interface{}

	if userException != "" {
		for _, val := range strings.Split(userException, ",") {
			userExceptionQueryValue = append(userExceptionQueryValue, val)
		}
		userExceptionQuery = fmt.Sprintf("AND b.nik NOT IN (%s)", helpers.Placeholders(len(userExceptionQueryValue)))
	} else {
		userExceptionQuery = "AND 1 = ?"
		userExceptionQueryValue = append(userExceptionQueryValue, 1)
	}

	query := fmt.Sprintf(`
	SELECT 
			a.user_id,
			b.name,
			b.nik,
			DATE_FORMAT(a.created_at, '%%Y-%%m') AS month,
			COUNT(DISTINCT DATE(a.created_at)) AS login_count_per_month
	FROM 
			%s a
	LEFT JOIN 
			%s b ON a.user_id = b.id
	WHERE 
			YEAR(a.created_at) = ?  AND MONTH(a.created_at) = ?
			AND b.deleted_at IS NULL
			%s
	GROUP BY 
			a.user_id, MONTH(a.created_at), YEAR(a.created_at), DATE_FORMAT(a.created_at, '%%Y-%%m')
	ORDER BY 
			login_count_per_month DESC LIMIT ?, ?;
	`, models.UserHistoryLoginTable, models.UserTable, userExceptionQuery)

	var args []interface{}
	args = append(args, year, month)
	args = append(args, userExceptionQueryValue...)
	args = append(args, skip, take)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersHistoryLogin []models.UserHistoryLogin

	for rows.Next() {
		var userHistoryLogin models.UserHistoryLogin
		var dummyMonth string
		var dummyLoginPerMonth int
		err := rows.Scan(
			&userHistoryLogin.UserId,
			&userHistoryLogin.UserDetail.Name,
			&userHistoryLogin.UserDetail.Nik,
			&dummyMonth,
			&dummyLoginPerMonth,
		)
		if err != nil {
			return nil, err
		}

		usersHistoryLogin = append(usersHistoryLogin, userHistoryLogin)
	}
	return usersHistoryLogin, nil
}
