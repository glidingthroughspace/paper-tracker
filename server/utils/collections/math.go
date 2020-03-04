package collections

// various helpers for collections, targeting statistcs

import "sort"

func MeanOf(values ...int) float64 {
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
	middleIndex := len(values) / 2

	sort.Ints(values)

	if IsOddAmountOfValues(values...) {
		return float64(values[middleIndex])
	}

	return (float64(values[middleIndex-1]) + float64(values[middleIndex])) / 2.0
}

func FirstQuartileOf(values ...int) float64 {
	sort.Ints(values)

	return MedianOf(values[:len(values)-1]...)
}

func ThirdQuartileOf(values ...int) float64 {
	sort.Ints(values)

	return MedianOf(values[len(values)-1:]...)
}

func MinOf(values ...int) int {
	sort.Ints(values)
	return values[0]
}

func MaxOf(values ...int) int {
	sort.Ints(values)
	return values[len(values)-1]
}
