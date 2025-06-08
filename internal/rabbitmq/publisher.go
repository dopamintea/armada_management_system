package rabbitmq

import (
	"encoding/json"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitPublisher() {
	var err error

	maxAttempts := 10
	rabbitURL := os.Getenv("RABBITMQ_URL")

	for i := 1; i <= maxAttempts; i++ {
		conn, err = amqp.Dial(rabbitURL)
		if err == nil {
			break
		}
		log.Printf("Retrying RabbitMQ connection (%d/%d): %v", i, maxAttempts, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("[error] failed to connect to RabbitMQ:", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal("RabbitMQ channel failed:", err)
	}

	err = ch.ExchangeDeclare(
		"fleet.events", "fanout", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("RabbitMQ exchange declare failed:", err)
	}

	log.Println("RabbitMQ connected and ready")
}

func PublishGeofenceEvent(payload any) {
	body, _ := json.Marshal(payload)

	err := ch.Publish(
		"fleet.events",
		"",
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Println("Failed to publish geofence event:", err)
	}
}
