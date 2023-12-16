package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webUser "github.com/backent/fra-golang/web/user"
	"github.com/go-playground/validator/v10"
)

type UserMiddleware struct {
	Validate *validator.Validate
	repositoriesUser.RepositoryUserInterface
}

func NewUserMiddleware(validator *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface) *UserMiddleware {
	return &UserMiddleware{
		Validate:                validator,
		RepositoryUserInterface: repositoriesUser,
	}
}

func (implementation *UserMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestCreate) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Nik)
	if err == nil {
		panic(exceptions.NewBadRequestError("user already exists"))
	}

}

func (implementation *UserMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestUpdate) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.CurrentPassword = user.Password
	request.Nik = user.Nik
}

func (implementation *UserMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestDelete) {

	_, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *UserMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestFindById) {

	_, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *UserMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestFindAll) {
}
