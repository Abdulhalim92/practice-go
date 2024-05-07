package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

const (
	// Адрес брокера AMQP
	amqpURI = "amqp://rmuser:rmpassword@localhost:5672/"
	// Название очереди для непотвержденных сообщений
	queueName = "queue"
	// Нвзвание обменника
	exchangeName = "exchange"
	// Название очереди для отложенных сообщений
	dlxQueueName = "dlx"
	// Время ожидания повторной попытки
	retryDelay = 5 * time.Second
)

func main() {
	// Создание подключения к брокеру AMQP
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()

	// Создание канала для связи с очередью
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// Определение параметров очереди
	queueArgs := amqp.Table{
		"x-dead-letter-exchange": "", // Пустая строка означает обмен по умолчанию
		"x-dead-routing-key":     dlxQueueName,
	}

	// Создание очереди
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		queueArgs,
	)
	if err != nil {
		panic(err)
	}

	// Создание обменника
	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Отправка сообщения в очередь
	err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			Body: []byte("Hello, world!"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сообщение отправлено в очередь")

	// Настройка обработчика для подтверждения сообщения
	confirm := make(chan amqp.Confirmation, 1)
	ch.NotifyPublish(confirm)

	// Ожидания подтверждения в течение 5 секунд
	select {
	case <-confirm:
		fmt.Println("Сообщение подтверждено")
	case <-time.After(5 * time.Second):
		fmt.Println("Сообщение не было подтверждено")
		// Отправка сообщения в очередь с отложенным повтором

	}

	// Создание очереди DLX
	_, err = ch.QueueDeclare(
		dlxQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Настройка обработчика для получения отложенных сообщений
	dlxMessages, err := ch.Consume(
		dlxQueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Ожидание и обрботка отложенных сообщений
	for msg := range dlxMessages {
		fmt.Printf("Сообщение отложенное: %s\n", msg.Body)
		_ = msg.Ack(false)
	}

	// Бесконечный цикл для поддержания работы программы
	select {}
}
