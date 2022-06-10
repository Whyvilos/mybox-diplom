package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

func (h *Handler) postCreateOrder(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input mybox.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error-JSON from create order: %s", err.Error()))
		return
	}

	id_order, err := h.services.Order.CreateOrder(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Createorder: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, "POST api/order/")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id_order": id_order,
	})
}

func (h *Handler) getOrders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.Order.GetOrders(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GET orders: %s", err.Error()))
		return
	}
	newAnswerResponse(http.StatusOK, "GET api/order/")
	c.JSON(http.StatusOK, list)
}

func (h *Handler) getOrdersForYou(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.Order.GetOrdersForYou(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GET orders/shop: %s", err.Error()))
		return
	}
	newAnswerResponse(http.StatusOK, "GET api/order/shop")
	c.JSON(http.StatusOK, list)
}

func (h *Handler) putOrderStatus(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_order, err := strconv.Atoi(c.Param("id_order"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	status := c.Param("status")
	if status == "check" {
		_, err := h.services.CreateChat(userId, id_order, "new")
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error: create chat - %s", err))
			return
		}
	}

	if status == "new" || status == "check" || status == "drop" || status == "taken" || status == "done" {
		err = h.services.Order.UpdateOrderStatus(userId, id_order, status)
	} else {
		newErrorResponse(c, http.StatusBadRequest, "invalid status param")
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("PUT orders/%d/%s: %s", id_order, status, err.Error()))
		return
	}
	newAnswerResponse(http.StatusOK, fmt.Sprintf("PUT orders/%d/%s", id_order, status))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}
