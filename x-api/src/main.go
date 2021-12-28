/**
 * @author: Duy Nguyen
 * @email: duynguyenngoc@hotmail.com
 * @create: 2021-12-08 int main
 * @update-1: 2021-12-28 add swagger
 */

package main

import (
	"log"
	// "github.com/jinzhu/gorm"
	"github.com/apot-group/golang-skeleton/x-api/src/docs"
	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
	"github.com/apot-group/golang-skeleton/x-api/src/routes"
	"github.com/apot-group/golang-skeleton/x-api/src/settings"
	"github.com/apot-group/golang-skeleton/x-api/src/tasks"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger x-api
// @version         1.0
// @description     This is a x-ml api server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   Duy Nguyen
// @contact.email  duynguyenngoc@hotmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
func main() {

	// INIT PROCESS ********************************************************************
	log.Println("Loading init app ......")
	loggers.InitLoggerConfig()
	loggers.InfoLogger.Println("Starting the application...")
	settings.InitConfig()

	// CONNECT CELERY ******************************************************************
	loggers.InfoLogger.Println("Init config variable completed! -> Init celery config")
	tasks.InitCeleryConfig()

	// HANDLE ROUTES *******************************************************************
	loggers.InfoLogger.Println("Init celery config completed! -> Start gin service ...")
	log.Println("Loading init app Completed! -> Loading gin router ...")
	gin.SetMode(gin.ReleaseMode)
	r := routes.SetupRoutes()

	// SWAGGER CONFIG PAGE *************************************************************
	log.Println("Loading router Completed! -> Loading swagger ...")
	docs.SwaggerInfo.Title = "Swagger x-API"
	docs.SwaggerInfo.Description = "This is x-api server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// START SERVER ********************************************************************
	log.Println("Loading swagger Completed! -> Server run on :8080")
	loggers.InfoLogger.Println("Server run on :8080")
	r.Run()

}
