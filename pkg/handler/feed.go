package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

type getAllPostsResponse struct {
	Data []mybox.Post `json:"data"`
}

func (h *Handler) createPost(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id_feed, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userId != id_feed {
		newErrorResponse(c, http.StatusInternalServerError, "it's not your feed")
		return
	}

	var input mybox.Post
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error-JSON: %s", err.Error()))
		return
	}

	id_post, err := h.services.Feed.CreatePost(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("CreateItem: %s", err.Error()))
		return
	}

	newAnswerResponse(http.StatusOK, "POST api/:id/feed")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id_post": id_post,
	})

}

func (h *Handler) getAllPosts(c *gin.Context) {

	id_feed, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	feed, err := h.services.Feed.GetAll(id_feed)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newAnswerResponse(http.StatusOK, fmt.Sprintf("GET api/:id/feed id_feed:%d", id_feed))
	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: feed,
	})
}

func (h *Handler) getPostById(c *gin.Context) {}

func (h *Handler) updatePost(c *gin.Context) {}

func (h *Handler) deletePost(c *gin.Context) {}
