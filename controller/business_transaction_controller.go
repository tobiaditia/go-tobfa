package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BusinessTransactionController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindByBusiness(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Stats(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
