package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"KotobaHub/config"
	"KotobaHub/session"
)

//go:embed assets/*
var assets embed.FS

//go:embed templates/*
var tpltfs embed.FS

func handlerAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(assets)).ServeHTTP(w, r)
}

func handlerHelthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	sid := session.Save(r, w)
	log.Println("Session ID:", sid)

	s := session.Get(sid)
	for k, v := range s {
		log.Println(k, v)
	}

	tplt, err := template.ParseFS(tpltfs,
		"templates/index.html",
		"templates/head.html",
		"templates/header.html",
		"templates/footer.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	err = tplt.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

	session.Load()

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", handlerHelthCheck)
	mux.HandleFunc("/assets/", handlerAssets)
	mux.HandleFunc("/", handlerMain)

	log.Println("Listening on", config.CFG.ListemAddress)
	err := http.ListenAndServe(config.CFG.ListemAddress, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
