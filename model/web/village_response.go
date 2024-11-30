package web

type VillageResponse struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	FullCode   string `json:"fullCode"`
	PosCode    string `json:"posCode"`
	DistrictId int    `json:"districtId"`
	Name       string `json:"name"`
	FullName   string `json:"fullName"`
}
