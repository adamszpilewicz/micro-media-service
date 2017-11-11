package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Version encapsulates the version number of the movie microservice
type Version struct {
	Version string `json:"version"`
}

var version Version

func init() {
	version = Version{"0.1.0"}
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	verJSON, err := json.Marshal(version)
	if err != nil {
		panic("Error marshaling version")
	}
	log.Printf("Request for version, the response is %+v", string(verJSON))
	fmt.Fprintf(w, string(verJSON))
}
