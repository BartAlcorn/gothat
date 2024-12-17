package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV         string // dev | prod
	ADDRESS     string // address, e.g. 127.0.0.1
	PORT        string // port, e.g. 9090
	OWM_API_KEY string // OpenWeaterherMap API Key
}

var Configs *Config

func init() {
	if err := godotenv.Load(); err != nil {
		// yes, this is fatal and is one of the few times it is proper for to force an exit
		log.Fatal("No .env file found")
	}
	Configs = &Config{
		ENV:         os.Getenv("ENV"),
		ADDRESS:     os.Getenv("ADDRESS"),
		PORT:        os.Getenv("PORT"),
		OWM_API_KEY: os.Getenv("OWM_API_KEY"),
	}
}
