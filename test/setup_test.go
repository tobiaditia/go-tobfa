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

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	businessRepository := repository.NewBusinessRepository()
	businessService := service.NewBusinessService(businessRepository, db, validate)
	businessController := controller.NewBusinessController(businessService)

	businessCategoryRepository := repository.NewBusinessCategoryRepository()
	businessCategoryService := service.NewBusinessCategoryService(businessCategoryRepository, db, validate)
	businessCategoryController := controller.NewBusinessCategoryController(businessCategoryService)

	router := app.NewRouter(userController, businessController, businessCategoryController)

	return middleware.NewAuthMiddleware(router)
}
