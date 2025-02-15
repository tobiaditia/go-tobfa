package middleware

import (
	"errors"
	"fmt"
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"net/http"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type AuthMiddleware struct {
	Handler http.Handler
	Config  *viper.Viper
}

func NewAuthMiddleware(handler http.Handler, config *viper.Viper) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, Config: config}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if middleware.withoutAuth(request) {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	tokenString := extractToken(request)

	if tokenString == "" {
		respondUnauthorized(writer)
		return
	}

	valid, err := validateToken(tokenString, middleware.Config)
	if err != nil || !valid {
		respondUnauthorized(writer)
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}

func (middleware *AuthMiddleware) withoutAuth(request *http.Request) bool {
	// Jika URL mengandung "swagger", maka bypass autentikasi
	if strings.Contains(request.URL.Path, "swagger") {
		return true
	}

	// fmt.Println(request.param)
	noAuthRoutes := map[string][]string{
		"POST": {
			"/api/authentication/login",
			"/api/users",
		},
		"GET": {
			"/swagger/*",
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
	matched, _ := regexp.MatchString(realPathRegex, checkPath)
	return matched
}

func extractToken(request *http.Request) string {
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}

func validateToken(tokenString string, config *viper.Viper) (bool, error) {
	secretKey := config.GetString("SECRET_KEY")
	if secretKey == "" {
		return false, errors.New("JWT secret key not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}

func respondUnauthorized(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
