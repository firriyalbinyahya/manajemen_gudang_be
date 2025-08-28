package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type JWT struct {
	AccessSecret      string
	AccessExpiryInSec int
}

var Config struct {
	JWT         JWT
	AdminAPIKey string
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading env variables directly")
	}

	accessExp, _ := strconv.Atoi(os.Getenv("ACCESS_EXPIRY_IN_SEC"))

	Config.JWT = JWT{
		AccessSecret:      os.Getenv("ACCESS_SECRET"),
		AccessExpiryInSec: accessExp,
	}
}
