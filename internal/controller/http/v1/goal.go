package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nordew/Strive/internal/dto"
)

func (c *Controller) initGoalRoutes() {
	goalGroup := c.router.Group("/goal")
	{
		goalGroup.POST("", c.createGoal)
	}
}

func (c *Controller) createGoal(ctx *gin.Context) {
	internalErr := errors.New("failed to create goal")

	var goalDTO dto.CreateGoalDTO
	if err := ctx.ShouldBindJSON(&goalDTO); err != nil {
		handleErr(ctx, 400, internalErr)
	}

	if err := c.goalService.Create(ctx, &goalDTO); err != nil {
		handleErr(ctx, 500, internalErr)
	}

	ctx.JSON(200, gin.H{"message": "goal created"})
}
