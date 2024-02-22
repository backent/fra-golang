package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webAuth "github.com/backent/fra-golang/web/auth"
	"github.com/go-playground/validator/v10"
)

type AuthMiddleware struct {
	*validator.Validate
	repositoriesUser.RepositoryUserInterface
	repositoriesAuth.RepositoryAuthInterface
}

func NewAuthMiddleware(validate *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface) *AuthMiddleware {
	return &AuthMiddleware{
		Validate:                validate,
		RepositoryUserInterface: repositoriesUser,
		RepositoryAuthInterface: repositoriesAuth,
	}
}

func (implementation *AuthMiddleware) Login(ctx context.Context, tx *sql.Tx, request *webAuth.LoginRequest) {
	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	user, err := implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Username)
	if err != nil || user.ApplyStatus != "approve" {
		panic(exceptions.NewBadRequestError("username or password is incorrect"))
	}

	passwordValid := helpers.CheckPassword(request.Password, user.Password)
	token, err := helpers.LoginLdap("402746", request.Password)

	if !passwordValid && (err != nil || token == "") {
		panic(exceptions.NewBadRequestError("wrong username or password."))
	} else {
		request.UserId = user.Id
	}

}
