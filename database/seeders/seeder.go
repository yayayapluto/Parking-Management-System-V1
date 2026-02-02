package seeders

import (
	"gorm.io/gorm"
	"log"
	"parking-management-system-v1/internal/models"
)

func RunSeeders(db *gorm.DB) {
	log.Println("Seeding database...")

	vehicleTypes := []models.VehicleType{
		{Code: "CAR", Name: "Mobil", Description: "Kendaraan Roda 4", IsActive: true},
		{Code: "MOTOR", Name: "Motor", Description: "Kendaraan Roda 2", IsActive: true},
		{Code: "TRUCK", Name: "Truk/Bus", Description: "Kendaraan Besar", IsActive: true},
	}
	for _, v := range vehicleTypes {
		db.FirstOrCreate(&v, models.VehicleType{Code: v.Code})
	}

	log.Println("Seeding completed successfully!")
}
