package main

import (
	"learning-project/config"

	"log"
)

func main() {
	appConfig := config.GetAppConfiguration()

	log.Printf("Starting Application: %v", appConfig.Name)
	log.Printf("At Environment: %v", appConfig.AppEnv)
}
