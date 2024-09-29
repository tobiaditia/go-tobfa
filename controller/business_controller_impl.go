package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BusinessControllerImpl struct {
	BusinessService service.BusinessService
}

func NewBusinessController(businessService service.BusinessService) BusinessController {
	return &BusinessControllerImpl{
		BusinessService: businessService,
	}
}

func (controller BusinessControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessCreateRequest := web.BusininessCreateRequest{}
	helper.ReadFromRequestBody(request, &businessCreateRequest)

	businessResponse := controller.BusinessService.Create(request.Context(), businessCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller BusinessControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessUpdateRequest := web.BusininessUpdateRequest{}
	helper.ReadFromRequestBody(request, &businessUpdateRequest)

	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)
	businessUpdateRequest.Id = id

	businessResponse := controller.BusinessService.Update(request.Context(), businessUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)

	controller.BusinessService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)

	businessResponse := controller.BusinessService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller BusinessControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessResponses := controller.BusinessService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
