package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nordew/Strive/internal/service"
	"log"
)

type Controller struct {
	userService service.UserService
}

func NewController(userService service.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (c *Controller) InitAndRun(httpPort int) {
	router := gin.New()

	c.initAuthRoutes(router)

	if err := router.Run(":" + fmt.Sprint(httpPort)); err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}
}

func handleErr(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
	return
}
