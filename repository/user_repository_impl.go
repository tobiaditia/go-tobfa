package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-tobfa/helper"
	"go-tobfa/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	sql := "select * from users where id = ?"
	user := domain.User{}
	err := tx.QueryRowContext(ctx, sql, userId).Scan(&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Handphone,
		&user.CreatedAt,
		&user.UpdatedAt)

	helper.PanicIfError(err)
	return user, err
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	sql := "select * from users where email = ?"
	rows, err := tx.QueryContext(ctx, sql, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Handphone,
			&user.CreatedAt,
			&user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("not found data")
	}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "insert into users (name, email, password, handphone, created_at, updated_at) values (?,?,?,?, NOW(), NOW())"
	result, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.Handphone)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "update users set name = ?, handphone = ?, updated_at = NOW() where id = ?"
	_, err := tx.ExecContext(ctx, sql, user.Name, user.Handphone, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	sql := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, sql, user.Id)
	helper.PanicIfError(err)
}
