package helper

import (
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
)

func ToBusinessResponse(business domain.Business) web.BusinessResponse {
	return web.BusinessResponse{
		Id:                 business.Id,
		UserId:             business.UserId,
		Name:               business.Name,
		Address:            business.Address,
		BusinessCategoryId: business.BusinessCategoryId,
		CountryId:          business.CountryId,
		ProvinceId:         business.ProvinceId,
		CityId:             business.CityId,
		DistrictId:         business.DistrictId,
		VillageId:          business.VillageId,
	}
}

func ToBusinessResponses(businesses []domain.Business) []web.BusinessResponse {
	var businessResponses []web.BusinessResponse
	for _, business := range businesses {
		businessResponses = append(businessResponses, ToBusinessResponse(business))
	}

	return businessResponses
}
