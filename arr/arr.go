// Package arr provides utility functions for array/slice manipulation.
package arr

import (
	"fmt"
	"github.com/gflydev/utils/num"
	"github.com/gflydev/utils/str"
	"math/rand/v2"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

// Chunk splits an array into groups of the specified size.
//
// Parameters:
//   - array: The array to split into chunks
//   - size: The size of each chunk
//
// Returns:
//   - [][]T: A new array containing chunks of the original array
//
// Example:
//
//	Chunk([]int{1, 2, 3, 4}, 2) -> [][]int{{1, 2}, {3, 4}}
//	Chunk([]string{"a", "b", "c", "d", "e"}, 2) -> [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
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
//
// Parameters:
//   - array: The array to compact
//
// Returns:
//   - []T: A new array with all falsey values removed
//
// Example:
//
//	Compact([]int{0, 1, 2, 0, 3}) -> []int{1, 2, 3}
//	Compact([]string{"", "a", "", "b"}) -> []string{"a", "b"}
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
//
// Parameters:
//   - arrays: Variable number of arrays to concatenate
//
// Returns:
//   - []T: A new array containing all elements from the input arrays
//
// Example:
//
//	Concat([]int{1, 2}, []int{3, 4}) -> []int{1, 2, 3, 4}
//	Concat([]string{"a", "b"}, []string{"c"}, []string{"d", "e"}) -> []string{"a", "b", "c", "d", "e"}
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
//
// Parameters:
//   - array: The base array to compare against
//   - others: Variable number of arrays to compare with the base array
//
// Returns:
//   - []T: A new array containing elements that are in the base array but not in any of the other arrays
//
// Example:
//
//	Difference([]int{1, 2, 3}, []int{2, 3, 4}) -> []int{1}
//	Difference([]string{"a", "b", "c"}, []string{"b"}, []string{"c", "d"}) -> []string{"a"}
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
//
// Parameters:
//   - array: The input array
//   - n: Number of elements to drop from the beginning
//
// Returns:
//   - []T: A new array with the first n elements removed
//
// Example:
//
//	Drop([]int{1, 2, 3, 4}, 2) -> []int{3, 4}
//	Drop([]string{"a", "b", "c"}, 1) -> []string{"b", "c"}
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
//
// Parameters:
//   - array: The input array
//   - n: Number of elements to drop from the end
//
// Returns:
//   - []T: A new array with the last n elements removed
//
// Example:
//
//	DropRight([]int{1, 2, 3, 4}, 2) -> []int{1, 2}
//	DropRight([]string{"a", "b", "c"}, 1) -> []string{"a", "b"}
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
//
// Parameters:
//   - array: The input array
//   - value: The value to fill the array with
//   - start: The starting index (inclusive)
//   - end: The ending index (exclusive)
//
// Returns:
//   - []T: A new array with elements filled with the specified value
//
// Example:
//
//	Fill([]int{1, 2, 3, 4}, 0, 1, 3) -> []int{1, 0, 0, 4}
//	Fill([]string{"a", "b", "c", "d"}, "x", 0, 2) -> []string{"x", "x", "c", "d"}
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
//
// Parameters:
//   - array: The input array
//   - predicate: A function that returns true for elements that satisfy the condition
//
// Returns:
//   - int: The index of the first element that satisfies the predicate, or -1 if none found
//
// Example:
//
//	FindIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 2
//	FindIndex([]string{"a", "b", "c"}, func(s string) bool { return s == "b" }) -> 1
func FindIndex[T any](array []T, predicate func(T) bool) int {
	for i, v := range array {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// FindLastIndex returns the index of the last element that satisfies the predicate function.
//
// Parameters:
//   - array: The input array
//   - predicate: A function that returns true for elements that satisfy the condition
//
// Returns:
//   - int: The index of the last element that satisfies the predicate, or -1 if none found
//
// Example:
//
//	FindLastIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 3
//	FindLastIndex([]string{"a", "b", "c", "b"}, func(s string) bool { return s == "b" }) -> 3
func FindLastIndex[T any](array []T, predicate func(T) bool) int {
	for i := len(array) - 1; i >= 0; i-- {
		if predicate(array[i]) {
			return i
		}
	}
	return -1
}

// First returns the first element of an array.
//
// Parameters:
//   - array: The input array
//
// Returns:
//   - T: The first element of the array
//   - bool: True if the array is not empty, false otherwise
//
// Example:
//
//	First([]int{1, 2, 3}) -> 1, true
//	First([]string{"a", "b"}) -> "a", true
//	First([]int{}) -> 0, false
func First[T any](array []T) (T, bool) {
	var zero T
	if len(array) == 0 {
		return zero, false
	}
	return array[0], true
}

// First returns the first element in an array passing a given truth test
/*func First[T any](array []T, callback func(T) bool) (T, bool) {
	for _, value := range array {
		if callback(value) {
			return value, true
		}
	}

	var zero T
	return zero, false
}*/

// Flatten flattens an array a single level deep.
//
// Parameters:
//   - array: The nested array to flatten
//
// Returns:
//   - []T: A new array with all nested elements combined into a single level
//
// Example:
//
//	Flatten([][]int{{1, 2}, {3, 4}}) -> []int{1, 2, 3, 4}
//	Flatten([][]string{{"a", "b"}, {"c"}}) -> []string{"a", "b", "c"}
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
//
// Parameters:
//   - array: The array to search in
//   - value: The value to search for
//
// Returns:
//   - bool: True if the value is found in the array, false otherwise
//
// Example:
//
//	Includes([]int{1, 2, 3}, 2) -> true
//	Includes([]string{"a", "b", "c"}, "d") -> false
func Includes[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of value in array.
//
// Parameters:
//   - array: The array to search in
//   - value: The value to search for
//
// Returns:
//   - int: The index of the first occurrence of the value, or -1 if not found
//
// Example:
//
//	IndexOf([]int{1, 2, 3, 2}, 2) -> 1
//	IndexOf([]string{"a", "b", "c"}, "c") -> 2
//	IndexOf([]int{1, 2, 3}, 4) -> -1
func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of value in array.
//
// Parameters:
//   - array: The array to search in
//   - value: The value to search for
//
// Returns:
//   - int: The index of the last occurrence of the value, or -1 if not found
//
// Example:
//
//	LastIndexOf([]int{1, 2, 3, 2}, 2) -> 3
//	LastIndexOf([]string{"a", "b", "c", "b"}, "b") -> 3
//	LastIndexOf([]int{1, 2, 3}, 4) -> -1
func LastIndexOf[T comparable](array []T, value T) int {
	for i := len(array) - 1; i >= 0; i-- {
		if array[i] == value {
			return i
		}
	}
	return -1
}

// Initial returns all but the last element of an array.
//
// Parameters:
//   - array: The input array
//
// Returns:
//   - []T: A new array containing all elements except the last one
//
// Example:
//
//	Initial([]int{1, 2, 3}) -> []int{1, 2}
//	Initial([]string{"a", "b", "c"}) -> []string{"a", "b"}
//	Initial([]int{1}) -> []int{}
func Initial[T any](array []T) []T {
	if len(array) <= 1 {
		return []T{}
	}
	return array[:len(array)-1]
}

// Intersection returns an array of unique values that are included in all given arrays.
//
// Parameters:
//   - arrays: Variable number of arrays to find common elements from
//
// Returns:
//   - []T: A new array containing elements that exist in all input arrays
//
// Example:
//
//	Intersection([]int{1, 2, 3}, []int{2, 3, 4}) -> []int{2, 3}
//	Intersection([]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "e"}) -> []string{"b"}
//	Intersection([]int{1, 2}, []int{3, 4}) -> []int{}
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
//
// Parameters:
//   - array: The array of elements to join
//   - separator: The string to insert between elements
//
// Returns:
//   - string: A string containing all array elements joined with the separator
//
// Example:
//
//	Join([]int{1, 2, 3}, ",") -> "1,2,3"
//	Join([]string{"a", "b", "c"}, "-") -> "a-b-c"
//	Join([]bool{true, false}, " and ") -> "true and false"
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
//
// Parameters:
//   - array: The input array
//
// Returns:
//   - T: The last element of the array
//   - bool: True if the array is not empty, false otherwise
//
// Example:
//
//	Last([]int{1, 2, 3}) -> 3, true
//	Last([]string{"a", "b"}) -> "b", true
//	Last([]int{}) -> 0, false
func Last[T any](array []T) (T, bool) {
	var zero T
	if len(array) == 0 {
		return zero, false
	}
	return array[len(array)-1], true
}

// Last returns the last element in an array passing a given truth test
/*func Last[T any](array []T, callback func(T) bool) (T, bool) {
	for i := len(array) - 1; i >= 0; i-- {
		if callback(array[i]) {
			return array[i], true
		}
	}

	var zero T
	return zero, false
}*/

// Nth returns the element at index n of array.
// If n is negative, the nth element from the end is returned.
//
// Parameters:
//   - array: The input array
//   - n: The index of the element to retrieve (can be negative)
//
// Returns:
//   - T: The element at the specified index
//   - bool: True if a valid element was found, false otherwise
//
// Example:
//
//	Nth([]int{1, 2, 3}, 1) -> 2, true
//	Nth([]string{"a", "b", "c"}, -1) -> "c", true
//	Nth([]int{1, 2, 3}, 5) -> 0, false
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
//
// Parameters:
//   - array: The input array
//   - values: Variable number of values to remove from the array
//
// Returns:
//   - []T: A new array with all specified values removed
//
// Example:
//
//	Pull([]int{1, 2, 3, 1, 2, 3}, 2, 3) -> []int{1, 1}
//	Pull([]string{"a", "b", "c", "a"}, "a") -> []string{"b", "c"}
//	Pull([]int{1, 2, 3}, 4) -> []int{1, 2, 3}
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

// Pull removes and returns an item from the array by key
/*func Pull[T any](array []T, index int) (T, []T) {
	if index < 0 || index >= len(array) {
		var zero T
		return zero, array
	}

	item := array[index]
	result := append(array[:index], array[index+1:]...)
	return item, result
}*/

// Reverse reverses the order of elements in array.
//
// Parameters:
//   - slice: The input array to reverse
//
// Returns:
//   - []T: A new array with elements in reverse order
//
// Example:
//
//	Reverse([]int{1, 2, 3}) -> []int{3, 2, 1}
//	Reverse([]string{"a", "b", "c"}) -> []string{"c", "b", "a"}
//	Reverse([]int{1}) -> []int{1}
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, j := 0, len(slice)-1; j >= 0; i, j = i+1, j-1 {
		result[i] = slice[j]
	}
	return result
}

// Shuffle returns a new slice with elements in random order.
//
// Parameters:
//   - slice: The input array to shuffle
//
// Returns:
//   - []T: A new array with elements randomly reordered
//
// Example:
//
//	Shuffle([]int{1, 2, 3, 4, 5}) -> [3, 1, 5, 2, 4]
//	Shuffle([]string{"a", "b", "c"}) -> ["c", "a", "b"]
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
//
// Parameters:
//   - slice: The input array to select elements from
//   - n: The number of random elements to return
//
// Returns:
//   - []T: A new array containing n randomly selected elements
//
// Example:
//
//	Random([]int{1, 2, 3, 4, 5}, 3) -> [2, 4, 1]
//	Random([]string{"a", "b", "c", "d"}, 2) -> ["c", "a"]
//	Random([]int{1, 2}, 3) -> [2, 1]
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

// Random returns a random value from an array
/*func Random[T any](array []T) (T, bool) {
	if len(array) == 0 {
		var zero T
		return zero, false
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return array[r.Intn(len(array))], true
}*/

// RandomChoice returns a random element from the given slice.
//
// Parameters:
//   - choices: The input array to select a random element from
//
// Returns:
//   - T: A randomly selected element from the array
//   - bool: True if a valid element was selected, false if the array is empty
//
// Example:
//
//	RandomChoice([]string{"a", "b", "c"}) -> "b", true
//	RandomChoice([]int{1, 2, 3, 4}) -> 3, true
//	RandomChoice([]int{}) -> 0, false
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
//
// Parameters:
//   - array: The input array
//   - start: The starting index (inclusive)
//   - end: The ending index (exclusive)
//
// Returns:
//   - []T: A new array containing elements from start index up to but not including end index
//
// Example:
//
//	Slice([]int{1, 2, 3, 4}, 1, 3) -> []int{2, 3}
//	Slice([]string{"a", "b", "c", "d"}, 0, 2) -> []string{"a", "b"}
//	Slice([]int{1, 2, 3}, 2, 2) -> []int{}
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
//
// Parameters:
//   - array: The sorted input array
//   - value: The value to determine insertion index for
//
// Returns:
//   - int: The index at which the value should be inserted to maintain sort order
//
// Example:
//
//	SortedIndex([]int{1, 3, 5, 7}, 4) -> 2
//	SortedIndex([]int{10, 20, 30, 40}, 25) -> 2
//	SortedIndex([]float64{1.5, 3.5, 5.5}, 0.5) -> 0
func SortedIndex[T int | int8 | int16 | int32 | int64 | float32 | float64](array []T, value T) int {
	for i, v := range array {
		if v >= value {
			return i
		}
	}
	return len(array)
}

// Tail returns all but the first element of array.
//
// Parameters:
//   - array: The input array
//
// Returns:
//   - []T: A new array containing all elements except the first one
//
// Example:
//
//	Tail([]int{1, 2, 3}) -> []int{2, 3}
//	Tail([]string{"a", "b", "c"}) -> []string{"b", "c"}
//	Tail([]int{1}) -> []int{}
func Tail[T any](array []T) []T {
	if len(array) <= 1 {
		return []T{}
	}
	return array[1:]
}

// Take creates a slice of array with n elements taken from the beginning.
//
// Parameters:
//   - array: The input array
//   - n: Number of elements to take from the beginning
//
// Returns:
//   - []T: A new array with the first n elements
//
// Example:
//
//	Take([]int{1, 2, 3, 4}, 2) -> []int{1, 2}
//	Take([]string{"a", "b", "c"}, 1) -> []string{"a"}
//	Take([]int{1, 2}, 3) -> []int{1, 2}
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
//
// Parameters:
//   - array: The input array
//   - n: Number of elements to take from the end
//
// Returns:
//   - []T: A new array with the last n elements
//
// Example:
//
//	TakeRight([]int{1, 2, 3, 4}, 2) -> []int{3, 4}
//	TakeRight([]string{"a", "b", "c"}, 1) -> []string{"c"}
//	TakeRight([]int{1, 2}, 3) -> []int{1, 2}
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
//
// Parameters:
//   - arrays: Variable number of arrays to combine
//
// Returns:
//   - []T: A new array containing all unique elements from the input arrays
//
// Example:
//
//	Union([]int{1, 2}, []int{2, 3}) -> []int{1, 2, 3}
//	Union([]string{"a", "b"}, []string{"b", "c"}, []string{"c", "d"}) -> []string{"a", "b", "c", "d"}
//	Union([]int{1, 1, 2}, []int{2, 2, 3}) -> []int{1, 2, 3}
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
//
// Parameters:
//   - array: The input array
//
// Returns:
//   - []T: A new array with duplicate elements removed
//
// Example:
//
//	Uniq([]int{1, 2, 1, 3}) -> []int{1, 2, 3}
//	Uniq([]string{"a", "b", "a", "c", "b"}) -> []string{"a", "b", "c"}
//	Uniq([]int{1, 1, 1}) -> []int{1}
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
//
// Parameters:
//   - array: The input array
//   - values: Variable number of values to exclude from the array
//
// Returns:
//   - []T: A new array with all specified values excluded
//
// Example:
//
//	Without([]int{1, 2, 3, 4}, 2, 4) -> []int{1, 3}
//	Without([]string{"a", "b", "c"}, "a", "c") -> []string{"b"}
//	Without([]int{1, 2, 3}, 4) -> []int{1, 2, 3}
func Without[T comparable](array []T, values ...T) []T {
	return Pull(array, values...)
}

// Zip creates an array of grouped elements.
//
// Parameters:
//   - arrays: Variable number of arrays to zip together
//
// Returns:
//   - [][]T: A new array of arrays where each inner array contains elements from the input arrays at the same index
//
// Example:
//
//	Zip([]int{1, 2}, []int{3, 4}) -> [][]int{{1, 3}, {2, 4}}
//	Zip([]string{"a", "b"}, []string{"c", "d"}, []string{"e", "f"}) -> [][]string{{"a", "c", "e"}, {"b", "d", "f"}}
//	Zip([]int{1, 2, 3}, []int{4, 5}) -> [][]int{{1, 4}, {2, 5}}
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

// SortBy sorts an array by the results of running each element through the iteratee function.
// It returns a new sorted array without modifying the original.
//
// Parameters:
//   - array: The input array to be sorted
//   - iteratee: A function that transforms each element into a comparable value
//
// Returns:
//   - A new sorted array
//
// Example:
//
//	// Sort numbers in ascending order
//	SortBy([]int{1, 3, 2}, func(n int) int { return n }) // Returns [1, 2, 3]
//
//	// Sort strings by length
//	SortBy([]string{"apple", "banana", "kiwi"}, func(s string) int { return len(s) }) // Returns ["kiwi", "apple", "banana"]
//
//	// Sort structs by a specific field
//	type Person struct { Age int }
//	people := []Person{{Age: 30}, {Age: 25}, {Age: 40}}
//	SortBy(people, func(p Person) int { return p.Age }) // Returns [{Age: 25}, {Age: 30}, {Age: 40}]
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

// Contains checks if a slice contains a specific element.
// It returns true if the element is found, false otherwise.
//
// Parameters:
//   - slice: The input slice to search in
//   - element: The element to search for
//
// Returns:
//   - true if the element is found in the slice, false otherwise
//
// Example:
//
//	// Check if a number exists in a slice
//	Contains([]int{1, 2, 3, 4}, 3) // Returns true
//	Contains([]int{1, 2, 3, 4}, 5) // Returns false
//
//	// Check if a string exists in a slice
//	Contains([]string{"apple", "banana", "orange"}, "banana") // Returns true
//	Contains([]string{"apple", "banana", "orange"}, "grape") // Returns false
func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
// It does not modify the original slice.
//
// Parameters:
//   - slice: The input slice to filter
//   - predicate: A function that returns true for elements to keep and false for elements to exclude
//
// Returns:
//   - A new slice containing only the elements for which the predicate returns true
//
// Example:
//
//	// Filter even numbers
//	Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }) // Returns [2, 4]
//
//	// Filter strings longer than 5 characters
//	Filter([]string{"apple", "banana", "kiwi", "strawberry"}, func(s string) bool {
//	    return len(s) > 5
//	}) // Returns ["banana", "strawberry"]
//
//	// Filter structs based on a condition
//	type Person struct {
//	    Name string
//	    Age int
//	}
//	people := []Person{
//	    {Name: "Alice", Age: 25},
//	    {Name: "Bob", Age: 17},
//	    {Name: "Charlie", Age: 30},
//	}
//	adults := Filter(people, func(p Person) bool { return p.Age >= 18 })
//	// Returns [{Name: "Alice", Age: 25}, {Name: "Charlie", Age: 30}]
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map applies a function to each element in a slice and returns a new slice with the results.
// It transforms each element from type T to type R using the provided mapping function.
//
// Parameters:
//   - slice: The input slice to transform
//   - mapFunc: A function that transforms each element from type T to type R
//
// Returns:
//   - A new slice containing the transformed elements
//
// Example:
//
//	// Double each number
//	Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // Returns [2, 4, 6]
//
//	// Convert numbers to strings
//	Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("Number: %d", n) })
//	// Returns ["Number: 1", "Number: 2", "Number: 3"]
//
//	// Extract a field from structs
//	type User struct {
//	    ID int
//	    Name string
//	}
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 3, Name: "Charlie"},
//	}
//	names := Map(users, func(u User) string { return u.Name })
//	// Returns ["Alice", "Bob", "Charlie"]
func Map[T any, R any](slice []T, mapFunc func(T) R) []R {
	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = mapFunc(item)
	}
	return result
}

// Find returns the first element in the slice that satisfies the predicate function
// and a boolean indicating whether such an element was found.
//
// Parameters:
//   - slice: The input slice to search in
//   - predicate: A function that returns true for the element to find
//
// Returns:
//   - The first element for which the predicate returns true
//   - A boolean indicating whether such an element was found (true if found, false otherwise)
//
// Example:
//
//	// Find the first even number
//	even, found := Find([]int{1, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
//	// even = 4, found = true
//
//	// Find a string with a specific prefix
//	str, found := Find([]string{"apple", "banana", "cherry"}, func(s string) bool {
//	    return strings.HasPrefix(s, "b")
//	})
//	// str = "banana", found = true
//
//	// Find a struct that matches a condition
//	type Product struct {
//	    Name string
//	    Price float64
//	}
//	products := []Product{
//	    {Name: "Laptop", Price: 1200},
//	    {Name: "Phone", Price: 800},
//	    {Name: "Tablet", Price: 500},
//	}
//	affordable, found := Find(products, func(p Product) bool { return p.Price < 1000 })
//	// affordable = {Name: "Phone", Price: 800}, found = true
//
//	// When no element is found
//	num, found := Find([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 })
//	// num = 0 (zero value), found = false
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, item := range slice {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// Unique returns a new slice with duplicate elements removed.
// It preserves the order of elements, keeping the first occurrence of each element.
//
// Parameters:
//   - slice: The input slice that may contain duplicates
//
// Returns:
//   - A new slice with duplicate elements removed
//
// Example:
//
//	// Remove duplicate integers
//	Unique([]int{1, 2, 2, 3, 1, 4, 5, 4}) // Returns [1, 2, 3, 4, 5]
//
//	// Remove duplicate strings
//	Unique([]string{"apple", "banana", "apple", "cherry", "banana"})
//	// Returns ["apple", "banana", "cherry"]
//
//	// Works with any comparable type
//	type User struct {
//	    ID int
//	    Name string
//	}
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 1, Name: "Alice"}, // Duplicate
//	    {ID: 3, Name: "Charlie"},
//	    {ID: 2, Name: "Bob"},   // Duplicate
//	}
//	// Note: For structs, all fields must match for it to be considered a duplicate
//	uniqueUsers := Unique(users)
//	// Returns [{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}, {ID: 3, Name: "Charlie"}]
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range slice {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// SortedCopy returns a sorted copy of the slice without modifying the original.
// It uses the provided less function to determine the order.
//
// Parameters:
//   - slice: The input slice to be sorted
//   - less: A function that returns true if element i should be ordered before element j
//
// Returns:
//   - A new sorted slice
//
// Example:
//
//	// Sort integers in ascending order
//	SortedCopy([]int{3, 1, 4, 2}, func(i, j int) bool { return i < j })
//	// Returns [1, 2, 3, 4]
//
//	// Sort integers in descending order
//	SortedCopy([]int{3, 1, 4, 2}, func(i, j int) bool { return i > j })
//	// Returns [4, 3, 2, 1]
//
//	// Sort strings by length
//	SortedCopy([]string{"apple", "banana", "kiwi", "orange"}, func(i, j string) bool {
//	    return len(i) < len(j)
//	})
//	// Returns ["kiwi", "apple", "orange", "banana"]
//
//	// Sort structs by a specific field
//	type Person struct {
//	    Name string
//	    Age int
//	}
//	people := []Person{
//	    {Name: "Alice", Age: 30},
//	    {Name: "Bob", Age: 25},
//	    {Name: "Charlie", Age: 35},
//	}
//	// Sort by age
//	SortedCopy(people, func(i, j Person) bool { return i.Age < j.Age })
//	// Returns [{Name: "Bob", Age: 25}, {Name: "Alice", Age: 30}, {Name: "Charlie", Age: 35}]
func SortedCopy[T any](slice []T, less func(i, j T) bool) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	return result
}

// Reduce applies a reducer function to each element in a slice, resulting in a single output value.
// It processes each element in the slice from left to right, accumulating a result.
//
// Parameters:
//   - slice: The input slice to reduce
//   - initialValue: The initial value for the accumulator
//   - reducer: A function that combines the accumulator with each element to produce a new accumulator value
//
// Returns:
//   - The final accumulated value
//
// Example:
//
//	// Sum all numbers in a slice
//	Reduce([]int{1, 2, 3, 4}, 0, func(acc int, item int) int {
//	    return acc + item
//	}) // Returns 10
//
//	// Find the maximum value
//	Reduce([]int{5, 2, 8, 3}, math.MinInt, func(acc int, item int) int {
//	    if item > acc {
//	        return item
//	    }
//	    return acc
//	}) // Returns 8
//
//	// Concatenate strings
//	Reduce([]string{"Hello", " ", "World", "!"}, "", func(acc string, item string) string {
//	    return acc + item
//	}) // Returns "Hello World!"
//
//	// Transform a slice of structs into a map
//	type User struct {
//	    ID int
//	    Name string
//	}
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 3, Name: "Charlie"},
//	}
//	userMap := Reduce(users, make(map[int]string), func(acc map[int]string, user User) map[int]string {
//	    acc[user.ID] = user.Name
//	    return acc
//	})
//	// Returns map[1:"Alice" 2:"Bob" 3:"Charlie"]
func Reduce[T any, R any](slice []T, initialValue R, reducer func(acc R, item T) R) R {
	result := initialValue
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// GroupBy groups elements in a slice by a key generated from each element.
// It creates a map where each key is the result of applying the keyFunc to an element,
// and each value is a slice of elements that produced that key.
//
// Parameters:
//   - slice: The input slice to group
//   - keyFunc: A function that generates a key for each element
//
// Returns:
//   - A map where keys are the generated keys and values are slices of elements
//
// Example:
//
//	// Group numbers by their remainder when divided by 3
//	GroupBy([]int{1, 2, 3, 4, 5, 6}, func(n int) int {
//	    return n % 3
//	})
//	// Returns map[0:[3, 6] 1:[1, 4] 2:[2, 5]]
//
//	// Group strings by their first letter
//	GroupBy([]string{"apple", "banana", "apricot", "blueberry"}, func(s string) string {
//	    return string(s[0])
//	})
//	// Returns map["a":["apple", "apricot"] "b":["banana", "blueberry"]]
//
//	// Group structs by a field
//	type Person struct {
//	    Name string
//	    Age int
//	}
//	people := []Person{
//	    {Name: "Alice", Age: 25},
//	    {Name: "Bob", Age: 30},
//	    {Name: "Charlie", Age: 25},
//	    {Name: "Dave", Age: 30},
//	}
//	byAge := GroupBy(people, func(p Person) int { return p.Age })
//	// Returns map[25:[{Name: "Alice", Age: 25}, {Name: "Charlie", Age: 25}]
//	//            30:[{Name: "Bob", Age: 30}, {Name: "Dave", Age: 30}]]
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range slice {
		key := keyFunc(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Accessible checks if the given value can be accessed as an array, slice, or map.
// It returns true if the value is an array, slice, or map, and false otherwise.
//
// Parameters:
//   - value: The value to check
//
// Returns:
//   - true if the value is an array, slice, or map, false otherwise
//
// Example:
//
//	// Check arrays and slices
//	Accessible([]int{1, 2, 3}) // Returns true
//	Accessible([3]string{"a", "b", "c"}) // Returns true
//
//	// Check maps
//	Accessible(map[string]int{"a": 1, "b": 2}) // Returns true
//
//	// Check other types
//	Accessible(42) // Returns false
//	Accessible("hello") // Returns false
//	Accessible(struct{}{}) // Returns false
//
//	// Check nil
//	Accessible(nil) // Returns false
//
//	// Check pointers to arrays/slices/maps
//	arr := []int{1, 2, 3}
//	Accessible(&arr) // Returns false (it's a pointer, not directly accessible)
func Accessible(value any) bool {
	if value == nil {
		return false
	}

	kind := reflect.TypeOf(value).Kind()
	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// Add adds a key/value pair to a map if the key doesn't already exist.
// It returns a new map without modifying the original.
//
// Parameters:
//   - array: The input map to add the key/value pair to
//   - key: The key to add
//   - value: The value to associate with the key
//
// Returns:
//   - A new map with the key/value pair added (if the key didn't exist)
//
// Example:
//
//	// Add a new key/value pair
//	original := map[string]any{"name": "John", "age": 30}
//	result := Add(original, "city", "New York")
//	// result = {"name": "John", "age": 30, "city": "New York"}
//	// original remains unchanged
//
//	// Try to add an existing key
//	original := map[string]any{"name": "John", "age": 30}
//	result := Add(original, "name", "Jane")
//	// result = {"name": "John", "age": 30}
//	// The key "name" already exists, so the value is not changed
//
//	// Add to an empty map
//	empty := map[string]any{}
//	result := Add(empty, "status", "active")
//	// result = {"status": "active"}
func Add(array map[string]any, key string, value any) map[string]any {
	result := make(map[string]any)
	for k, v := range array {
		result[k] = v
	}

	if _, exists := result[key]; !exists {
		result[key] = value
	}

	return result
}

// Collapse collapses a slice of slices into a single slice.
// It flattens a two-dimensional slice into a one-dimensional slice.
//
// Parameters:
//   - arrays: A slice of slices to collapse
//
// Returns:
//   - A single slice containing all elements from the input slices
//
// Example:
//
//	// Collapse multiple slices into one
//	Collapse([][]any{
//	    {1, 2, 3},
//	    {4, 5},
//	    {6, 7, 8, 9},
//	}) // Returns [1, 2, 3, 4, 5, 6, 7, 8, 9]
//
//	// Collapse with mixed types
//	Collapse([][]any{
//	    {"a", "b"},
//	    {1, 2},
//	    {true, false},
//	}) // Returns ["a", "b", 1, 2, true, false]
//
//	// Collapse with empty slices
//	Collapse([][]any{
//	    {1, 2},
//	    {},
//	    {3, 4},
//	}) // Returns [1, 2, 3, 4]
//
//	// Collapse an empty slice
//	Collapse([][]any{}) // Returns []
func Collapse(arrays [][]any) []any {
	totalLen := 0
	for _, arr := range arrays {
		totalLen += len(arr)
	}

	result := make([]any, 0, totalLen)
	for _, arr := range arrays {
		result = append(result, arr...)
	}

	return result
}

// CrossJoin cross joins the given arrays, returning a cartesian product with all possible permutations.
// It generates all possible combinations by taking one element from each input array.
//
// Parameters:
//   - arrays: Variable number of slices to cross join
//
// Returns:
//   - A slice of slices where each inner slice contains one element from each input array
//
// Example:
//
//	// Cross join two arrays
//	CrossJoin([]int{1, 2}, []int{3, 4})
//	// Returns [[1, 3], [1, 4], [2, 3], [2, 4]]
//
//	// Cross join three arrays
//	CrossJoin([]string{"a", "b"}, []string{"c", "d"}, []string{"e", "f"})
//	// Returns [
//	//   ["a", "c", "e"], ["a", "c", "f"],
//	//   ["a", "d", "e"], ["a", "d", "f"],
//	//   ["b", "c", "e"], ["b", "c", "f"],
//	//   ["b", "d", "e"], ["b", "d", "f"]
//	// ]
//
//	// Cross join with a single array
//	CrossJoin([]int{1, 2, 3})
//	// Returns [[1], [2], [3]]
//
//	// Cross join with an empty array
//	CrossJoin([]int{}, []int{1, 2})
//	// Returns [] (empty result because one array is empty)
//
//	// Cross join with no arrays
//	CrossJoin[int]()
//	// Returns [] (empty result)
func CrossJoin[T any](arrays ...[]T) [][]T {
	if len(arrays) == 0 {
		return [][]T{}
	}

	if len(arrays) == 1 {
		result := make([][]T, len(arrays[0]))
		for i, item := range arrays[0] {
			result[i] = []T{item}
		}
		return result
	}

	// Get the cartesian product of all but the first array
	subResult := CrossJoin[T](arrays[1:]...)

	// Combine the first array with the sub-result
	result := make([][]T, 0, len(arrays[0])*len(subResult))
	for _, item := range arrays[0] {
		for _, subItem := range subResult {
			newItem := make([]T, 1+len(subItem))
			newItem[0] = item
			copy(newItem[1:], subItem)
			result = append(result, newItem)
		}
	}

	return result
}

// Divide returns two slices, one containing the keys, and the other containing the values of the original map.
// It separates a map into its keys and values while preserving the corresponding order.
//
// Parameters:
//   - array: The input map to divide
//
// Returns:
//   - A slice containing all the keys from the map
//   - A slice containing all the values from the map
//
// Example:
//
//	// Divide a map into keys and values
//	keys, values := Divide(map[string]any{
//	    "name": "John",
//	    "age": 30,
//	    "city": "New York",
//	})
//	// keys could be ["name", "age", "city"] (order may vary)
//	// values could be ["John", 30, "New York"] (in the same order as keys)
//
//	// Divide an empty map
//	keys, values := Divide(map[string]any{})
//	// keys = [] (empty slice)
//	// values = [] (empty slice)
//
//	// Note: The order of keys and values is not guaranteed to be the same across different runs
//	// due to the non-deterministic iteration order of Go maps
func Divide(array map[string]any) ([]string, []any) {
	keys := make([]string, 0, len(array))
	values := make([]any, 0, len(array))

	for k, v := range array {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}

// Dot flattens a multi-dimensional map into a single level map with "dot" notation.
// It converts nested maps into a flat map where keys are paths to values using dot separators.
//
// Parameters:
//   - array: The input nested map to flatten
//
// Returns:
//   - A flattened map with dot notation keys
//
// Example:
//
//	// Flatten a nested map
//	nested := map[string]any{
//	    "user": map[string]any{
//	        "name": "John",
//	        "address": map[string]any{
//	            "city": "New York",
//	            "zip": 10001,
//	        },
//	    },
//	    "status": "active",
//	}
//
//	flat := Dot(nested)
//	// Returns:
//	// {
//	//    "user.name": "John",
//	//    "user.address.city": "New York",
//	//    "user.address.zip": 10001,
//	//    "status": "active"
//	// }
//
//	// Flatten an empty map
//	Dot(map[string]any{}) // Returns an empty map
//
//	// Flatten a map with no nested structures
//	Dot(map[string]any{"a": 1, "b": 2}) // Returns the same map {"a": 1, "b": 2}
func Dot(array map[string]any) map[string]any {
	result := make(map[string]any)
	dotRecursive(array, result, "")
	return result
}

// dotRecursive is a helper function for Dot
func dotRecursive(array, result map[string]any, prepend string) {
	for key, value := range array {
		if prepend != "" {
			key = prepend + "." + key
		}

		if subArray, ok := value.(map[string]any); ok {
			dotRecursive(subArray, result, key)
		} else {
			result[key] = value
		}
	}
}

// Except returns a new map with the specified keys removed from the original map.
// It creates a copy of the input map excluding the specified keys.
//
// Parameters:
//   - array: The input map to filter
//   - keys: Variable number of keys to exclude from the result
//
// Returns:
//   - A new map with the specified keys removed
//
// Example:
//
//	// Remove specific keys from a map
//	original := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	    "city": "New York",
//	    "country": "USA",
//	}
//	result := Except(original, "age", "country")
//	// Returns {"name": "John", "city": "New York"}
//
//	// Remove keys that don't exist
//	original := map[string]any{"a": 1, "b": 2}
//	result := Except(original, "c", "d")
//	// Returns {"a": 1, "b": 2} (unchanged since keys don't exist)
//
//	// Remove all keys
//	original := map[string]any{"a": 1, "b": 2}
//	result := Except(original, "a", "b")
//	// Returns {} (empty map)
//
//	// No keys to remove
//	original := map[string]any{"a": 1, "b": 2}
//	result := Except(original)
//	// Returns {"a": 1, "b": 2} (unchanged)
func Except(array map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)

	// Create a map for faster lookup
	keysMap := make(map[string]struct{})
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}

	for key, value := range array {
		if _, exists := keysMap[key]; !exists {
			result[key] = value
		}
	}

	return result
}

// Exists checks if the given key exists in the map.
// It returns true if the key exists, false otherwise.
//
// Parameters:
//   - array: The map to check
//   - key: The key to look for
//
// Returns:
//   - true if the key exists in the map, false otherwise
//
// Example:
//
//	// Check if a key exists
//	user := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	}
//	Exists(user, "name") // Returns true
//	Exists(user, "email") // Returns false
//
//	// Check in an empty map
//	Exists(map[string]any{}, "key") // Returns false
//
//	// Check with a nil map
//	var nilMap map[string]any
//	Exists(nilMap, "key") // Returns false (safe to use with nil maps)
func Exists(array map[string]any, key string) bool {
	_, exists := array[key]
	return exists
}

// FirstOrDefault returns the first element in the array, or a default value if the array is empty.
// It safely handles empty arrays by returning the provided default value.
//
// Parameters:
//   - array: The input array to get the first element from
//   - defaultValue: The value to return if the array is empty
//
// Returns:
//   - The first element of the array if it exists, otherwise the default value
//
// Example:
//
//	// Get the first element from a non-empty array
//	FirstOrDefault([]int{1, 2, 3}, 0) // Returns 1
//
//	// Get the default value from an empty array
//	FirstOrDefault([]int{}, 0) // Returns 0
//
//	// Works with any type
//	FirstOrDefault([]string{"apple", "banana"}, "default") // Returns "apple"
//	FirstOrDefault([]string{}, "default") // Returns "default"
//
//	// Works with structs
//	type User struct {
//	    Name string
//	}
//	users := []User{{Name: "Alice"}, {Name: "Bob"}}
//	defaultUser := User{Name: "Unknown"}
//	FirstOrDefault(users, defaultUser) // Returns {Name: "Alice"}
//	FirstOrDefault([]User{}, defaultUser) // Returns {Name: "Unknown"}
func FirstOrDefault[T any](array []T, defaultValue T) T {
	if len(array) > 0 {
		return array[0]
	}
	return defaultValue
}

// Forget removes the given key/value pairs from the map.
// It returns a new map with the specified keys removed, without modifying the original map.
//
// Parameters:
//   - array: The input map to remove keys from
//   - keys: Variable number of keys to remove
//
// Returns:
//   - A new map with the specified keys removed
//
// Example:
//
//	// Remove a single key
//	original := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	    "city": "New York",
//	}
//	result := Forget(original, "age")
//	// Returns {"name": "John", "city": "New York"}
//
//	// Remove multiple keys
//	original := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	    "city": "New York",
//	    "country": "USA",
//	}
//	result := Forget(original, "age", "country")
//	// Returns {"name": "John", "city": "New York"}
//
//	// Remove keys that don't exist
//	original := map[string]any{"a": 1, "b": 2}
//	result := Forget(original, "c")
//	// Returns {"a": 1, "b": 2} (unchanged since key doesn't exist)
//
//	// Note: Forget is similar to Except, but with a different parameter order
func Forget(array map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)
	for k, v := range array {
		result[k] = v
	}

	for _, key := range keys {
		delete(result, key)
	}

	return result
}

// Get retrieves a value from a map using "dot" notation for accessing nested values.
// It allows accessing deeply nested values in a map using dot-separated keys.
// If the key doesn't exist, it returns the provided default value.
//
// Parameters:
//   - array: The input map to retrieve the value from
//   - key: The key to look for, using dot notation for nested keys
//   - defaultValue: The value to return if the key doesn't exist
//
// Returns:
//   - The value associated with the key if it exists, otherwise the default value
//
// Example:
//
//	// Simple key access
//	data := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	}
//	Get(data, "name", "Unknown") // Returns "John"
//	Get(data, "email", "N/A") // Returns "N/A" (key doesn't exist)
//
//	// Nested key access
//	nested := map[string]any{
//	    "user": map[string]any{
//	        "name": "John",
//	        "address": map[string]any{
//	            "city": "New York",
//	            "zip": 10001,
//	        },
//	    },
//	    "status": "active",
//	}
//
//	Get(nested, "user.name", "Unknown") // Returns "John"
//	Get(nested, "user.address.city", "Unknown") // Returns "New York"
//	Get(nested, "user.address.country", "USA") // Returns "USA" (key doesn't exist)
//	Get(nested, "user.email", nil) // Returns nil (key doesn't exist)
//
//	// Empty key returns the entire map
//	Get(nested, "", nil) // Returns the entire nested map
func Get(array map[string]any, key string, defaultValue any) any {
	if array == nil {
		return defaultValue
	}

	if key == "" {
		return array
	}

	keys := strings.Split(key, ".")
	current := array

	for i, segment := range keys {
		if i == len(keys)-1 {
			if val, exists := current[segment]; exists {
				return val
			}
			return defaultValue
		}

		if val, exists := current[segment]; exists {
			if nextMap, ok := val.(map[string]any); ok {
				current = nextMap
			} else {
				return defaultValue
			}
		} else {
			return defaultValue
		}
	}

	return defaultValue
}

// Has determines if all of the specified keys exist in the map using "dot" notation.
// It checks if every key in the provided list exists in the map, including nested keys.
// Returns true only if all keys exist.
//
// Parameters:
//   - array: The input map to check
//   - keys: Variable number of keys to check for existence
//
// Returns:
//   - true if all specified keys exist in the map, false otherwise
//
// Example:
//
//	// Check simple keys
//	data := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	}
//	Has(data, "name") // Returns true
//	Has(data, "email") // Returns false
//	Has(data, "name", "age") // Returns true (both keys exist)
//	Has(data, "name", "email") // Returns false (not all keys exist)
//
//	// Check nested keys
//	nested := map[string]any{
//	    "user": map[string]any{
//	        "name": "John",
//	        "address": map[string]any{
//	            "city": "New York",
//	        },
//	    },
//	}
//
//	Has(nested, "user.name") // Returns true
//	Has(nested, "user.address.city") // Returns true
//	Has(nested, "user.address.country") // Returns false
//	Has(nested, "user.name", "user.address.city") // Returns true (both keys exist)
//
//	// Empty keys list
//	Has(data) // Returns false (no keys to check)
func Has(array map[string]any, keys ...string) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		if !hasDot(array, key) {
			return false
		}
	}

	return true
}

