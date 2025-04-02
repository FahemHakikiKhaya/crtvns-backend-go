package service

import (
	"context"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
)

type UserHistoryService interface {
	CreateUserHistory(ctx context.Context, request request.CreateUserHistoryRequest) 
	GetUserHistories(ctx context.Context, request request.GetUserHistories) ([]response.GetUserHistoriesResponse, error)
} 