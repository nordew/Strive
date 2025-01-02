package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
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

func RateLimiter(rps int, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, please try again later",
			})
			return
		}
		ctx.Next()
	}
}

func CORSProtection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func SecurityHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("X-Frame-Options", "DENY")
		ctx.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		ctx.Next()
	}
}

func RequestSizeLimiter(maxBytes int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxBytes)
		ctx.Next()
	}
}

func applyMiddlewares(router *gin.Engine) {
	router.Use(RateLimiter(10, 20)) // 10 requests per second with a burst of 20
	router.Use(CORSProtection())
	router.Use(SecurityHeaders())
	router.Use(RequestSizeLimiter(1048576)) // 1 MB max request body size
}
