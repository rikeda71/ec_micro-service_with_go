package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

const amqpURI = "amqp://user:bitnami@rabbitmq:5672"

// error handling of mq
func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// SendMessage メッセージ送信
func SendMessage(queueName string, payload []byte) {
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "failed to connect RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open channel")

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "failed to call queue")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		},
	)
	failOnError(err, "failed to sent message")
}
