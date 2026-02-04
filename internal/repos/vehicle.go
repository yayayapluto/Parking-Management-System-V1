package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type VehicleRepository interface {
	BaseRepository[models.Vehicle]
}

type vehicleRepository struct {
	BaseRepository[models.Vehicle]
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{
		BaseRepository: NewBaseRepository[models.Vehicle](db),
	}
}
