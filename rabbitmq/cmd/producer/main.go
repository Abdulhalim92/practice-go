package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"practice_go/rabbitmq/internal"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("abduhalim", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()

	// -->> rpc <<--
	// Never use the same Connection for Consume and Publish
	consumeConn, err := internal.ConnectRabbitMQ("abduhalim", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumeConn.Close() }()
	// -->> rpc <<--

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// -->> rpc <<--
	consumeClient, err := internal.NewRabbitMQClient(consumeConn)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumeClient.Close() }()
	// -->> rpc <<--

	// Create context to manage timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//direct(ctx, client)
	//fanOut(ctx, client)
	rpc(ctx, client, consumeClient)

	log.Println(client)
}

func direct(ctx context.Context, client internal.RabbitClient) {
	_, err := client.CreateQueue("customers_created", true, false)
	if err != nil {
		panic(err)
	}

	_, err = client.CreateQueue("customers_test", false, true)
	if err != nil {
		panic(err)
	}

	// Create binding between the customer_events exchange and the customers-created queue
	err = client.CreateBinding("customers_created", "customers.created.*", "customer_events")
	if err != nil {
		panic(err)
	}

	// Created binding between the customer_events exchange and the customers-test queue
	err = client.CreateBinding("customers_test", "customers.*", "customer_events")
	if err != nil {
		panic(err)
	}

	// Create customer from Sweden
	for i := 0; i < 10; i++ {
		if err := client.Send(ctx, "customer_events", "customers.created.se", amqp091.Publishing{
			ContentType:  "text/plain",       // The payload we send is plaintext, could be JSON or others...
			DeliveryMode: amqp091.Persistent, // This tells rabbitMQ that this manage should be Saved if no resources accepts in before a restart (durable)
			Body:         []byte("An cool message between services"),
			ReplyTo:      "customers_created",
		}); err != nil {
			panic(err)
		}
	}

	if err := client.Send(ctx, "customers_events", "customers.test", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Transient, // This tell rabbitMQ that this message can be deleted if no resources accepts it before a restart (non-durable)
		Body:         []byte("A second cool message"),
	}); err != nil {
		panic(err)
	}

}

func fanOut(ctx context.Context, client internal.RabbitClient) {
	// Create customer from Sweden
	for i := 0; i < 10; i++ {
		if err := client.Send(ctx, "customer_events", "customers.created.se", amqp091.Publishing{
			ContentType:  "text/plain",       // The payload we send is plaintext, could be JSON or others...
			DeliveryMode: amqp091.Persistent, // This tells rabbitMQ that this manage should be Saved if no resources accepts in before a restart (durable)
			Body:         []byte("An cool message between services"),
		}); err != nil {
			panic(err)
		}
	}
}

func rpc(ctx context.Context, produceClient, consumeClient internal.RabbitClient) {
	// Create Unnamed Queue which will generate a random name, set AutoDelete to True
	queue, err := consumeClient.CreateQueue("", true, true)
	if err != nil {
		panic(err)
	}

	err = consumeClient.CreateBinding(queue.Name, queue.Name, "customer_callbacks")
	if err != nil {
		panic(err)
	}

	messageBus, err := consumeClient.Consume(queue.Name, "customer-api", true)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range messageBus {
			log.Printf("Message ID %s\n", message.MessageId)
			log.Printf("Message Callback %s\n", message.CorrelationId)
		}
	}()

	// Create customer from sweden
	for i := 0; i < 10; i++ {
		if err := produceClient.Send(ctx, "customer_events", "customers.created.se", amqp091.Publishing{
			ContentType:  "text/plain",       // The payload we send is plaintext, could be JSON or others..
			DeliveryMode: amqp091.Persistent, // This tells rabbitMQ that this message should be Saved if no resources accepts it before a restart (durable)
			Body:         []byte("An cool message between services"),
			// We add a REPLY TO which defines the
			ReplyTo: queue.Name,
			// CorrelationId can be used to know which Event this relates to
			CorrelationId: fmt.Sprintf("customer_created_%d", i),
		}); err != nil {
			panic(err)
		}
	}

	var blocking chan struct{}

	log.Println("Waiting on Callbacks, to close the program press CTRL+C")
	// This will block forever
	<-blocking
}
