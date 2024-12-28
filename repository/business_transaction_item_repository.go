package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type BusinessTransactionItemRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BusinessTransactionItem
}
