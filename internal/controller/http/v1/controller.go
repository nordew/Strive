package v1

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nordew/Strive/internal/service"
)

type Controller struct {
	userService service.UserService
	goalService service.GoalService
	router      *gin.Engine
}

func NewController(userService service.UserService) *Controller {
	controller := &Controller{
		userService: userService,
		router:      gin.New(),
	}

	controller.initRoutes()
	return controller
}

func (c *Controller) Init() *gin.Engine {
	return c.router
}

func (c *Controller) Run(httpPort int) {
	if err := c.router.Run(":" + fmt.Sprint(httpPort)); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}

func (c *Controller) initRoutes() {
	applyMiddlewares(c.router)
	c.initAuthRoutes(c.router)
}

func handleErr(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"error": err.Error()})
}
