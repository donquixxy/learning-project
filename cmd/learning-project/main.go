package main

import (
	"learning-project/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.LearningCmd.Execute(); err != nil {
		log.Println("Stopping learning project . . . | ", err)
		os.Exit(1)
	}
}
