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
	contextFile, exists := os.LookupEnv("ABOUT_CONTEXT_FILE")
	if exists {
		log.Printf("Using environment variable ABOUT_CONTEXT_FILE: %v", contextFile)
	} else {
		contextFile = "/etc/about/ace-context.yaml"
		log.Printf("ABOUT_CONTEXT_FILE environment variable not set, using default value: %v", contextFile)
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
	env := env.NewEnv(contextFile)
	mux.HandleFunc("/env/context", env.EnvContextHandler)
	mux.HandleFunc("/env/version", env.EnvVersionHandler)
	mux.HandleFunc("/env/version/badge", env.EnvVersionBadgeHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("Listening on: %v", s.Addr)
	log.Fatal(s.ListenAndServe())
}
