package app

import (
	"os"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/database"
	"parking-management-system-v1/pkg/logger"
)

func InitDatabase(cfg *config.Config) *database.PostgresConnection {
	dbConn, err := database.NewPostgresConnection(cfg.Postgres)
	if err != nil {
		logger.Error("Could not initialize database", "app-init", map[string]interface{}{"error": err.Error()})
		os.Exit(1) // Keluar aplikasi dengan status error
	}

	return dbConn
}
