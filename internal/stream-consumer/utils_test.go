package stream_consumer

import (
	"testing"
)

func TestCalculatePercentile(t *testing.T) {
	numbers := []float64{10, 20, 30, 40, 50}

	p50 := calculatePercentile(numbers, 50)
	if p50 != 30 {
		t.Errorf("Expected P50 to be %f, but got %f", 30.0, p50)
	}

	p90 := calculatePercentile(numbers, 90)
	if p90 != 50 {
		t.Errorf("Expected P90 to be %f, but got %f", 50.0, p90)
	}

	p99 := calculatePercentile(numbers, 99)
	if p99 != 50 {
		t.Errorf("Expected P99 to be %f, but got %f", 50.0, p99)
	}
}

func TestCalculatePercentileEmpty(t *testing.T) {
	var emptyNumbers []float64
	result := calculatePercentile(emptyNumbers, 50)
	if result != 0 {
		t.Errorf("Expected result to be %f for empty slice, but got %f", 0.0, result)
	}
}
