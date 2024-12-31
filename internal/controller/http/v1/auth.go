package v1

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nordew/Strive/internal/dto"
	"log"
	"strconv"
)

func (c *Controller) initAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", c.login)

		auth.Use(TelegramAuthMiddleware())
		auth.POST("/authorize", c.authorize)
	}
}

func (c *Controller) login(gCtx *gin.Context) {
	internalErr := errors.New("failed to login")

	var loginDTO dto.LoginUserDTO
	if err := gCtx.ShouldBindJSON(&loginDTO); err != nil {
		handleErr(gCtx, 400, internalErr)
		return
	}

	authResp, err := c.userService.Login(context.Background(), &loginDTO)
	if err != nil {
		handleErr(gCtx, 500, internalErr)
		return
	}

	gCtx.JSON(200, authResp)
}

func (c *Controller) authorize(gCtx *gin.Context) {
	internalErr := errors.New("failed to authorize")

	telegramID, ok := gCtx.Get(TelegramIDKey)
	if !ok {
		log.Printf("telegram_id not found in context")
		handleErr(gCtx, 400, internalErr)
		return
	}

	var telegramIDInt int64
	switch v := telegramID.(type) {
	case int64:
		telegramIDInt = v
	case string:
		var err error
		telegramIDInt, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Printf("failed to convert telegram_id string to int64: %v", err)
			handleErr(gCtx, 400, internalErr)
			return
		}
	default:
		log.Printf("failed to cast telegram_id to int64, unexpected type: %T", telegramID)
		handleErr(gCtx, 400, internalErr)
		return
	}

	var authDTO dto.AuthorizeUserRequest
	if err := gCtx.ShouldBindJSON(&authDTO); err != nil {
		log.Println("failed to bind json")
		handleErr(gCtx, 400, internalErr)
		return
	}

	if err := c.userService.Authorize(context.Background(), telegramIDInt, &authDTO); err != nil {
		handleErr(gCtx, 500, internalErr)
		return
	}

	gCtx.JSON(200, gin.H{"message": "authorized"})
}
