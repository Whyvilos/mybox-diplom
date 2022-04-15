package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

func (h *Handler) signUp(c *gin.Context) {
	var input mybox.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("BindJSON: %s", err.Error()))
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("CreateUser: %s", err.Error()))
		return
	}
	newAnswerResponse(http.StatusOK, "POST auth/sign-up")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("BindJSON: %s", err.Error()))
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GenerateToken: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, "POST auth/sign-in")
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
func (h *Handler) testAuth(c *gin.Context) {

	id, _ := c.Get("userId")
	usId := id.(int)
	newAnswerResponse(http.StatusOK, fmt.Sprintf("GET api/test id:%d", usId))

	c.JSON(http.StatusOK, map[string]interface{}{
		"ping": "АВТОРИЗОВАН",
	})
}

func (h *Handler) getId(c *gin.Context) {

	id, _ := c.Get("userId")
	usId := id.(int)
	newAnswerResponse(http.StatusOK, fmt.Sprintf("GET api/get-id id:%d", usId))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": usId,
	})
}
