package routes

import (
	"net/http"

	"github.com/apot-group/golang-skeleton/x-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	api := r.Group("/api")
	v1 := api.Group("v1")

	// // HOME PAGE**********************************************************************************
	ml := v1.Group("home")
	{
		ml.POST("ml/face/detection", controllers.FaceDeteaction)
		ml.POST("ml/face/recognition", controllers.FaceRecognition)
		ml.GET("ml/face/:task_id", controllers.StatusCelery)
	}

	return r
}
