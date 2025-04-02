package controller

import (
	"fmt"
	"net/http"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/middleware"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/service"
	"github.com/julienschmidt/httprouter"
)

type UserHistoryController struct {
	UserHistoryService service.UserHistoryService
}

func NewUserHistoryController(userHistoryService service.UserHistoryService) *UserHistoryController {
	return &UserHistoryController{UserHistoryService: userHistoryService}
}

func (controller *UserHistoryController) GetUserHistories(writter http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	userId := requests.Context().Value(middleware.UserIDKey).(int)

	request := request.GetUserHistories{UserId: userId}

	data, err := controller.UserHistoryService.GetUserHistories(requests.Context(), request)

	fmt.Println(data, err)

	if err != nil {
		helper.WriteResponseBody(writter, response.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error fetching user histories",
			Data:   nil,
		})
		return
	}

	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: data,
	}

	helper.WriteResponseBody(writter, webResponse)
}

func (controller *UserHistoryController) CreateUserHistory(writter http.ResponseWriter, requests *http.Request, _ httprouter.Params) {
	userId := requests.Context().Value(middleware.UserIDKey).(int)

	createUserHistoryRequest := request.CreateUserHistoryRequest{}
	helper.ReadRequestBody(requests, &createUserHistoryRequest)
	createUserHistoryRequest.UserId = userId

	controller.UserHistoryService.CreateUserHistory(requests.Context(), createUserHistoryRequest)

	webResponse := response.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: nil,
	}

	helper.WriteResponseBody(writter, webResponse)
}