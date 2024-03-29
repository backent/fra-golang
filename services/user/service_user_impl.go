package user

import (
	"context"
	"database/sql"

	"github.com/backent/fra-golang/exceptions"
	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/middlewares"
	"github.com/backent/fra-golang/models"
	repositoriesUser "github.com/backent/fra-golang/repositories/user"
	webUser "github.com/backent/fra-golang/web/user"
)

type ServiceUserImpl struct {
	DB *sql.DB
	repositoriesUser.RepositoryUserInterface
	*middlewares.UserMiddleware
}

func NewServiceUserImpl(db *sql.DB, repositoriesUser repositoriesUser.RepositoryUserInterface, userMiddleware *middlewares.UserMiddleware) ServiceUserInterface {
	return &ServiceUserImpl{
		DB:                      db,
		RepositoryUserInterface: repositoriesUser,
		UserMiddleware:          userMiddleware,
	}
}

func (implementation *ServiceUserImpl) Create(ctx context.Context, request webUser.UserRequestCreate) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.Create(ctx, tx, &request)

	hashedPassword, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	user := models.User{
		Nik:      request.Nik,
		Name:     request.Name,
		Password: hashedPassword,
	}

	user, err = implementation.RepositoryUserInterface.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToUserResponse(user)
}
func (implementation *ServiceUserImpl) Update(ctx context.Context, request webUser.UserRequestUpdate) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.Update(ctx, tx, &request)

	userPassword := request.CurrentPassword

	if request.IsEmptyPassword {
		userPassword = ""
	} else if request.Password != "" {
		userPassword, err = helpers.HashPassword(request.Password)
		helpers.PanicIfError(err)
	}

	user := request.User
	user.Password = userPassword
	user.Unit = request.Unit
	user.Role = request.Role

	user, err = implementation.RepositoryUserInterface.Update(ctx, tx, user)
	helpers.PanicIfError(err)

	return webUser.UserModelToUserResponse(user)
}
func (implementation *ServiceUserImpl) Delete(ctx context.Context, request webUser.UserRequestDelete) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.Delete(ctx, tx, &request)

	err = implementation.RepositoryUserInterface.Delete(ctx, tx, request.Id)
	helpers.PanicIfError(err)

}
func (implementation *ServiceUserImpl) FindById(ctx context.Context, request webUser.UserRequestFindById) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.FindById(ctx, tx, &request)

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	return webUser.UserModelToUserResponse(user)
}
func (implementation *ServiceUserImpl) FindAll(ctx context.Context, request webUser.UserRequestFindAll) ([]webUser.UserResponse, int) {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.FindAll(ctx, tx, &request)

	users, total, err := implementation.RepositoryUserInterface.FindAll(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection(), request.QueryStatus, request.QuerySearch)
	helpers.PanicIfError(err)

	if total > 0 {
		return webUser.BulkUserModelToUserResponse(users), total
	}

	return []webUser.UserResponse{}, total
}

func (implementation *ServiceUserImpl) FindAllWithRisksDetail(ctx context.Context, request webUser.UserRequestFindAll) []webUser.UserResponseWithRisksDetail {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.FindAll(ctx, tx, &request)

	users, err := implementation.RepositoryUserInterface.FindAllWithRisksDetail(ctx, tx, request.GetTake(), request.GetSkip(), request.GetOrderBy(), request.GetOrderDirection())
	helpers.PanicIfError(err)

	return webUser.BulkUserModelToUserResponseWithRisksDetail(users)
}
func (implementation *ServiceUserImpl) CurrentUser(ctx context.Context, request webUser.UserRequestCurrentUser) webUser.UserResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.UserMiddleware.CurrentUser(ctx, tx, &request)

	user, err := implementation.RepositoryUserInterface.FindById(ctx, tx, request.UserId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return webUser.UserModelToUserResponse(user)
}