// hasDot is a helper function for Has
func hasDot(array map[string]any, key string) bool {
	if array == nil {
		return false
	}

	if key == "" {
		return false
	}

	keys := strings.Split(key, ".")
	current := array

	for i, segment := range keys {
		if val, exists := current[segment]; exists {
			if i == len(keys)-1 {
				return true
			}

			if nextMap, ok := val.(map[string]any); ok {
				current = nextMap
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return false
}

// HasAny determines if any of the specified keys exist in the map using "dot" notation.
// It checks if at least one key in the provided list exists in the map, including nested keys.
// Returns true if at least one key exists.
//
// Parameters:
//   - array: The input map to check
//   - keys: Variable number of keys to check for existence
//
// Returns:
//   - true if at least one of the specified keys exists in the map, false otherwise
//
// Example:
//
//	// Check simple keys
//	data := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	}
//	HasAny(data, "name") // Returns true
//	HasAny(data, "email") // Returns false
//	HasAny(data, "name", "email") // Returns true (at least one key exists)
//	HasAny(data, "email", "phone") // Returns false (none of the keys exist)
//
//	// Check nested keys
//	nested := map[string]any{
//	    "user": map[string]any{
//	        "name": "John",
//	        "address": map[string]any{
//	            "city": "New York",
//	        },
//	    },
//	}
//
//	HasAny(nested, "user.name") // Returns true
//	HasAny(nested, "user.address.city") // Returns true
//	HasAny(nested, "user.address.country") // Returns false
//	HasAny(nested, "user.address.country", "user.name") // Returns true (at least one key exists)
//
//	// Empty keys list
//	HasAny(data) // Returns false (no keys to check)
func HasAny(array map[string]any, keys ...string) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		if hasDot(array, key) {
			return true
		}
	}

	return false
}

// IsAssoc determines if a value is an associative array/map (has string keys).
// It checks if the value is a map with string keys.
//
// Parameters:
//   - array: The value to check
//
// Returns:
//   - true if the value is a map with string keys, false otherwise
//
// Example:
//
//	// Check associative arrays (maps with string keys)
//	IsAssoc(map[string]any{"name": "John", "age": 30}) // Returns true
//	IsAssoc(map[string]int{"a": 1, "b": 2}) // Returns true
//
//	// Check non-associative arrays
//	IsAssoc([]int{1, 2, 3}) // Returns false (slice, not a map)
//	IsAssoc(map[int]string{1: "a", 2: "b"}) // Returns false (map with non-string keys)
//
//	// Check other types
//	IsAssoc(42) // Returns false
//	IsAssoc("hello") // Returns false
//
//	// Check nil
//	IsAssoc(nil) // Returns false
func IsAssoc(array any) bool {
	if array == nil {
		return false
	}

	value := reflect.ValueOf(array)
	if value.Kind() != reflect.Map {
		return false
	}

	// Check if all keys are strings
	for _, key := range value.MapKeys() {
		if key.Kind() != reflect.String {
			return false
		}
	}

	return true
}

// IsList determines if a value is a list (slice or array).
// It checks if the value is a slice or array type.
//
// Parameters:
//   - array: The value to check
//
// Returns:
//   - true if the value is a slice or array, false otherwise
//
// Example:
//
//	// Check slices
//	IsList([]int{1, 2, 3}) // Returns true
//	IsList([]string{"a", "b", "c"}) // Returns true
//	IsList([]any{1, "a", true}) // Returns true
//
//	// Check arrays
//	IsList([3]int{1, 2, 3}) // Returns true
//	IsList([2]string{"a", "b"}) // Returns true
//
//	// Check non-list types
//	IsList(map[string]int{"a": 1, "b": 2}) // Returns false (map, not a slice/array)
//	IsList(42) // Returns false
//	IsList("hello") // Returns false
//
//	// Check nil
//	IsList(nil) // Returns false
//
//	// Check empty slice
//	IsList([]int{}) // Returns true (empty slice is still a list)
func IsList(array any) bool {
	if array == nil {
		return false
	}

	value := reflect.ValueOf(array)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return false
	}

	return true
}

