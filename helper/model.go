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

func ToBusinessTransactionResponse(businessTransaction domain.BusinessTransaction) web.BusinessTransactionResponse {
	return web.BusinessTransactionResponse{
		Id:                        businessTransaction.Id,
		BusinessId:                businessTransaction.BusinessId,
		BusinessTransactionTypeId: businessTransaction.BusinessTransactionTypeId,
		BusinessTransactionItemId: businessTransaction.BusinessTransactionItemId,
		Total:                     businessTransaction.Total,
		Multiplier:                businessTransaction.Multiplier,
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
