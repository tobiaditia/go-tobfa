package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	fmt.Println("from request body")
	fmt.Println(request.Body)
	decoder := json.NewDecoder(request.Body)
	fmt.Println(decoder)
	err := decoder.Decode(result)
	fmt.Println(result)
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
