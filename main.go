package main

import (
	"log"
	"net/http"

	ascii "ascii/functions"
)

func main() {
	http.HandleFunc("/", ascii.HomeHandler)
	http.HandleFunc("/ascii-art", ascii.ArtHandler)

	http.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./stylesheets/style.css")
	})
	http.HandleFunc("/css/error.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./stylesheets/error.css")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
