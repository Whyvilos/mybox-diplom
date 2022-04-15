package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

type getAllItemsResponse struct {
	Data []mybox.SimpleItem `json:"data"`
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_catalog, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userId != id_catalog {
		newErrorResponse(c, http.StatusInternalServerError, "it's not your catalog")
		return
	}

	var input mybox.Item
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error-JSON: %s", err.Error()))
		return
	}

	id_item, err := h.services.Catalog.CreateItem(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("CreateItem: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, "POST api/:id/catalog")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id_item": id_item,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {

	id_catalog, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	catalog, err := h.services.Catalog.GetAll(id_catalog)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("GET api/:id/catalog id_catalog:%d", id_catalog))
	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: catalog,
	})

}

func (h *Handler) getItemById(c *gin.Context) {}

func (h *Handler) updateItem(c *gin.Context) {}

func (h *Handler) deleteItem(c *gin.Context) {}
