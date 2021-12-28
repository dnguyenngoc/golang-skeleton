package controllers

import (
	_ "net/http"

	_ "github.com/apot-group/golang-skeleton/x-api/src/tasks"
	"github.com/gin-gonic/gin"
	// "github.com/apot-group/golang-skeleton/x-api/src/loggers"
)

func FaceDeteaction(c *gin.Context) {
	// tasks.SentTask("", "x-ml", "x-ml.test", []string{}, nil)

}

func FaceRecognition(c *gin.Context) {
	// tasks.SentTask("", "x-ml", "x-ml.test", []string{}, nil)
}
