package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controllers "github.com/apot-group/golang-skeleton/x-api/src/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	// API AND V1 ********************************************************************************
	api := r.Group("/api")
	v1 := api.Group("v1")

	// HOME PAGE**********************************************************************************
	ml := v1.Group("ml")
	{
		ml.POST("face/detection", controllers.FaceDeteaction)
		ml.POST("face/recognition", controllers.FaceRecognition)
		ml.GET("face/:task_id", controllers.StatusCelery)
	}

	return r
}
