package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoEnv() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}
	return os.Getenv("MONGOURI")
}

func LoadEnvVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading env file")
	}
	return os.Getenv(key)

}
