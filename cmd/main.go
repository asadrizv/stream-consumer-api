package main

import (
	"log"
	"os"
	"stream-consumer-api/api"
)

func main() {
	streamingURL := os.Getenv("STREAMING_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	if streamingURL == "" {
		log.Fatal("Please set the STREAMING_URL environment variable.")
	}

	server := api.NewServer(port, streamingURL)
	server.Start()
}
