package config

import (
	"log"

	"github.com/vandyahmad24/alat-bantu/config"
)

type Config struct {
	// Specific Apps
	PORT                string
	MIDTRANS_SERVER_KEY string
	MIDTRANS_ACCESS_KEY string
	JWT_SECRET          string
	DB_HOST             string
	DB_PORT             string
	DB_USERNAME         string
	DB_PASSWORD         string
	DB_NAME             string
	DB_MYSQL            string
}

func NewConfig() *Config {
	if err := config.LoadEnv(configPath); err != nil {
		log.Fatalf("Application dimissed. Application cannot find %s to run this application", err.Error())
	}
	return &Config{
		PORT:                config.GetEnv("PORT", ""),
		MIDTRANS_SERVER_KEY: config.GetEnv("MIDTRANS_SERVER_KEY", ""),
		MIDTRANS_ACCESS_KEY: config.GetEnv("MIDTRANS_ACCESS_KEY", ""),
		JWT_SECRET:          config.GetEnv("JWT_SECRET", ""),
		DB_HOST:             config.GetEnv("DB_HOST", ""),
		DB_PORT:             config.GetEnv("DB_PORT", ""),
		DB_USERNAME:         config.GetEnv("DB_USERNAME", ""),
		DB_PASSWORD:         config.GetEnv("DB_PASSWORD", ""),
		DB_NAME:             config.GetEnv("DB_NAME", ""),
		DB_MYSQL:            config.GetEnv("DB_MYSQL", ""),
	}
}
