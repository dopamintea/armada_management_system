package service

import (
	"armada_management_system/internal/config"
	"math"
)

func isWithinGeofence(lat, lon float64) bool {
	centerLat := config.GetEnvFloat("GEOFENCE_LAT", 0.0000)
	centerLon := config.GetEnvFloat("GEOFENCE_LON", 0.0000)
	radius := config.GetEnvFloat("GEOFENCE_RADIUS_METERS", 50)

	const R = 6371000
	dLat := deg2rad(centerLat - lat)
	dLon := deg2rad(centerLon - lon)

	lat1 := deg2rad(lat)
	lat2 := deg2rad(centerLat)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return distance <= radius
}

func deg2rad(d float64) float64 {
	return d * math.Pi / 180
}
