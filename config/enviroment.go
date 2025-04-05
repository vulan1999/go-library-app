package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Loading enviroment variables failed")
	} else {
		log.Println("Loading enviroment variables successfully")
	}
}
