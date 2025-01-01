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

// @Tags         Business
// @Summary      Create Business
// @Description  Create Business
// @Accept       json
// @Produce      json
// @Param        body	body		web.BusinessCreateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.BusinessResponse}
// @Router       /businesses [post]
func (controller BusinessControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessCreateRequest := web.BusinessCreateRequest{}
	helper.ReadFromRequestBody(request, &businessCreateRequest)

	businessResponse := controller.BusinessService.Create(request.Context(), businessCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

// @Tags         Business
// @Summary      Update Business
// @Description  Update Business
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Param        body	body		web.BusinessUpdateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.BusinessResponse}
// @Router       /businesses/{id} [put]
func (controller BusinessControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessUpdateRequest := web.BusinessUpdateRequest{}
	helper.ReadFromRequestBody(request, &businessUpdateRequest)

	businessId := params.ByName("id")
	id, err := strconv.Atoi(businessId)
	helper.PanicIfError(err)

	businessResponse := controller.BusinessService.Update(request.Context(), id, businessUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         Business
// @Summary      Delete Business
// @Description  Delete Business
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse
// @Router       /businesses/{id} [delete]
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

// @Tags         Business
// @Summary      Find Business
// @Description  Find Business
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse{data=web.BusinessResponse}
// @Router       /businesses/{id} [get]
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

// @Tags         Business
// @Summary      Get Business
// @Description  Get Business
// @Accept       json
// @Produce      json
// @Success      200  {object}  web.WebResponse{data=[]web.BusinessResponse}
// @Router       /businesses [get]
func (controller BusinessControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessResponses := controller.BusinessService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         Stat
// @Summary      Stats Business
// @Description  Stats Business
// @Accept       json
// @Produce      json
// @Param        dateStarted   path      string  true  "dateStarted"
// @Param        dateEnded   path      string  true  "dateEnded"
// @Param        businessIds   path      string  false  "businessIds, sparated by coma"
// @Param        businessTransactionTypeIds   path      string  false  "businessTransactionTypeIds, sparated by coma"
// @Param        businessTransactionItemIds   path      string  false  "businessTransactionItemIds, sparated by coma"
// @Success      200  {object}  web.WebResponse{data=web.BusinessStatsResponse}
// @Router       /stats/business [get]
func (controller BusinessControllerImpl) Stats(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessStatsGetRequest := web.BusinessStatsGetRequest{}
	helper.ReadFromURLQuery(request, &businessStatsGetRequest)

	businessResponses := controller.BusinessService.Stats(request.Context(), businessStatsGetRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
