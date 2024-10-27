package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BusinessCategoryController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
