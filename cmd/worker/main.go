package main

import (
	"armada_management_system/internal/config"
	"armada_management_system/internal/worker"
)

func main() {
	config.LoadEnv()
	worker.StartWorker()

	select {}
}
