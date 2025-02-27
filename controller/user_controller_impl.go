package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// @Tags         User
// @Summary      Create User
// @Description  Create User
// @Accept       json
// @Produce      json
// @Param        body	body		web.UserCreateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.UserResponse}
// @Failure      400  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Failure      404  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Failure      500  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Router       /users [post]
func (controller UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

// @Tags         User
// @Summary      Update User
// @Description  Update User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Param        body	body		web.UserUpdateRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.UserResponse}
// @Router       /users/{id} [put]
// @Security     BearerAuth
func (controller UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(request.Context(), id, userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         User
// @Summary      Update Password User
// @Description  Update Password User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Param        body	body		web.UserUpdatePasswordRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.UserResponse}
// @Router       /users/{id}/password [put]
// @Security     BearerAuth
func (controller UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdatePasswordRequest := web.UserUpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &userUpdatePasswordRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.UpdatePassword(request.Context(), id, userUpdatePasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         User
// @Summary      Delete User
// @Description  Delete User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse
// @Router       /users/{id} [delete]
// @Security     BearerAuth
func (controller UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Tags         User
// @Summary      Find User
// @Description  Find User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  web.WebResponse{data=web.UserResponse}
// @Router       /users/{id} [get]
// @Security     BearerAuth
func (controller UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
