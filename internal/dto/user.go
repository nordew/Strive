package dto

type LoginUserDTO struct {
	TelegramID int64 `json:"telegram_id" binding:"required"`
}

type AuthorizeUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
