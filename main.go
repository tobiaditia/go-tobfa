package main

import (
	"go-tobfa/app"
	"go-tobfa/controller"
	"go-tobfa/helper"
	"go-tobfa/middleware"
	"go-tobfa/repository"
	"go-tobfa/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
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

	businessTransactionRepository := repository.NewBusinessTransactionRepository()
	businessTransactionService := service.NewBusinessTransactionService(businessTransactionRepository, businessCategoryRepository, db, validate)
	businessTransactionController := controller.NewBusinessTransactionController(businessTransactionService)

	router := app.NewRouter(userController, businessController, businessCategoryController, businessTransactionController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
