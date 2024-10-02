package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
)

type BusinessRepositoryImpl struct {
}

func NewBusinessRepository() BusinessRepository {
	return &BusinessRepositoryImpl{}
}

func (repository *BusinessRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.Business {
	sql := "select * from businesses"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var businesses []domain.Business
	for rows.Next() {
		business := domain.Business{}
		err := rows.Scan(
			&business.Id,
			&business.UserId,
			&business.Name,
			&business.Address,
			&business.BusinessCategoryId,
			&business.CountryId,
			&business.ProvinceId,
			&business.CityId,
			&business.DistrictId,
			&business.VillageId,
			&business.CreatedAt,
			&business.UpdatedAt)
		helper.PanicIfError(err)
		businesses = append(businesses, business)
	}

	return businesses
}

func (repository *BusinessRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, businessId int) (domain.Business, error) {
	sql := "select * from businesses where id = ?"
	rows, err := tx.QueryContext(ctx, sql, businessId)
	helper.PanicIfError(err)
	defer rows.Close()

	business := domain.Business{}

	if rows.Next() {
		err := rows.Scan(
			&business.Id,
			&business.UserId,
			&business.Name,
			&business.Address,
			&business.BusinessCategoryId,
			&business.CountryId,
			&business.ProvinceId,
			&business.CityId,
			&business.DistrictId,
			&business.VillageId,
			&business.CreatedAt,
			&business.UpdatedAt)
		helper.PanicIfError(err)
		return business, nil
	} else {
		return business, errors.New("not found data")
	}
}

func (repository *BusinessRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, business domain.Business) domain.Business {
	sql := "insert into businesses (user_id, name, address, business_category_id, country_id, province_id, city_id, district_id, village_id, created_at, updated_at) values (?,?,?,?,?,?,?,?,?, NOW(), NOW())"
	result, err := tx.ExecContext(ctx, sql, business.UserId, business.Name, business.Address, business.BusinessCategoryId, business.CountryId, business.ProvinceId, business.CityId, business.DistrictId, business.VillageId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	business.Id = int(id)
	return business
}

func (repository *BusinessRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, business domain.Business) domain.Business {
	sql := "update businesses set user_id = ?, name = ?, address = ?, business_category_id = ?, country_id = ?, province_id = ?, city_id = ?, district_id = ?, village_id = ?, updated_at = NOW() where id = ?"
	_, err := tx.ExecContext(ctx, sql, business.UserId, business.Name, business.Address, business.BusinessCategoryId, business.CountryId, business.ProvinceId, business.CityId, business.DistrictId, business.VillageId, business.Id)
	helper.PanicIfError(err)

	return business
}

func (repository *BusinessRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, business domain.Business) {
	sql := "delete from businesses where id = ?"
	_, err := tx.ExecContext(ctx, sql, business.Id)
	helper.PanicIfError(err)
}
