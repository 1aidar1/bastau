package v1

import (
	"1aidar1/bastau/go-api/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "user"
)

func Auth(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Vary", "Authorization")
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			contextSetUser(c, entity.EmptyUser)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			//app.invalidAuthenticationTokenResponse(w, r)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		token := headerParts[1]
		user, err := h.services.Users.GetUserByToken(c, []byte(token))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		contextSetUser(c, user)
		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := contextGetUser(c)
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		c.Next()

	}
}

func contextGetUser(c *gin.Context) entity.User {

	data, ok := c.Get(userCtx)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "missing user value in request context"})
	}
	user, ok := data.(entity.User)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "missing user value in request context"})
	}
	return user
}

func contextSetUser(c *gin.Context, user entity.User) {
	c.Set(userCtx, user)
	//ctx := context.WithValue(r.Context(), userContextKey, user)
	//return r.WithContext(ctx)
}