// KeyBy creates a map from an array, using the result of the key function as the map key.
// It transforms a slice into a map where each element is indexed by a key derived from the element itself.
//
// Parameters:
//   - array: The input slice to transform into a map
//   - keyFunc: A function that generates a key for each element
//
// Returns:
//   - A map where keys are generated by the keyFunc and values are elements from the input array
//
// Example:
//
//	// Key numbers by themselves
//	KeyBy([]int{1, 2, 3}, func(n int) int { return n })
//	// Returns map[1:1 2:2 3:3]
//
//	// Key strings by their length
//	KeyBy([]string{"apple", "banana", "kiwi"}, func(s string) int {
//	    return len(s)
//	})
//	// Returns map[5:"apple" 6:"banana" 4:"kiwi"]
//
//	// Key structs by a field
//	type User struct {
//	    ID int
//	    Name string
//	}
//	users := []User{
//	    {ID: 1, Name: "Alice"},
//	    {ID: 2, Name: "Bob"},
//	    {ID: 3, Name: "Charlie"},
//	}
//	userMap := KeyBy(users, func(u User) int { return u.ID })
//	// Returns map[1:{ID:1 Name:"Alice"} 2:{ID:2 Name:"Bob"} 3:{ID:3 Name:"Charlie"}]
//
//	// Note: If multiple elements produce the same key, later elements will overwrite earlier ones
func KeyBy[T any, K comparable](array []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range array {
		key := keyFunc(item)
		result[key] = item
	}
	return result
}

