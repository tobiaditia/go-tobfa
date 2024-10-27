package app

import (
	"go-tobfa/controller"
	"go-tobfa/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, businessController controller.BusinessController, businessCategoryController controller.BusinessCategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:id", userController.Update)
	router.GET("/api/users/:id", userController.FindById)

	router.GET("/api/businesses", businessController.FindAll)
	router.GET("/api/businesses/:id", businessController.FindById)
	router.POST("/api/businesses", businessController.Create)
	router.PUT("/api/businesses/:id", businessController.Update)
	router.DELETE("/api/businesses/:id", businessController.Delete)

	router.GET("/api/businessCategories", businessCategoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
