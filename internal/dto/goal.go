package dto

import "time"

type (
	CreateGoalDTO struct {
		UserID      string    `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description omitempty"`
		Tags        []string  `json:"tags"`
		Deadline    time.Time `json:"deadline"`
	}
)
