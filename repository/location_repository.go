package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type LocationRepository interface {
	Provinces(ctx context.Context, tx *sql.Tx) []domain.Province
	ProvinceFindById(ctx context.Context, tx *sql.Tx, id int) (domain.Province, error)
	Cities(ctx context.Context, tx *sql.Tx, provinceId int) []domain.City
	CityFindById(ctx context.Context, tx *sql.Tx, id int) (domain.City, error)
	Districts(ctx context.Context, tx *sql.Tx, cityId int) []domain.District
	DistrictFindById(ctx context.Context, tx *sql.Tx, id int) (domain.District, error)
	Villages(ctx context.Context, tx *sql.Tx, districtId int) []domain.Village
	VillageFindById(ctx context.Context, tx *sql.Tx, id int) (domain.Village, error)
	Search(ctx context.Context, tx *sql.Tx, search string) []domain.Village
}
