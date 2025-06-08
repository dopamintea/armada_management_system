package repository

import (
	"armada_management_system/internal/database"
	"armada_management_system/internal/model"
)

func SaveLocation(loc model.VehicleLocation) error {
	return database.DB.Create(&loc).Error
}
