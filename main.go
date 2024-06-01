package main

import (
	"KotobaHub/config"
	"log"
)

func main() {
	cfg := config.New()
	log.Println("Starting KotobaHub")

	if cfg.Debug {
		log.Println("Debug mode enabled")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	log.Println("Database:", cfg.DBPath)

}
