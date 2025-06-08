package service

import (
	"armada_management_system/internal/dto"
	"armada_management_system/internal/model"
	"armada_management_system/internal/rabbitmq"
	"armada_management_system/internal/repository"
	"log"
	"time"
)

func ProcessIncomingPayload(payload dto.LocationPayload) {
	if !ValidatePayload(payload) {
		log.Println("Invalid payload:", payload)
		return
	}

	loc := model.VehicleLocation{
		VehicleID: payload.VehicleID,
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Timestamp: time.Unix(payload.Timestamp, 0),
	}

	if err := repository.SaveLocation(loc); err != nil {
		log.Println("Failed to store to DB:", err)
		return
	}

	log.Printf("Stored: %+v\n", payload)

	if isWithinGeofence(loc.Latitude, loc.Longitude) {
	event := map[string]any{
		"vehicle_id": loc.VehicleID,
		"event":      "geofence_entry",
		"location": map[string]float64{
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
		},
		"timestamp": loc.Timestamp.Unix(),
	}
	rabbitmq.PublishGeofenceEvent(event)
	log.Println("Geofence event published:", event)
	}

}
