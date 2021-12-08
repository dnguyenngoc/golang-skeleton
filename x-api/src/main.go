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

	"github.com/apot-group/golang-skeleton/x-api/src/routes"
	"github.com/apot-group/golang-skeleton/x-api/src/settings"

	// "github.com/bsphere/celery"
	"github.com/gin-gonic/gin"
	// "github.com/streadway/amqp"
	// "github.com/jinzhu/gorm"
)

func main() {

	// init config
	settings.InitConfig()

	// int celery

	// conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-service:5672/")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// ch, err := conn.Channel()
	// if err != nil {
	// 	panic(err)
	// }

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

	r.Run()
}
