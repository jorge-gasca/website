// Simple HTTP server

package main

import (
	"net/http"
	"website/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
