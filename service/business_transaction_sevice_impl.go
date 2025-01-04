package service

import (
	"context"
	"database/sql"
	"fmt"
	"go-tobfa/exception"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
	"go-tobfa/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type BusinessTransactionServiceImpl struct {
	BusinessTransactionRepository repository.BusinessTransactionRepository
	BusinessCategoryRepository    repository.BusinessCategoryRepository
	DB                            *sql.DB
	Validate                      *validator.Validate
}

func NewBusinessTransactionService(businessTransactionRepository repository.BusinessTransactionRepository, businessCategoryRepository repository.BusinessCategoryRepository, DB *sql.DB, validate *validator.Validate) BusinessTransactionService {
	return &BusinessTransactionServiceImpl{
		BusinessTransactionRepository: businessTransactionRepository,
		BusinessCategoryRepository:    businessCategoryRepository,
		DB:                            DB,
		Validate:                      validate,
	}
}

func (service *BusinessTransactionServiceImpl) FindAll(ctx context.Context) []web.BusinessTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessTransactiones := service.BusinessTransactionRepository.FindAll(ctx, tx)

	return helper.ToBusinessTransactionResponses(businessTransactiones)
}

func (service *BusinessTransactionServiceImpl) FindById(ctx context.Context, id int) web.BusinessTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessTransaction, err := service.BusinessTransactionRepository.FindById(ctx, tx, id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBusinessTransactionResponse(businessTransaction)
}

func (service *BusinessTransactionServiceImpl) FindByBusiness(ctx context.Context, businessId int) []web.BusinessTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessTransactions := service.BusinessTransactionRepository.FindByBusiness(ctx, tx, businessId)

	return helper.ToBusinessTransactionResponses(businessTransactions)
}

func (service *BusinessTransactionServiceImpl) Stats(ctx context.Context, request web.BusinessTransactionStatsGetRequest) []web.BusinessTransactionStatsResponse {

	err := service.Validate.Struct(request)
	fmt.Println(request)
	fmt.Println(err)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessCategories := service.BusinessCategoryRepository.FindAll(ctx, tx)

	var businessTransactionStatsResponses []web.BusinessTransactionStatsResponse

	for _, businessCategory := range businessCategories {

		request.BusinessCategoryId = businessCategory.Id

		businessTransactions := service.BusinessTransactionRepository.FindForStats(ctx, tx, request)

		var stats []domain.Stat
		totalBusinessTransaction := 0

		for _, businessTransaction := range businessTransactions {

			stats = append(stats, domain.Stat{Date: businessTransaction.Date, Total: businessTransaction.Total})
			totalBusinessTransaction += businessTransaction.Total
		}

		averange := totalBusinessTransaction / len(businessTransactions)
		businessTransactionStatsResponses = append(businessTransactionStatsResponses, helper.ToBusinessTransactionStatResponse(averange, stats, businessCategory))
	}

	return businessTransactionStatsResponses
}

func (service *BusinessTransactionServiceImpl) Create(ctx context.Context, request web.BusinessTransactionCreateRequest) web.BusinessTransactionResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	date, err := time.Parse("2024-01-01", request.Date)
	helper.PanicIfError(err)

	businessTransaction := domain.BusinessTransaction{
		BusinessId:                request.BusinessId,
		BusinessTransactionTypeId: request.BusinessTransactionTypeId,
		BusinessTransactionItemId: request.BusinessTransactionItemId,
		Total:                     request.Total,
		Multiplier:                request.Multiplier,
		Date:                      date,
		Description:               request.Description,
	}

	businessTransaction = service.BusinessTransactionRepository.Create(ctx, tx, businessTransaction)

	return helper.ToBusinessTransactionResponse(businessTransaction)

}

func (service *BusinessTransactionServiceImpl) Update(ctx context.Context, request web.BusinessTransactionUpdateRequest) web.BusinessTransactionResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessTransaction, err := service.BusinessTransactionRepository.FindById(ctx, tx, request.Id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	date, err := time.Parse("2024-01-01", request.Date)
	helper.PanicIfError(err)

	businessTransaction.BusinessId = request.BusinessId
	businessTransaction.BusinessTransactionTypeId = request.BusinessTransactionTypeId
	businessTransaction.BusinessTransactionItemId = request.BusinessTransactionItemId
	businessTransaction.Total = request.Total
	businessTransaction.Multiplier = request.Multiplier
	businessTransaction.Date = date
	businessTransaction.Description = request.Description

	businessTransaction = service.BusinessTransactionRepository.Update(ctx, tx, businessTransaction)

	return helper.ToBusinessTransactionResponse(businessTransaction)
}

func (service *BusinessTransactionServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businessTransaction, err := service.BusinessTransactionRepository.FindById(ctx, tx, id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BusinessTransactionRepository.Delete(ctx, tx, businessTransaction)
}
