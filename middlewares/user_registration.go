package middlewares

import (
	"context"
	"database/sql"
	"strings"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	repositoriesAuth "github.com/backent/fra-golang/repositories/auth"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	repositoriesUserRegistration "github.com/backent/fra-golang/repositories/user_registration"
	webUserRegistration "github.com/backent/fra-golang/web/user_registration"
	"github.com/go-playground/validator/v10"
)

type UserRegistrationMiddleware struct {
	Validate *validator.Validate
	repositoriesUserRegistration.RepositoryUserRegistrationInterface
	repositoriesAuth.RepositoryAuthInterface
	repositoriesUser.RepositoryUserInterface
}

func NewUserRegistrationMiddleware(validator *validator.Validate, repositoriesUserRegistration repositoriesUserRegistration.RepositoryUserRegistrationInterface, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, repositoriesUser repositoriesUser.RepositoryUserInterface) *UserRegistrationMiddleware {
	return &UserRegistrationMiddleware{
		Validate:                            validator,
		RepositoryUserRegistrationInterface: repositoriesUserRegistration,
		RepositoryAuthInterface:             repositoriesAuth,
		RepositoryUserInterface:             repositoriesUser,
	}
}

func (implementation *UserRegistrationMiddleware) Apply(ctx context.Context, tx *sql.Tx, request *webUserRegistration.UserRegistrationRequestApply) {

	err := implementation.Validate.Struct(request)
	helpers.PanicIfError(err)

	_, err = implementation.RepositoryUserInterface.FindByNik(ctx, tx, request.Nik)
	if err == nil {
		panic(exceptions.NewBadRequestError("nik already requested or exists"))
	}

	token, err := helpers.LoginLdap("402746", "T3lk0mDCS24")
	helpers.PanicIfError(err)

	userLdap, err := helpers.GetUserLdap(request.Nik, token)
	helpers.PanicIfError(err)

	request.Name = userLdap.Name
	request.Email = userLdap.Email

}

func (implementation *UserRegistrationMiddleware) FindAll(ctx context.Context, tx *sql.Tx, request *webUserRegistration.UserRegistrationRequestFindAll) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)
}

func (implementation *UserRegistrationMiddleware) Approve(ctx context.Context, tx *sql.Tx, request *webUserRegistration.UserRegistrationRequestApprove) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, "user-approval")

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.User.Id = user.Id
	request.User.Nik = user.Nik
	request.User.Name = user.Name
	request.User.Email = user.Email
	request.User.Status = user.ApplyStatus
	request.User.ApproveBy = userId

	request.Unit = strings.ToLower(request.Unit)
}

func (implementation *UserRegistrationMiddleware) Reject(ctx context.Context, tx *sql.Tx, request *webUserRegistration.UserRegistrationRequestReject) {
	userId := ValidateToken(ctx, implementation.RepositoryAuthInterface)

	ValidateUserPermission(ctx, tx, implementation.RepositoryUserInterface, userId, "user-approval")

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	request.User.Id = user.Id
	request.User.Nik = user.Nik
	request.User.Name = user.Name
	request.User.Email = user.Email
	request.User.Status = user.ApplyStatus
	request.User.RejectBy = userId
}
