package main

import (
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger"
)

func main() {
	// 1. Generate ID unik untuk sesi ini
	traceID := logger.GenerateTraceID()

	// 2. Buat logger dengan ID tersebut
	log := logger.WithTraceID(traceID)

	// 3. Gunakan logger tersebut
	log.Info("Application Starting", "main")

	cfg, _ := config.LoadConfig()
	db := app.InitDatabase(cfg)
	defer db.Close()

	_ = app.NewContainer(db.DB)
}
