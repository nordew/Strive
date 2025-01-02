package storage

import (
	"context"
	"github.com/nordew/Strive/internal/model"
)

type (
	UserStorage interface {
		Create(ctx context.Context, user *model.User) error
		GetByID(ctx context.Context, id string) (*model.User, error)
		GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error)
		Update(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, id string) error
	}

	GoalStorage interface {
		Create(ctx context.Context, goal *model.Goal) error
		CreateChapter(ctx context.Context, chapter *model.Chapter) error
		CreateComment(ctx context.Context, comment *model.Comment) error
		GetByID(ctx context.Context, id string) (*model.Goal, error)
		GetByUserID(ctx context.Context, userID string) ([]*model.Goal, error)
		GetChapterByID(ctx context.Context, id string) (*model.Chapter, error)
		GetCommentByID(ctx context.Context, id string) (*model.Comment, error)
		Update(ctx context.Context, goal *model.Goal) error
		UpdateChapter(ctx context.Context, chapter *model.Chapter) error
		UpdateComment(ctx context.Context, comment *model.Comment) error
		Delete(ctx context.Context, id string) error
		DeleteChapter(ctx context.Context, id string) error
		DeleteComment(ctx context.Context, id string) error
	}
)
