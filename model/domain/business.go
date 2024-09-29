package domain

type Business struct {
	Id                 int
	UserId             int
	Name               string
	Address            string
	BusinessCategoryId int
	CountryId          int
	ProvinceId         int
	CityId             int
	DistrictId         int
	VillageId          int
	CreatedAt          string
	UpdatedAt          string
}
