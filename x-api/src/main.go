/**
 * @author Duy Nguyen
 * @email duynguyenngoc@hotmail.com
 * @create date 2021-12-08
 * @desc go-main init
 */

package main

import (
	"log"

	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
	"github.com/apot-group/golang-skeleton/x-api/src/routes"
	"github.com/apot-group/golang-skeleton/x-api/src/settings"
	"github.com/apot-group/golang-skeleton/x-api/src/tasks"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
)

func main() {

	log.Println("Loading init app ......")

	loggers.InitLoggerConfig()

	loggers.InfoLogger.Println("Starting the application...")

	settings.InitConfig()

	loggers.InfoLogger.Println("Init config variable completed! -> Init celery config")

	tasks.InitCeleryConfig()

	log.Println("Loading init app Completed! -> Loading gin router ...")

	gin.SetMode(gin.ReleaseMode)

	r := routes.SetupRoutes()

	// r.GET("/test", func(c *gin.Context) {

	// })
	log.Println("Loading gin router Completed! -> Server run on :8080")
	r.Run()

}
