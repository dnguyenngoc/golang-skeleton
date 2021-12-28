package routes

import (
	"github.com/apot-group/golang-skeleton/x-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("v1")

	celery := v1.Group("celery")
	{
		celery.GET("status/:task_id", controllers.StatusCelery)
	}

	ml := v1.Group("ml")
	{
		ml.POST("face/detection", controllers.FaceDeteaction)
		ml.POST("face/recognition", controllers.FaceRecognition)
	}

	return r
}
