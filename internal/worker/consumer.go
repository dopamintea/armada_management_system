package worker

import (
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func StartWorker() {
	var conn *amqp091.Connection
	var ch *amqp091.Channel
	var err error

	maxAttempts := 10
	rabbitURL := os.Getenv("RABBITMQ_URL")

	for i := 1; i <= maxAttempts; i++ {
		conn, err = amqp091.Dial(rabbitURL)
		if err == nil {
			break
		}
		log.Printf("Retrying RabbitMQ connection for worker (%d/%d): %v", i, maxAttempts, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("[error] Worker failed to connect to RabbitMQ:", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal("Worker failed to open channel:", err)
	}

	exchangeName := "fleet.events"
	err = ch.ExchangeDeclare(
		"fleet.events", 
		"fanout",       
		true,           
		false,          
		false,          
		false,          
		nil,            
	)
	if err != nil {
		log.Fatal("Failed to declare exchange in worker:", err)
	}

	queueName := "geofence_alerts"
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	err = ch.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		log.Fatal("Failed to bind queue to exchange:", err)
	}

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Worker failed to consume from queue:", err)
	}

	log.Println("Worker listening for geofence alerts...")

	for msg := range msgs {
		log.Println("Alert received:")
		log.Println(string(msg.Body))
	}
}
