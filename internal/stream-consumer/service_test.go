package stream_consumer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStreamForDuration(t *testing.T) {
	// Create a test SSE handler that generates data
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)

		// Generate test SSE data
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data: {\"instagram_media\": {\"timestamp\": 123456.789, \"likes\": %d}}\n\n", i*10)
			w.(http.Flusher).Flush() // Flush the response to simulate streaming
			time.Sleep(500 * time.Millisecond)
		}
	}))
	defer testServer.Close()

	svc := StreamConsumerService{StreamingURL: testServer.URL}

	duration := 5 * time.Second
	dimension := "likes"

	_, err := svc.StreamForDuration(duration, dimension)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
