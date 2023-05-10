package repository

import (
	"context"
	"vk-intern-bot/internal/models"
)

type UsersRepo interface {
	Set(ctx context.Context, input *models.SetUserDataInput) error
	Get(ctx context.Context, input *models.GetUserDataInput) (*models.UserData, error)
	Delete(ctx context.Context, input *models.DeleteUserDataInput) error
}
