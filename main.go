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

	buseinessRepository := repository.NewBusinessRepository()
	buseinessService := service.NewBusinessService(buseinessRepository, db, validate)
	businessController := controller.NewBusinessController(buseinessService)

	buseinessCategoryRepository := repository.NewBusinessCategoryRepository()
	buseinessCategoryService := service.NewBusinessCategoryService(buseinessCategoryRepository, db, validate)
	buseinessCategoryController := controller.NewBusinessCategoryController(buseinessCategoryService)

	router := app.NewRouter(userController, businessController, buseinessCategoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
