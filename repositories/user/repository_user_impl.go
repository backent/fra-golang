package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/backent/fra-golang/helpers"
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
	query := fmt.Sprintf("UPDATE  %s SET name = ?, password = ?, unit = ?, role = ? WHERE id = ?", models.UserTable)
	_, err := tx.ExecContext(ctx, query, user.Name, user.Password, user.Unit, user.Role, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (implementation *RepositoryUserImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE id = ?", models.UserTable)
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
func (implementation *RepositoryUserImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, nik, name, email, role, unit, password FROM %s WHERE deleted_at IS NULL AND id = ?", models.UserTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {

		var nullPassword sql.NullString
		var nullEmail sql.NullString
		var nullUnit sql.NullString

		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &nullEmail, &user.Role, &nullUnit, &nullPassword)
		if err != nil {
			return user, err
		}
		user.Password = nullPassword.String
		user.Email = nullEmail.String
		user.Unit = nullUnit.String
	} else {
		return user, errors.New("not found user")
	}

	return user, nil
}
func (implementation *RepositoryUserImpl) FindAll(ctx context.Context, tx *sql.Tx, take int, skip int, orderBy string, orderDirection string, applyStatus string, search string) ([]models.User, int, error) {
	var users []models.User
	var totalDocument int

	var conditionalQueryStatus string
	var conditionalQueryStatusValue []interface{}
	if applyStatus == "" {
		conditionalQueryStatus = "AND 1 = ?"
		conditionalQueryStatusValue = append(conditionalQueryStatusValue, "1")
	} else {
		for _, val := range strings.Split(applyStatus, ",") {
			conditionalQueryStatusValue = append(conditionalQueryStatusValue, val)
		}
		helpers.Placeholders(len(conditionalQueryStatusValue))
		conditionalQueryStatus = fmt.Sprintf("AND apply_status IN (%s)", helpers.Placeholders(len(conditionalQueryStatusValue)))
	}

	var conditionalQuerySearch string
	var conditionalQuerySearchValue string
	if search == "" {
		conditionalQuerySearch = "AND 1 = ?"
		conditionalQuerySearchValue = "1"
	} else {
		conditionalQuerySearch = "AND name LIKE ?"
		conditionalQuerySearchValue = "%" + search + "%"
	}

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s WHERE 1 = 1 AND deleted_at IS NULL %s %s
		)
		SELECT
		a.id,
		a.nik,
		a.name,
		a.email,
		a.role,
		a.unit,
		a.apply_status,
		a.password,
		b.count
		FROM (SELECT * FROM main_table ORDER BY %s %s LIMIT ?, ?) a
		LEFT JOIN (SELECT COUNT(*) as count FROM main_table) b ON true
	`, models.UserTable, conditionalQueryStatus, conditionalQuerySearch, orderBy, orderDirection)

	var args []interface{}
	args = append(args, conditionalQueryStatusValue...)
	args = append(args, conditionalQuerySearchValue)
	args = append(args, skip, take)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, totalDocument, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.NullAbleUser
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Email, &user.Role, &user.Unit, &user.ApplyStatus, &user.Password, &totalDocument)
		if err != nil {
			return nil, totalDocument, err
		}
		users = append(users, models.NullAbleUserToUser(user))
	}

	return users, totalDocument, nil
}

func (implementation *RepositoryUserImpl) FindByNik(ctx context.Context, tx *sql.Tx, nik string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, nik, name, role, password, apply_status FROM %s WHERE nik = ? AND deleted_at IS NULL", models.UserTable)
	rows, err := tx.QueryContext(ctx, query, nik)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		var nullPassword sql.NullString
		var nullApplyStatus sql.NullString
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Role, &nullPassword, &nullApplyStatus)
		if err != nil {
			return user, err
		}
		user.Password = nullPassword.String
		user.ApplyStatus = nullApplyStatus.String
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
