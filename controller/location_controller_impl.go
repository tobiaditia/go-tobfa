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

// @Tags         Location
// @Summary      Provinces
// @Description  Provinces
// @Accept       json
// @Produce      json
// @Success      200  {object}  web.WebResponse{data=[]web.ProvinceResponse}
// @Router       /provinces [get]
func (controller LocationControllerImpl) Provinces(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responses := controller.LocationService.Provinces(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         Location
// @Summary      Cities
// @Description  Cities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Province ID"
// @Success      200  {object}  web.WebResponse{data=[]web.CityResponse}
// @Router       /cities/{id} [get]
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

// @Tags         Location
// @Summary      Districts
// @Description  Districts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "City ID"
// @Success      200  {object}  web.WebResponse{data=[]web.DistrictResponse}
// @Router       /districts/{id} [get]
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

// @Tags         Location
// @Summary      Villages
// @Description  Villages
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "District ID"
// @Success      200  {object}  web.WebResponse{data=[]web.VillageResponse}
// @Router       /villages/{id} [get]
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

// @Tags         Location
// @Summary      Search
// @Description  Search
// @Accept       json
// @Produce      json
// @Param        search   path      string  true  "Search"
// @Success      200  {object}  web.WebResponse{data=[]web.VillageResponse}
// @Router       /location/search/{search} [get]
func (controller LocationControllerImpl) Search(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	search := params.ByName("search")

	responses := controller.LocationService.Search(request.Context(), search)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
