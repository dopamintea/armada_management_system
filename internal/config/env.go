package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found.")
	}
}

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvFloat(key string, fallback float64) float64 {
	val := os.Getenv(key)
	if f, err := strconv.ParseFloat(val, 64); err == nil {
		return f
	}
	return fallback
}
