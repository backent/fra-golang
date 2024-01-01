package middlewares

import (
	"context"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/repositories/auth"
)

func ValidateToken(ctx context.Context, repositoriesAuth auth.RepositoryAuthInterface) int {
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
