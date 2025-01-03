package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nordew/Strive/internal/dto"
	"github.com/nordew/Strive/internal/model"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/auth"
	"github.com/nordew/Strive/pkg/logger"
	"time"
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

func (s *userService) Login(ctx context.Context, loginDTO *dto.LoginUserDTO) (*AuthResponse, error) {
	const op = "userService.Login"

	user, err := s.userStorage.GetByTelegramID(ctx, loginDTO.TelegramID)
	if err != nil {
		if errors.Is(err, storage.ErrorUserNotFound) {
			now := time.Now()

			user, err := model.NewUser(uuid.NewString(), loginDTO.TelegramID, "", "", 0, now, now)
			if err != nil {
				s.logger.Errorf("[%s] failed to create new user: %v", op, err)
				return nil, fmt.Errorf("failed to create new user: %w", err)
			}

			if err := s.userStorage.Create(ctx, user); err != nil {
				s.logger.Errorf("[%s] failed to create new user: %v", op, err)
				return nil, fmt.Errorf("failed to create new user: %w", err)
			}
		} else {
			s.logger.Errorf("[%s] failed to get user by bots id: %v", op, err)
			return nil, fmt.Errorf("failed to get user by bots id: %w", err)
		}
	}

	accessToken, refreshToken, err := s.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.GetID(),
		Role:   user.GetRole(),
	})
	if err != nil {
		s.logger.Errorf("[%s] failed to generate tokens: %v", op, err)
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsAuthorized: user.IsAuthorized,
	}, nil
}

func (s *userService) Authorize(ctx context.Context, telegramID int64, authDTO *dto.AuthorizeUserRequest) error {
	const op = "userService.Authorize"

	user, err := s.userStorage.GetByTelegramID(ctx, telegramID)
	if err != nil {
		if errors.Is(err, storage.ErrorUserNotFound) {
			s.logger.Infof("[%s] user not found by bots id: %d", op, telegramID)
			return err
		}

		s.logger.Errorf("[%s] failed to get user by bots id: %v", op, err)
		return fmt.Errorf("failed to get user by bots id: %w", err)
	}

	_, err = user.SetFirstName(authDTO.FirstName)
	if err != nil {
		s.logger.Errorf("[%s] failed to set first name: %v", op, err)
		return fmt.Errorf("failed to set first name: %w", err)
	}

	_, err = user.SetLastName(authDTO.LastName)
	if err != nil {
		s.logger.Errorf("[%s] failed to set last name: %v", op, err)
		return fmt.Errorf("failed to set last name: %w", err)
	}

	_, err = user.SetIsAuthorized(true)
	if err != nil {
		s.logger.Errorf("[%s] failed to set is authorized: %v", op, err)
		return fmt.Errorf("failed to set is authorized: %w", err)
	}

	_, err = user.SetUpdatedAt(time.Now())
	if err != nil {
		s.logger.Errorf("[%s] failed to set updated at: %v", op, err)
		return fmt.Errorf("failed to set updated at: %w", err)
	}

	if err := s.userStorage.Update(ctx, user); err != nil {
		s.logger.Errorf("[%s] failed to update user: %v", op, err)
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
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
