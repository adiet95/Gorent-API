package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Godotenv(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file dari helpers")
	}

	return os.Getenv(key)
}
