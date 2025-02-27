package app

import (
	"go-tobfa/controller"
	"go-tobfa/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(controllers controller.Controller) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/authentication/login", controllers.Authentication.Login)

	router.POST("/api/users", controllers.User.Create)
	router.PUT("/api/users/:id", controllers.User.Update)
	router.PUT("/api/users/:id/password", controllers.User.UpdatePassword)
	router.GET("/api/users/:id", controllers.User.FindById)
	router.DELETE("/api/users/:id", controllers.User.Delete)

	router.GET("/api/provinces", controllers.Location.Provinces)
	router.GET("/api/cities/:id", controllers.Location.Cities)
	router.GET("/api/districts/:id", controllers.Location.Districts)
	router.GET("/api/villages/:id", controllers.Location.Villages)
	router.GET("/api/location/search/:search", controllers.Location.Villages)

	router.GET("/api/businesses", controllers.Business.FindAll)
	router.GET("/api/businesses/:id", controllers.Business.FindById)
	router.POST("/api/businesses", controllers.Business.Create)
	router.PUT("/api/businesses/:id", controllers.Business.Update)
	router.DELETE("/api/businesses/:id", controllers.Business.Delete)

	router.GET("/api/businessCategories", controllers.BusinessCategory.FindAll)

	router.GET("/api/businessTransactions", controllers.BusinessTransaction.FindAll)
	router.GET("/api/businessTransactions/:id", controllers.BusinessTransaction.FindById)
	router.GET("/api/businessTransactions/:id/business", controllers.BusinessTransaction.FindByBusiness)
	router.POST("/api/businessTransactions", controllers.BusinessTransaction.Create)
	router.PUT("/api/businessTransactions/:id", controllers.BusinessTransaction.Update)
	router.DELETE("/api/businessTransactions/:id", controllers.BusinessTransaction.Delete)
	router.GET("/api/stats/businessTransactions", controllers.BusinessTransaction.Stats)

	router.GET("/api/stats/business", controllers.Business.Stats)

	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	router.PanicHandler = exception.ErrorHandler

	return router
}
