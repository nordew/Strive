package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TelegramIDKey is a key to store telegram_id in the context
const TelegramIDKey = "telegram_id"

var (
	ErrMissingAuthHeader = errors.New("authorization header is required")
	ErrInvalidAuthHeader = errors.New("invalid authorization header format")
	ErrInvalidTelegramID = errors.New("invalid telegram_id")
)

func TelegramAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrMissingAuthHeader.Error()})
			c.Abort()
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidAuthHeader.Error()})
			c.Abort()
			return
		}

		telegramID := strings.TrimPrefix(authHeader, prefix)
		if telegramID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrInvalidTelegramID.Error()})
			c.Abort()
			return
		}

		c.Set(TelegramIDKey, telegramID)
		c.Next()
	}
}
