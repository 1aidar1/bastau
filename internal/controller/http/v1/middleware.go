package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	return getIdByContext(c, userCtx)
}
func getIdByContext(c *gin.Context, context string) (int, error) {
	idFromCtx, ok := c.Get(context)
	if !ok {
		return 0, errors.New("studentCtx not found")
	}
	id := idFromCtx.(int)

	return id, nil
}
