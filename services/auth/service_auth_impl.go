package auth

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	webAuth "github.com/backent/fra-golang/web/auth"
)

type ServiceAuthImpl struct {
	*sql.DB
	repositoriesAuth.RepositoryAuthInterface
	*middlewares.AuthMiddleware
}

func NewServiceAuthImpl(db *sql.DB, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, authMiddleware *middlewares.AuthMiddleware) ServiceAuthInterface {
	return &ServiceAuthImpl{
		DB:                      db,
		RepositoryAuthInterface: repositoriesAuth,
		AuthMiddleware:          authMiddleware,
	}
}

func (implementation *ServiceAuthImpl) Login(ctx context.Context, request webAuth.LoginRequest) webAuth.LoginResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.AuthMiddleware.Login(ctx, tx, &request)

	stringUserId := strconv.Itoa(request.UserId)
	token, err := implementation.RepositoryAuthInterface.Issue(stringUserId)
	helpers.PanicIfError(err)

	return webAuth.LoginResponse{
		Token: token,
	}
}
func (implementation *ServiceAuthImpl) Register(ctx context.Context, request webAuth.RegisterRequest) webAuth.RegisterResponse {
	panic("implement me")
}
