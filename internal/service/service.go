package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/nordew/Strive/internal/model"
)

var (
	ErrValidation = errors.New("validation error")
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *userService) handleError(op string, msg string, err error, telegramID int64) error {
	s.logger.Errorf("[%s]: %s: telegramID=%d, err=%v", op, msg, telegramID, err)
	return fmt.Errorf("[%s]: %w", op, err)
}

type UserService interface {
	Login(ctx context.Context, telegramID int64) (*AuthResponse, error)

	// Get supports id and telegramID
	Get(ctx context.Context, id int) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int) error
}
