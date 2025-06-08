package utils

import (
	"os"
	"github.com/joho/godotenv"
)

func SetupEnviroment() {
	enviroment := os.Getenv("ENV")

	if enviroment == "DEV" {
		println("Development enviroment detected!!")
		godotenv.Load()
	} else {
		println("Production enviroment detected!!")
	}
}
