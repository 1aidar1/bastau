package v1

import (
	"1aidar1/bastau/go-api/internal/service"
	"1aidar1/bastau/go-api/pkg/auth"
	"github.com/gin-gonic/gin"
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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
	}
}

func parseIdFromPath(c *gin.Context, param string) (int, error) {

	return 0, nil
}
