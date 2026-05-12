package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	DB  DBConfig
	App AppConfig
	S3  S3Config
}

type DBConfig struct {
	URL string
}

type AppConfig struct {
	Port int
}

type S3Config struct {
	Bucket          string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

func LoadConfig() *Config {
	return &Config{
		DB: DBConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		App: AppConfig{
			Port: getEnvInt("PORT", 8080),
		},
		S3: S3Config{
			Bucket:          os.Getenv("AWS_S3_BUCKET"),
			Region:          os.Getenv("AWS_REGION"),
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
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
