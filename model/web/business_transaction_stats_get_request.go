package web

type BusinessTransactionStatsGetRequest struct {
	DateStarted               string `validate:"required" json:"dateStarted"`
	DateEnded                 string `validate:"required" json:"dateEnded"`
	BusinessTransactionTypeId int    `validate:"required" json:"businessTransactionTypeId"`
	BusinessTransactionItemId int    `validate:"required" json:"businessTransactionItemId"`
	ProvinceId                int    `validate:"required" json:"provinceId"`
	CityId                    int    `validate:"required" json:"cityId"`
	DistrictId                int    `validate:"required" json:"districtId"`
	VillageId                 int    `validate:"required" json:"villageId"`
	BusinessCategoryId        int    ` json:"BusinessCategoryId"`
}
