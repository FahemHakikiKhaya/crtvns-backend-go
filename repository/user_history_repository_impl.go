package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/data/response"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
)

type UserHistoryRepositoryImpl struct {
	DB *sql.DB
}

func (u *UserHistoryRepositoryImpl) FindAll(ctx context.Context, userId int) ([]response.GetUserHistoriesResponse, error) {
	SQL := "SELECT id, userId, text, rate, pitch, volume, voiceName, created_at FROM user_histories WHERE userId = $1"

	rows, err := u.DB.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userHistories []response.GetUserHistoriesResponse

	for rows.Next() {
		var userHistory response.GetUserHistoriesResponse
		err := rows.Scan(&userHistory.Id, &userHistory.UserId, &userHistory.Text, &userHistory.Rate, &userHistory.Pitch, &userHistory.Volume, &userHistory.VoiceName, &userHistory.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		userHistories = append(userHistories, userHistory)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userHistories, nil
}

func (u *UserHistoryRepositoryImpl) Save(ctx context.Context, userHistory model.UserHistory) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into user_histories(userId, text, rate, pitch, volume, voiceName) values($1, $2, $3, $4, $5, $6)"
	tx.ExecContext(ctx, SQL, userHistory.UserId, userHistory.Text, userHistory.Rate, userHistory.Pitch, userHistory.Volume, userHistory.VoiceName)
}

func NewUserHistoryRepository(Db *sql.DB) UserHistoryRepository {
	return &UserHistoryRepositoryImpl{DB: Db}
}
