package tasks

import (
	"log"

	"github.com/streadway/amqp"
)

var ch *amqp.Channel

func InitCelery() {
	
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-service:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
}
