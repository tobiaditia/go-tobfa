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
