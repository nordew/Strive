package service

import (
	"context"
	"errors"
	"github.com/nordew/Strive/internal/model"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/auth"
	"github.com/nordew/Strive/pkg/logger"
)

type userService struct {
	userStorage storage.UserStorage
	auth        auth.Authenticator
	logger      logger.Logger
}

// NewUserService creates a new UserService instance.
func NewUserService(userStorage storage.UserStorage, auth auth.Authenticator, logger logger.Logger) UserService {
	if userStorage == nil {
		panic("userStorage cannot be nil")
	}
	if auth == nil {
		panic("auth cannot be nil")
	}
	if logger == nil {
		panic("logger cannot be nil")
	}

	return &userService{
		userStorage: userStorage,
		auth:        auth,
		logger:      logger,
	}
}

func (s *userService) Login(ctx context.Context, telegramID int64) (*AuthResponse, error) {
	const op = "userService.Login"

	user, err := s.userStorage.GetByTelegramID(ctx, telegramID)
	if err != nil {
		if errors.Is(err, storage.ErrorUserNotFound) {
			user = &model.User{
				TelegramID: telegramID,
			}
			if err := s.userStorage.Create(ctx, user); err != nil {
				return nil, s.handleError(op, "failed to create user", err, telegramID)
			}
		} else {
			return nil, s.handleError(op, "failed to get user by telegramID", err, telegramID)
		}
	}

	accessToken, refreshToken, err := s.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.ID,
		Role:   user.Role,
	})
	if err != nil {
		return nil, s.handleError(op, "failed to generate tokens", err, telegramID)
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshTokens returns new access and refresh tokens
func (s *userService) RefreshTokens(ctx context.Context, refreshToken string) (*AuthResponse, error) {
	return nil, nil
}

// Get returns user by id or telegramID
func (s *userService) Get(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

// Update updates user data
func (s *userService) Update(ctx context.Context, user *model.User) error {
	return nil
}

// Delete deletes user from the system
func (s *userService) Delete(ctx context.Context, id int) error {
	return nil
}
