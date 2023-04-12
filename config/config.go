package config

import (
	"main/models"
	"os"
)

func ConfigDB() *models.ConfigDB {
	return &models.ConfigDB{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func MainConfig() *models.MainConfig {
	return &models.MainConfig{
		PORT:         os.Getenv("PORT"),
		ENDPOINT_URL: os.Getenv("ENDPOINT_URL"),
	}
}
