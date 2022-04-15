package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserProfile(c *gin.Context) {

	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	selected_user, err := h.services.UserProfile.GetById(you_id, id_user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if you_id == selected_user.Id_user {
		newInfoResponse("It's you!")
	}

	newAnswerResponse(http.StatusOK, "GET api/:id")
	c.JSON(http.StatusOK, selected_user)
}
