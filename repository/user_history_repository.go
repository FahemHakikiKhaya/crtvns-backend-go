package repository

import (
	"context"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
)

type UserHistoryRepository interface {
	FindAll(ctx context.Context, userId int) ([]response.GetUserHistoriesResponse, error)
	Save(ctx context.Context, userHistory model.UserHistory)
}