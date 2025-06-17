package main

import (
	"net/http"
)

func main() {
	// Step 1: Create a new http.ServeMux
	mux := http.NewServeMux()

	// Add readiness endpoint at /healthz
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// 1. Write the Content-Type header
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// 2. Write the status code
		w.WriteHeader(http.StatusOK)
		// 3. Write the body text
		w.Write([]byte("OK"))
	})

	// Use http.FileServer to serve files from current directory at /app/ path
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	// Step 2: Create a new http.Server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Step 3: Use the server's ListenAndServe method to start the server
	server.ListenAndServe()
}