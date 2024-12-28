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

func (repository *BusinessRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Business {
	userId := 1
	sql := "select * from businesses where user_id = ?"
	rows, err := tx.QueryContext(ctx, sql, userId)
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

func (repository *BusinessRepositoryImpl) FindByIds(ctx context.Context, tx *sql.Tx, businessIds *[]int) []domain.Business {
	if businessIds == nil || len(*businessIds) == 0 {
		return repository.FindAll(ctx, tx)
	} else {
		userId := 1
		var rows *sql.Rows
		var err error
		// Convert businessIds slice to a string suitable for SQL IN clause
		businessIdsParse := helper.ConvertListOfIntToString(*businessIds)
		sql := "SELECT * FROM businesses WHERE user_id = ? AND id IN (?)"
		rows, err = tx.QueryContext(ctx, sql, userId, businessIdsParse)

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
