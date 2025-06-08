package seeder

import (
	"encoding/json"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type LocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func InitialSeed() {
	_ = godotenv.Load()

	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTT_BROKER")).
		SetClientID("initial_publisher")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Failed to connect to MQTT:", token.Error())
	}
	log.Println("Connected to MQTT broker")

	vehicleID := "B1234XYZ"
	startTime := int64(1715003456)

	first := LocationPayload{
		VehicleID: vehicleID,
		Latitude:  -6.2088,
		Longitude: 106.8456,
		Timestamp: startTime,
	}

	second := LocationPayload{
		VehicleID: vehicleID,
		Latitude:  -6.2086,        
		Longitude: 106.8458,       
		Timestamp: startTime + 2,  
	}

	publishLocation(client, first)
	time.Sleep(2 * time.Second)
	publishLocation(client, second)
}

func publishLocation(client mqtt.Client, payload LocationPayload) {
	data, _ := json.Marshal(payload)
	topic := "/fleet/vehicle/" + payload.VehicleID + "/location"

	token := client.Publish(topic, 0, false, data)
	token.Wait()

	if token.Error() != nil {
		log.Fatalf("Failed to publish: %v", token.Error())
	}

	log.Printf("Published to %s: %s", topic, data)
}
