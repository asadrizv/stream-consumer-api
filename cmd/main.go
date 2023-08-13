package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	streamConsumer "social-media-analyser/internal/stream-consumer"
	"time"
)

var streamer streamConsumer.StreamConsumerService

func main() {
	// Get the streaming URL and port from environment variables
	streamingURL := os.Getenv("STREAMING_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089" // Default port if not specified
	}

	if streamingURL == "" {
		fmt.Println("Please set the STREAMING_URL environment variable.")
		return
	}

	// Configure the Stream Consumer service
	streamer = streamConsumer.StreamConsumerService{StreamingURL: streamingURL}

	// Set up the HTTP server
	http.HandleFunc("/analysis", analysisHandler)
	serverAddr := ":" + port
	fmt.Printf("Server is listening on %s...\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func analysisHandler(w http.ResponseWriter, r *http.Request) {
	durationParam := r.URL.Query().Get("duration")
	dimensionParam := r.URL.Query().Get("dimension")

	duration, err := time.ParseDuration(durationParam)
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	if dimensionParam == "" {
		http.Error(w, "Invalid dimension", http.StatusBadRequest)
		return
	}

	stats, err := streamer.StreamForDuration(duration, dimensionParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := AnalysisResponse{
		TotalPosts:       stats.TotalPosts,
		MinimumTimestamp: stats.MinimumTimestamp,
		MaximumTimestamp: stats.MaxTimestamp,
		P50:              stats.P50,
		P90:              stats.P90,
		P99:              stats.P99,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type AnalysisResponse struct {
	TotalPosts       int     `json:"total_posts"`
	MinimumTimestamp float64 `json:"minimum_timestamp"`
	MaximumTimestamp float64 `json:"maximum_timestamp"`
	P50              float64 `json:"p50"`
	P90              float64 `json:"p90"`
	P99              float64 `json:"p99"`
}
