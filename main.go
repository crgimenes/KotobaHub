package main

import (
	"log"
	"net/http"

	"KotobaHub/config"
	"KotobaHub/db"
)

func handlerMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handlerHelthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "ok"}`))
}

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

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerMain)
	mux.HandleFunc("/healthcheck", handlerHelthCheck)

	log.Println("Listening on", config.CFG.ListemAddress)
	err = http.ListenAndServe(config.CFG.ListemAddress, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
