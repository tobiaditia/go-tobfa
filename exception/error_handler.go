package exception

import (
	"database/sql"
	"fmt"
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	// panic(err)
	// fmt.Println(err)
	if notFoundError(writer, err) {
		return
	}

	if alreadyExistError(writer, err) {
		return
	}

	if validationError(writer, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors) //cek apakah err nya bertipe ValidationErrors, ValidationErrors adalah bawaan dari package validator, jadi tidak perlu buat func baru
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   reformatErrMsg(exception),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func reformatErrMsg(validationErrors validator.ValidationErrors) []string {
	var errors []string
	for _, err := range validationErrors {
		var message string
		field := err.Field() // The struct field name
		tag := err.Tag()     // The validation tag (e.g., "required", "min", "max")
		param := err.Param() // Additional parameter for the tag (e.g., min=3)

		switch tag {
		case "required":
			message = fmt.Sprintf("Field '%s' is required", field)
		case "min":
			message = fmt.Sprintf("Field '%s' must have at least %s characters", field, param)
		case "max":
			message = fmt.Sprintf("Field '%s' must have at most %s characters", field, param)
		default:
			message = fmt.Sprintf("Field '%s' failed on the '%s' tag", field, tag)
		}

		errors = append(errors, message)
	}
	return errors
}

func alreadyExistError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(AlreadyExistsError) //cek apakah err nya bertipe ValidationErrors, ValidationErrors adalah bawaan dari package validator, jadi tidak perlu buat func baru
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, err interface{}) bool {
	exception, ok := err.(NotFoundError)  //cek apakah err nya bertipe NotFoundError
	errorSqlNoRow := err == sql.ErrNoRows //cek apakah err nya bertipe ErrNoRows, ErrNoRows adalah kembalian dari QueryRowContext jika data kosong
	if ok || errorSqlNoRow {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	logger := logrus.New()
	file, _ := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.WithFields(logrus.Fields{
		"path": request.RequestURI,
	}).Error(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
