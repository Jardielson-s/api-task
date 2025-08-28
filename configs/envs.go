package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func Envs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to loading .env file")
	}
}
