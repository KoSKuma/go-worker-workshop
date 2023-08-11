package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/koskuma/go-worker-workshop/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQAdapter interface {
	Publish(message []byte, exchangeName string, routingKey string) error // Publish message to queue
	CreateAndBindQueue(queueConfig config.QueueConfig) error              // Create and bind queue
	Subscribe(queueName string, callback func(message []byte)) error      // Subscribe to queue
}

type queue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewQueue(uri string) IRabbitMQAdapter {
	conn, err := amqp.Dial(uri)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q := queue{conn: conn, ch: ch}
	return q
}

func (q queue) CreateAndBindQueue(queueConfig config.QueueConfig) error {
	err := q.declareExchange(queueConfig.ExchangeName, queueConfig.ExchangeType)
	if err != nil {
		return err
	}
	err = q.declareQueue(queueConfig.QueueName)
	if err != nil {
		return err
	}
	routingKey := ""
	if queueConfig.ExchangeType == "direct" {
		routingKey = queueConfig.QueueName
	}
	err = q.bindQueue(queueConfig.QueueName, routingKey, queueConfig.ExchangeName)
	if err != nil {
		return err
	}
	return nil
}

func (q queue) declareExchange(exchangeName string, exchangeType string) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	return err
}

func (q queue) declareQueue(queueName string) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}

func (q queue) bindQueue(queueName string, routingKey string, exchangeName string) error {
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil)
	return err
}

func (q queue) Publish(message []byte, exchangeName string, routingKey string) error {
	// fmt.Println("Publishing message to queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.PublishWithContext(ctx,
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        message,
		})
	return err
}

func (q queue) Subscribe(queueName string, callback func(message []byte)) error {
	fmt.Println("Subscribing to queue")
	ch, err := q.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			callback(d.Body)
			ch.Ack(d.DeliveryTag, false)
		}
	}()
	<-forever
	return nil
}
