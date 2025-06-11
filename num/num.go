// Package num provides utility functions for number manipulation.
package num

import (
	"math"
	"math/rand/v2"
)

// Clamp clamps a number between lower and upper bounds.
// Example: Clamp(10, 0, 5) -> 5
func Clamp(n, lower, upper float64) float64 {
	return math.Min(math.Max(n, lower), upper)
}

// InRange checks if a number is between start and end (inclusive).
// Example: InRange(3, 2, 4) -> true
func InRange(n, start, end float64) bool {
	if start > end {
		start, end = end, start
	}
	return n >= start && n <= end
}

// Random returns a random number between min and max (inclusive).
// Example: Random(1, 10) -> a random number between 1 and 10
func Random(min, max int) int {
	if min > max {
		min, max = max, min // Swap if min > max
	}

	if min == max {
		return min
	}

	// Use time-based seed for better randomness
	return rand.IntN(max-min+1) + min
}

// Round rounds a number to the nearest integer or to the specified precision.
// Example: Round(4.7) -> 5, Round(4.7, 1) -> 4.7
func Round(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Round(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Round(n*p) / p
}

// Floor rounds a number down to the nearest integer or to the specified precision.
// Example: Floor(4.7) -> 4, Floor(4.78, 1) -> 4.7
func Floor(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Floor(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Floor(n*p) / p
}

// Ceil rounds a number up to the nearest integer or to the specified precision.
// Example: Ceil(4.3) -> 5, Ceil(4.78, 1) -> 4.8
func Ceil(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Ceil(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Ceil(n*p) / p
}

// Max returns the maximum of the given numbers.
// Example: Max(1, 2, 3) -> 3
func Max(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	m := numbers[0]
	for _, n := range numbers[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

// MaxBy returns the maximum value of a slice using the provided iteratee function.
// Example: MaxBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * n) }) -> 3
func MaxBy[T any](collection []T, iteratee func(T) float64) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}

	maxIdx := 0
	maxValue := iteratee(collection[0])

	for i := 1; i < len(collection); i++ {
		value := iteratee(collection[i])
		if value > maxValue {
			maxValue = value
			maxIdx = i
		}
	}

	return collection[maxIdx]
}

// Min returns the minimum of the given numbers.
// Example: Min(1, 2, 3) -> 1
func Min(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	m := numbers[0]
	for _, n := range numbers[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

// MinBy returns the minimum value of a slice using the provided iteratee function.
// Example: MinBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * n) }) -> 1
func MinBy[T any](collection []T, iteratee func(T) float64) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}

	minIdx := 0
	minValue := iteratee(collection[0])

	for i := 1; i < len(collection); i++ {
		value := iteratee(collection[i])
		if value < minValue {
			minValue = value
			minIdx = i
		}
	}

	return collection[minIdx]
}

// Sum returns the sum of the given numbers.
// Example: Sum(1, 2, 3) -> 6
func Sum(numbers ...float64) float64 {
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// SumBy returns the sum of values in a slice after applying the iteratee function.
// Example: SumBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * 2) }) -> 12
func SumBy[T any](collection []T, iteratee func(T) float64) float64 {
	var sum float64
	for _, item := range collection {
		sum += iteratee(item)
	}
	return sum
}

// Mean returns the arithmetic mean of the given numbers.
// Example: Mean(1, 2, 3) -> 2
func Mean(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return Sum(numbers...) / float64(len(numbers))
}

// MeanBy returns the mean of values in a slice after applying the iteratee function.
// Example: MeanBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * 2) }) -> 4
func MeanBy[T any](collection []T, iteratee func(T) float64) float64 {
	if len(collection) == 0 {
		return 0
	}

	return SumBy(collection, iteratee) / float64(len(collection))
}

// Abs returns the absolute value of a number.
// Example: Abs(-5) -> 5
func Abs(n float64) float64 {
	return math.Abs(n)
}

// Pow returns the base raised to the exponent power.
// Example: Pow(2, 3) -> 8
func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Sqrt returns the square root of a number.
// Example: Sqrt(9) -> 3
func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}
