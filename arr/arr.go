// Package arr provides utility functions for array/slice manipulation.
package arr

import (
	"github.com/gflydev/utils/num"
	"github.com/gflydev/utils/str"
	"math/rand/v2"
	"sort"
)

// Chunk splits an array into groups of the specified size.
// Example: Chunk([]int{1, 2, 3, 4}, 2) -> [][]int{{1, 2}, {3, 4}}
func Chunk[T any](array []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	length := len(array)
	chunks := make([][]T, 0, (length+size-1)/size)

	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		chunks = append(chunks, array[i:end])
	}

	return chunks
}

// Compact removes falsey values from an array.
// In Go, we consider nil, zero values, and empty collections as falsey.
// Example: Compact([]int{0, 1, 2, 0, 3}) -> []int{1, 2, 3}
func Compact[T comparable](array []T) []T {
	var zero T
	result := make([]T, 0)

	for _, v := range array {
		if v != zero {
			result = append(result, v)
		}
	}

	return result
}

// Concat concatenates arrays together.
// Example: Concat([]int{1, 2}, []int{3, 4}) -> []int{1, 2, 3, 4}
func Concat[T any](arrays ...[]T) []T {
	var totalLen int
	for _, arr := range arrays {
		totalLen += len(arr)
	}

	result := make([]T, 0, totalLen)
	for _, arr := range arrays {
		result = append(result, arr...)
	}

	return result
}

// Difference returns an array of elements that are in the first array but not in the others.
// Example: Difference([]int{1, 2, 3}, []int{2, 3, 4}) -> []int{1}
func Difference[T comparable](array []T, others ...[]T) []T {
	if len(array) == 0 {
		return []T{}
	}

	// Create a map of all values in others
	excludeMap := make(map[T]bool)
	for _, other := range others {
		for _, v := range other {
			excludeMap[v] = true
		}
	}

	result := make([]T, 0)
	for _, v := range array {
		if !excludeMap[v] {
			result = append(result, v)
		}
	}

	return result
}

// Drop creates a slice with n elements dropped from the beginning.
// Example: Drop([]int{1, 2, 3, 4}, 2) -> []int{3, 4}
func Drop[T any](array []T, n int) []T {
	if n <= 0 {
		return array
	}
	if n >= len(array) {
		return []T{}
	}
	return array[n:]
}

// DropRight creates a slice with n elements dropped from the end.
// Example: DropRight([]int{1, 2, 3, 4}, 2) -> []int{1, 2}
func DropRight[T any](array []T, n int) []T {
	if n <= 0 {
		return array
	}
	if n >= len(array) {
		return []T{}
	}
	return array[:len(array)-n]
}

// Fill fills elements of array with value from start up to, but not including, end.
// Example: Fill([]int{1, 2, 3, 4}, 0, 1, 3) -> []int{1, 0, 0, 4}
func Fill[T any](array []T, value T, start, end int) []T {
	length := len(array)
	if length == 0 {
		return array
	}

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return array
	}

	result := make([]T, length)
	copy(result, array)

	for i := start; i < end; i++ {
		result[i] = value
	}

	return result
}

