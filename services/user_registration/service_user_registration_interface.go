package user_registration

import (
	"context"

	"github.com/backent/fra-golang/web/user_registration"
)

type ServiceUserRegistrationInterface interface {
	Apply(ctx context.Context, request user_registration.UserRegistrationRequestApply)
}
