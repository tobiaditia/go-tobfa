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

	controllers := controller.Controller{
		User:                controller.NewUserController(service.NewUserService(repository.NewUserRepository(), db, validate)),
		Location:            controller.NewLocationController(service.NewLocationService(repository.NewLocationRepository(), db, validate)),
		Business:            controller.NewBusinessController(service.NewBusinessService(repository.NewBusinessRepository(), repository.NewBusinessTransactionItemRepository(), db, validate)),
		BusinessCategory:    controller.NewBusinessCategoryController(service.NewBusinessCategoryService(repository.NewBusinessCategoryRepository(), db, validate)),
		BusinessTransaction: controller.NewBusinessTransactionController(service.NewBusinessTransactionService(repository.NewBusinessTransactionRepository(), repository.NewBusinessCategoryRepository(), db, validate)),
	}

	router := app.NewRouter(controllers)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
