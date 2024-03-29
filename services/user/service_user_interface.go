package user

import (
	"context"

	"github.com/backent/fra-golang/web/user"
)

type ServiceUserInterface interface {
	Create(ctx context.Context, request user.UserRequestCreate) user.UserResponse
	Update(ctx context.Context, request user.UserRequestUpdate) user.UserResponse
	Delete(ctx context.Context, request user.UserRequestDelete)
	FindById(ctx context.Context, request user.UserRequestFindById) user.UserResponse
	FindAll(ctx context.Context, request user.UserRequestFindAll) ([]user.UserResponse, int)
	FindAllWithRisksDetail(ctx context.Context, request user.UserRequestFindAll) []user.UserResponseWithRisksDetail
	CurrentUser(ctx context.Context, request user.UserRequestCurrentUser) user.UserResponse
}
