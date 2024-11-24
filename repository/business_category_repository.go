package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type BusinessCategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BusinessCategory
}
