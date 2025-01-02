package model

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

// User TODO: Add subscription field
type User struct {
	ID           string    `json:"id"`
	TelegramID   int64     `json:"telegram_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Role         int       `json:"role"`
	IsAuthorized bool      `json:"is_authorized"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// NewUser creates a new User instance with validation and safety checks
func NewUser(
	id string,
	telegramID int64,
	firstName,
	lastName string,
	role int,
	createdAt,
	updatedAt time.Time) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	if telegramID <= 0 {
		return nil, errors.New("telegram_id must be a positive integer")
	}

	//firstName = strings.TrimSpace(firstName)
	//if firstName == ""{
	//
	//lastName = strings.TrimSpace(lastName)
	//
	if role < 0 {
		return nil, errors.New("role must be a positive integer")
	}

	if createdAt.IsZero() {
		return nil, errors.New("created_at cannot be zero")
	}
	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	}
	if updatedAt.Before(createdAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	user := &User{
		ID:         id,
		TelegramID: telegramID,
		FirstName:  firstName,
		LastName:   lastName,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	return user, nil
}
func (u *User) SetID(id string) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.New("id must be a valid UUID")
	}

	u.ID = id
	return u, nil
}

func (u *User) SetTelegramID(telegramID int64) (*User, error) {
	if telegramID <= 0 {
		return nil, errors.New("telegram_id must be a positive integer")
	}

	u.TelegramID = telegramID
	return u, nil
}

func (u *User) SetFirstName(firstName string) (*User, error) {
	firstName = strings.TrimSpace(firstName)
	if firstName == "" {
		return nil, errors.New("first_name cannot be empty")
	}

	u.FirstName = firstName
	return u, nil
}

func (u *User) SetLastName(lastName string) (*User, error) {
	lastName = strings.TrimSpace(lastName)
	u.LastName = lastName
	return u, nil
}

func (u *User) SetRole(role int) (*User, error) {
	if role < 0 {
		return nil, errors.New("role must be a positive integer")
	}

	u.Role = role
	return u, nil
}

func (u *User) SetUpdatedAt(updatedAt time.Time) (*User, error) {
	if updatedAt.IsZero() {
		return nil, errors.New("updated_at cannot be zero")
	}

	if updatedAt.Before(u.CreatedAt) {
		return nil, errors.New("updated_at cannot be before created_at")
	}

	u.UpdatedAt = updatedAt
	return u, nil
}

func (u *User) SetIsAuthorized(isAuthorized bool) (*User, error) {
	u.IsAuthorized = isAuthorized
	return u, nil
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetTelegramID() int64 {
	return u.TelegramID
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetRole() int {
	return u.Role
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
