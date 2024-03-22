package user_registration

import (
	"context"
	"database/sql"
	"log"

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

func (implementation *ServiceUserRegistrationImpl) Approve(ctx context.Context, request webUserRegistration.UserRegistrationRequestApprove) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserRegistrationMiddleware.Approve(ctx, tx, &request)

	user_registration := models.UserRegistration{
		Id:        request.Id,
		Nik:       request.User.Nik,
		Name:      request.User.Name,
		Email:     request.User.Email,
		ApproveBy: request.User.ApproveBy,
		Unit:      request.Unit,
		Status:    "approve",
	}

	_, err = implementation.RepositoryUserRegistrationInterface.Update(ctx, tx, user_registration)
	helpers.PanicIfError(err)

	go func() {
		recipient := helpers.RecipientRegistration{
			Name:    user_registration.Name,
			Email:   user_registration.Email,
			Status:  "approved",
			Subject: "FRA - Registration Approved",
		}

		err = helpers.SendMail(recipient)
		if err != nil {
			log.Println("Error while sending email to :", recipient, err)
		}
	}()
}

func (implementation *ServiceUserRegistrationImpl) Reject(ctx context.Context, request webUserRegistration.UserRegistrationRequestReject) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserRegistrationMiddleware.Reject(ctx, tx, &request)

	user_registration := models.UserRegistration{
		Id:       request.Id,
		Nik:      request.User.Nik,
		Name:     request.User.Name,
		Email:    request.User.Email,
		RejectBy: request.User.RejectBy,
		Status:   "reject",
	}

	_, err = implementation.RepositoryUserRegistrationInterface.Update(ctx, tx, user_registration)
	helpers.PanicIfError(err)

	go func() {
		recipient := helpers.RecipientRegistration{
			Name:    user_registration.Name,
			Email:   user_registration.Email,
			Status:  "rejected",
			Subject: "FRA - Registration Rejected",
		}
		err = helpers.SendMail(recipient)
		if err != nil {
			log.Println("Error while sending email to :", recipient, err)
		}
	}()

}

func (implementation *ServiceUserRegistrationImpl) FindAll(ctx context.Context, request webUserRegistration.UserRegistrationRequestFindAll) ([]webUserRegistration.UserRegistrationResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserRegistrationMiddleware.FindAll(ctx, tx, &request)

	userRegistrations, total, err := implementation.RepositoryUserRegistrationInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.QueryStatus)
	helpers.PanicIfError(err)

	if len(webUserRegistration.BulkUserRegistrationModelToUserRegistrationResponse(userRegistrations)) == 0 {
		return []webUserRegistration.UserRegistrationResponse{}, total
	}

	return webUserRegistration.BulkUserRegistrationModelToUserRegistrationResponse(userRegistrations), total

}
