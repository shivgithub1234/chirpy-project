package main

import (
	"net/http"
)

func main() {
	// Step 1: Create a new http.ServeMux
	mux := http.NewServeMux()

	// Serve a specific HTML file for the root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html") // Replace with your filename
		}else if r.URL.Path == "/assets/logo.png" {
			http.ServeFile(w, r, "img.png")
		}else {
			http.NotFound(w, r)
		}
	})

	// Optional: Still serve other static files from /static/ path
	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	// Step 2: Create a new http.Server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Step 3: Use the server's ListenAndServe method to start the server
	server.ListenAndServe()
}