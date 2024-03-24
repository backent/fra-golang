package middlewares

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webUser "github.com/backent/fra-golang/web/user"
	"github.com/go-playground/validator/v10"
)

type UserMiddleware struct {
	Validate *validator.Validate
	repositoriesUser.RepositoryUserInterface
	repositoriesAuth.RepositoryAuthInterface
}

func NewUserMiddleware(validator *validator.Validate, repositoriesUser repositoriesUser.RepositoryUserInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface) *UserMiddleware {
	return &UserMiddleware{
		Validate:                validator,
		RepositoryUserInterface: repositoriesUser,
		RepositoryAuthInterface: repositoriesAuth,
	}
}

func (implementation *UserMiddleware) Create(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestCreate) {
	// Temp turn command to comment for creating user if needed
	// ValidateToken(ctx, implementation.RepositoryAuthInterface)

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Nik)
	if err == nil {
		panic(exceptions.NewBadRequestError("user already exists"))
	}

}

func (implementation *UserMiddleware) Update(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestUpdate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.CurrentPassword = user.Password
	request.Nik = user.Nik
	request.User = user
}

func (implementation *UserMiddleware) Delete(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestDelete) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *UserMiddleware) FindById(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestFindById) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	_, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
}

func (implementation *UserMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestFindAll) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

	if request.QueryStatus == "" {
		request.QueryStatus = "approve,reject"
	}
}

func (implementation *UserMiddleware) CurrentUser(ctx context.Context, tx *sql.Tx, request *webUser.UserRequestCurrentUser) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)
	request.UserId = userId
}
