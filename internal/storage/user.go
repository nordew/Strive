package storage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nordew/Strive/internal/model"
)

const usersTable = "users"

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserExists   = errors.New("user already exists")
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) UserStorage {
	return &userStorage{db: db}
}

func (s *userStorage) Create(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (id, telegram_id, first_name, last_name, role, is_authorized) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := s.db.QueryRow(ctx, query, user.ID, user.TelegramID, user.FirstName, user.LastName, user.Role, user.IsAuthorized).Scan(&user.ID)
	return err
}

func (s *userStorage) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := "SELECT id, telegram_id, first_name, last_name, role, is_authorized, created_at, updated_at FROM users WHERE id = $1"
	return scanUser(s.db.QueryRow(ctx, query, id))
}

func (s *userStorage) GetByTelegramID(ctx context.Context, telegramID int64) (*model.User, error) {
	query := "SELECT id, telegram_id, first_name, last_name, role, is_authorized, created_at, updated_at FROM users WHERE telegram_id = $1"
	return scanUser(s.db.QueryRow(ctx, query, telegramID))
}

func (s *userStorage) Update(ctx context.Context, user *model.User) error {
	query := "UPDATE users SET first_name = $1, last_name = $2, role = $3, is_authorized = $4, updated_at = now() WHERE id = $5"
	result, err := s.db.Exec(ctx, query, user.FirstName, user.LastName, user.Role, user.IsAuthorized, user.ID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrorUserNotFound
	}

	return nil
}

func (s *userStorage) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrorUserNotFound
	}

	return nil
}

func scanUser(row pgx.Row) (*model.User, error) {
	user := &model.User{}
	err := row.Scan(&user.ID, &user.TelegramID, &user.FirstName, &user.LastName, &user.Role, &user.IsAuthorized, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}
	return user, nil
}
