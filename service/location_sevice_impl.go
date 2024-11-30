package service

import (
	"context"
	"database/sql"
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/repository"

	"github.com/go-playground/validator/v10"
)

type LocationServiceImpl struct {
	LocationRepository repository.LocationRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewLocationService(locationRepository repository.LocationRepository, DB *sql.DB, validate *validator.Validate) LocationService {
	return &LocationServiceImpl{
		LocationRepository: locationRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *LocationServiceImpl) Provinces(ctx context.Context) []web.ProvinceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	response := service.LocationRepository.Provinces(ctx, tx)

	return helper.ToProvinceResponses(response)
}

func (service *LocationServiceImpl) Cities(ctx context.Context, provinceId int) []web.CityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	response := service.LocationRepository.Cities(ctx, tx, provinceId)

	return helper.ToCityResponses(response)
}

func (service *LocationServiceImpl) Districts(ctx context.Context, cityId int) []web.DistrictResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	response := service.LocationRepository.Districts(ctx, tx, cityId)

	return helper.ToDistrictResponses(response)
}

func (service *LocationServiceImpl) Villages(ctx context.Context, districtId int) []web.VillageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	response := service.LocationRepository.Villages(ctx, tx, districtId)

	return helper.ToVillageResponses(response)
}

func (service *LocationServiceImpl) Search(ctx context.Context, search string) []web.VillageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	response := service.LocationRepository.Search(ctx, tx, search)

	return helper.ToVillageResponses(response)
}
