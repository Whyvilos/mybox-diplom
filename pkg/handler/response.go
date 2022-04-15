package handler

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newInfoResponse(message string) {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.Warn(message)
}

func newAnswerResponse(statusCode int, message string) {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.Info(color.Ize(color.Blue, fmt.Sprintf("[%d] ", statusCode)) + message)
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.Error(color.Ize(color.Red, fmt.Sprintf("[%d] ", statusCode)) + message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
