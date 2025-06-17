package main

import (
	"net/http"
)

func main() {
	// Step 1: Create a new http.ServeMux
	mux := http.NewServeMux()

	// Step 2: Create a new http.Server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Step 3: Use the server's ListenAndServe method to start the server
	server.ListenAndServe()
}