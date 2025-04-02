package repository

import (
	"context"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/model"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User)
	Delete(ctx context.Context, userId int)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindAll(ctx context.Context) []model.User
	FindById(ctx context.Context, userId int) (model.User, error)
}