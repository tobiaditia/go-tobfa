package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
)

type BusinessTransactionRepositoryImpl struct {
}

func NewBusinessTransactionRepository() BusinessTransactionRepository {
	return &BusinessTransactionRepositoryImpl{}
}

func (repository *BusinessTransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.BusinessTransaction {
	sql := "select * from business_transactions"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var businessTransactions []domain.BusinessTransaction
	for rows.Next() {
		businessTransaction := domain.BusinessTransaction{}
		err := rows.Scan(
			&businessTransaction.Id,
			&businessTransaction.BusinessId,
			&businessTransaction.BusinessTransactionTypeId,
			&businessTransaction.BusinessTransactionItemId,
			&businessTransaction.Total,
			&businessTransaction.Multiplier,
			&businessTransaction.Date,
			&businessTransaction.Description,
			&businessTransaction.CreatedAt,
			&businessTransaction.UpdatedAt)
		helper.PanicIfError(err)
		businessTransactions = append(businessTransactions, businessTransaction)
	}

	return businessTransactions
}

func (repository *BusinessTransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, businessTransactionId int) (domain.BusinessTransaction, error) {
	sql := "select * from business_transactions where id = ?"
	rows, err := tx.QueryContext(ctx, sql, businessTransactionId)
	helper.PanicIfError(err)
	defer rows.Close()

	businessTransaction := domain.BusinessTransaction{}

	if rows.Next() {
		err := rows.Scan(
			&businessTransaction.Id,
			&businessTransaction.BusinessId,
			&businessTransaction.BusinessTransactionTypeId,
			&businessTransaction.BusinessTransactionItemId,
			&businessTransaction.Total,
			&businessTransaction.Multiplier,
			&businessTransaction.Date,
			&businessTransaction.Description,
			&businessTransaction.CreatedAt,
			&businessTransaction.UpdatedAt)
		helper.PanicIfError(err)
		return businessTransaction, nil
	} else {
		return businessTransaction, errors.New("not found data")
	}
}

func (repository *BusinessTransactionRepositoryImpl) FindByBusiness(ctx context.Context, tx *sql.Tx, businessId int) []domain.BusinessTransaction {
	sql := "select * from business_transactions where business_id = ? "
	rows, err := tx.QueryContext(ctx, sql, businessId)
	helper.PanicIfError(err)
	defer rows.Close()

	var businessTransactions []domain.BusinessTransaction
	for rows.Next() {
		businessTransaction := domain.BusinessTransaction{}
		err := rows.Scan(
			&businessTransaction.Id,
			&businessTransaction.BusinessId,
			&businessTransaction.BusinessTransactionTypeId,
			&businessTransaction.BusinessTransactionItemId,
			&businessTransaction.Total,
			&businessTransaction.Multiplier,
			&businessTransaction.Date,
			&businessTransaction.Description,
			&businessTransaction.CreatedAt,
			&businessTransaction.UpdatedAt)
		helper.PanicIfError(err)
		businessTransactions = append(businessTransactions, businessTransaction)
	}

	return businessTransactions
}

func (repository *BusinessTransactionRepositoryImpl) FindForStats(ctx context.Context, tx *sql.Tx, param web.BusinessTransactionStatsGetRequest) []domain.BusinessTransaction {

	sql := `SELECT 
				bt.id, 
				bt.business_id, 
				bt.business_transaction_type_id, 
				bt.business_transaction_item_id, 
				bt.total, 
				bt.multiplier, 
				bt.date, 
				bt.description, 
				bt.created_at, 
				bt.updated_at
			FROM business_transactions bt
			JOIN businesses b ON bt.business_id = b.id
			WHERE bt.date BETWEEN ? AND ?
			AND b.business_category_id = ?
			AND bt.business_transaction_item_id = ?
			AND bt.business_transaction_type_id = ?
			AND (? = 0 OR b.province_id = ?)
			AND (? = 0 OR b.city_id = ?)
			AND (? = 0 OR b.district_id = ?)
			AND (? = 0 OR b.village_id = ?);
	`

	rows, err := tx.QueryContext(
		ctx,
		sql,
		param.DateStarted,
		param.DateEnded,
		param.BusinessCategoryId,
		param.BusinessTransactionItemId,
		param.BusinessTransactionTypeId,
		param.ProvinceId,
		param.ProvinceId,
		param.CityId,
		param.CityId,
		param.DistrictId,
		param.DistrictId,
		param.VillageId,
		param.VillageId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	var businessTransactions []domain.BusinessTransaction
	for rows.Next() {
		businessTransaction := domain.BusinessTransaction{}
		err := rows.Scan(
			&businessTransaction.Id,
			&businessTransaction.BusinessId,
			&businessTransaction.BusinessTransactionTypeId,
			&businessTransaction.BusinessTransactionItemId,
			&businessTransaction.Total,
			&businessTransaction.Multiplier,
			&businessTransaction.Date,
			&businessTransaction.Description,
			&businessTransaction.CreatedAt,
			&businessTransaction.UpdatedAt)

		helper.PanicIfError(err)
		businessTransactions = append(businessTransactions, businessTransaction)
	}

	return businessTransactions
}

func (repository *BusinessTransactionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, businessTransaction domain.BusinessTransaction) domain.BusinessTransaction {
	sql := "insert into business_transactions (user_id, name, address, business_category_id, country_id, province_id, city_id, district_id, village_id, created_at, updated_at) values (?,?,?,?,?,?,?,?,?, NOW(), NOW())"
	result, err := tx.ExecContext(ctx, sql, businessTransaction.BusinessId, businessTransaction.BusinessTransactionTypeId, businessTransaction.BusinessTransactionItemId, businessTransaction.Total, businessTransaction.Multiplier, businessTransaction.Date, businessTransaction.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	businessTransaction.Id = int(id)
	return businessTransaction
}

func (repository *BusinessTransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, businessTransaction domain.BusinessTransaction) domain.BusinessTransaction {
	sql := "update business_transactions set user_id = ?, name = ?, address = ?, business_category_id = ?, country_id = ?, province_id = ?, city_id = ?, district_id = ?, village_id = ?, updated_at = NOW() where id = ?"
	_, err := tx.ExecContext(ctx, sql, businessTransaction.BusinessId, businessTransaction.BusinessTransactionTypeId, businessTransaction.BusinessTransactionItemId, businessTransaction.Total, businessTransaction.Multiplier, businessTransaction.Date, businessTransaction.Description, businessTransaction.Id)
	helper.PanicIfError(err)

	return businessTransaction
}

func (repository *BusinessTransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, businessTransaction domain.BusinessTransaction) {
	sql := "delete from business_transactions where id = ?"
	_, err := tx.ExecContext(ctx, sql, businessTransaction.Id)
	helper.PanicIfError(err)
}
