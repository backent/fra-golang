package users_history_login

import (
	"context"
	"database/sql"
	"fmt"

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
) ([]models.UserHistoryLogin, error) {

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
	GROUP BY 
			a.user_id, MONTH(a.created_at), YEAR(a.created_at), DATE_FORMAT(a.created_at, '%%Y-%%m')
	ORDER BY 
			login_count_per_month DESC LIMIT ?, ?;
	`, models.UserHistoryLoginTable, models.UserTable)

	rows, err := tx.QueryContext(ctx, query, year, month, skip, take)
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
