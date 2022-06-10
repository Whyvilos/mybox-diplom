package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

type getFavoriteResponse struct {
	Data []mybox.Item `json:"data"`
}

func (h *Handler) getFavorite(c *gin.Context) {
	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.UserProfile.LoadFavorite(you_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newAnswerResponse(http.StatusOK, "GET api/line")

	c.JSON(http.StatusOK, getFavoriteResponse{
		Data: list,
	})

}

func (h *Handler) postAddFavorite(c *gin.Context) {
	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostAddFavorite: "+err.Error())
		return
	}

	id_item, err := strconv.Atoi(c.Param("id_item"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "PostAddFavorite: "+"invalid id param")
		return
	}

	err = h.services.Catalog.AddFavorite(you_id, id_item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostAddFavorite: "+err.Error())
		return
	}
	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/%d/favorite/%d", you_id, id_item))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}

func (h *Handler) postDeleteFavorite(c *gin.Context) {
	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostUnFavorite: "+err.Error())
		return
	}

	id_item, err := strconv.Atoi(c.Param("id_item"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "PostUnFavorite: "+"invalid id param")
		return
	}

	err = h.services.Catalog.DeleteFavorite(you_id, id_item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostUnFavorite: "+err.Error())
		return
	}
	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/%d/unfavorite/%d", you_id, id_item))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}

func (h *Handler) postCheckFavorite(c *gin.Context) {
	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostCheckFavorite: "+err.Error())
		return
	}
	id_item, err := strconv.Atoi(c.Param("id_item"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "PostCheckFavorite: "+"invalid id param")
		return
	}

	check_flag, err := h.services.Catalog.CheckFavorite(you_id, id_item)
	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/%d/favorite/%d/check", you_id, id_item))
	c.JSON(http.StatusOK, map[string]interface{}{
		"check_flag": check_flag,
	})
}
