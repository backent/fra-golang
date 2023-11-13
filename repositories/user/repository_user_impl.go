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
	query := fmt.Sprintf("UPDATE  %s SET nik = ?, name = ?, password = ? WHERE id = ?", models.UserTable)
	_, err := tx.ExecContext(ctx, query, user.Nik, user.Name, user.Password, user.Id)
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

	query := fmt.Sprintf("SELECT id, nik, name, password FROM %s WHERE id = ?", models.UserTable)
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Password)
		if err != nil {
			return user, err
		}
	} else {
		return user, errors.New("not found user")
	}

	return user, nil
}
func (implementation *RepositoryUserImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]models.User, error) {
	var users []models.User

	query := fmt.Sprintf("SELECT id, nik, name, password FROM %s ORDER BY id DESC", models.UserTable)
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Nik, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
