package main

import (
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/internal/container"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger"
)

// @title           Parking API
// @version         1.0
// @description     Dokumentasi API Sistem Manajemen Parkir.
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	traceID := logger.GenerateTraceID()
	log := logger.WithTraceID(traceID)
	log.Info("Application Starting", "main")

	cfg, _ := config.LoadConfig()
	db := app.InitDatabase(cfg)
	defer db.Close()

	cont := container.NewContainer(db.DB, cfg)
	server := app.NewServer(cfg, cont)
	server.Run()
}
