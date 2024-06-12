package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SERVER_PORT string
var DATABASE_URL string

func LoadEnvVars() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
}
