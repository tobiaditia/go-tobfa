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

	"github.com/go-playground/validator/v10"
)

type BusinessServiceImpl struct {
	BusinessRepository                repository.BusinessRepository
	BusinessTransactionItemRepository repository.BusinessTransactionItemRepository
	DB                                *sql.DB
	Validate                          *validator.Validate
}

func NewBusinessService(businessRepository repository.BusinessRepository, businessTransactionItemRepository repository.BusinessTransactionItemRepository, DB *sql.DB, validate *validator.Validate) BusinessService {
	return &BusinessServiceImpl{
		BusinessRepository:                businessRepository,
		BusinessTransactionItemRepository: businessTransactionItemRepository,
		DB:                                DB,
		Validate:                          validate,
	}
}

func (service *BusinessServiceImpl) Create(ctx context.Context, request web.BusinessCreateRequest) web.BusinessResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	address := sql.NullString{Valid: false}
	if request.Address != nil {
		address = sql.NullString{String: *request.Address, Valid: true}
	}

	business := domain.Business{
		UserId:             request.UserId,
		Name:               request.Name,
		Address:            address,
		BusinessCategoryId: request.BusinessCategoryId,
		CountryId:          request.CountryId,
		ProvinceId:         request.ProvinceId,
		CityId:             request.CityId,
		DistrictId:         request.DistrictId,
		VillageId:          request.VillageId,
	}

	business = service.BusinessRepository.Create(ctx, tx, business)

	return helper.ToBusinessResponse(business)

}

func (service *BusinessServiceImpl) Update(ctx context.Context, id int, request web.BusinessUpdateRequest) web.BusinessResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	business, err := service.BusinessRepository.Find(ctx, tx, id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	business.Address = sql.NullString{Valid: false}
	if request.Address != nil {
		business.Address = sql.NullString{String: *request.Address, Valid: true}
	}
	business.UserId = request.UserId
	business.Name = request.Name
	business.BusinessCategoryId = request.BusinessCategoryId
	business.CountryId = request.CountryId
	business.ProvinceId = request.ProvinceId
	business.CityId = request.CityId
	business.DistrictId = request.DistrictId
	business.VillageId = request.VillageId

	business = service.BusinessRepository.Update(ctx, tx, business)

	return helper.ToBusinessResponse(business)
}

func (service *BusinessServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	business, err := service.BusinessRepository.Find(ctx, tx, id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BusinessRepository.Delete(ctx, tx, business)
}

func (service *BusinessServiceImpl) FindById(ctx context.Context, id int) web.BusinessResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	business, err := service.BusinessRepository.Find(ctx, tx, id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBusinessResponse(business)
}

func (service *BusinessServiceImpl) FindAll(ctx context.Context) []web.BusinessResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businesses := service.BusinessRepository.FindAll(ctx, tx)

	return helper.ToBusinessResponses(businesses)
}

func (service *BusinessServiceImpl) Stats(ctx context.Context, request web.BusinessStatsGetRequest) web.BusinessStatsResponse {
	fmt.Println(request)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	// panic(nil)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	businesses := service.BusinessRepository.FindByIds(ctx, tx, &request.BusinessIds)
	businessTransactionItems := service.BusinessTransactionItemRepository.FindAll(ctx, tx)
	fmt.Println(businesses)
	// panic(nil)

	var businessResponses []web.Business
	for _, business := range businesses {
		var transactionItemStats []web.TransactionItemStat

		for _, businessTransactionItem := range businessTransactionItems {
			transactionItemStats = append(transactionItemStats, web.TransactionItemStat{
				Name:                     businessTransactionItem.Name,
				BusinessTransactionTypes: []web.TransactionTypeStat{},
				Total:                    0,
				TotalSell:                0,
				TotalBuy:                 0,
			})
		}

		businessResponses = append(businessResponses, web.Business{
			Id:                       business.Id,
			Name:                     business.Name,
			BusinessTransactionItems: transactionItemStats,
			Total:                    0,
			TotalSell:                0,
			TotalBuy:                 0,
		})

	}

	return web.BusinessStatsResponse{
		Name:                       "Stats of transactions",
		DateStarted:                request.DateStarted,
		DateEnded:                  request.DateEnded,
		BusinessTransactionTypeIds: request.BusinessTransactionItemIds,
		BusinessTransactionItemIds: request.BusinessTransactionItemIds,
		Businesses:                 businessResponses,
		BusinessTransactionType:    []web.TransactionTypeStat{},
		Total:                      0,
		TotalSell:                  0,
		TotalBuy:                   0,
	}

	// var businessTransactionStatsResponses []web.BusinessTransactionStatsResponse

	// for _, businessCategory := range businessCategories {

	// 	request.BusinessCategoryId = businessCategory.Id

	// 	businessTransactions := service.BusinessTransactionRepository.FindForStats(ctx, tx, request)

	// 	var stats []domain.Stat
	// 	totalBusinessTransaction := 0

	// 	for _, businessTransaction := range businessTransactions {

	// 		stats = append(stats, domain.Stat{businessTransaction.Date, businessTransaction.Total})
	// 		totalBusinessTransaction += businessTransaction.Total
	// 	}

	// 	averange := totalBusinessTransaction / len(businessTransactions)
	// 	businessTransactionStatsResponses = append(businessTransactionStatsResponses, helper.ToBusinessTransactionStatResponse(averange, stats, businessCategory))
	// }

	// return businessTransactionStatsResponses
}
