package config

import (
	"github.com/joho/godotenv"
	"os"
	"parking-management-system-v1/pkg/logger" // import logger
	"strconv"
)

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	if os.Getenv("APP_ENV") == "development" {
		paths := []string{".env", "../.env", "../../.env", "../../../.env"}
		for _, path := range paths {
			if err := godotenv.Load(path); err == nil {
				logger.Info("Environment file loaded", "config-load", map[string]interface{}{"path": path})
				break
			}
		}
	}

	config := &Config{
		Postgres: PostgresConfig{
			URI:             getEnv("POSTGRES_URI", ""),
			Database:        getEnv("POSTGRES_DATABASE", ""),
			ConnectTimeout:  getEnvAsInt("POSTGRES_CONNECT_TIMEOUT", 5),
			MaxIdleTime:     getEnvAsInt("POSTGRES_MAX_IDLE_TIME", 2),
			MaxConnLifetime: getEnvAsInt("POSTGRES_MAX_CONN_LIFETIME", 30),
			MaxPoolSize:     uint64(getEnvAsInt("POSTGRES_MAX_POOL_SIZE", 20)),
			MinPoolSize:     uint64(getEnvAsInt("POSTGRES_MIN_POOL_SIZE", 2)),
		},
	}
	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
