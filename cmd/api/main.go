package main

import (
	"armada_management_system/internal/config"
	"armada_management_system/internal/database"
	"armada_management_system/internal/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectAndMigrate()

	r := gin.Default()
	r.GET("/vehicles/:vehicle_id/location", handler.GetLatestLocation)
	r.GET("/vehicles/:vehicle_id/history", handler.GetLocationHistory)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
