package middleware

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"net/http"
	"regexp"
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
	// fmt.Println(request.param)
	noAuthRoutes := map[string][]string{
		"POST": {"/api/users"},
		"GET": {
			"/api/businessCategories",
			"/api/provinces",
			"/api/cities/1",
			"/api/districts/1",
			"/api/villages/1",
		},
	}

	/**
	 * Penjelasan loop range.
	 * checkPath berisi masing2 list dari {"/api/users"}
	 * _ adalah index, karena indexnya tidak digunakan, maka diganti dengan _
	 */
	for _, checkPath := range noAuthRoutes[request.Method] {
		if isRouteMatch(checkPath, request.URL.Path) {
			return true
		}
	}

	return false
}

func isRouteMatch(checkPath string, realPath string) bool {
	// Regex to match integers
	re := regexp.MustCompile(`/\d+`)
	realPathRegex := re.ReplaceAllString(realPath, "/1")

	// Match the generated regex with the request path
	// fmt.Println(realPathRegex)
	// fmt.Println(checkPath)
	matched, _ := regexp.MatchString(realPathRegex, checkPath)
	return matched
}
