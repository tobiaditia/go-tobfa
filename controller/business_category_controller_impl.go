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

// @Tags         Business Category
// @Summary      Get Business Category
// @Description  Get Business Category
// @Accept       json
// @Produce      json
// @Success      200  {object}  web.WebResponse{data=[]web.BusinessCategoryResponse}
// @Router       /businessCategories [get]
func (controller BusinessCategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	businessCategoryResponses := controller.BusinessCategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   businessCategoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
