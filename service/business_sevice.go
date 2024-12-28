package service

import (
	"context"
	"go-tobfa/model/web"
)

type BusinessService interface {
	Create(ctx context.Context, request web.BusinessCreateRequest) web.BusinessResponse
	Update(ctx context.Context, request web.BusinessUpdateRequest) web.BusinessResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.BusinessResponse
	FindAll(ctx context.Context) []web.BusinessResponse
	Stats(ctx context.Context, request web.BusinessStatsGetRequest) web.BusinessStatsResponse
}
