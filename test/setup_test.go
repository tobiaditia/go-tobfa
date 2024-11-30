package test

import (
	"database/sql"
	"go-tobfa/app"
	"go-tobfa/controller"
	"go-tobfa/helper"
	"go-tobfa/middleware"
	"go-tobfa/repository"
	"go-tobfa/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/tobfa_go_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {

	validate := validator.New()

	controllers := controller.Controller{
		User:                controller.NewUserController(service.NewUserService(repository.NewUserRepository(), db, validate)),
		Location:            controller.NewLocationController(service.NewLocationService(repository.NewLocationRepository(), db, validate)),
		Business:            controller.NewBusinessController(service.NewBusinessService(repository.NewBusinessRepository(), db, validate)),
		BusinessCategory:    controller.NewBusinessCategoryController(service.NewBusinessCategoryService(repository.NewBusinessCategoryRepository(), db, validate)),
		BusinessTransaction: controller.NewBusinessTransactionController(service.NewBusinessTransactionService(repository.NewBusinessTransactionRepository(), repository.NewBusinessCategoryRepository(), db, validate)),
	}

	router := app.NewRouter(controllers)

	return middleware.NewAuthMiddleware(router)
}
