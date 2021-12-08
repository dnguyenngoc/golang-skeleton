/**
 * @author Duy Nguyen
 * @email duynguyenngoc@hotmail.com
 * @create date 2021-12-08
 * @desc go-main init
 */

package main

import (
	_ "encoding/json"
	_ "fmt"
	"log"

	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
	"github.com/apot-group/golang-skeleton/x-api/src/routes"
	"github.com/apot-group/golang-skeleton/x-api/src/settings"

	// "github.com/bsphere/celery"
	"github.com/gin-gonic/gin"
	// "github.com/streadway/amqp"
	// "github.com/jinzhu/gorm"
)

func main() {

	log.Println("Loading init app ......")

	loggers.InitLoggerConfig()

	loggers.InfoLogger.Println("Starting the application...")
	settings.InitConfig()
	loggers.InfoLogger.Println("Init config variable completed! -> Init celery config")

	// int celery
	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-service:5672/")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// // if err != nil {
	// // 	panic(err)
	// // }

	log.Println("Loading init app Completed! -> Loading gin router ...")

	gin.SetMode(gin.ReleaseMode)

	r := routes.SetupRoutes()

	// r.GET("/test", func(c *gin.Context) {
	// 	task, err := celery.NewTask("x-ml.test", []string{}, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	err = task.Publish(ch, "", "x-ml")
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// })
	log.Println("Loading gin router Completed! -> Server run on :8080")
	r.Run()

}
