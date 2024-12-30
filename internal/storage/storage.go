package storage

import (
	"context"
	"github.com/nordew/Strive/internal/model"
)

type UserStorage interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
