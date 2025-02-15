package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationController interface {
	Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
