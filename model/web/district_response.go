package web

type DistrictResponse struct {
	Id       int    `json:"id"`
	Code     string `json:"code"`
	FullCode string `json:"fullCode"`
	CityId   int    `json:"cityId"`
	Name     string `json:"name"`
}
