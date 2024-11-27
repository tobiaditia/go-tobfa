package middleware

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if middleware.withoutAuth(request) {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	if request.Header.Get("X-API-Key") == "sosecret" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}

func (middleware *AuthMiddleware) withoutAuth(request *http.Request) bool {
	noAuthRoutes := map[string][]string{
		"POST": {"/api/users"},
		"GET":  {"/api/businessCategories"},
	}

	/**
	 * Penjelasan loop range.
	 * path berisi masing2 list dari {"/api/users"}
	 * _ adalah index, karena indexnya tidak digunakan, maka diganti dengan _
	 */
	for _, path := range noAuthRoutes[request.Method] {
		if request.URL.Path == path {
			return true
		}
	}

	return false
}
