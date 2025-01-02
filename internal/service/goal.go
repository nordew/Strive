package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nordew/Strive/internal/dto"
	"github.com/nordew/Strive/internal/model"
	"github.com/nordew/Strive/internal/storage"
	"github.com/nordew/Strive/pkg/logger"
	"time"
)

type goalService struct {
	goalStorage storage.GoalStorage
	logger      logger.Logger
}

func NewGoalService(
	goalStorage storage.GoalStorage,
	logger logger.Logger,
) GoalService {
	return &goalService{
		goalStorage: goalStorage,
		logger:      logger,
	}
}

func (s *goalService) Create(ctx context.Context, createDTO *dto.CreateGoalDTO) error {
	const op = "goalService.Create"

	now := time.Now()

	goalID := uuid.NewString()
	goal, err := model.NewGoal(
		goalID,
		createDTO.UserID,
		createDTO.Title,
		createDTO.Description,
		nil,
		0,
		false,
		createDTO.Deadline,
		0,
		nil,
		nil,
		now,
		now,
	)

	if err != nil {
		s.logger.Errorf("%s: failed to create goal: %v", op, err)
		return fmt.Errorf("failed to create goal: %w", err)
	}

	if err := s.goalStorage.Create(ctx, goal); err != nil {
		s.logger.Errorf("%s: failed to create goal: %v", op, err)
		return fmt.Errorf("failed to create goal: %w", err)
	}

	return nil
}
