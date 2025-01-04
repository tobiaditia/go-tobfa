package web

import "time"

type BusinessResponse struct {
	Id                 int       `json:"id"`
	UserId             int       `json:"userId"`
	Name               string    `json:"name"`
	Address            *string   `json:"address"`
	BusinessCategoryId int       `json:"businessCategoryId"`
	CountryId          int       `json:"countryId"`
	ProvinceId         int       `json:"provinceId"`
	CityId             int       `json:"cityId"`
	DistrictId         int       `json:"districtId"`
	VillageId          int       `json:"villageId"`
	CreatedAt          time.Time `json:"createdAt"`
}
