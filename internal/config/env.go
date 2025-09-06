package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("❌ Failed to load .env file: %v", err)
	}

	log.Println("✅ .env file loaded")
}
