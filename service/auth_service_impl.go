package service

import (
	"context"
	"fmt"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/repository"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(ctx context.Context, request request.AuthRegisterRequest) response.AuthRegisterResponse {
	existingUser, err := a.UserRepository.FindByEmail(ctx, request.Email)

	helper.PanicIfError(err)

	if (existingUser != model.User{}) {
		return response.AuthRegisterResponse{}
	}

	hashedPassword, err := helper.HashedPassword(request.Password)

	helper.PanicIfError(err)

	user := model.User{
		Name: request.Name,
		Email: request.Email,
		Password: hashedPassword,
	}

	result, err  := a.UserRepository.Save(ctx, user)

	helper.PanicIfError(err)
	
	return response.AuthRegisterResponse{Id: result.Id,Name: result.Name, Email: result.Email}

}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(ctx context.Context, request request.AuthLoginRequest) response.AuthLoginResponse {
	existingUser, err := a.UserRepository.FindByEmail(ctx, request.Email)
	
	helper.PanicIfError(err)

	passwordMatch := helper.ComparePasswords(request.Password, existingUser.Password)

	fmt.Println(request.Password,existingUser.Password,request.Email)

	if !passwordMatch {
		return response.AuthLoginResponse{}
	}

	generateToken, err := helper.GenerateJWT(existingUser.Id)

	helper.PanicIfError(err)

	userLoginResponse := response.AuthLoginResponse{
		Id: existingUser.Id,
		Name: existingUser.Name,
		Email: existingUser.Email,
		Token: generateToken,
	}

	return userLoginResponse
}


func NewAuthServiceImpl(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepository: userRepository}
}
