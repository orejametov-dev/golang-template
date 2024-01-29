package api

import (
	"experiment/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
}

func NewHandler() *Router {
	return &Router{}
}

func (h *Router) Init(cfg *config.AppConfig) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Router) initAPI(router *gin.Engine) {
	h.initUserRoutes(router)
}
