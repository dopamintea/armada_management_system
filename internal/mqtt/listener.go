package mqtt

import (
	"armada_management_system/internal/dto"
	"armada_management_system/internal/service"
	"encoding/json"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartListener() {
	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTT_BROKER")).
		SetClientID(os.Getenv("MQTT_CLIENT_ID"))

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT connection failed:", token.Error())
	}
	log.Println("MQTT connected")

	topic := "/fleet/vehicle/+/location"

	token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Message received on topic: %s", msg.Topic())

		var payload dto.LocationPayload
		if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
			log.Println("Invalid JSON:", err)
			return
		}

		service.ProcessIncomingPayload(payload)
	})

	if token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, token.Error())
	}

	log.Printf("Subscribed to topic: %s", topic)
}