// LastOrDefault returns the last element in the array, or a default value if the array is empty.
// It safely handles empty arrays by returning the provided default value.
//
// Parameters:
//   - array: The input array to get the last element from
//   - defaultValue: The value to return if the array is empty
//
// Returns:
//   - The last element of the array if it exists, otherwise the default value
//
// Example:
//
//	// Get the last element from a non-empty array
//	LastOrDefault([]int{1, 2, 3}, 0) // Returns 3
//
//	// Get the default value from an empty array
//	LastOrDefault([]int{}, 0) // Returns 0
//
//	// Works with any type
//	LastOrDefault([]string{"apple", "banana"}, "default") // Returns "banana"
//	LastOrDefault([]string{}, "default") // Returns "default"
//
//	// Works with structs
//	type User struct {
//	    Name string
//	}
//	users := []User{{Name: "Alice"}, {Name: "Bob"}}
//	defaultUser := User{Name: "Unknown"}
//	LastOrDefault(users, defaultUser) // Returns {Name: "Bob"}
//	LastOrDefault([]User{}, defaultUser) // Returns {Name: "Unknown"}
func LastOrDefault[T any](array []T, defaultValue T) T {
	if len(array) > 0 {
		return array[len(array)-1]
	}
	return defaultValue
}

// Only returns a new map containing only the specified keys from the original map.
// It creates a filtered copy of the input map with just the requested keys.
//
// Parameters:
//   - array: The input map to filter
//   - keys: Variable number of keys to include in the result
//
// Returns:
//   - A new map containing only the specified keys and their values
//
// Example:
//
//	// Keep only specific keys
//	original := map[string]any{
//	    "name": "John",
//	    "age": 30,
//	    "city": "New York",
//	    "country": "USA",
//	}
//	result := Only(original, "name", "city")
//	// Returns {"name": "John", "city": "New York"}
//
//	// Request keys that don't exist
//	original := map[string]any{"a": 1, "b": 2}
//	result := Only(original, "a", "c")
//	// Returns {"a": 1} (only existing keys are included)
//
//	// Request no keys
//	original := map[string]any{"a": 1, "b": 2}
//	result := Only(original)
//	// Returns {} (empty map)
//
//	// Note: Only is the opposite of Except - it keeps only the specified keys
//	// while Except removes the specified keys
func Only(array map[string]any, keys ...string) map[string]any {
	result := make(map[string]any)

	for _, key := range keys {
		if value, exists := array[key]; exists {
			result[key] = value
		}
	}

	return result
}

