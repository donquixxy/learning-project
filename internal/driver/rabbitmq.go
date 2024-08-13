package driver

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"learning-project/config"
	"log"
)

func NewRabbitMQConnection() *amqp.Connection {
	rabbitConfig := config.GetRabbitConfig()
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d", rabbitConfig.Username,
		rabbitConfig.Password,
		rabbitConfig.RabbitHost,
		rabbitConfig.RabbitPort)

	conn, err := amqp.Dial(connStr)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	log.Println("Connected to RabbitMQ")
	return conn
}

func DeclareNewExchange(conn *amqp.Connection, exchangeName string, queues ...string) (err error) {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	kind := "fanout"
	err = ch.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false, true, nil,
	)

	if err != nil {
		return err
	}

	for _, q := range queues {
		routingKey := exchangeName + "_" + q

		if err = ch.QueueBind(q, routingKey, exchangeName, false, nil); err != nil {
			log.Printf("Failed to bind queue %s to exchange: %s", q, err)
			return err
		}
	}

	return nil
}

func PublishMessage(conn *amqp.Connection, exchange, queue string, msg amqp.Publishing) error {

	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	return ch.Publish(exchange, queue, false, false, msg)
}
