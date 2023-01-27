package main

import (
	"log"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/app/apiserver"
	"github.com/MikhailMishutkin/Test_MediaSoft/internal/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config := config.NewConfig()

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}