// Pluck extracts a specific property from each element in a slice.
// It creates a new slice containing the values of a specified property from each element.
//
// Parameters:
//   - array: The input slice of elements
//   - key: A function that extracts a value from each element
//
// Returns:
//   - A slice containing the extracted values
//
// Example:
//
//	// Extract a property from structs
//	type User struct {
//	    ID int
//	    Name string
//	    Age int
//	}
//	users := []User{
//	    {ID: 1, Name: "Alice", Age: 25},
//	    {ID: 2, Name: "Bob", Age: 30},
//	    {ID: 3, Name: "Charlie", Age: 35},
//	}
//
//	// Get all names
//	names := Pluck(users, func(u User) string { return u.Name })
//	// Returns ["Alice", "Bob", "Charlie"]
//
//	// Get all ages
//	ages := Pluck(users, func(u User) int { return u.Age })
//	// Returns [25, 30, 35]
//
//	// Works with maps too
//	people := []map[string]any{
//	    {"name": "Alice", "age": 25},
//	    {"name": "Bob", "age": 30},
//	    {"name": "Charlie", "age": 35},
//	}
//	names := Pluck(people, func(p map[string]any) string {
//	    return p["name"].(string)
//	})
//	// Returns ["Alice", "Bob", "Charlie"]
func Pluck[T any, V any](array []T, key func(T) V) []V {
	result := make([]V, len(array))
	for i, item := range array {
		result[i] = key(item)
	}
	return result
}

