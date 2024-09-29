package main

import (
	"go-tobfa/app"
	"go-tobfa/controller"
	"go-tobfa/helper"
	"go-tobfa/repository"
	"go-tobfa/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	buseinessRepository := repository.NewBusinessRepository()
	buseinessService := service.NewBusinessService(buseinessRepository, db, validate)
	businessController := controller.NewBusinessController(buseinessService)

	router := httprouter.New()
	router.GET("/api/businesses", businessController.FindAll)
	router.GET("/api/businesses/:id", businessController.FindById)
	router.POST("/api/businesses", businessController.Create)
	router.PUT("/api/businesses/:id", businessController.Update)
	router.DELETE("/api/businesses/:id", businessController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
