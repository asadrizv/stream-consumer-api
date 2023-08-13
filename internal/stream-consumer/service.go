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

func (svc *StreamConsumerService) StreamForDuration(duration time.Duration, dimension string) (MetricsStatistics, error) {
	stats := MetricsStatistics{}
	dimensions := make([]float64, 0)
	maxTimestamp := 0.0
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
			break Loop
		default:
			data := strings.TrimPrefix(scanner.Text(), "data:")

			var event StreamEvents
			if err := json.Unmarshal([]byte(data), &event); err != nil {
				continue
			}

			for _, v := range event {
				timestamp := v["timestamp"].(float64)
				maxTimestamp = math.Max(maxTimestamp, timestamp)
				minTimestamp = math.Min(minTimestamp, timestamp)

				if dim, ok := v[dimension]; ok {
					dimensions = append(dimensions, dim.(float64))
				}
				stats.TotalPosts++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stream:", err)
	}

	stats.MaxTimestamp = maxTimestamp
	stats.MinimumTimestamp = minTimestamp
	stats.Dimensions = dimensions
	stats.P99 = calculatePercentile(dimensions, 99)
	stats.P90 = calculatePercentile(dimensions, 90)
	stats.P50 = calculatePercentile(dimensions, 50)

	return stats, nil
}
