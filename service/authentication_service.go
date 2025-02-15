package service

import (
	"context"
	"go-tobfa/model/web"
)

type AuthenticationService interface {
	Login(ctx context.Context, request web.LoginRequest) web.AuthResponse
}
