package mq

import (
	"encoding/json"
	"fmt"
	"github.com/s14t284/ec_micro-service_with_go/handler"
	"github.com/s14t284/ec_micro-service_with_go/infra"
	"github.com/streadway/amqp"
)

const amqpURI string = "amqp://user:bitnami@rabbitmq:5672"

// error handling of mq
func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// ReceiveMessage メッセージ受信
func ReceiveMessage() {
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "failed to connect RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "failed to open channel")
	q, err := ch.QueueDeclare(
		"order-complete", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            //no wait
		nil,              // arguments
	)
	failOnError(err, "failed to call queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  //no-local
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "failed to setup receiving message")

	forever := make(chan bool)
	// 並行処理でメッセージ受信
	go func() {
		for data := range messages {
			fmt.Printf("%s\n", data.Body)
			user := new(handler.User)
			if err := json.Unmarshal(data.Body, user); err != nil {
				failOnError(err, "failed to parse json")
			}

			var cartConn []infra.Cart
			infra.DB.Where("user_id = ?", user.UserId).Find(&cartConn)
			ids := make([]int, len(cartConn))
			for i, c := range cartConn {
				ids[i] = int(c.ID)
			}
			infra.DB.Delete(&cartConn, ids)
			fmt.Println("complete to delete cart")
		}
	}()
	fmt.Println("start to receive message [To exit press Ctrl+C]")
	<-forever
}
