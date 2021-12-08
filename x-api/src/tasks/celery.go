package tasks

import (
	"log"

	"github.com/apot-group/golang-skeleton/x-api/src/loggers"
	"github.com/bsphere/celery"
	"github.com/streadway/amqp"
)

// var conn *amqp.Connection
// var ch *amqp.Channel

const rabbitmqUrl = `amqp://guest:guest@rabbitmq-service:5672/`

func InitCeleryConfig() {

	loggers.InfoLogger.Println("Connect to celery broker ...")

	conn, err := amqp.Dial(rabbitmqUrl)

	if err != nil {
		log.Fatal(err)
		loggers.ErrorLogger.Fatal(err)
	}

	defer conn.Close()

	loggers.InfoLogger.Println("Init Celery connect completed!")

}

// need improve this fuction when have lagre sent task to rabbitmq in save time
func SentTask(exchange string, key string, taskName string, args []string, kwargs map[string]interface{}) {
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}

	defer ch.Close()

	task, err := celery.NewTask(taskName, args, kwargs)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}
	err = task.Publish(ch, exchange, key)
	if err != nil {
		loggers.ErrorLogger.Fatal(err)
	}
}
