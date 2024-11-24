package service

import (
	"context"
	"go-tobfa/model/web"
)

type BusinessTransactionService interface {
	FindAll(ctx context.Context) []web.BusinessTransactionResponse
	FindById(ctx context.Context, id int) web.BusinessTransactionResponse
	FindByBusiness(ctx context.Context, businessId int) []web.BusinessTransactionResponse
	Stats(ctx context.Context, param web.BusinessTransactionStatsGetRequest) []web.BusinessTransactionStatsResponse
	Create(ctx context.Context, request web.BusinessTransactionCreateRequest) web.BusinessTransactionResponse
	Update(ctx context.Context, request web.BusinessTransactionUpdateRequest) web.BusinessTransactionResponse
	Delete(ctx context.Context, id int)
}
