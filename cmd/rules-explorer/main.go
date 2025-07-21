package main

import (
	"log"

	"rules-explorer/internal/app"
)

func main() {
	application := app.New()
	
	if err := application.Initialize(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	
	if err := application.Run(); err != nil {
		log.Fatalf("App failed: %v", err)
	}
}