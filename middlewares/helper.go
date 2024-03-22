package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/config"
	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
)

func ValidateToken(ctx context.Context, repositoriesAuth repositoriesAuth.RepositoryAuthInterface) int {
	defer func() {
		validateFail := recover()
		if validateFail != nil {
			helpers.PanicIfError(exceptions.NewUnAuthorized("authorization invalid"))
		}
	}()
	var tokenString string
	token := ctx.Value(helpers.ContextKey("token"))

	tokenString, ok := token.(string)
	if !ok || tokenString == "" {
		helpers.PanicIfError(exceptions.NewUnAuthorized("authorization required"))
	}

	idInt, isValid := repositoriesAuth.Validate(tokenString)
	if !isValid {
		helpers.PanicIfError(exceptions.NewUnAuthorized("authorization invalid"))
	}
	return idInt
}

func ValidateUserPermission(ctx context.Context, tx *sql.Tx, repositoriesUser repositoriesUser.RepositoryUserInterface, userId int, action string) {
	user, _ := repositoriesUser.FindById(ctx, tx, userId)
	valid, _ := config.ValidatePermission(user, action)
	if !valid {
		helpers.PanicIfError(exceptions.NewForbidden("role not permitted"))
	}
}
