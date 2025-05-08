package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8082" 

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend on port %s!\n", port)
	})

	log.Printf("Backend server listening on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
