package repository

import (
	"context"
	"database/sql"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
)

type BusinessTransactionItemRepositoryImpl struct {
}

func NewBusinessTransactionItemRepository() BusinessTransactionItemRepository {
	return &BusinessTransactionItemRepositoryImpl{}
}

func (repository *BusinessTransactionItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.BusinessTransactionItem {
	sql := "select * from business_transaction_items"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var businessTransactionItems []domain.BusinessTransactionItem
	for rows.Next() {
		businessTransactionItem := domain.BusinessTransactionItem{}
		err := rows.Scan(
			&businessTransactionItem.Id,
			&businessTransactionItem.Name,
			&businessTransactionItem.CreatedAt,
			&businessTransactionItem.UpdatedAt)
		helper.PanicIfError(err)
		businessTransactionItems = append(businessTransactionItems, businessTransactionItem)
	}

	return businessTransactionItems
}
