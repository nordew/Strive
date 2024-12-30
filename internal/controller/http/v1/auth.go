package v1

import "github.com/gin-gonic/gin"

func (c *Controller) initAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", c.login)
	}
}

type loginRequest struct {
	TelegramID int64 `json:"telegram_id" binding:"required"`
}

func (c *Controller) login(ctx *gin.Context) {
	const op = "controller.login"

	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleErr(ctx, 400, err)
		return
	}

	authResp, err := c.userService.Login(ctx, req.TelegramID)
	if err != nil {
		handleErr(ctx, 500, err)
		return
	}

	ctx.JSON(200, authResp)
}
