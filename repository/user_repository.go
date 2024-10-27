package repository

import (
	"context"
	"database/sql"
	"go-tobfa/model/domain"
)

type UserRepository interface {
	FindById(ctx context.Context, tx *sql.Tx, businessId int) (domain.User, error)
	Create(ctx context.Context, tx *sql.Tx, business domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, business domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, business domain.User)
}
