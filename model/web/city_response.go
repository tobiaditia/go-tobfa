package web

type CityResponse struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	FullCode   string `json:"fullCode"`
	ProvinceId int    `json:"provinceId"`
	Type       int    `json:"type"`
	Name       string `json:"name"`
}
