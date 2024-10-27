package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type BusinessCategoryRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.BusinessCategory
}
