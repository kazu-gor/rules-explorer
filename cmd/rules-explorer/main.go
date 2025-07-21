package main

import (
	"log"

	"rules-explorer/internal/ui"
)

func main() {
	app := ui.NewApp()
	
	if err := app.Initialize(); err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	
	if err := app.Run(); err != nil {
		log.Fatalf("App failed: %v", err)
	}
}