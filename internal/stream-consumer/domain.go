package stream_consumer

type StreamEvents map[string]map[string]interface{}

type MetricsStatistics struct {
	TotalPosts       int
	MinimumTimestamp float64
	MaxTimestamp     float64
	Dimensions       []float64
	P99              float64
	P90              float64
	P50              float64
}
