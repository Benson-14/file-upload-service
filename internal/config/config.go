package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

type DBConfig struct {
	URL string
}

type AppConfig struct {
	Port    int
	Storage string
}

func LoadConfig() *Config {
	return &Config{
		DB: DBConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		App: AppConfig{
			Port:    getEnvInt("PORT", 8080),
			Storage: os.Getenv("STORAGE"),
		},
	}
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		} else {
			slog.Error("error converting to int: " + err.Error())
		}
	}
	return defaultValue
}
