package collections

// various helpers for collections, targeting statistcs

import (
	"sort"

	"github.com/montanaflynn/stats"
)

func MeanOf(values ...int) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum float64 = 0.0
	for _, v := range values {
		sum += float64(v)
	}
	return sum / float64(len(values))
}

func IsOddAmountOfValues(values ...int) bool {
	return len(values)%2 == 1
}

func MedianOf(values ...int) float64 {
	data := stats.LoadRawData(values)
	median, err := stats.Median(data)
	// TODO: Do we really want to do this?
	if err != nil {
		return 0
	}
	return median
}

func FirstQuartileOf(values ...int) float64 {
	// This is a special case for our use and is not "standard math"
	if len(values) == 1 {
		return float64(values[0])
	}
	data := stats.LoadRawData(values)
	quartiles, err := stats.Quartile(data)
	// TODO: Do we really want to do this?
	if err != nil {
		return 0
	}
	return quartiles.Q1
}

func ThirdQuartileOf(values ...int) float64 {
	// This is a special case for our use and is not "standard math"
	if len(values) == 1 {
		return float64(values[0])
	}
	data := stats.LoadRawData(values)
	quartiles, err := stats.Quartile(data)
	// TODO: Do we really want to do this?
	if err != nil {
		return 0
	}
	return quartiles.Q3
}

func MinOf(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	sort.Ints(values)
	return values[0]
}

func MaxOf(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	sort.Ints(values)
	return values[len(values)-1]
}
