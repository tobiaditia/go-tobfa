package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
)

type LocationRepositoryImpl struct {
}

func NewLocationRepository() LocationRepository {
	return &LocationRepositoryImpl{}
}

func (repository *LocationRepositoryImpl) Provinces(ctx context.Context, tx *sql.Tx) []domain.Province {
	sql := "select * from provinces"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var provinces []domain.Province
	for rows.Next() {
		province := domain.Province{}
		err := rows.Scan(
			&province.Id,
			&province.Code,
			&province.Name)
		helper.PanicIfError(err)
		provinces = append(provinces, province)
	}

	return provinces
}

func (repository *LocationRepositoryImpl) ProvinceFindById(ctx context.Context, tx *sql.Tx, id int) (domain.Province, error) {
	sql := "select * from provinces where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var province domain.Province
	if rows.Next() {
		err := rows.Scan(
			&province.Id,
			&province.Code,
			&province.Name)
		helper.PanicIfError(err)
		return province, nil
	} else {
		return province, errors.New("not found data")
	}
}

func (repository *LocationRepositoryImpl) Cities(ctx context.Context, tx *sql.Tx, provinceId int) []domain.City {
	sql := "select * from cities where province_id = ?"
	rows, err := tx.QueryContext(ctx, sql, provinceId)
	helper.PanicIfError(err)
	defer rows.Close()

	var cities []domain.City
	for rows.Next() {
		city := domain.City{}
		err := rows.Scan(
			&city.Id,
			&city.Code,
			&city.FullCode,
			&city.ProvinceId,
			&city.Type,
			&city.Name)
		helper.PanicIfError(err)
		cities = append(cities, city)
	}

	return cities
}

func (repository *LocationRepositoryImpl) CityFindById(ctx context.Context, tx *sql.Tx, id int) (domain.City, error) {
	sql := "select * from cities where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var city domain.City
	if rows.Next() {
		err := rows.Scan(
			&city.Id,
			&city.Code,
			&city.FullCode,
			&city.ProvinceId,
			&city.Type,
			&city.Name)
		helper.PanicIfError(err)
		return city, nil
	} else {
		return city, errors.New("not found data")
	}
}

func (repository *LocationRepositoryImpl) Districts(ctx context.Context, tx *sql.Tx, cityId int) []domain.District {
	sql := "select * from districts where city_id = ?"
	rows, err := tx.QueryContext(ctx, sql, cityId)
	helper.PanicIfError(err)
	defer rows.Close()

	var districts []domain.District
	for rows.Next() {
		district := domain.District{}
		err := rows.Scan(
			&district.Id,
			&district.Code,
			&district.FullCode,
			&district.CityId,
			&district.Name)
		helper.PanicIfError(err)
		districts = append(districts, district)
	}

	return districts
}

func (repository *LocationRepositoryImpl) DistrictFindById(ctx context.Context, tx *sql.Tx, id int) (domain.District, error) {
	sql := "select * from districts where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var district domain.District
	if rows.Next() {
		err := rows.Scan(
			&district.Id,
			&district.Code,
			&district.FullCode,
			&district.CityId,
			&district.Name)
		helper.PanicIfError(err)
		return district, nil
	} else {
		return district, errors.New("not found data")
	}
}

func (repository *LocationRepositoryImpl) Villages(ctx context.Context, tx *sql.Tx, districtId int) []domain.Village {
	sql := "select * from villages where district_id = ?"
	rows, err := tx.QueryContext(ctx, sql, districtId)
	helper.PanicIfError(err)
	defer rows.Close()

	var villages []domain.Village
	for rows.Next() {
		village := domain.Village{}
		err := rows.Scan(
			&village.Id,
			&village.Code,
			&village.FullCode,
			&village.PosCode,
			&village.DistrictId,
			&village.Name)
		helper.PanicIfError(err)
		villages = append(villages, village)
	}

	return villages
}

func (repository *LocationRepositoryImpl) VillageFindById(ctx context.Context, tx *sql.Tx, id int) (domain.Village, error) {
	sql := "select * from villages where id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var village domain.Village
	if rows.Next() {
		err := rows.Scan(
			&village.Id,
			&village.Code,
			&village.FullCode,
			&village.PosCode,
			&village.DistrictId,
			&village.Name)
		helper.PanicIfError(err)
		return village, nil
	} else {
		return village, errors.New("not found data")
	}
}

func (repository *LocationRepositoryImpl) Search(ctx context.Context, tx *sql.Tx, search string) []domain.Village {
	sql := "select * from villages where name LIKE ?"
	searchParam := "%" + search + "%"
	rows, err := tx.QueryContext(ctx, sql, searchParam)
	helper.PanicIfError(err)
	defer rows.Close()

	var villages []domain.Village
	for rows.Next() {
		village := domain.Village{}
		err := rows.Scan(
			&village.Id,
			&village.Code,
			&village.FullCode,
			&village.PosCode,
			&village.DistrictId,
			&village.Name)
		helper.PanicIfError(err)
		villages = append(villages, village)
	}

	return villages
}
