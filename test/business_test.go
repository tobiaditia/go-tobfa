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

func truncateBusiness(db *sql.DB) {
	db.Exec("TRUNCATE businesses")
}

func TestCreateBusinessSuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"userId": 1,
		"name": "111",
		"address": "1aa",
		"businessCategoryId": 1,
		"countryId": 1,
		"provinceId": 1,
		"cityId": 1,
		"districtId": 1,
		"villageId": 1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/businesses", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	responseBodyData := responseBody["data"]
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "111", responseBodyData.(map[string]interface{})["name"])
}

func TestCreateBusinessErrorValidation(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"userId": 1,
		"name": "1",
		"address": "1aa",
		"businessCategoryId": 1,
		"countryId": 1,
		"provinceId": 1,
		"cityId": 1,
		"districtId": 1,
		"villageId": 1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/businesses", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateBusinessSuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)

	tx, _ := db.Begin()
	buseinessRepository := repository.NewBusinessRepository()
	business := buseinessRepository.Create(context.Background(), tx, domain.Business{
		UserId:             1,
		Name:               "111",
		Address:            "111",
		BusinessCategoryId: 1,
		CountryId:          1,
		ProvinceId:         1,
		CityId:             1,
		DistrictId:         1,
		VillageId:          1,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"userId": 1,
		"name": "1112",
		"address": "1aa",
		"businessCategoryId": 1,
		"countryId": 1,
		"provinceId": 1,
		"cityId": 1,
		"districtId": 1,
		"villageId": 1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/businesses/"+strconv.Itoa(business.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	responseBodyData := responseBody["data"]
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "1112", responseBodyData.(map[string]interface{})["name"])
}

func TestUpdateBusinessErrorValidation(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)

	tx, _ := db.Begin()
	buseinessRepository := repository.NewBusinessRepository()
	business := buseinessRepository.Create(context.Background(), tx, domain.Business{
		UserId:             1,
		Name:               "111",
		Address:            "111",
		BusinessCategoryId: 1,
		CountryId:          1,
		ProvinceId:         1,
		CityId:             1,
		DistrictId:         1,
		VillageId:          1,
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{
		"userId": 1,
		"name": "12",
		"address": "1aad",
		"businessCategoryId": 1,
		"countryId": 1,
		"provinceId": 1,
		"cityId": 1,
		"districtId": 1,
		"villageId": 1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/businesses/"+strconv.Itoa(business.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetBusinessSuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)

	tx, _ := db.Begin()
	buseinessRepository := repository.NewBusinessRepository()
	buseinessRepository.Create(context.Background(), tx, domain.Business{
		UserId:             1,
		Name:               "111",
		Address:            "111",
		BusinessCategoryId: 1,
		CountryId:          1,
		ProvinceId:         1,
		CityId:             1,
		DistrictId:         1,
		VillageId:          1,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/businesses", nil)
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

	responseBodyData := responseBody["data"].([]interface{})
	assert.Equal(t, "111", responseBodyData[0].(map[string]interface{})["name"])
}

func TestFindBusinessSuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)

	tx, _ := db.Begin()
	buseinessRepository := repository.NewBusinessRepository()
	business := buseinessRepository.Create(context.Background(), tx, domain.Business{
		UserId:             1,
		Name:               "111",
		Address:            "111",
		BusinessCategoryId: 1,
		CountryId:          1,
		ProvinceId:         1,
		CityId:             1,
		DistrictId:         1,
		VillageId:          1,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/businesses/"+strconv.Itoa(business.Id), nil)
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

func TestFindBusinessFailed(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/businesses/1", nil)
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

func TestDeleteBusinessSuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)

	tx, _ := db.Begin()
	buseinessRepository := repository.NewBusinessRepository()
	business := buseinessRepository.Create(context.Background(), tx, domain.Business{
		UserId:             1,
		Name:               "111",
		Address:            "111",
		BusinessCategoryId: 1,
		CountryId:          1,
		ProvinceId:         1,
		CityId:             1,
		DistrictId:         1,
		VillageId:          1,
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/businesses/"+strconv.Itoa(business.Id), nil)
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
	assert.Equal(t, nil, responseBodyData)
}

func TestDeleteBusinessFailed(t *testing.T) {
	db := setupTestDB()
	truncateBusiness(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/businesses/1", nil)
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
