package main

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/errgroup"
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
	publishConn, err := internal.ConnectRabbitMQ("abduhalim", "secret", "localhost:5672", "customers")
	if err != nil {
		panic(err)
	}
	defer func() { _ = publishConn.Close() }()
	// -->> rpc <<--

	mqClient, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}

	// -->> rpc <<--
	publishClient, err := internal.NewRabbitMQClient(publishConn)
	if err != nil {
		panic(err)
	}
	// -->> rpc <<--

	// Set a timeout for 15 secs
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// Create an ErrGroup to manage concurrent tasks
	g, ctx := errgroup.WithContext(ctx)
	// Set amount of concurrent tasks
	g.SetLimit(10)

	//direct(ctx, g, mqClient)
	//fanOut(ctx, g, mqClient)
	rpc(ctx, g, mqClient, publishClient)
}

func direct(ctx context.Context, g *errgroup.Group, client internal.RabbitClient) {
	// blocking is used to block forever
	var blocking chan struct{}

	messageBus, err := client.Consume("customers_created", "email-service", false)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range messageBus {
			// Spawn a worker
			msg := message
			g.Go(func() error {
				log.Printf("New Message: %s", msg.Body)

				time.Sleep(10 * time.Second)
				// Multiple means that we acknowledge a batch of messages, leave false for now
				if err := msg.Ack(false); err != nil {
					log.Printf("Acknowledged message failed: Retry ? Handle manually %s\n", msg.MessageId)
					return err
				}
				log.Printf("Acknowledged message %s\n", msg.MessageId)
				return nil
			})
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")
	// This will block forever
	<-blocking

}

func fanOut(ctx context.Context, g *errgroup.Group, client internal.RabbitClient) {
	// blocking is used to block forever
	var blocking chan struct{}

	// Create Unnamed Queue which will generate a random name, set AutoDelete to True
	queue, err := client.CreateQueue("", true, true)
	if err != nil {
		panic(err)
	}
	// Create binding between the customer_events exchange and the new Random Queue
	// Can skip Binding key since fan-out will skip that rule
	if err := client.CreateBinding(queue.Name, "", "customer_events"); err != nil {
		panic(err)
	}

	messageBus, err := client.Consume(queue.Name, "email-service", false)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range messageBus {
			// Spawn a worker
			msg := message
			g.Go(func() error {
				log.Printf("New Message: %s", msg.Body)

				time.Sleep(10 * time.Second)
				// Multiple means that we acknowledge a batch of messages, leave false for now
				if err := msg.Ack(false); err != nil {
					log.Printf("Acknowledged message failed: Retry ? Handle manually %s\n", msg.MessageId)
					return err
				}
				log.Printf("Acknowledged message %s\n", msg.MessageId)
				return nil
			})
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")
	// This will block forever
	<-blocking
}

func rpc(ctx context.Context, g *errgroup.Group, consumeClient, publishClient internal.RabbitClient) {
	// blocking is used to block forever
	var blocking chan struct{}

	// Create Unnamed Queue which will generate a random name, set AutoDelete to True
	queue, err := consumeClient.CreateQueue("", true, true)
	if err != nil {
		panic(err)
	}
	// Create binding between the customer_events exchange and the new Random Queue
	// Can skip Binding key since fan-out will skip that rule
	err = consumeClient.CreateBinding(queue.Name, "", "customer_events")
	if err != nil {
		panic(err)
	}

	messageBus, err := consumeClient.Consume(queue.Name, "email-service", false)
	if err != nil {
		panic(err)
	}

	go func() {
		for message := range messageBus {
			// Spawn a worker
			g.Go(func() error {
				log.Printf("New Message: %s", message.Body)
				// Multiple means that we acknowledge a batch of messages, leave false for now
				if err := message.Ack(false); err != nil {
					log.Printf("Acknowledged message failed: Retry ? Handle manually %s\n", message.MessageId)
					return err
				}

				log.Printf("Acknowledged message %s\n", message.ReplyTo)

				// Use the msg.ReplyTo to send the message to the proper Queue
				err = publishClient.Send(ctx, "customer_callbacks", message.ReplyTo, amqp091.Publishing{})
			})
		}
	}()
}
