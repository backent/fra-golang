package middlewares

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	fmt.Println("START validating login for user : ", request.Username)
	defer fmt.Println("END validating login for user : ", request.Username)

	user, err := implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Username)
	if err != nil || user.ApplyStatus != "approve" {
		panic(exceptions.NewBadRequestError("username or password is incorrect"))
	}

	chanPasswordValid := make(chan bool)
	go func() {
		chanPasswordValid <- helpers.CheckPassword(request.Password, user.Password)
		close(chanPasswordValid)
	}()

	chanLdapValid := make(chan bool)

	go func() {
		token, err := helpers.LoginLdap(request.Username, request.Password)
		if err != nil {
			log.Println("login ldap error : ", err)
		}
		chanLdapValid <- token != ""
		close(chanLdapValid)
	}()

	passwordValid := <-chanPasswordValid
	ldapValid := <-chanLdapValid

	if !passwordValid && !ldapValid {
		panic(exceptions.NewBadRequestError("wrong username or password."))
	} else {
		request.UserId = user.Id
	}

}
