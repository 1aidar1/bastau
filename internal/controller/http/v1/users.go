package v1

import (
	"1aidar1/bastau/go-api/internal/repository"
	"1aidar1/bastau/go-api/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
	}
	{
		users.Use(Auth(h))
		users.Use(RequireAuth())
		users.POST("/sign-out", h.userSignOut)
		users.GET("/by-token", h.getUserByToken)
		users.GET("/:id", h.getUser)
	}

}

var (
	ErrUserNotFound = "user not found"
)

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=13"`
	Role     string `json:"role" binding:"required,oneof=worker client all"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) userSignUp(c *gin.Context) {
	var inp userSignUpInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		invalidInputMsg(c, err.Error())
		return
	}

	if err := h.services.Users.Register(c.Request.Context(), service.UserSignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Phone:    inp.Phone,
		Password: inp.Password,
		Role:     inp.Role,
	}); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) getUserByToken(c *gin.Context) {
	var input struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		invalidInputMsg(c, err)
		return
	}

	user, err := h.services.Users.GetUserByToken(c, []byte(input.Token))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			notFoundMsg(c, ErrUserNotFound)
			return
		default:
			serverErrorMsg(c)
			return
		}
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) userSignOut(c *gin.Context) {
	//var input struct {
	//}
}

func (h *Handler) getUser(c *gin.Context) {
	userId, err := intParam(c, "id")
	if err != nil {
		writeJSONMsg(c, http.StatusBadRequest, "Bad request. Can't parse params.")
		return
	}
	user, err := h.services.Users.GetUserById(c, userId)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			notFoundMsg(c, ErrUserNotFound)
			return
		default:
			serverErrorMsg(c)
			return
		}
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) userSignIn(c *gin.Context) {
	var inp signInInput
	if err := c.ShouldBindJSON(&inp); err != nil {
		invalidInputMsg(c, err)
		return
	}

	res, err := h.services.Users.SignIn(c.Request.Context(), service.UserSignInInput{
		Email:    inp.Email,
		Password: inp.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, res)
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

//func (h *Handler) userRefresh(c *gin.Context) {
//	var inp refreshInput
//	if err := c.ShouldBindJSON(&inp); err != nil {
//		invalidInputMsg(c, err)
//
//		return
//	}
//
//	res, err := h.services.Users.RefreshTokens(c.Request.Context(), inp.Token)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err.Error())
//
//		return
//	}
//
//	c.JSON(http.StatusOK, tokenResponse{
//		AccessToken:  res.AccessToken,
//		RefreshToken: res.RefreshToken,
//	})
//}

//func (h *Handler) userVerify(c *gin.Context) {
//	code := c.Param("code")
//	if code == "" {
//		c.JSON(http.StatusBadRequest, "code is empty")
//
//		return
//	}
//
//	id, err := getUserId(c)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, err.Error())
//
//		return
//	}
//
//	if err := h.services.Users.Verify(c.Request.Context(), id, code); err != nil {
//		c.JSON(http.StatusInternalServerError, err.Error())
//
//		return
//	}
//
//	c.JSON(http.StatusOK, "success")
//}
