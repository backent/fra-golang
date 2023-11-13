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
	helpers.PanifIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Nik)
	if err == nil {
		panic(exceptions.NewBadRequestError("user already exists"))
	}

}
