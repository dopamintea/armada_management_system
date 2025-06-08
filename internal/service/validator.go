package service

import (
	"armada_management_system/internal/dto"
	"regexp"
	"time"
)

var vehicleIDPattern = regexp.MustCompile(`^[A-Za-z]+\d+[A-Za-z\d]*$`)

func ValidatePayload(p dto.LocationPayload) bool {
	if !vehicleIDPattern.MatchString(p.VehicleID) {
		return false
	}
	if time.Unix(p.Timestamp, 0).After(time.Now().Add(24 * time.Hour)) {
		return false
	}
	return true
}
