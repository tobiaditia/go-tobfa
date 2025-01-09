package helper

import (
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
)

func ToBusinessResponse(business domain.Business) web.BusinessResponse {
	address := func() *string {
		if business.Address.Valid {
			return &business.Address.String
		}
		return nil
	}()
	return web.BusinessResponse{
		Id:                 business.Id,
		UserId:             business.UserId,
		Name:               business.Name,
		Address:            address,
		BusinessCategoryId: business.BusinessCategoryId,
		CountryId:          business.CountryId,
		ProvinceId:         business.ProvinceId,
		CityId:             business.CityId,
		DistrictId:         business.DistrictId,
		VillageId:          business.VillageId,
		CreatedAt:          business.CreatedAt,
	}
}

func ToBusinessResponses(businesses []domain.Business) []web.BusinessResponse {
	var businessResponses []web.BusinessResponse
	for _, business := range businesses {
		businessResponses = append(businessResponses, ToBusinessResponse(business))
	}

	return businessResponses
}

func ToBusinessCategoryResponse(businessCategory domain.BusinessCategory) web.BusinessCategoryResponse {
	return web.BusinessCategoryResponse{
		Id:   businessCategory.Id,
		Name: businessCategory.Name,
	}
}

func ToBusinessCategoryResponses(businesses []domain.BusinessCategory) []web.BusinessCategoryResponse {
	var businessCategoryResponses []web.BusinessCategoryResponse
	for _, business := range businesses {
		businessCategoryResponses = append(businessCategoryResponses, ToBusinessCategoryResponse(business))
	}

	return businessCategoryResponses
}

func ToUserResponse(business domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        business.Id,
		Name:      business.Name,
		Email:     business.Email,
		Handphone: business.Handphone,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToBusinessTransactionResponse(businessTransaction domain.BusinessTransaction) web.BusinessTransactionResponse {
	return web.BusinessTransactionResponse{
		Id:                        businessTransaction.Id,
		BusinessId:                businessTransaction.BusinessId,
		BusinessTransactionTypeId: businessTransaction.BusinessTransactionTypeId,
		BusinessTransactionItemId: businessTransaction.BusinessTransactionItemId,
		Total:                     businessTransaction.Total,
		Quantity:                  businessTransaction.Quantity,
		Date:                      businessTransaction.Date,
		Description:               businessTransaction.Description,
	}
}

func ToBusinessTransactionResponses(businessTransactions []domain.BusinessTransaction) []web.BusinessTransactionResponse {
	var businessTransactionResponses []web.BusinessTransactionResponse
	for _, businessTransaction := range businessTransactions {
		businessTransactionResponses = append(businessTransactionResponses, ToBusinessTransactionResponse(businessTransaction))
	}

	return businessTransactionResponses
}

func ToStatResponse(stat domain.Stat) web.StatsResponse {
	return web.StatsResponse{
		Date:  stat.Date,
		Total: stat.Total,
	}
}

func ToStatResponses(stats []domain.Stat) []web.StatsResponse {
	var statResponses []web.StatsResponse
	for _, stat := range stats {
		statResponses = append(statResponses, ToStatResponse(stat))
	}

	return statResponses
}

func ToBusinessTransactionStatResponse(averange int, stats []domain.Stat, businessCategory domain.BusinessCategory) web.BusinessTransactionStatsResponse {
	return web.BusinessTransactionStatsResponse{
		Averange:         averange,
		BusinessCategory: ToBusinessCategoryResponse(businessCategory),
		Stats:            ToStatResponses(stats),
	}
}

func ToProvinceResponse(province domain.Province) web.ProvinceResponse {
	return web.ProvinceResponse{
		Id:   province.Id,
		Name: province.Name,
	}
}

func ToProvinceResponses(provinces []domain.Province) []web.ProvinceResponse {
	var provinceResponses []web.ProvinceResponse
	for _, province := range provinces {
		provinceResponses = append(provinceResponses, ToProvinceResponse(province))
	}

	return provinceResponses
}

func ToCityResponse(city domain.City) web.CityResponse {
	return web.CityResponse{
		Id:         city.Id,
		Code:       city.Code,
		FullCode:   city.FullCode,
		ProvinceId: city.ProvinceId,
		Type:       city.Type,
		Name:       city.Name,
	}
}

func ToCityResponses(cities []domain.City) []web.CityResponse {
	var cityResponses []web.CityResponse
	for _, city := range cities {
		cityResponses = append(cityResponses, ToCityResponse(city))
	}

	return cityResponses
}

func ToDistrictResponse(district domain.District) web.DistrictResponse {
	return web.DistrictResponse{
		Id:       district.Id,
		Code:     district.Code,
		FullCode: district.FullCode,
		CityId:   district.CityId,
		Name:     district.Name,
	}
}

func ToDistrictResponses(districts []domain.District) []web.DistrictResponse {
	var districtResponses []web.DistrictResponse
	for _, district := range districts {
		districtResponses = append(districtResponses, ToDistrictResponse(district))
	}

	return districtResponses
}

func ToVillageResponse(village domain.Village) web.VillageResponse {
	return web.VillageResponse{
		Id:         village.Id,
		Code:       village.Code,
		FullCode:   village.FullCode,
		PosCode:    village.PosCode,
		DistrictId: village.DistrictId,
		Name:       village.Name,
		FullName:   "a",
	}
}

func ToVillageResponses(villages []domain.Village) []web.VillageResponse {
	var villageResponses []web.VillageResponse
	for _, village := range villages {
		villageResponses = append(villageResponses, ToVillageResponse(village))
	}

	return villageResponses
}
