package auth

import (
	"context"

	"github.com/backent/fra-golang/web/auth"
)

type ServiceAuthInterface interface {
	Login(ctx context.Context, request auth.LoginRequest) auth.LoginResponse
	Register(ctx context.Context, request auth.RegisterRequest) auth.RegisterResponse
}
