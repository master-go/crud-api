package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVeriables() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
