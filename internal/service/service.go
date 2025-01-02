package service

import (
	"context"
	"errors"
	"github.com/nordew/Strive/internal/dto"
	"github.com/nordew/Strive/internal/model"
)

var (
	ErrValidation = errors.New("validation error")
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IsAuthorized bool   `json:"is_authorized"`
}

type (
	UserService interface {
		Login(ctx context.Context, loginDTO *dto.LoginUserDTO) (*AuthResponse, error)
		Authorize(ctx context.Context, telegramID int64, authDTO *dto.AuthorizeUserRequest) error

		// Get supports id and telegramID
		Get(ctx context.Context, id int) (*model.User, error)
		Update(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, id int) error
	}

	GoalService interface {
		Create(ctx context.Context, createDTO *dto.CreateGoalDTO) error
	}
)
