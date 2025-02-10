package service

import (
	"context"
	"go-tobfa/model/web"
)

type LocationService interface {
	Provinces(ctx context.Context) []web.ProvinceResponse
	Cities(ctx context.Context, provinceId int) []web.CityResponse
	Districts(ctx context.Context, cityId int) []web.DistrictResponse
	Villages(ctx context.Context, districtId int) []web.VillageResponse
	Search(ctx context.Context, search string) []web.VillageResponse
}
