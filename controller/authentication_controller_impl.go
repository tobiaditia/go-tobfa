package controller

import (
	"go-tobfa/helper"
	"go-tobfa/model/web"
	"go-tobfa/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationControllerImpl struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(authenticationService service.AuthenticationService) AuthenticationController {
	return &AuthenticationControllerImpl{
		AuthenticationService: authenticationService,
	}
}

// @Tags         Authentication
// @Summary      Login
// @Description  Login
// @Accept       json
// @Produce      json
// @Param        body	body		web.LoginRequest	true	"Body"
// @Success      200  {object}  web.WebResponse{data=web.UserResponse}
// @Failure      400  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Failure      404  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Failure      500  {object}  web.WebResponse{code=integer,status=string,data=[]string}
// @Router       /authentication/login [post]
func (controller AuthenticationControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	userResponse := controller.AuthenticationService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
