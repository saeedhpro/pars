package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type mq struct {
	rabbitmq *amqp.Connection
	channel  *amqp.Channel
}

var MQ mq

func Init() {
	rabbitmqConnection()
}

func rabbitmqConnection() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	MQ.rabbitmq = conn
	//createChannel()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func createChannel() {
	ch, err := MQ.rabbitmq.Channel()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ch.Close()
	MQ.channel = ch
	messages, err := ch.Consume(
		"FileUploaded",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for message := range messages {
			fmt.Printf("Received %s\n", message.Body)
		}
	}()
	<-forever
}
