package controllers

import (
	_ "net/http"

	"github.com/apot-group/golang-skeleton/x-api/src/tasks"
	"github.com/gin-gonic/gin"
)

func StatusCelery(c *gin.Context) {

	tasks.SentTask("", "x-ml", "x-ml.test", []string{}, nil)

}
