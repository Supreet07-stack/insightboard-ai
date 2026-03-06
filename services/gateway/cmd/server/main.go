package main

import (
	"log"
	"net/http"

	"github.com/Supreet07-stack/insightboard-ai/services/gateway/internal/http/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)

	addr := ":8080"

	log.Printf("gateway service starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}