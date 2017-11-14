package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const port = 8086

var verboseFlag bool

func init() {
	flag.BoolVar(&verboseFlag, "verbose", false, "Enable verbose mode")
	flag.Parse()
	if verboseEnv := strings.ToLower(os.Getenv("MOVIE_VERBOSE")); verboseEnv == "true" || verboseFlag {
		log.SetLevel(log.DebugLevel)
	}
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
}

func main() {
	http.HandleFunc("/api/v1/version", handleVersion)
	http.HandleFunc("/api/v1/movies", handleMovies)
	http.HandleFunc("/api/v1/movies/", handleMovieFromID)

	log.Printf("Starting movies microservice on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
