package user_registration

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesUserRegistration "github.com/backent/fra-golang/repositories/user_registration"
	webUserRegistration "github.com/backent/fra-golang/web/user_registration"
)

type ServiceUserRegistrationImpl struct {
	DB *sql.DB
	repositoriesUserRegistration.RepositoryUserRegistrationInterface
	*middlewares.UserRegistrationMiddleware
}

func NewServiceUserRegistrationImpl(db *sql.DB, repositoriesUserRegistration repositoriesUserRegistration.RepositoryUserRegistrationInterface, user_registrationMiddleware *middlewares.UserRegistrationMiddleware) ServiceUserRegistrationInterface {
	return &ServiceUserRegistrationImpl{
		DB:                                  db,
		RepositoryUserRegistrationInterface: repositoriesUserRegistration,
		UserRegistrationMiddleware:          user_registrationMiddleware,
	}
}

func (implementation *ServiceUserRegistrationImpl) Apply(ctx context.Context, request webUserRegistration.UserRegistrationRequestApply) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserRegistrationMiddleware.Apply(ctx, tx, &request)

	user_registration := models.UserRegistration{
		Nik:    request.Nik,
		Name:   request.Name,
		Email:  request.Email,
		Status: "pending",
	}

	_, err = implementation.RepositoryUserRegistrationInterface.Create(ctx, tx, user_registration)
	helpers.PanicIfError(err)
}

func (implementation *ServiceUserRegistrationImpl) FindAll(ctx context.Context, request webUserRegistration.UserRegistrationRequestFindAll) ([]webUserRegistration.UserRegistrationResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserRegistrationMiddleware.FindAll(ctx, tx, &request)

	userRegistrations, total, err := implementation.RepositoryUserRegistrationInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.QueryStatus)
	helpers.PanicIfError(err)

	return webUserRegistration.BulkUserRegistrationModelToUserRegistrationResponse(userRegistrations), total

}
