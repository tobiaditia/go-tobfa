package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
	"go-tobfa/model/web"
)

type BusinessTransactionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.BusinessTransaction
	FindById(ctx context.Context, tx *sql.Tx, businessTransactionId int) (domain.BusinessTransaction, error)
	FindByBusiness(ctx context.Context, tx *sql.Tx, businessId int) []domain.BusinessTransaction
	FindForStats(ctx context.Context, tx *sql.Tx, param web.BusinessTransactionStatsGetRequest) []domain.BusinessTransaction
	Create(ctx context.Context, tx *sql.Tx, business domain.BusinessTransaction) domain.BusinessTransaction
	Update(ctx context.Context, tx *sql.Tx, business domain.BusinessTransaction) domain.BusinessTransaction
	Delete(ctx context.Context, tx *sql.Tx, business domain.BusinessTransaction)
}