// Prepend adds one or more items to the beginning of a slice.
// It returns a new slice with the values added at the beginning, without modifying the original slice.
//
// Parameters:
//   - array: The original slice
//   - values: One or more values to add to the beginning of the slice
//
// Returns:
//   - A new slice with the values prepended
//
// Example:
//
//	// Prepend a single value
//	Prepend([]int{2, 3, 4}, 1) // Returns [1, 2, 3, 4]
//
//	// Prepend multiple values
//	Prepend([]int{3, 4, 5}, 1, 2) // Returns [1, 2, 3, 4, 5]
//
//	// Prepend to an empty slice
//	Prepend([]string{}, "hello") // Returns ["hello"]
//
//	// Works with any type
//	Prepend([]string{"world"}, "hello") // Returns ["hello", "world"]
//
//	// Prepend no values (returns a copy of the original)
//	Prepend([]int{1, 2, 3}) // Returns [1, 2, 3]
func Prepend[T any](array []T, values ...T) []T {
	result := make([]T, len(values)+len(array))
	copy(result, values)
	copy(result[len(values):], array)
	return result
}

// Query builds a URL query string from a map.
// It converts a map into a URL-encoded query string suitable for HTTP requests.
//
// Parameters:
//   - array: The input map to convert to a query string
//
// Returns:
//   - A URL-encoded query string
//
// Example:
//
//	// Simple key-value pairs
//	Query(map[string]any{
//	    "name": "John Doe",
//	    "age": 30,
//	}) // Returns "age=30&name=John+Doe" (order may vary)
//
//	// With array values
//	Query(map[string]any{
//	    "colors": []string{"red", "blue", "green"},
//	    "id": 123,
//	}) // Returns "colors%5B%5D=red&colors%5B%5D=blue&colors%5B%5D=green&id=123" (order may vary)
//	// Decoded: "colors[]=red&colors[]=blue&colors[]=green&id=123"
//
//	// With special characters
//	Query(map[string]any{
//	    "search": "hello world",
//	    "filter": "price>100",
//	}) // Returns "filter=price%3E100&search=hello+world" (order may vary)
//	// Decoded: "filter=price>100&search=hello world"
//
//	// Empty map
//	Query(map[string]any{}) // Returns "" (empty string)
func Query(array map[string]any) string {
	values := url.Values{}
	for key, value := range array {
		switch v := value.(type) {
		case string:
			values.Add(key, v)
		case []string:
			for _, item := range v {
				values.Add(key+"[]", item)
			}
		default:
			// Convert to string using fmt.Sprint
			values.Add(key, fmt.Sprint(v))
		}
	}

	return values.Encode()
}

// RandomOrDefault returns a random value from a slice or a default value if the slice is empty.
// It safely handles empty slices by returning the provided default value.
//
// Parameters:
//   - array: The input slice to get a random element from
//   - defaultValue: The value to return if the slice is empty
//
// Returns:
//   - A random element from the slice if it's not empty, otherwise the default value
//
// Example:
//
//	// Get a random element from a non-empty slice
//	// Note: The actual returned value will vary due to randomness
//	RandomOrDefault([]int{1, 2, 3, 4, 5}, 0) // Returns one of 1, 2, 3, 4, or 5
//
//	// Get the default value from an empty slice
//	RandomOrDefault([]int{}, 0) // Returns 0
//
//	// Works with any type
//	RandomOrDefault([]string{"apple", "banana", "cherry"}, "default")
//	// Returns one of "apple", "banana", or "cherry"
//
//	RandomOrDefault([]string{}, "default") // Returns "default"
//
//	// Works with structs
//	type User struct {
//	    Name string
//	}
//	users := []User{{Name: "Alice"}, {Name: "Bob"}, {Name: "Charlie"}}
//	defaultUser := User{Name: "Unknown"}
//	RandomOrDefault(users, defaultUser) // Returns one of the users randomly
func RandomOrDefault[T any](array []T, defaultValue T) T {
	if len(array) == 0 {
		return defaultValue
	}

	return array[rand.IntN(len(array))]
}

// Set sets a value within a nested map using "dot" notation.
//
// Parameters:
//   - array: The source map to modify
//   - key: The key in dot notation (e.g., "user.address.city")
//   - value: The value to set at the specified key
//
// Returns:
//   - A new map with the value set at the specified key
//
// Example:
//
//	data := map[string]any{"user": map[string]any{"name": "John"}}
//	result := arr.Set(data, "user.age", 30)
//	// result: {"user": {"name": "John", "age": 30}}
func Set(array map[string]any, key string, value any) map[string]any {
	result := make(map[string]any)
	for k, v := range array {
		result[k] = v
	}

	if key == "" {
		return result
	}

	keys := strings.Split(key, ".")
	current := result

	for i, segment := range keys {
		if i == len(keys)-1 {
			current[segment] = value
			break
		}

		if val, exists := current[segment]; exists {
			if nextMap, ok := val.(map[string]any); ok {
				current = nextMap
			} else {
				// Convert to map if it's not already
				nextMap = make(map[string]any)
				current[segment] = nextMap
				current = nextMap
			}
		} else {
			nextMap := make(map[string]any)
			current[segment] = nextMap
			current = nextMap
		}
	}

	return result
}

