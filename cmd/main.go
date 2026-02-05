package main

import (
	"parking-management-system-v1/internal/app"
	"parking-management-system-v1/internal/container"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger"
)

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
