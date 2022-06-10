package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

func (h *Handler) postCreateNotice(c *gin.Context) {

}

type getNoticeResponse struct {
	Data     []mybox.Notice `json:"data"`
	CountNew int            `json:"count_new"`
}

func (h *Handler) getNotices(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.services.UserProfile.GetNotices(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GET notice: %s", err.Error()))
		return
	}
	count := 0
	for _, x := range list {
		if x.Status == "new" {
			count++
		}
	}
	newAnswerResponse(http.StatusOK, "GET api/notice/")
	c.JSON(http.StatusOK, getNoticeResponse{
		Data:     list,
		CountNew: count,
	})
}

func (h *Handler) getCheckNotice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.UserProfile.NoticeCheck(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("GET check notice: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, "GET api/notice/check")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}
