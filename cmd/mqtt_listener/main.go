package main

import (
	"armada_management_system/internal/config"
	"armada_management_system/internal/database"
	"armada_management_system/internal/mqtt"
	"armada_management_system/internal/rabbitmq"
)

func main() {
	config.LoadEnv()
	database.ConnectAndMigrate()
	rabbitmq.InitPublisher()
	mqtt.StartListener()

	select {}
}
