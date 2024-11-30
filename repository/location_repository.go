package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type LocationRepository interface {
	Provinces(ctx context.Context, tx *sql.Tx) []domain.Province
	Cities(ctx context.Context, tx *sql.Tx, provinceId int) []domain.City
	Districts(ctx context.Context, tx *sql.Tx, cityId int) []domain.District
	Villages(ctx context.Context, tx *sql.Tx, districtId int) []domain.Village
	Search(ctx context.Context, tx *sql.Tx, search string) []domain.Village
}
