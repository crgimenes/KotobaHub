package main

import (
	"log"

	"KotobaHub/config"
	"KotobaHub/db"
)

func main() {
	config.Load()
	log.Println("Starting KotobaHub")

	if config.CFG.Debug {
		log.Println("Debug mode enabled")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
