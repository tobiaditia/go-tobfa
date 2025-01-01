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

// @Tags         Business Transaction
// @Summary      Find Business Transaction
// @Description  Find Business Transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse{data=web.BusinessTransactionResponse}
// @Router       /businessTransactions/{id} [get]
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

// @Tags         Business Transaction
// @Summary      Find Business Transaction
// @Description  Find Business Transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Business ID"
// @Success      200  {object}  web.WebResponse{data=[]web.BusinessTransactionResponse}
// @Router       /businessTransactions/{id}/business [get]
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

// @Tags         Stat
// @Summary      Stats Business Transaction
// @Description  Stats Business Transaction
// @Accept       json
// @Produce      json
// @Param        dateStarted   path      string  true  "dateStarted"
// @Param        dateEnded   path      string  true  "dateEnded"
// @Param        businessTransactionTypeId   path      int  true  "businessTransactionTypeId"
// @Param        businessTransactionItemId   path      int  true  "businessTransactionItemId"
// @Param        provinceId   path      int  true  "provinceId"
// @Param        cityId   path      int  true  "cityId"
// @Param        districtId   path      int  true  "districtId"
// @Param        villageId   path      int  true  "villageId"
// @Param        businessCategoryId   path      int  true  "businessCategoryId"
// @Success      200  {object}  web.WebResponse{data=[]web.BusinessTransactionStatsResponse}
// @Router       /stats/businessTransactions [get]
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

// @Tags         Business Transaction
// @Summary      Create Business Transaction
// @Description  Create Business Transaction
// @Accept       json
// @Produce      json
// @Param        body	body		web.BusinessTransactionCreateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.BusinessTransactionResponse}
// @Router       /businessTransactions [post]
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

// @Tags         Business Transaction
// @Summary      Update Business Transaction
// @Description  Update Business Transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Param        body	body		web.BusinessTransactionUpdateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.BusinessTransactionResponse}
// @Router       /businessTransactions/{id} [put]
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

// @Tags         Business Transaction
// @Summary      Delete Business Transaction
// @Description  Delete Business Transaction
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse
// @Router       /businessTransactions/{id} [delete]
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
