package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	getenv := os.Getenv(key)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return getenv
}
