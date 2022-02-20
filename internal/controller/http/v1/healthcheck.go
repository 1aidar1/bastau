package v1

import (
	"1aidar1/bastau/go-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthcheckRoutes struct {
	l logger.LoggerI
}

func newHealthcheckRoutes(handler *gin.RouterGroup, l logger.LoggerI) {
	r := &healthcheckRoutes{l}

	h := handler.Group("/healthcheck")
	{
		h.GET("/ping", r.ping)
	}
}

// @Summary     ping
// @Description check if server is running
// @ID          ping
// @Tags  	    health-check
// @Accept      json
// @Produce     json
// @Success     200 {object} pong
// @Failure     500 {object}
// @Router      /healthcheck/ping [get]
func (r *healthcheckRoutes) ping(c *gin.Context) {
	c.JSON(http.StatusOK, "PONG")
}
