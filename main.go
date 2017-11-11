package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8086

func main() {
	http.HandleFunc("/api/v1/version", handleVersion)
	http.HandleFunc("/api/v1/movies", handleMovies)
	http.HandleFunc("/api/v1/movies/", handleMovieFromID)

	log.Printf("Starting movies microservice on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
