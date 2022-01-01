package routes

import (
	"net/http"

	"github.com/apot-group/golang-skeleton/x-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/storage", http.Dir("storage"))
	api := r.Group("/api")
	v1 := api.Group("v1")

	celery := v1.Group("celery")
	{
		celery.GET("status/:task_id", controllers.StatusCelery)
	}

	ml := v1.Group("ml")
	{
		ml.POST("face/detection", controllers.FaceDetection)
		ml.POST("face/recognition", controllers.FaceRecognition)
	}
	return r
}
