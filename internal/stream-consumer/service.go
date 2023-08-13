package stream_consumer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
)

type StreamConsumerService struct {
	StreamingURL string
}

func (svc *StreamConsumerService) StreamForDuration(duration time.Duration,
	dimension string) (MetricsStatistics, error) {
	var dimensions []float64
	maxTimestamp := 0.0
	stats := MetricsStatistics{}
	totalPosts := 0
	minTimestamp := math.MaxFloat64

	resp, err := http.Get(svc.StreamingURL)
	if err != nil {
		fmt.Println("Error connecting to the stream:", err)
		return stats, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Stream server returned status:", resp.Status)
		return stats, err
	}

	scanner := bufio.NewScanner(resp.Body)
	timer := time.NewTimer(duration)
	defer timer.Stop()

Loop:
	for scanner.Scan() {
		select {
		case <-timer.C:
			stats = MetricsStatistics{
				MaxTimestamp:     maxTimestamp,
				MinimumTimestamp: minTimestamp,
				TotalPosts:       totalPosts,
				Dimensions:       dimensions,
				P99:              calculatePercentile(dimensions, 99),
				P90:              calculatePercentile(dimensions, 90),
				P50:              calculatePercentile(dimensions, 50),
			}
			break Loop
		default:
			data := scanner.Text() // Read the line as a string

			// Remove the "data: " prefix
			data = strings.TrimPrefix(data, "data:")

			// Unmarshal the JSON data from the line
			var event StreamEvents
			if err := json.Unmarshal([]byte(data), &event); err != nil {
				continue
			}

			for _, v := range event {
				if v["timestamp"].(float64) > maxTimestamp {
					maxTimestamp = v["timestamp"].(float64)
				}

				if v["timestamp"].(float64) < minTimestamp {
					minTimestamp = v["timestamp"].(float64)
				}

				if dim, ok := v[dimension]; ok {
					totalPosts++
					dimensions = append(dimensions, dim.(float64))
				}

			}

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stream:", err)
	}

	return stats, nil
}
