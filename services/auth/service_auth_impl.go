package auth

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUserHistoryLogin "github.com/backent/fra-golang/repositories/users_history_login"
	webAuth "github.com/backent/fra-golang/web/auth"
)

type ServiceAuthImpl struct {
	*sql.DB
	repositoriesAuth.RepositoryAuthInterface
	repositoriesUserHistoryLogin.RepositoryUserHistoryLoginInterface
	*middlewares.AuthMiddleware
}

func NewServiceAuthImpl(db *sql.DB, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, repositoriesUserHistoryLogin repositoriesUserHistoryLogin.RepositoryUserHistoryLoginInterface, authMiddleware *middlewares.AuthMiddleware) ServiceAuthInterface {
	return &ServiceAuthImpl{
		DB:                                  db,
		RepositoryAuthInterface:             repositoriesAuth,
		RepositoryUserHistoryLoginInterface: repositoriesUserHistoryLogin,
		AuthMiddleware:                      authMiddleware,
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

	userHistoryLogin := models.UserHistoryLogin{
		UserId: request.UserId,
	}

	err = implementation.RepositoryUserHistoryLoginInterface.Create(ctx, tx, userHistoryLogin)
	helpers.PanicIfError(err)

	return webAuth.LoginResponse{
		Token: token,
	}
}
func (implementation *ServiceAuthImpl) Register(ctx context.Context, request webAuth.RegisterRequest) webAuth.RegisterResponse {
	panic("implement me")
}