// SortByKey sorts a map by keys in ascending alphabetical order.
//
// Parameters:
//   - array: The source map to sort
//
// Returns:
//   - A new map with keys sorted in ascending order
//
// Example:
//
//	data := map[string]any{"c": 3, "a": 1, "b": 2}
//	result := arr.SortByKey(data)
//	// result: {"a": 1, "b": 2, "c": 3}
func SortByKey(array map[string]any) map[string]any {
	// Get the keys and sort them
	keys := make([]string, 0, len(array))
	for k := range array {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create a new map with the sorted keys
	result := make(map[string]any)
	for _, k := range keys {
		result[k] = array[k]
	}

	return result
}

// SortByKeyDesc sorts a map by keys in descending alphabetical order.
//
// Parameters:
//   - array: The source map to sort
//
// Returns:
//   - A new map with keys sorted in descending order
//
// Example:
//
//	data := map[string]any{"a": 1, "b": 2, "c": 3}
//	result := arr.SortByKeyDesc(data)
//	// result: {"c": 3, "b": 2, "a": 1}
func SortByKeyDesc(array map[string]any) map[string]any {
	// Get the keys and sort them in descending order
	keys := make([]string, 0, len(array))
	for k := range array {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	// Create a new map with the sorted keys
	result := make(map[string]any)
	for _, k := range keys {
		result[k] = array[k]
	}

	return result
}

// SortRecursive recursively sorts maps by keys and nested arrays/maps.
//
// Parameters:
//   - array: The source data structure to sort (can be a map, slice, or any other value)
//
// Returns:
//   - A new data structure with all nested maps sorted by keys
//
// Example:
//
//	data := map[string]any{
//	    "c": 3,
//	    "a": map[string]any{"z": 26, "x": 24},
//	    "b": []any{2, 1, 3}
//	}
//	result := arr.SortRecursive(data)
//	// result: {
//	//   "a": {"x": 24, "z": 26},
//	//   "b": [2, 1, 3], // Note: array order is preserved
//	//   "c": 3
//	// }
func SortRecursive(array any) any {
	switch arr := array.(type) {
	case map[string]any:
		// Sort the map by keys
		keys := make([]string, 0, len(arr))
		for k := range arr {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// Create a new map with the sorted keys and recursively sorted values
		result := make(map[string]any)
		for _, k := range keys {
			result[k] = SortRecursive(arr[k])
		}

		return result

	case []any:
		// Create a new slice with recursively sorted values
		result := make([]any, len(arr))
		for i, v := range arr {
			result[i] = SortRecursive(v)
		}

		return result

	default:
		// Return the value as is
		return array
	}
}

// Undot expands a flattened map with "dot" notation keys back into a nested map structure.
//
// Parameters:
//   - array: The flattened map with dot notation keys
//
// Returns:
//   - A new nested map structure
//
// Example:
//
//	data := map[string]any{
//	    "user.name": "John",
//	    "user.address.city": "New York",
//	    "user.address.zip": "10001"
//	}
//	result := arr.Undot(data)
//	// result: {
//	//   "user": {
//	//     "name": "John",
//	//     "address": {
//	//       "city": "New York",
//	//       "zip": "10001"
//	//     }
//	//   }
//	// }
func Undot(array map[string]any) map[string]any {
	result := make(map[string]any)

	for key, value := range array {
		parts := strings.Split(key, ".")

		// Reference to the current level in the result
		current := result

		// Traverse the parts of the key
		for i, part := range parts {
			// If this is the last part, set the value
			if i == len(parts)-1 {
				current[part] = value
				continue
			}

			// If the next level doesn't exist, create it
			if _, exists := current[part]; !exists {
				current[part] = make(map[string]any)
			}

			// Move to the next level
			current = current[part].(map[string]any)
		}
	}

	return result
}

// WhereNotNull filters an array by removing nil values.
//
// Parameters:
//   - array: The source array to filter
//
// Returns:
//   - A new array with all nil values removed
//
// Example:
//
//	type User struct {
//	    Name string
//	}
//	var u1 = &User{Name: "Alice"}
//	var u2 *User = nil
//	var u3 = &User{Name: "Bob"}
//	users := []*User{u1, u2, u3}
//	result := arr.WhereNotNull(users)
//	// result: [&User{Name: "Alice"}, &User{Name: "Bob"}]
func WhereNotNull[T any](array []T) []T {
	result := make([]T, 0)

	for _, item := range array {
		// Check if the item is nil
		if !isNil(item) {
			result = append(result, item)
		}
	}

	return result
}

// isNil checks if a value is nil. This is a helper function used internally.
//
// Parameters:
//   - value: The value to check for nil
//
// Returns:
//   - true if the value is nil, false otherwise
//
// Note:
//   - This function handles nil checks for pointers, interfaces, maps, slices, and channels
func isNil(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	kind := v.Kind()

	// Check for nil pointers, interfaces, maps, slices, and channels
	return (kind == reflect.Ptr || kind == reflect.Interface ||
		kind == reflect.Map || kind == reflect.Slice ||
		kind == reflect.Chan) && v.IsNil()
}

// Wrap ensures a value is contained in a slice. If the value is already a slice or array,
// it converts it to []any. Otherwise, it creates a new slice containing the value.
//
// Parameters:
//   - value: The value to wrap in a slice
//
// Returns:
//   - A slice containing the value or the converted slice
//
// Example:
//
//	// Wrapping a single value
//	result1 := arr.Wrap(42)
//	// result1: []any{42}
//
//	// Wrapping an existing slice
//	nums := []int{1, 2, 3}
//	result2 := arr.Wrap(nums)
//	// result2: []any{1, 2, 3}
//
//	// Handling nil
//	result3 := arr.Wrap(nil)
//	// result3: []any{}
func Wrap(value any) []any {
	if value == nil {
		return []any{}
	}

	v := reflect.ValueOf(value)

	// If it's already a slice or array, convert it to []any
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		result := make([]any, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = v.Index(i).Interface()
		}
		return result
	}

	// Otherwise, wrap it in a slice
	return []any{value}
}

// MapMerge combines multiple maps into a single new map.
//
// Parameters:
//   - maps: Variable number of maps to merge
//
// Returns:
//   - A new map containing all key-value pairs from the input maps
//
// Notes:
//   - If there are duplicate keys, values from later maps will overwrite earlier ones
//
// Example:
//
//	map1 := map[string]int{"a": 1, "b": 2}
//	map2 := map[string]int{"b": 3, "c": 4}
//	result := arr.MapMerge(map1, map2)
//	// result: {"a": 1, "b": 3, "c": 4}
func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// MapKeys extracts all keys from a map into a slice.
//
// Parameters:
//   - m: The source map
//
// Returns:
//   - A slice containing all keys from the map
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := arr.MapKeys(data)
//	// keys: []string{"a", "b", "c"} (order may vary)
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues extracts all values from a map into a slice.
//
// Parameters:
//   - m: The source map
//
// Returns:
//   - A slice containing all values from the map
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := arr.MapValues(data)
//	// values: []int{1, 2, 3} (order may vary)
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// MapValuesFn transforms all values in a map using a mapping function.
//
// Parameters:
//   - m: The source map
//   - mapFunc: A function that transforms values of type V to type R
//
// Returns:
//   - A new map with the same keys but transformed values
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	doubled := arr.MapValuesFn(data, func(v int) int {
//	    return v * 2
//	})
//	// doubled: {"a": 2, "b": 4, "c": 6}
//
//	// Converting types
//	toString := arr.MapValuesFn(data, func(v int) string {
//	    return fmt.Sprintf("value-%d", v)
//	})
//	// toString: {"a": "value-1", "b": "value-2", "c": "value-3"}
func MapValuesFn[K comparable, V any, R any](m map[K]V, mapFunc func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = mapFunc(v)
	}
	return result
}

// MapFindKey finds the first key in a map that corresponds to a specific value.
//
// Parameters:
//   - m: The source map to search in
//   - value: The value to search for
//
// Returns:
//   - The first key that maps to the specified value
//   - A boolean indicating whether such a key was found
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 1}
//	key, found := arr.MapFindKey(data, 1)
//	// key: "a" (or "c" depending on map iteration order), found: true
//
//	key, found = arr.MapFindKey(data, 5)
//	// key: "" (zero value for string), found: false
func MapFindKey[K comparable, V comparable](m map[K]V, value V) (K, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// MapFilterMap creates a new map containing only the key-value pairs that satisfy a predicate function.
//
// Parameters:
//   - m: The source map to filter
//   - predicate: A function that takes a key and value and returns true if the pair should be included
//
// Returns:
//   - A new map containing only the key-value pairs that satisfy the predicate
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//
//	// Keep only entries with even values
//	evens := arr.MapFilterMap(data, func(k string, v int) bool {
//	    return v%2 == 0
//	})
//	// evens: {"b": 2, "d": 4}
//
//	// Keep only entries where key is "a" or "c"
//	filtered := arr.MapFilterMap(data, func(k string, v int) bool {
//	    return k == "a" || k == "c"
//	})
//	// filtered: {"a": 1, "c": 3}
func MapFilterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// MapInvertMap creates a new map by swapping the keys and values of the original map.
//
// Parameters:
//   - m: The source map to invert
//
// Returns:
//   - A new map with the keys and values swapped
//
// Notes:
//   - If multiple keys map to the same value in the original map, only one key-value pair will be in the result
//   - The last key-value pair processed will be the one that appears in the result
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	inverted := arr.MapInvertMap(data)
//	// inverted: {1: "a", 2: "b", 3: "c"}
//
//	// With duplicate values
//	data2 := map[string]int{"a": 1, "b": 2, "c": 1}
//	inverted2 := arr.MapInvertMap(data2)
//	// inverted2 might be {1: "c", 2: "b"} or {1: "a", 2: "b"} depending on map iteration order
func MapInvertMap[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// MapGetOrDefault safely retrieves a value from a map, returning a default value if the key doesn't exist.
//
// Parameters:
//   - m: The source map
//   - key: The key to look up
//   - defaultValue: The value to return if the key doesn't exist
//
// Returns:
//   - The value associated with the key, or the default value if the key doesn't exist
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//
//	value := arr.MapGetOrDefault(data, "b", 0)
//	// value: 2
//
//	value = arr.MapGetOrDefault(data, "d", 0)
//	// value: 0 (default value)
func MapGetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	return defaultValue
}

// MapGetOrInsert retrieves a value from a map, or inserts a default value if the key doesn't exist.
//
// Parameters:
//   - m: The source map (will be modified if the key doesn't exist)
//   - key: The key to look up
//   - defaultValue: The value to insert and return if the key doesn't exist
//
// Returns:
//   - The value associated with the key, or the default value if the key didn't exist
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2}
//
//	// Key exists
//	value := arr.MapGetOrInsert(data, "b", 0)
//	// value: 2, data unchanged: {"a": 1, "b": 2}
//
//	// Key doesn't exist
//	value = arr.MapGetOrInsert(data, "c", 3)
//	// value: 3, data modified: {"a": 1, "b": 2, "c": 3}
func MapGetOrInsert[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, ok := m[key]; ok {
		return value
	}
	m[key] = defaultValue
	return defaultValue
}

// MapToSlice converts a map to a slice of key-value pair structs.
//
// Parameters:
//   - m: The source map to convert
//
// Returns:
//   - A slice of structs, each containing a Key and Value field
//
// Example:
//
//	data := map[string]int{"a": 1, "b": 2, "c": 3}
//	pairs := arr.MapToSlice(data)
//	// pairs is a slice of struct{Key string; Value int} with entries like:
//	// [{Key: "a", Value: 1}, {Key: "b", Value: 2}, {Key: "c", Value: 3}]
//	// (order may vary due to map iteration)
//
//	// You can iterate over the pairs:
//	for _, pair := range pairs {
//	    fmt.Printf("Key: %s, Value: %d\n", pair.Key, pair.Value)
//	}
func MapToSlice[K comparable, V any](m map[K]V) []struct {
	Key   K
	Value V
} {
	result := make([]struct {
		Key   K
		Value V
	}, 0, len(m))

	for k, v := range m {
		result = append(result, struct {
			Key   K
			Value V
		}{k, v})
	}

	return result
}

// MapSliceToMap converts a slice of key-value pair structs to a map.
// This is the inverse operation of MapToSlice.
//
// Parameters:
//   - slice: A slice of structs, each containing a Key and Value field
//
// Returns:
//   - A map with keys and values from the input slice
//
// Example:
//
//	pairs := []struct {
//	    Key   string
//	    Value int
//	}{
//	    {Key: "a", Value: 1},
//	    {Key: "b", Value: 2},
//	    {Key: "c", Value: 3},
//	}
//
//	data := arr.MapSliceToMap(pairs)
//	// data: map[string]int{"a": 1, "b": 2, "c": 3}
//
//	// If there are duplicate keys, the last one wins:
//	pairs2 := []struct {
//	    Key   string
//	    Value int
//	}{
//	    {Key: "a", Value: 1},
//	    {Key: "a", Value: 10},
//	}
//	data2 := arr.MapSliceToMap(pairs2)
//	// data2: map[string]int{"a": 10}
func MapSliceToMap[K comparable, V any](slice []struct {
	Key   K
	Value V
}) map[K]V {
	result := make(map[K]V, len(slice))
	for _, item := range slice {
		result[item.Key] = item.Value
	}
	return result
}

// MapEqualMaps checks if two maps contain exactly the same key-value pairs.
//
// Parameters:
//   - m1: The first map to compare
//   - m2: The second map to compare
//
// Returns:
//   - true if both maps have the same keys with the same values, false otherwise
//
// Example:
//
//	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
//	map2 := map[string]int{"c": 3, "b": 2, "a": 1} // Same content, different order
//	map3 := map[string]int{"a": 1, "b": 2}         // Missing a key
//	map4 := map[string]int{"a": 1, "b": 2, "c": 4} // Different value
//
//	arr.MapEqualMaps(map1, map2) // Returns: true
//	arr.MapEqualMaps(map1, map3) // Returns: false
//	arr.MapEqualMaps(map1, map4) // Returns: false
func MapEqualMaps[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// MapDiffMaps identifies the differences between two maps.
//
// Parameters:
//   - m1: The first map (considered the "original" map)
//   - m2: The second map (considered the "new" map)
//
// Returns:
//   - added: A map containing keys in m2 that are not in m1 (with their values from m2)
//   - removed: A map containing keys in m1 that are not in m2 (with their values from m1)
//   - changed: A map containing keys that exist in both maps but have different values (with their values from m2)
//
// Example:
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	new := map[string]int{"b": 20, "c": 3, "d": 4}
//
//	added, removed, changed := arr.MapDiffMaps(original, new)
//	// added: {"d": 4}        - key "d" is in new but not in original
//	// removed: {"a": 1}      - key "a" is in original but not in new
//	// changed: {"b": 20}     - key "b" is in both but values differ (2 vs 20)
//	// Note: key "c" is not in any result map because it's unchanged
func MapDiffMaps[K comparable, V comparable](m1, m2 map[K]V) (added, removed, changed map[K]V) {
	added = make(map[K]V)
	removed = make(map[K]V)
	changed = make(map[K]V)

	// Find keys in m1 that are not in m2 or have different values
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok {
			removed[k] = v1
		} else if v1 != v2 {
			changed[k] = v2
		}
	}

	// Find keys in m2 that are not in m1
	for k, v2 := range m2 {
		if _, ok := m1[k]; !ok {
			added[k] = v2
		}
	}

	return added, removed, changed
}

// SetContains checks if a set (implemented as map[T]struct{}) contains a specific element.
//
// Parameters:
//   - set: The set to check
//   - item: The element to look for
//
// Returns:
//   - true if the set contains the element, false otherwise
//
// Example:
//
//	// Create a set
//	set := map[string]struct{}{
//	    "apple":  {},
//	    "banana": {},
//	    "cherry": {},
//	}
//
//	arr.SetContains(set, "banana") // Returns: true
//	arr.SetContains(set, "orange") // Returns: false
func SetContains[T comparable](set map[T]struct{}, item T) bool {
	_, ok := set[item]
	return ok
}

// SetToSlice converts a set (implemented as map[T]struct{}) to a slice.
//
// Parameters:
//   - set: The set to convert
//
// Returns:
//   - A slice containing all elements from the set
//
// Example:
//
//	// Create a set
//	set := map[string]struct{}{
//	    "apple":  {},
//	    "banana": {},
//	    "cherry": {},
//	}
//
//	slice := arr.SetToSlice(set)
//	// slice: []string{"apple", "banana", "cherry"} (order may vary)
func SetToSlice[T comparable](set map[T]struct{}) []T {
	result := make([]T, 0, len(set))
	for item := range set {
		result = append(result, item)
	}
	return result
}

// SliceToSet converts a slice to a set (implemented as map[T]struct{}).
// This is the inverse operation of SetToSlice.
//
// Parameters:
//   - slice: The slice to convert
//
// Returns:
//   - A set containing all unique elements from the slice
//
// Example:
//
//	// Create a slice with duplicate elements
//	slice := []string{"apple", "banana", "apple", "cherry", "banana"}
//
//	set := arr.SliceToSet(slice)
//	// set: map[string]struct{}{"apple": {}, "banana": {}, "cherry": {}}
//	// Note: duplicates are automatically removed
func SliceToSet[T comparable](slice []T) map[T]struct{} {
	result := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		result[item] = struct{}{}
	}
	return result
}

// SetUnion creates a new set containing all elements from both input sets.
//
// Parameters:
//   - set1: The first set
//   - set2: The second set
//
// Returns:
//   - A new set containing all elements that appear in either set1 or set2
//
// Example:
//
//	set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
//	set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}
//
//	union := arr.SetUnion(set1, set2)
//	// union: {"a": {}, "b": {}, "c": {}, "d": {}}
func SetUnion[T comparable](set1, set2 map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{}, len(set1)+len(set2))
	for item := range set1 {
		result[item] = struct{}{}
	}
	for item := range set2 {
		result[item] = struct{}{}
	}
	return result
}

