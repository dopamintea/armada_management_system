package model

import "time"

type VehicleLocation struct {
	ID        uint      `gorm:"primaryKey"`
	VehicleID string    `gorm:"column:vehicle_id;type:varchar(50);index"`
	Latitude  float64   `gorm:"type:decimal(9,6)"`
	Longitude float64   `gorm:"type:decimal(9,6)"`
	Timestamp time.Time `gorm:"index"`
}

func (VehicleLocation) TableName() string {
	return "vehicle_loctions"
}