package service

import (
	"context"
	"go-tobfa/model/web"
)

type BusinessService interface {
	Create(ctx context.Context, request web.BusininessCreateRequest) web.BusinessResponse
	Update(ctx context.Context, request web.BusininessUpdateRequest) web.BusinessResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.BusinessResponse
	FindAll(ctx context.Context) []web.BusinessResponse
}
