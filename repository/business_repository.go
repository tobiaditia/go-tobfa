package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type BusinessRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Business
	Find(ctx context.Context, tx *sql.Tx, businessId int) (domain.Business, error)
	Create(ctx context.Context, tx *sql.Tx, business domain.Business) domain.Business
	Update(ctx context.Context, tx *sql.Tx, business domain.Business) domain.Business
	Delete(ctx context.Context, tx *sql.Tx, business domain.Business)
}
