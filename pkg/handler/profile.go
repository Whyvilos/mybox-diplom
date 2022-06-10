package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox"
)

func (h *Handler) getUserProfile(c *gin.Context) {

	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "GetUser1: "+err.Error())
		return
	}

	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "GetUser2: "+" - invalid id param")
		return
	}

	selected_user, err := h.services.UserProfile.GetById(you_id, id_user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "GetUser3: "+err.Error())
		return
	}

	//flagYou := (you_id == selected_user.Id_user)
	if you_id == selected_user.Id_user {
		selected_user.IsYou = true
		newInfoResponse("It's you!")
	}

	//TODO проверка на подписку

	newAnswerResponse(http.StatusOK, "GET api/:id")
	c.JSON(http.StatusOK, selected_user)
}

func (h *Handler) postFollow(c *gin.Context) {
	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostFollow: "+err.Error())
		return
	}

	id_user, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if id_user == you_id {
		newErrorResponse(c, http.StatusBadRequest, "you can't follow yourself")
		return
	}

	err = h.services.UserProfile.Follow(you_id, id_user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//TODO проверка на подписку

	newAnswerResponse(http.StatusOK, "POST api/:id/follow")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "followed",
	})

}

func (h *Handler) postUnFollow(c *gin.Context) {
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

	if id_user == you_id {
		newErrorResponse(c, http.StatusBadRequest, "you can't")
		return
	}

	err = h.services.UserProfile.UnFollow(you_id, id_user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//TODO проверка на подписку

	newAnswerResponse(http.StatusOK, "POST api/:id/unfollow")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "unfollowed",
	})

}

func (h *Handler) getCheckFollow(c *gin.Context) {

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

	if id_user == you_id {
		newErrorResponse(c, http.StatusBadRequest, "it's you")
		return
	}

	checkFlag, err := h.services.UserProfile.CheckFollow(you_id, id_user)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	newAnswerResponse(http.StatusOK, "GET api/check_foolow")

	c.JSON(http.StatusOK, map[string]interface{}{
		"check_flag": checkFlag,
	})
}

type getLineResponse struct {
	Data []mybox.Post `json:"data"`
}

func (h *Handler) getLine(c *gin.Context) {

	you_id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	line, err := h.services.UserProfile.LoadLine(you_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newAnswerResponse(http.StatusOK, "GET api/line")

	c.JSON(http.StatusOK, getLineResponse{
		Data: line,
	})

}

func (h *Handler) postUploadAvatar(c *gin.Context) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	c.Request.ParseMultipartForm(10 << 20)

	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newInfoResponse(fmt.Sprintf("User ID: %d", userId))
	err = h.services.Media.SaveUrlAvatar(userId, strings.Split(tempFile.Name(), "\\")[1])
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newInfoResponse("Upload file: " + strings.Split(tempFile.Name(), "\\")[1])
	newAnswerResponse(http.StatusOK, "POST api/:id/upload-avatar")
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successful",
	})
}

func (h *Handler) postUploadPostMedia(c *gin.Context) {

	c.Request.ParseMultipartForm(10 << 20)

	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		newErrorResponse(c, http.StatusInternalServerError, "PostPostMedia: "+err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "post-*.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostPostMedia: "+err.Error())
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostPostMedia: "+err.Error())
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostPostMedia: "+err.Error())
		return
	}
	newInfoResponse("Upload file: " + strings.Split(tempFile.Name(), "\\")[1])
	newAnswerResponse(http.StatusOK, "POST api/:id/upload-post-media")
	c.JSON(http.StatusOK, map[string]interface{}{
		"url_media": strings.Split(tempFile.Name(), "\\")[1],
	})
}

func (h *Handler) postUploadItemMedia(c *gin.Context) {
	c.Request.ParseMultipartForm(10 << 20)

	file, handler, err := c.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		newErrorResponse(c, http.StatusInternalServerError, "PostItemMedia: "+err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "item-*.png")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostItemMedia: "+err.Error())
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostItemMedia: "+err.Error())
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "PostItemMedia: "+err.Error())
		return
	}
	newInfoResponse("Upload file: " + strings.Split(tempFile.Name(), "\\")[1])
	newAnswerResponse(http.StatusOK, "POST api/:id/upload-post-media")
	c.JSON(http.StatusOK, map[string]interface{}{
		"url_media": strings.Split(tempFile.Name(), "\\")[1],
	})
}
