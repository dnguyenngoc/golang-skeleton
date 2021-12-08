package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "app/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})
	// api := r.Group("/api")
	// v1 := api.Group("v1")

	

	// // AUTH***************************************************************************************
	// account := v1.Group("account")
	// {
	// 	account.POST("/login/token", controller.LoginAccessToken)
	// }

	// // USER CRUD**********************************************************************************
	// users := v1.Group("users")
	// {
	// 	users.GET("/get/:id", controller.GetUserById)
	// 	users.GET("/uuid/get/:uuid", controller.GetUserByUUID)
	// 	users.POST("/create", controller.CreateUser)
	// 	users.PUT("/update", controller.EditUserByID)
	// 	users.DELETE("/delete/:id", controller.DeleteUserByID)
	// }


	// // HOME PAGE**********************************************************************************
	// home:= v1.Group("home")
	// {
	// 	home.GET("intagramer", controller.GetIntagramerAvatar)
	// 	home.GET("intagramer/:name", controller.GetIntagramerImageByName)
	// 	home.GET("intagram/:name/:page", controller.FetchImageFromHashTagIntagram)
	// }

	// // NEWSPAPERS**********************************************************************************
	// news:= v1.Group("news")
	// {
	// 	news.GET("/24h", controller.FetchNews24h)
	// }

	return r
}
