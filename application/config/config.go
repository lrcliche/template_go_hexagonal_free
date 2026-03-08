package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv              string
	HTTPPort            string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
}

func MustLoad() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		AppEnv:              getEnv("APP_ENV", "development"),
		HTTPPort:            getEnv("HTTP_PORT", "8080"),
		ReadTimeoutSeconds:  getEnvAsInt("READ_TIMEOUT_SECONDS", 10),
		WriteTimeoutSeconds: getEnvAsInt("WRITE_TIMEOUT_SECONDS", 10),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
