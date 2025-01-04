package web

type BusinessUpdateRequest struct {
	UserId             int     `validate:"required" json:"userId"`
	Name               string  `validate:"required,max=100,min=3" json:"name"`
	Address            *string `validate:"" json:"address"`
	BusinessCategoryId int     `validate:"required" json:"businessCategoryId"`
	CountryId          int     `validate:"required" json:"countryId"`
	ProvinceId         int     `validate:"required" json:"provinceId"`
	CityId             int     `validate:"required" json:"cityId"`
	DistrictId         int     `validate:"required" json:"districtId"`
	VillageId          int     `validate:"required" json:"villageId"`
}
