package controller

import (
	"net/http"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/service"
	"github.com/julienschmidt/httprouter"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}


func(controller *AuthController) Register(writter http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	authRegisterRequest := request.AuthRegisterRequest{}
	helper.ReadRequestBody(requests, &authRegisterRequest)

	data := controller.AuthService.Register(requests.Context(), authRegisterRequest)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: data,
	}

	helper.WriteResponseBody(writter, webResponse)
}


func(controller *AuthController) Login(writter http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	authLoginRequest := request.AuthLoginRequest{}
	helper.ReadRequestBody(requests, &authLoginRequest)

	data := controller.AuthService.Login(requests.Context(), authLoginRequest)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: data,
	}

	helper.WriteResponseBody(writter, webResponse)
}