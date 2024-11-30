package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LocationControllerImpl struct {
	LocationService service.LocationService
}

func NewLocationController(locationService service.LocationService) LocationController {
	return &LocationControllerImpl{
		LocationService: locationService,
	}
}

func (controller LocationControllerImpl) Provinces(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responses := controller.LocationService.Provinces(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller LocationControllerImpl) Cities(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	provinceId := params.ByName("id")
	id, err := strconv.Atoi(provinceId)
	helper.PanicIfError(err)
	responses := controller.LocationService.Cities(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller LocationControllerImpl) Districts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cityId := params.ByName("id")
	id, err := strconv.Atoi(cityId)
	helper.PanicIfError(err)
	responses := controller.LocationService.Districts(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller LocationControllerImpl) Villages(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	districtId := params.ByName("id")
	id, err := strconv.Atoi(districtId)
	helper.PanicIfError(err)
	responses := controller.LocationService.Villages(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller LocationControllerImpl) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	search := params.ByName("searh")

	responses := controller.LocationService.Search(request.Context(), search)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
