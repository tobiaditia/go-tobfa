package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func ReadFromURLQuery(request *http.Request, result interface{}) {
	decoder := schema.NewDecoder()
	err := decoder.Decode(result, request.URL.Query())
	PanicIfError(err)

	// Parse JSON body (if necessary)
	// if request.Header.Get("Content-Type") == "application/json" {
	// 	if err := json.NewDecoder(request.Body).Decode(result); err != nil {
	// 		http.Error(writer, err.Error(), http.StatusBadRequest)
	// 		return
	// 	}
	// }
}
