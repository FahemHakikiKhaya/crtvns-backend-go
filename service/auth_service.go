package service

import (
	"context"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
)

type AuthService interface {
	Login(ctx context.Context, request request.AuthLoginRequest) response.AuthLoginResponse
	Register(ctx context.Context, request request.AuthRegisterRequest) response.AuthRegisterResponse
} 