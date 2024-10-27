package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func truncateBusinessCategory(db *sql.DB) {
	db.Exec("TRUNCATE business_categories")
}

func TestGetBusinessCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateBusinessCategory(db)

	// tx, _ := db.Begin()
	// buseinessCategoryRepository := repository.NewBusinessCategoryRepository()
	// buseinessCategoryRepository.Create(context.Background(), tx, domain.BusinessCategory{
	// 	Name:                       "111",
	// })
	// tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/businessCategories", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "sosecret")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)

	// body, _ := io.ReadAll(response.Body)
	// var responseBody map[string]interface{}
	// json.Unmarshal(body, &responseBody)

	// assert.Equal(t, "OK", responseBody["status"])

	// responseBodyData := responseBody["data"].([]interface{})
	// assert.Equal(t, "111", responseBodyData[0].(map[string]interface{})["name"])
}
