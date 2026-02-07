package main

import (
	"log"
	"os"
	"parking-management-system-v1/database/seeders"
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger"
)

func main() {
	traceID := logger.GenerateTraceID()
	l := logger.WithTraceID(traceID)

	cfg, _ := config.LoadConfig()

	// Proteksi lingkungan produksi
	if os.Getenv("APP_ENV") == "production" {
		l.Error("Refresh denied: Production environment detected", "db-refresh")
		return
	}

	dbConn := app.InitDatabase(cfg)
	defer dbConn.Close()
	db := dbConn.DB

	// Daftar model untuk di-reset
	allModels := []interface{}{
		&models.VehicleType{},
		&models.CustomerRegistSource{},
		&models.PaymentMethod{},
		&models.Holiday{},
		&models.ZoneType{},
		// Tambahkan model lain di sini sesuai perkembangan project
	}

	log.Println("Database Refresh: Starting process")

	// Drop tables
	for _, m := range allModels {
		if err := db.Migrator().DropTable(m); err != nil {
			log.Printf("Warning: Could not drop table for model %T: %v", m, err)
		}
	}

	// Auto Migrate ulang
	log.Println("Database Refresh: Re-migrating tables")
	if err := db.AutoMigrate(allModels...); err != nil {
		log.Fatalf("Critical: Migration failed: %v", err)
	}

	// Jalankan Seeder
	log.Println("Database Refresh: Running seeders")
	seeders.RunSeeders(db)

	log.Println("Database Refresh: Success")
}
