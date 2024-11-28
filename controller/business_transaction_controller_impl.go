package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BusinessTransactionControllerImpl struct {
	BusinessTransactionService service.BusinessTransactionService
}

func NewBusinessTransactionController(businessService service.BusinessTransactionService) BusinessTransactionController {
	return &BusinessTransactionControllerImpl{
		BusinessTransactionService: businessService,
	}
}

func (controller BusinessTransactionControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessResponses := controller.BusinessTransactionService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessTransactionControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessTransactionId := params.ByName("id")
	id, err := strconv.Atoi(businessTransactionId)
	helper.PanicIfError(err)

	businessTransactionResponse := controller.BusinessTransactionService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessTransactionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessTransactionControllerImpl) FindByBusiness(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)

	businessResponse := controller.BusinessTransactionService.FindByBusiness(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessTransactionControllerImpl) Stats(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessTransactionStatGetRequest := web.BusinessTransactionStatsGetRequest{}
	helper.ReadFromURLQuery(request, &businessTransactionStatGetRequest)

	businessResponses := controller.BusinessTransactionService.Stats(request.Context(), businessTransactionStatGetRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessTransactionControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessTransactionCreateRequest := web.BusinessTransactionCreateRequest{}
	helper.ReadFromRequestBody(request, &businessTransactionCreateRequest)

	businessResponse := controller.BusinessTransactionService.Create(request.Context(), businessTransactionCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller BusinessTransactionControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessTransactionUpdateRequest := web.BusinessTransactionUpdateRequest{}
	helper.ReadFromRequestBody(request, &businessTransactionUpdateRequest)

	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)
	businessTransactionUpdateRequest.Id = id

	businessResponse := controller.BusinessTransactionService.Update(request.Context(), businessTransactionUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessTransactionControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)

	controller.BusinessTransactionService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
