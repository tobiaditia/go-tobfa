package service

import (
	"context"
	"go-tobfa/model/web"
)

type BusinessCategoryService interface {
	FindAll(ctx context.Context) []web.BusinessCategoryResponse
}
