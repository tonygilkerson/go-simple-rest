package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

)

// Main
func main() {

	// Log to the console with date, time and filename prepended
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//
	// Get environment Variables
	//
	dataDir, exists := os.LookupEnv("ABOUT_DATA_DIR")
	if exists {
		log.Printf("Using ABOUT_DATA_DIR to set data directory: %v",dataDir)
	} else {
		dataDir = "/etc/about"
		log.Printf("ABOUT_DATA_DIR environment variable not set, using default value: %v",dataDir)
	}

	//
	// Server up API endpoints
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		about(w,r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

// /////////////////////////////////////////////////////////////////////////////
//
//	Functions
//
// /////////////////////////////////////////////////////////////////////////////

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

