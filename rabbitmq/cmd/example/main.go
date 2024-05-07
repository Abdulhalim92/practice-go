package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"practice_go/rabbitmq/internal"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("abduhalim", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()

	// Создаем канал
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer func() { _ = ch.Close() }()

	// Объявляем очередь
	q, err := ch.QueueDeclare(
		"test_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Отправляем сообщение
	msg := amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, RabbitMQ!"),
	}
	err = ch.PublishWithContext(
		context.Background(),
		"",
		q.Name,
		false,
		false,
		msg,
	)
	if err != nil {
		panic(err)
	}

	// Получаем сообщение
	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		fmt.Printf("Received a message with ID: %s and Body %s\n", msg.MessageId, msg.Body)
	}
}
