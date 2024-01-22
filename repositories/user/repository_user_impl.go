package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/backent/fra-golang/models"
)

type RepositoryUserImpl struct {
}

func NewRepositoryUserImpl() RepositoryUserInterface {
	return &RepositoryUserImpl{}
}

func (implementation *RepositoryUserImpl) Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (nik, name, password) VALUES (?, ?, ?)", models.UserTable)
	result, err := tx.ExecContext(ctx, query, user.Nik, user.Name, user.Password)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = int(id)

	return user, nil

}
func (implementation *RepositoryUserImpl) Update(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := fmt.Sprintf("UPDATE  %s SET name = ?, password = ? WHERE id = ?", models.UserTable)
	_, err := tx.ExecContext(ctx, query, user.Name, user.Password, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (implementation *RepositoryUserImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("DELETE FROM  %s  WHERE id = ?", models.UserTable)
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (implementation *RepositoryUserImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, nik, name, role, password FROM %s WHERE id = ?", models.UserTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Role, &user.Password)
		if err != nil {
			return user, err
		}
	} else {
		return user, errors.New("not found user")
	}

	return user, nil
}
func (implementation *RepositoryUserImpl) FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, nik, name, role, password FROM %s ORDER BY %s %s LIMIT ?, ?", models.UserTable, orderBy, orderDirection)
	rows, err := tx.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Role, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (implementation *RepositoryUserImpl) FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, nik, name, role, password FROM %s WHERE nik = ?", models.UserTable)
	rows, err := tx.QueryContext(ctx, query, nik)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Role, &user.Password)
		if err != nil {
			return user, err
		}
	} else {
		return user, errors.New("not found user")
	}

	return user, nil
}

func (implementation *RepositoryUserImpl) FindAllWithRisksDetail(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string) ([]models.User, error) {
	query := fmt.Sprintf(`SELECT
		a.id,
		a.nik,
		a.name,
		a.role,
		a.password,
		b.id,
		b.document_id,
		b.user_id,
		b.risk_name
		FROM
		(SELECT * FROM %s ORDER BY %s %s LIMIT ?, ?) a 
		LEFT JOIN %s b
		ON a.id = b.user_id`, models.UserTable, orderBy, orderDirection, models.RiskTable)
	rows, err := tx.QueryContext(ctx, query, skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	usersMap := make(map[int]*models.User)

	for rows.Next() {
		var user models.User
		var risk struct {
			Id         sql.NullInt32
			DocumentId sql.NullInt32
			UserId     sql.NullInt32
			RiskName   sql.NullString
		}
		err = rows.Scan(
			&user.Id,
			&user.Nik,
			&user.Name,
			&user.Role,
			&user.Password,
			&risk.Id,
			&risk.DocumentId,
			&risk.UserId,
			&risk.RiskName,
		)
		if err != nil {
			return nil, err
		}
		item, found := usersMap[user.Id]
		if !found {
			item = &user
			usersMap[user.Id] = item
			users = append(users, item)
		}
		if risk.Id.Valid {
			validRisk := models.Risk{
				Id:         int(risk.Id.Int32),
				DocumentId: int(risk.DocumentId.Int32),
				RiskName:   risk.RiskName.String,
			}
			if item.RisksDetail == nil {
				item.RisksDetail = []models.Risk{validRisk}
			} else {
				item.RisksDetail = append(item.RisksDetail, validRisk)
			}
		}
	}

	var returnedUser []models.User
	for _, item := range users {
		returnedUser = append(returnedUser, *item)
	}

	return returnedUser, nil
}
