package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tonygilkerson/go-simple-rest/internal/acectx"
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

	http.HandleFunc("/ctx", func(w http.ResponseWriter, r *http.Request){
		aceCtxHandler(w,r,dataDir)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

// /////////////////////////////////////////////////////////////////////////////
//
//	Functions
//
// /////////////////////////////////////////////////////////////////////////////

func aceCtxHandler(w http.ResponseWriter, r *http.Request,dataDir string) {
	aceCtx := acectx.New()
	aceContextYamlFile := dataDir + "/ace-context.yaml"
	aceCtx.LoadAceContext(aceContextYamlFile)
	
	out := fmt.Sprintf("ACE Contex: %v",aceCtx)
	fmt.Fprint(w, out)

}

