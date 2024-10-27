package repository

import (
	"context"
	"database/sql"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
)

type BusinessCategoryRepositoryImpl struct {
}

func NewBusinessCategoryRepository() BusinessCategoryRepository {
	return &BusinessCategoryRepositoryImpl{}
}

func (repository *BusinessCategoryRepositoryImpl) Get(ctx context.Context, tx *sql.Tx) []domain.BusinessCategory {
	sql := "select * from business_categories"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var businessCategories []domain.BusinessCategory
	for rows.Next() {
		businessCategory := domain.BusinessCategory{}
		err := rows.Scan(
			&businessCategory.Id,
			&businessCategory.Name,
			&businessCategory.CreatedAt,
			&businessCategory.UpdatedAt)
		helper.PanicIfError(err)
		businessCategories = append(businessCategories, businessCategory)
	}

	return businessCategories
}
