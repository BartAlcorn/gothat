package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV         string
	ADDRESS     string
	PORT        string
	OWM_API_KEY string
}

var Configs *Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	Configs = &Config{
		ENV:         os.Getenv("ENV"),
		ADDRESS:     os.Getenv("ADDRESS"),
		PORT:        os.Getenv("PORT"),
		OWM_API_KEY: os.Getenv("OWM_API_KEY"),
	}
}
