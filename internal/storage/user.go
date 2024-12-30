package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/nordew/Strive/internal/model"
)

const (
	usersTable = "users"
)

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserExists   = errors.New("user already exists")
)

type userStorage struct {
	db *pgx.Conn
}

func NewUserStorage(db *pgx.Conn) UserStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Create(ctx context.Context, user *model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (telegram_id, first_name, last_name) VALUES ($1, $2, $3) RETURNING id", usersTable)
	err := s.db.QueryRow(ctx, query, user.TelegramID, user.FirstName, user.LastName).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *userStorage) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := fmt.Sprintf("SELECT id, telegram_id, first_name, last_name, created_at, updated_at FROM %s WHERE id = $1", usersTable)
	row := s.db.QueryRow(ctx, query, id)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.TelegramID, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrorUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *userStorage) GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error) {
	query := fmt.Sprintf("SELECT id, telegram_id, first_name, last_name, created_at, updated_at FROM %s WHERE telegram_id = $1", usersTable)
	row := s.db.QueryRow(ctx, query, telegramID)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.TelegramID, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrorUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *userStorage) Update(ctx context.Context, user *model.User) error {
	query := fmt.Sprintf("UPDATE %s SET first_name = $1, last_name = $2, updated_at = now() WHERE id = $3", usersTable)
	_, err := s.db.Exec(ctx, query, user.FirstName, user.LastName, user.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrorUserExists
		}

		return err
	}

	return nil
}

func (s *userStorage) Delete(ctx context.Context, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrorUserExists
		}

		return err
	}

	return nil
}
