package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LocationController interface {
	Provinces(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Cities(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Districts(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Villages(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Search(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
