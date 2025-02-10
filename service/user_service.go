package service

import (
	"context"
	"go-tobfa/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, id int, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.UserResponse
}
