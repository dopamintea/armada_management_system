package main

import (
	"armada_management_system/internal/config"
	"encoding/json"
	"log"
	"math"
	"math/rand"
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

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	godotenv.Load()
	opts := mqtt.NewClientOptions().
		AddBroker(os.Getenv("MQTT_BROKER")).
		SetClientID("mock_publisher_" + randomString(rnd, 5))

	centerLat := config.GetEnvFloat("GEOFENCE_LAT", 0.0000)
	centerLon := config.GetEnvFloat("GEOFENCE_LON", 0.0000)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Println("Connected to MQTT broker")

	for {
		vehicleID := generateVehicleID(rnd)
		lat, lon := randomLatLonAround(rnd, centerLat, centerLon, 0, 300)

		payload := LocationPayload{
			VehicleID: vehicleID,
			Latitude:  lat,
			Longitude: lon,
			Timestamp: time.Now().Unix(),
		}

		data, _ := json.Marshal(payload)
		topic := "/fleet/vehicle/" + vehicleID + "/location"
		token := client.Publish(topic, 0, false, data)
		token.Wait()

		log.Printf("Published to %s: %s", topic, data)
		time.Sleep(2 * time.Second)
	}
}

func randomLatLonAround(r *rand.Rand, centerLat, centerLon float64, minDist, maxDist float64) (float64, float64) {
	dist := minDist + r.Float64()*(maxDist-minDist)
	angle := r.Float64() * 2 * math.Pi

	latOffset := (dist * math.Cos(angle)) / 111000
	lonOffset := (dist * math.Sin(angle)) / (111000 * math.Cos(centerLat*math.Pi/180))

	return centerLat + latOffset, centerLon + lonOffset
}

func randomString(r *rand.Rand, n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}

func generateVehicleID(r *rand.Rand) string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")

	prefix := string(letters[r.Intn(len(letters))])
	digits := make([]rune, r.Intn(3)+2) 
	for i := range digits {
		digits[i] = numbers[r.Intn(len(numbers))]
	}
	suffix := make([]rune, r.Intn(2)+2) 
	for i := range suffix {
		suffix[i] = letters[r.Intn(len(letters))]
	}

	return prefix + string(digits) + string(suffix)
}