// FindIndex returns the index of the first element that satisfies the predicate function.
// Example: FindIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 2
func FindIndex[T any](array []T, predicate func(T) bool) int {
	for i, v := range array {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element that satisfies the predicate function.
// Example: FindLastIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 3
func FindLastIndex[T any](array []T, predicate func(T) bool) int {
	for i := len(array) - 1; i >= 0; i-- {
		if predicate(array[i]) {
			return i
		}
	}
	return -1
}

// First returns the first element of an array.
// Example: First([]int{1, 2, 3}) -> 1
func First[T any](array []T) (T, bool) {
	var zero T
	if len(array) == 0 {
		return zero, false
	}
	return array[0], true
}

// Flatten flattens an array a single level deep.
// Example: Flatten([][]int{{1, 2}, {3, 4}}) -> []int{1, 2, 3, 4}
func Flatten[T any](array [][]T) []T {
	var totalLen int
	for _, arr := range array {
		totalLen += len(arr)
	}

	result := make([]T, 0, totalLen)
	for _, arr := range array {
		result = append(result, arr...)
	}

	return result
}

// Includes checks if a value is in the array.
// Example: Includes([]int{1, 2, 3}, 2) -> true
func Includes[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of value in array.
// Example: IndexOf([]int{1, 2, 3, 2}, 2) -> 1
func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of value in array.
// Example: LastIndexOf([]int{1, 2, 3, 2}, 2) -> 3
func LastIndexOf[T comparable](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Initial returns all but the last element of an array.
// Example: Initial([]int{1, 2, 3}) -> []int{1, 2}
func Initial[T any](array []T) []T {
	if len(array) <= 1 {
		return []T{}
	}
	return array[:len(array)-1]
}

// Intersection returns an array of unique values that are included in all given arrays.
// Example: Intersection([]int{1, 2, 3}, []int{2, 3, 4}) -> []int{2, 3}
func Intersection[T comparable](arrays ...[]T) []T {
	if len(arrays) == 0 {
		return []T{}
	}
	if len(arrays) == 1 {
		return Uniq(arrays[0])
	}

	// Count occurrences of each value
	counts := make(map[T]int)
	for _, arr := range arrays {
		// Use a set to count each value only once per array
		seen := make(map[T]bool)
		for _, v := range arr {
			if !seen[v] {
				counts[v]++
				seen[v] = true
			}
		}
	}

	// Keep values that appear in all arrays
	result := make([]T, 0)
	for v, count := range counts {
		if count == len(arrays) {
			result = append(result, v)
		}
	}

	return result
}

// Join joins all elements of an array into a string.
// Example: Join([]int{1, 2, 3}, ",") -> "1,2,3"
func Join[T any](array []T, separator string) string {
	if len(array) == 0 {
		return ""
	}

	// Convert each element to string
	strArr := make([]string, len(array))
	for i, v := range array {
		// Have to use str.ToString to get real data instead of use reflect.ValueOf(v).String()
		// to get key of <int value>, <string value>,....
		strArr[i] = str.ToString(v)
	}

	// Join the strings
	var result string
	for i, s := range strArr {
		if i > 0 {
			result += separator
		}
		result += s
	}

	return result
}

// Last returns the last element of an array.
// Example: Last([]int{1, 2, 3}) -> 3
func Last[T any](array []T) (T, bool) {
	var zero T
	if len(array) == 0 {
		return zero, false
	}
	return array[len(array)-1], true
}

// Nth returns the element at index n of array.
// If n is negative, the nth element from the end is returned.
// Example: Nth([]int{1, 2, 3}, 1) -> 2
func Nth[T any](array []T, n int) (T, bool) {
	var zero T
	length := len(array)
	if length == 0 {
		return zero, false
	}

	if n < 0 {
		n = length + n
	}

	if n < 0 || n >= length {
		return zero, false
	}

	return array[n], true
}

// Pull removes all given values from array.
// Example: Pull([]int{1, 2, 3, 1, 2, 3}, 2, 3) -> []int{1, 1}
func Pull[T comparable](array []T, values ...T) []T {
	if len(array) == 0 || len(values) == 0 {
		return array
	}

	// Create a map of values to remove
	removeMap := make(map[T]bool)
	for _, v := range values {
		removeMap[v] = true
	}

	result := make([]T, 0)
	for _, v := range array {
		if !removeMap[v] {
			result = append(result, v)
		}
	}

	return result
}

// Reverse reverses the order of elements in array.
// Example: Reverse([]int{1, 2, 3}) -> []int{3, 2, 1}
func Reverse[T any](array []T) []T {
	length := len(array)
	if length <= 1 {
		return array
	}

	result := make([]T, length)
	copy(result, array)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Shuffle returns a new slice with elements in random order.
// Example: Shuffle([]int{1, 2, 3, 4, 5}) -> [3, 1, 5, 2, 4]
func Shuffle[T any](slice []T) []T {
	if len(slice) <= 1 {
		result := make([]T, len(slice))
		copy(result, slice)
		return result
	}

	// Create a copy to avoid modifying original
	result := make([]T, len(slice))
	copy(result, slice)

	// Fisher-Yates shuffle
	for i := len(result) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Random returns n random elements from the given slice without replacement.
// Example: Random([]int{1, 2, 3, 4, 5}, 3) -> [2, 4, 1]
func Random[T any](slice []T, n int) []T {
	if n <= 0 || len(slice) == 0 {
		return []T{}
	}

	if n >= len(slice) {
		// Return shuffled copy of entire slice
		return Shuffle(slice)
	}

	// Create a copy to avoid modifying original
	copied := make([]T, len(slice))
	copy(copied, slice)

	// Shuffle and take first n elements
	shuffled := Shuffle(copied)
	return shuffled[:n]
}

// RandomChoice returns a random element from the given slice.
// Example: RandomChoice([]string{"a", "b", "c"}) -> "b"
func RandomChoice[T any](choices []T) (T, bool) {
	var zero T
	if len(choices) == 0 {
		return zero, false
	}

	if len(choices) == 1 {
		return choices[0], true
	}

	index := num.Random(0, len(choices)-1)
	return choices[index], true
}

// Slice returns a slice of array from start up to, but not including, end.
// Example: Slice([]int{1, 2, 3, 4}, 1, 3) -> []int{2, 3}
func Slice[T any](array []T, start, end int) []T {
	length := len(array)
	if length == 0 {
		return []T{}
	}

	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return []T{}
	}

	return array[start:end]
}

// SortedIndex returns the index at which value should be inserted into array to maintain its sort order.
// Example: SortedIndex([]int{1, 3, 5, 7}, 4) -> 2
func SortedIndex[T int | int8 | int16 | int32 | int64 | float32 | float64](array []T, value T) int {
	for i, v := range array {
		if v >= value {
			return i
		}
	}
	return len(array)
}

// Tail returns all but the first element of array.
// Example: Tail([]int{1, 2, 3}) -> []int{2, 3}
func Tail[T any](array []T) []T {
	if len(array) <= 1 {
		return []T{}
	}
	return array[1:]
}

// Take creates a slice of array with n elements taken from the beginning.
// Example: Take([]int{1, 2, 3, 4}, 2) -> []int{1, 2}
func Take[T any](array []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(array) {
		return array
	}
	return array[:n]
}

// TakeRight creates a slice of array with n elements taken from the end.
// Example: TakeRight([]int{1, 2, 3, 4}, 2) -> []int{3, 4}
func TakeRight[T any](array []T, n int) []T {
	length := len(array)
	if n <= 0 {
		return []T{}
	}
	if n >= length {
		return array
	}
	return array[length-n:]
}

// Union creates an array of unique values from all given arrays.
// Example: Union([]int{1, 2}, []int{2, 3}) -> []int{1, 2, 3}
func Union[T comparable](arrays ...[]T) []T {
	if len(arrays) == 0 {
		return []T{}
	}

	// Use a map to track unique values
	uniqueMap := make(map[T]bool)
	for _, arr := range arrays {
		for _, v := range arr {
			uniqueMap[v] = true
		}
	}

	// Convert map keys to slice
	result := make([]T, 0, len(uniqueMap))
	for v := range uniqueMap {
		result = append(result, v)
	}

	return result
}

// Uniq creates an array of unique values.
// Example: Uniq([]int{1, 2, 1, 3}) -> []int{1, 2, 3}
func Uniq[T comparable](array []T) []T {
	if len(array) <= 1 {
		return array
	}

	// Use a map to track unique values
	uniqueMap := make(map[T]bool)
	result := make([]T, 0)

	for _, v := range array {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Without creates an array excluding all given values.
// Example: Without([]int{1, 2, 3, 4}, 2, 4) -> []int{1, 3}
func Without[T comparable](array []T, values ...T) []T {
	return Pull(array, values...)
}

// Zip creates an array of grouped elements.
// Example: Zip([]int{1, 2}, []string{"a", "b"}) -> [][]interface{}{{1, "a"}, {2, "b"}}
func Zip[T any](arrays ...[]T) [][]T {
	if len(arrays) == 0 {
		return [][]T{}
	}

	// Find the minimum length of all arrays
	minLen := len(arrays[0])
	for _, arr := range arrays[1:] {
		if len(arr) < minLen {
			minLen = len(arr)
		}
	}

	if minLen == 0 {
		return [][]T{}
	}

	// Create the result
	result := make([][]T, minLen)
	for i := 0; i < minLen; i++ {
		result[i] = make([]T, len(arrays))
		for j, arr := range arrays {
			result[i][j] = arr[i]
		}
	}

	return result
}

// SortBy sorts an array by the results of running each element through iteratee.
// Example: SortBy([]int{1, 3, 2}, func(n int) int { return n })
func SortBy[T any, U int | int8 | int16 | int32 | int64 | float32 | float64 | string](array []T, iteratee func(T) U) []T {
	if len(array) <= 1 {
		return array
	}

	result := make([]T, len(array))
	copy(result, array)

	sort.Slice(result, func(i, j int) bool {
		return iteratee(result[i]) < iteratee(result[j])
	})

	return result
}
