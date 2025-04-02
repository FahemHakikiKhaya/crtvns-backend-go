package service

import (
	"context"
	"fmt"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/request"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/repository"
)

type UserHistoryServiceImpl struct {
	UserHistoryRepository repository.UserHistoryRepository
}

// FindAll implements UserHistoryService.
func (u *UserHistoryServiceImpl) GetUserHistories(ctx context.Context, request request.GetUserHistories) ([]response.GetUserHistoriesResponse, error) {
	userHistories, err :=  u.UserHistoryRepository.FindAll(ctx, request.UserId)

	return userHistories, err
}

// Save implements UserHistoryService.
func (u *UserHistoryServiceImpl) CreateUserHistory(ctx context.Context, request request.CreateUserHistoryRequest) {
	fmt.Println(request.UserId)
	userHistory := model.UserHistory{
		UserId: request.UserId,
		Text: request.Text,
		Rate: request.Rate,
		Pitch: request.Pitch,
		Volume: request.Volume,
		VoiceName: request.VoiceName,
	}

	u.UserHistoryRepository.Save(ctx, userHistory)
}

func NewUserHistoryService(userHistoryRepository repository.UserHistoryRepository) UserHistoryService {
	return &UserHistoryServiceImpl{UserHistoryRepository: userHistoryRepository}
}
