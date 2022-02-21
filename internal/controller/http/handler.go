package http

import (
	"1aidar1/bastau/go-api/config"
	v1 "1aidar1/bastau/go-api/internal/controller/http/v1"
	"1aidar1/bastau/go-api/internal/service"
	"1aidar1/bastau/go-api/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//if cfg.Env != config.Prod {
	//	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func (h *Handler) initAPI(router *gin.Engine) {

	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	//router.Use(JSONMiddleware())
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
