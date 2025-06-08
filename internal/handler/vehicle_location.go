package handler

import (
	"armada_management_system/internal/database"
	"armada_management_system/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetLatestLocation(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	var loc model.VehicleLocation
	err := database.DB.
		Where("vehicle_id = ?", vehicleID).
		Order("timestamp DESC").
		First(&loc).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicle_id": loc.VehicleID,
		"latitude":   loc.Latitude,
		"longitude":  loc.Longitude,
		"timestamp":  loc.Timestamp.Unix(),
	})
}

func GetLocationHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	startStr := c.Query("start")
	endStr := c.Query("end")

	startUnix, err1 := strconv.ParseInt(startStr, 10, 64)
	endUnix, err2 := strconv.ParseInt(endStr, 10, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start or end timestamp"})
		return
	}

	var history []model.VehicleLocation
	err := database.DB.
		Where("vehicle_id = ? AND timestamp BETWEEN ? AND ?", vehicleID, time.Unix(startUnix, 0), time.Unix(endUnix, 0)).
		Order("timestamp ASC").
		Find(&history).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history"})
		return
	}

	results := make([]gin.H, len(history))
	for i, loc := range history {
		results[i] = gin.H{
			"vehicle_id": loc.VehicleID,
			"latitude":   loc.Latitude,
			"longitude":  loc.Longitude,
			"timestamp":  loc.Timestamp.Unix(),
		}
	}

	c.JSON(http.StatusOK, results)
}
