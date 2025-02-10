package service

import (
	"context"
	"database/sql"
	"fmt"
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

	villages := service.LocationRepository.Villages(ctx, tx, districtId)
	district, err := service.LocationRepository.DistrictFindById(ctx, tx, districtId)
	helper.PanicIfError(err)
	city, err := service.LocationRepository.CityFindById(ctx, tx, district.CityId)
	helper.PanicIfError(err)
	province, err := service.LocationRepository.ProvinceFindById(ctx, tx, city.ProvinceId)
	helper.PanicIfError(err)

	var responses []web.VillageResponse
	for _, village := range villages {
		villageResponse := helper.ToVillageResponse(village)
		villageResponse.FullName = fmt.Sprintf("%s - %s - %s - %s ( %s )", province.Name, city.Name, district.Name, village.Name, village.Code)
		responses = append(responses, villageResponse)
	}

	return responses
}

func (service *LocationServiceImpl) Search(ctx context.Context, search string) []web.VillageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	villages := service.LocationRepository.Search(ctx, tx, search)

	var responses []web.VillageResponse
	for _, village := range villages {

		district, err := service.LocationRepository.DistrictFindById(ctx, tx, village.DistrictId)
		helper.PanicIfError(err)
		city, err := service.LocationRepository.CityFindById(ctx, tx, district.CityId)
		helper.PanicIfError(err)
		province, err := service.LocationRepository.ProvinceFindById(ctx, tx, city.ProvinceId)
		helper.PanicIfError(err)

		villageResponse := helper.ToVillageResponse(village)
		villageResponse.FullName = fmt.Sprintf("%s - %s - %s - %s ( %s )", province.Name, city.Name, district.Name, village.Name, village.Code)
		responses = append(responses, villageResponse)
	}

	return responses
}
