package main

import (
	"log"
	"net/http"

	"round-robin-reverse-proxy/proxy"
)

func main() {
	// List of backend servers
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	// Create a reverse proxy instance
	rp := proxy.NewReverseProxy(backends)

	// Start HTTP server
	log.Println("Reverse proxy listening on :8000")
	err := http.ListenAndServe(":8000", rp)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
