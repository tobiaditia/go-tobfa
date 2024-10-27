package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BusinessCategoryControllerImpl struct {
	BusinessCategoryService service.BusinessCategoryService
}

func NewBusinessCategoryController(businessCategoryService service.BusinessCategoryService) BusinessCategoryController {
	return &BusinessCategoryControllerImpl{
		BusinessCategoryService: businessCategoryService,
	}
}

func (controller BusinessCategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessCategoryResponses := controller.BusinessCategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessCategoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
