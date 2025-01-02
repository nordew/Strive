package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nordew/Strive/internal/model"
)

const (
	goalsTable    = "goals"
	chaptersTable = "chapters"
	commentsTable = "comments"
)

var (
	ErrGoalNotFound    = fmt.Errorf("goal not found")
	ErrChapterNotFound = fmt.Errorf("chapter not found")
	ErrCommentNotFound = fmt.Errorf("comment not found")
)

type goalStorage struct {
	db *pgxpool.Pool
}

func NewGoalStorage(db *pgxpool.Pool) GoalStorage {
	return &goalStorage{db: db}
}

func (s *goalStorage) Create(ctx context.Context, goal *model.Goal) error {
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", goalsTable)

	_, err := s.db.Exec(ctx, query, goal.ID, goal.UserID, goal.Title, goal.Description, goal.CreatedAt, goal.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create goal: %w", err)
	}

	return nil
}

func (s *goalStorage) CreateChapter(ctx context.Context, chapter *model.Chapter) error {
	query := fmt.Sprintf("INSERT INTO %s (id, goal_id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", chaptersTable)

	_, err := s.db.Exec(ctx, query, chapter.ID, chapter.GoalID, chapter.Title, chapter.Description, chapter.CreatedAt, chapter.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create chapter: %w", err)
	}

	return nil
}

func (s *goalStorage) CreateComment(ctx context.Context, comment *model.Comment) error {
	query := fmt.Sprintf("INSERT INTO %s (id, goal_id, chapter_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", commentsTable)

	_, err := s.db.Exec(ctx, query, comment.ID, comment.GoalID, comment.ChapterID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}

	return nil
}

func (s *goalStorage) GetByID(ctx context.Context, id string) (*model.Goal, error) {
	var goal model.Goal

	query := fmt.Sprintf("SELECT id, user_id, title, description, created_at, updated_at FROM %s WHERE id = $1", goalsTable)

	err := s.db.QueryRow(ctx, query, id).Scan(&goal.ID, &goal.UserID, &goal.Title, &goal.Description, &goal.CreatedAt, &goal.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrGoalNotFound
		}

		return nil, fmt.Errorf("failed to get goal by id: %w", err)
	}

	return &goal, nil
}

func (s *goalStorage) GetByUserID(ctx context.Context, userID string) ([]*model.Goal, error) {
	var goals []*model.Goal

	query := fmt.Sprintf("SELECT id, user_id, title, description, created_at, updated_at FROM %s WHERE user_id = $1", goalsTable)

	rows, err := s.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get goals by user id: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var goal model.Goal

		if err := rows.Scan(&goal.ID, &goal.UserID, &goal.Title, &goal.Description, &goal.CreatedAt, &goal.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan goal: %w", err)
		}

		goals = append(goals, &goal)
	}

	return goals, nil
}

func (s *goalStorage) GetChapterByID(ctx context.Context, id string) (*model.Chapter, error) {
	var chapter model.Chapter

	query := fmt.Sprintf("SELECT id, goal_id, title, description, created_at, updated_at FROM %s WHERE id = $1", chaptersTable)

	err := s.db.QueryRow(ctx, query, id).Scan(&chapter.ID, &chapter.GoalID, &chapter.Title, &chapter.Description, &chapter.CreatedAt, &chapter.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrChapterNotFound
		}

		return nil, fmt.Errorf("failed to get chapter by id: %w", err)
	}

	return &chapter, nil
}

func (s *goalStorage) GetCommentByID(ctx context.Context, id string) (*model.Comment, error) {
	var comment model.Comment

	query := fmt.Sprintf("SELECT id, goal_id, chapter_id, content, created_at, updated_at FROM %s WHERE id = $1", commentsTable)

	err := s.db.QueryRow(ctx, query, id).Scan(&comment.ID, &comment.GoalID, &comment.ChapterID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCommentNotFound
		}

		return nil, fmt.Errorf("failed to get comment by id: %w", err)
	}

	return &comment, nil
}

func (s *goalStorage) Update(ctx context.Context, goal *model.Goal) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, updated_at = $3 WHERE id = $4", goalsTable)

	_, err := s.db.Exec(ctx, query, goal.Title, goal.Description, goal.UpdatedAt, goal.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrGoalNotFound
		}

		return fmt.Errorf("failed to update goal: %w", err)
	}

	return nil
}

func (s *goalStorage) UpdateChapter(ctx context.Context, chapter *model.Chapter) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, updated_at = $3 WHERE id = $4", chaptersTable)

	_, err := s.db.Exec(ctx, query, chapter.Title, chapter.Description, chapter.UpdatedAt, chapter.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrChapterNotFound
		}

		return fmt.Errorf("failed to update chapter: %w", err)
	}

	return nil
}

func (s *goalStorage) UpdateComment(ctx context.Context, comment *model.Comment) error {
	query := fmt.Sprintf("UPDATE %s SET content = $1, updated_at = $2 WHERE id = $3", commentsTable)

	_, err := s.db.Exec(ctx, query, comment.Content, comment.UpdatedAt, comment.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrCommentNotFound
		}

		return fmt.Errorf("failed to update comment: %w", err)
	}

	return nil
}

func (s *goalStorage) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", goalsTable)

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrGoalNotFound
		}

		return fmt.Errorf("failed to delete goal: %w", err)
	}

	return nil
}

func (s *goalStorage) DeleteChapter(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", chaptersTable)

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrChapterNotFound
		}

		return fmt.Errorf("failed to delete chapter: %w", err)
	}

	return nil
}

func (s *goalStorage) DeleteComment(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", commentsTable)

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrCommentNotFound
		}

		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}
