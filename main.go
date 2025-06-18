package main

import (
	"net/http"
	"fmt"
	"sync/atomic"

)

type apiConfig struct {
	fileserverHits atomic.Int32
}

// middlewareMetricsInc increments the fileserverHits counter for each request
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// handlerMetrics returns the current hit count as plain text
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	hits := cfg.fileserverHits.Load()
	fmt.Fprintf(w, "Hits: %d\n", hits)
}

// handlerReset resets the hit counter back to 0
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Hits reset to 0\n")
}

func main() {
	// Step 1: Create a new http.ServeMux
	mux := http.NewServeMux()

	// Add readiness endpoint at /healthz
	// mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
	// 	// 1. Write the Content-Type header
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	// 2. Write the status code
	// 	w.WriteHeader(http.StatusOK)
	// 	// 3. Write the body text
	// 	w.Write([]byte("OK"))
	// })

	
	apiCfg := &apiConfig{}

	// Create serve mux
	// mux := http.NewServeMux()

	// File server handler with metrics middleware
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	// Metrics and reset endpoints
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	// Step 3: Use the server's ListenAndServe method to start the server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}