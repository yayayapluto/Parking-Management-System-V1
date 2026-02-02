package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type VehicleTypeRepository interface {
	BaseRepository[models.VehicleType]
}

type vehicleTypeRepository struct {
	BaseRepository[models.VehicleType]
}

func NewVehicleTypeRepository(db *gorm.DB) VehicleTypeRepository {
	return &vehicleTypeRepository{
		BaseRepository: NewBaseRepository[models.VehicleType](db),
	}
}
