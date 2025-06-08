package database

import (
	"armada_management_system/internal/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectAndMigrate() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	var db *gorm.DB
	var err error

	maxAttempts := 10
	for i := 1; i <= maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Retrying DB connection (%d/%d): %v", i, maxAttempts, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("[error] failed to initialize database, got error", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute)
	sqlDB.SetMaxOpenConns(10)

	log.Println("Database connected.")

	err = db.AutoMigrate(&model.VehicleLocation{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed.")
	DB = db
}
