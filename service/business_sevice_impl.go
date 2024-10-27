package service

import (
	"context"
	"database/sql"
	"go-tobfa/exception"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
	"go-tobfa/repository"

	"github.com/go-playground/validator/v10"
)

type BusinessServiceImpl struct {
	BusinessRepository repository.BusinessRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBusinessService(businessRepository repository.BusinessRepository, DB *sql.DB, validate *validator.Validate) BusinessService {
	return &BusinessServiceImpl{
		BusinessRepository: businessRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *BusinessServiceImpl) Create(ctx context.Context, request web.BusinessCreateRequest) web.BusinessResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	business := domain.Business{
		UserId:             request.UserId,
		Name:               request.Name,
		Address:            request.Address,
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

func (service *BusinessServiceImpl) Update(ctx context.Context, request web.BusinessUpdateRequest) web.BusinessResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	business, err := service.BusinessRepository.Find(ctx, tx, request.Id)
	if nil != err {
		panic(exception.NewNotFoundError(err.Error()))
	}

	business.UserId = request.UserId
	business.Name = request.Name
	business.Address = request.Address
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

	businesses := service.BusinessRepository.Get(ctx, tx)

	return helper.ToBusinessResponses(businesses)
}
