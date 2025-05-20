package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load() // загрузка вирт окружения
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     getEnv("DBHOST", "localhost"),
		DBPort:     getEnv("DBPORT", "5433"),
		DBUser:     getEnv("DBUSER", "postgres"),
		DBPassword: getEnv("DBPASSWORD", "postgres"),
		DBName:     getEnv("DBNAME", "postgres"),
		DBSSLMode:  getEnv("DBSSLMODE", "disable"),
	}, nil
}

// НЕОБЯЗАТЕЛЬНО, НО УДОБНО
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
