package api

import (
	"experiment/internal/ui/handler"
	"github.com/gin-gonic/gin"
)

func (h *Router) initUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", handler.GetUser)
		// Добавьте другие роуты для пользователя
	}
}