// SetIntersection creates a new set containing only elements that exist in both input sets.
//
// Parameters:
//   - set1: The first set
//   - set2: The second set
//
// Returns:
//   - A new set containing only elements that appear in both set1 and set2
//
// Example:
//
//	set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
//	set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}
//
//	intersection := arr.SetIntersection(set1, set2)
//	// intersection: {"b": {}, "c": {}}
//
// Note:
//   - The function optimizes performance by iterating over the smaller set
func SetIntersection[T comparable](set1, set2 map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{})

	// Use the smaller set for iteration
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	for item := range set1 {
		if _, ok := set2[item]; ok {
			result[item] = struct{}{}
		}
	}
	return result
}

// SetDifference creates a new set containing elements that are in the first set but not in the second set.
//
// Parameters:
//   - set1: The first set (source set)
//   - set2: The second set (elements to exclude)
//
// Returns:
//   - A new set containing elements that are in set1 but not in set2
//
// Example:
//
//	set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
//	set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}
//
//	difference := arr.SetDifference(set1, set2)
//	// difference: {"a": {}}
//
//	// Note that the difference is not symmetric:
//	difference2 := arr.SetDifference(set2, set1)
//	// difference2: {"d": {}}
func SetDifference[T comparable](set1, set2 map[T]struct{}) map[T]struct{} {
	result := make(map[T]struct{})
	for item := range set1 {
		if _, ok := set2[item]; !ok {
			result[item] = struct{}{}
		}
	}
	return result
}
