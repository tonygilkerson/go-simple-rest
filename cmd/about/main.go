package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonygilkerson/go-simple-rest/internal/env"
)

func main() {

	// Log to the console with date, time and filename prepended
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//
	// Get environment Variables
	//
	dataDir, exists := os.LookupEnv("ABOUT_DATA_DIR")
	if exists {
		log.Printf("Using ABOUT_DATA_DIR to set data directory: %v", dataDir)
	} else {
		dataDir = "/etc/about"
		log.Printf("ABOUT_DATA_DIR environment variable not set, using default value: %v", dataDir)
	}

	//
	// Create Server
	//
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//
	// Define routes
	//
	env := env.NewEnv(dataDir)
	mux.HandleFunc("/env/context", env.EnvContextHandler)
	mux.HandleFunc("/env/version", env.EnvVersionHandler)
	mux.HandleFunc("/env/version/badge", env.EnvVersionBadgeHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("Listening on: %v", s.Addr)
	log.Fatal(s.ListenAndServe())
}
