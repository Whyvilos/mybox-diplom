package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

func (h *Handler) postSendMessage(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_chat, err := strconv.Atoi(c.Param("id_chat"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input mybox.Messaage
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error-JSON: %s", err.Error()))
		return
	}

	id_message, err := h.services.Chat.SendMassage(userId, id_chat, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Send message: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/chat/%d/", id_chat))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id_message,
	})
}

func (h *Handler) getChat(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_chat, err := strconv.Atoi(c.Param("id_chat"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	chat, err := h.services.Chat.GetAllMessage(userId, id_chat)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GET message: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("GET api/chat/%d/", id_chat))
	c.JSON(http.StatusOK, chat)
}

func (h *Handler) getFindChat(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_order, err := strconv.Atoi(c.Param("id_order"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id_chat, err := h.services.Chat.FindChat(userId, id_order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Find chat: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/chat/find/%d/", id_order))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id_chat,
	})
}

func (h *Handler) getFindChat2(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_order, err := strconv.Atoi(c.Param("id_order"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id_chat, err := h.services.Chat.FindChat2(userId, id_order)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Find chat2: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("POST api/chat/find/%d/", id_order))
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id_chat,
	})
}
