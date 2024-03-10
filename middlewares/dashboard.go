package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webDashboard "github.com/backent/fra-golang/web/dashboard"
	"github.com/go-playground/validator/v10"
)

type DashboardMiddleware struct {
	Validate *validator.Validate
	repositoriesAuth.RepositoryAuthInterface
	repositoriesUser.RepositoryUserInterface
}

func NewDashboardMiddleware(validator *validator.Validate, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, repositoriesUser repositoriesUser.RepositoryUserInterface) *DashboardMiddleware {
	return &DashboardMiddleware{
		Validate:                validator,
		RepositoryAuthInterface: repositoriesAuth,
		RepositoryUserInterface: repositoriesUser,
	}
}

func (implementation *DashboardMiddleware) Summary(ctx context.Context, tx *sql.Tx, request *webDashboard.DashboardRequestSummary) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, userId)
	helpers.PanicIfError(err)

	request.User = user

}
