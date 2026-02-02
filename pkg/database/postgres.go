package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger" // alias agar tidak bentrok

	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/pkg/config"
	"parking-management-system-v1/pkg/logger" // import logger kita
)

type PostgresConnection struct {
	DB *gorm.DB
}

func NewPostgresConnection(cfg config.PostgresConfig) (*PostgresConnection, error) {
	logger.Info("Initializing PostgreSQL with GORM", "database-init")

	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	}

	db, err := gorm.Open(postgres.Open(cfg.URI), gormConfig)
	if err != nil {
		logger.Error("Failed to connect database", "database-init", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(int(cfg.MaxPoolSize))
	sqlDB.SetMaxIdleConns(int(cfg.MinPoolSize))
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Minute)

	logger.Info("Running AutoMigrate", "database-migration")
	err = db.AutoMigrate(
		&models.VehicleType{},
		&models.PaymentMethod{},
		// ... tambahkan semua model Anda ...
	)

	if err != nil {
		logger.Error("Auto migration failed", "database-migration", map[string]interface{}{"error": err.Error()})
		return nil, err
	}

	logger.Info("Database is ready and migrated", "database-init")

	return &PostgresConnection{DB: db}, nil
}

func (pc *PostgresConnection) Close() {
	if pc.DB != nil {
		sqlDB, err := pc.DB.DB()
		if err == nil {
			err := sqlDB.Close()
			if err != nil {
				return
			}
			logger.Info("Database connection closed", "database-close")
		}
	}
}
