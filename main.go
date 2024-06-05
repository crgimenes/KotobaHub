package main

import (
	"embed"
	"log"
	"net/http"

	"KotobaHub/config"
)

//go:embed assets/*
var assets embed.FS

func handlerAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(assets)).ServeHTTP(w, r)
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

	/*
		err := db.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
	*/
	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", handlerHelthCheck)
	mux.HandleFunc("/assets/", handlerAssets)

	log.Println("Listening on", config.CFG.ListemAddress)
	err := http.ListenAndServe(config.CFG.ListemAddress, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
