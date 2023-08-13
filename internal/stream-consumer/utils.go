package stream_consumer

import "sort"

func calculatePercentile(numbers []float64, percentile int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Float64s(numbers)
	index := int(float64(percentile) / 100 * float64(len(numbers)-1))

	if index == len(numbers)-1 {
		return numbers[len(numbers)-1]
	}

	lowerValue := numbers[index]
	upperValue := numbers[index+1]

	fraction := float64(percentile%100) / 100
	percentileValue := lowerValue + fraction*(upperValue-lowerValue)
	return percentileValue
}
