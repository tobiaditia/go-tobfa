package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-tobfa/model/domain"
	"go-tobfa/repository"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "111",
		"email": "tobiaditia549@gmail.com",
		"password": "12345678",
		"handphone": "085895402090"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody)

	responseBodyData := responseBody["data"]
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "111", responseBodyData.(map[string]interface{})["name"])
}

func TestCreateUserErrorValidation(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "111",
		"email": "tobiaditia549gmail.com",
		"password": "12345678",
		"handphone": "085895402090"
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody)

	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:      "111",
		Email:     "tobiaditia549@gmail.com",
		Password:  "12345678",
		Handphone: "085895402090",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "1112",
		"handphone": "0858954020902"
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody)

	responseBodyData := responseBody["data"]
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "1112", responseBodyData.(map[string]interface{})["name"])
}

func TestUpdateUserErrorValidation(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:      "111",
		Email:     "tobiaditia549@gmail.com",
		Password:  "12345678",
		Handphone: "085895402090",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"name": "1",
		"handphone": "0858954020902"
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	// fmt.Println(responseBody)

	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestFindUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:      "111",
		Email:     "tobiaditia549@gmail.com",
		Password:  "12345678",
		Handphone: "085895402090",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])

	responseBodyData := responseBody["data"]
	assert.Equal(t, "111", responseBodyData.(map[string]interface{})["name"])
}

func TestFindUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/1", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

// func TestDeleteUserSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateUser(db)

// 	tx, _ := db.Begin()
// 	userRepository := repository.NewUserRepository()
// 	user := userRepository.Create(context.Background(), tx, domain.User{
// 		UserId:             1,
// 		Name:               "111",
// 		Address:            "111",
// 		UserCategoryId: 1,
// 		CountryId:          1,
// 		ProvinceId:         1,
// 		CityId:             1,
// 		DistrictId:         1,
// 		VillageId:          1,
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "sosecret")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()

// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	fmt.Println(responseBody)

// 	assert.Equal(t, "OK", responseBody["status"])

// 	responseBodyData := responseBody["data"]
// 	assert.Equal(t, nil, responseBodyData)
// }

// func TestDeleteUserFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateUser(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/1", nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "sosecret")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()

// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, "NOT FOUND", responseBody["status"])
// }
