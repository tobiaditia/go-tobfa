package service

import (
	"context"
	"database/sql"
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/repository"

	"github.com/go-playground/validator/v10"
)

type BusinessCategoryServiceImpl struct {
	BusinessCategoryRepository repository.BusinessCategoryRepository
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewBusinessCategoryService(businessCategoryRepository repository.BusinessCategoryRepository, DB *sql.DB, validate *validator.Validate) BusinessCategoryService {
	return &BusinessCategoryServiceImpl{
		BusinessCategoryRepository: businessCategoryRepository,
		DB:                         DB,
		Validate:                   validate,
	}
}

func (service *BusinessCategoryServiceImpl) FindAll(ctx context.Context) []web.BusinessCategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessCategories := service.BusinessCategoryRepository.FindAll(ctx, tx)

	return helper.ToBusinessCategoryResponses(businessCategories)
}
