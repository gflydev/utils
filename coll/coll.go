// Package coll provides utility functions for collection manipulation.
// In Go, collections can be slices or maps.
package coll

import (
	"github.com/gflydev/utils/arr"
	"github.com/gflydev/utils/num"
	"sort"
)

// CountBy counts the occurrences of elements in a collection based on the iteratee function.
// Example: CountBy([]int{1, 2, 3, 4}, func(n int) string { return "even" if n % 2 == 0 else "odd" })
// -> map[string]int{"odd": 2, "even": 2}
func CountBy[T any, K comparable](collection []T, iteratee func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range collection {
		key := iteratee(item)
		result[key]++
	}
	return result
}

// Every checks if all elements in the collection satisfy the predicate.
// Example: Every([]int{2, 4, 6}, func(n int) bool { return n % 2 == 0 }) -> true
func Every[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Filter filters elements of a collection that satisfy the predicate.
// Example: Filter([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) -> []int{2, 4}
func Filter[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Find finds the first element in the collection that satisfies the predicate.
// Example: Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 3, true
func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	var zero T
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}
	return zero, false
}

// FindLast finds the last element in the collection that satisfies the predicate.
// Example: FindLast([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) -> 4, true
func FindLast[T any](collection []T, predicate func(T) bool) (T, bool) {
	var zero T
	for i := len(collection) - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], true
		}
	}
	return zero, false
}

// ForEach iterates over elements of a collection and invokes iteratee for each element.
// Example: ForEach([]int{1, 2, 3}, func(n int) { fmt.Println(n) })
func ForEach[T any](collection []T, iteratee func(T)) {
	for _, item := range collection {
		iteratee(item)
	}
}

// ForEachWithIndex iterates over elements of a collection and invokes iteratee for each element with its index.
// Example: ForEachWithIndex([]int{1, 2, 3}, func(n int, i int) { fmt.Println(i, n) })
func ForEachWithIndex[T any](collection []T, iteratee func(T, int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// GroupBy creates an object composed of keys generated from the results of running each element of collection through iteratee.
// Example: GroupBy([]int{1, 2, 3, 4}, func(n int) string { return "even" if n % 2 == 0 else "odd" })
// -> map[string][]int{"odd": {1, 3}, "even": {2, 4}}
func GroupBy[T any, K comparable](collection []T, iteratee func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Includes checks if a value is in the collection.
// Example: Includes([]int{1, 2, 3}, 2) -> true
func Includes[T comparable](collection []T, value T) bool {
	return arr.Includes(collection, value)
}

// KeyBy creates an object composed of keys generated from the results of running each element of collection through iteratee.
// Example: KeyBy([]struct{ID int, Name string}{{1, "Alice"}, {2, "Bob"}}, func(u User) int { return u.ID })
// -> map[int]struct{ID int, Name string}{1: {1, "Alice"}, 2: {2, "Bob"}}
func KeyBy[T any, K comparable](collection []T, iteratee func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range collection {
		key := iteratee(item)
		result[key] = item
	}
	return result
}

// Map creates an array of values by running each element in collection through iteratee.
// Example: Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) -> []int{2, 4, 6}
func Map[T any, R any](collection []T, iteratee func(T) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item)
	}
	return result
}

// MapWithIndex creates an array of values by running each element in collection through iteratee with its index.
// Example: MapWithIndex([]int{1, 2, 3}, func(n, i int) int { return n * i }) -> []int{0, 2, 6}
func MapWithIndex[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Partition splits a collection into two groups, the first of which contains elements that satisfy the predicate.
// Example: Partition([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) -> [][]int{{2, 4}, {1, 3}}
func Partition[T any](collection []T, predicate func(T) bool) [][]T {
	trueResult := make([]T, 0)
	falseResult := make([]T, 0)

	for _, item := range collection {
		if predicate(item) {
			trueResult = append(trueResult, item)
		} else {
			falseResult = append(falseResult, item)
		}
	}

	return [][]T{trueResult, falseResult}
}

// Reduce reduces a collection to a value by iterating through the collection and applying an accumulator function.
// Example: Reduce([]int{1, 2, 3}, func(sum, n int) int { return sum + n }, 0) -> 6
func Reduce[T any, R any](collection []T, iteratee func(R, T) R, accumulator R) R {
	result := accumulator
	for _, item := range collection {
		result = iteratee(result, item)
	}
	return result
}

// ReduceRight reduces a collection to a value by iterating through the collection from right to left.
// Example: ReduceRight([]int{1, 2, 3}, func(result, n int) int { return result * 10 + n }, 0) -> 321
func ReduceRight[T any, R any](collection []T, iteratee func(R, T) R, accumulator R) R {
	result := accumulator
	for i := len(collection) - 1; i >= 0; i-- {
		result = iteratee(result, collection[i])
	}
	return result
}

// Reject is the opposite of Filter; it returns elements that don't satisfy the predicate.
// Example: Reject([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) -> []int{1, 3}
func Reject[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range collection {
		if !predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Sample gets a random element from a collection.
// Example: Sample([]int{1, 2, 3, 4}) -> a random element from the collection
func Sample[T any](collection []T) (T, bool) {
	var zero T
	if len(collection) == 0 {
		return zero, false
	}

	// Use the num package's RandomInt function to get a random index
	randomIndex := num.Random(0, len(collection)-1)
	return collection[randomIndex], true
}

// SampleSize gets n random elements from a collection.
// Example: SampleSize([]int{1, 2, 3, 4}, 2) -> []int{2, 4} (random elements)
func SampleSize[T any](collection []T, n int) []T {
	if len(collection) == 0 || n <= 0 {
		return []T{}
	}

	// If n is greater than the collection size, return a shuffled copy of the collection
	if n >= len(collection) {
		return arr.Shuffle(collection)
	}

	// Create a copy of the collection to avoid modifying the original
	result := make([]T, len(collection))
	copy(result, collection)

	// Shuffle the copy
	result = arr.Shuffle(result)

	// Return the first n elements
	return result[:n]
}

// Size returns the size of a collection.
// Example: Size([]int{1, 2, 3}) -> 3
func Size[T any](collection []T) int {
	return len(collection)
}

// Some checks if any element in the collection satisfies the predicate.
// Example: Some([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 }) -> true
func Some[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// SortBy sorts a collection by the results of running each element through iteratee.
// Example: SortBy([]int{1, 3, 2}, func(n int) int { return n }) -> []int{1, 2, 3}
func SortBy[T any, U int | int8 | int16 | int32 | int64 | float32 | float64 | string](collection []T, iteratee func(T) U) []T {
	return arr.SortBy(collection, iteratee)
}

// OrderBy sorts a collection by multiple iteratees.
// Example: OrderBy([]User{{Name: "fred", Age: 48}, {Name: "barney", Age: 34}}, []string{"age", "name"}, []bool{true, false})
// This is a complex function that's hard to implement in Go due to type system limitations.
// For now, we'll just provide a simplified version that sorts by a single iteratee.
func OrderBy[T any, U int | int8 | int16 | int32 | int64 | float32 | float64 | string](collection []T, iteratee func(T) U, ascending bool) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	sort.Slice(result, func(i, j int) bool {
		if ascending {
			return iteratee(result[i]) < iteratee(result[j])
		}
		return iteratee(result[i]) > iteratee(result[j])
	})

	return result
}

// ForEachMap iterates over elements of a map and invokes iteratee for each element.
// Example: ForEachMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) { fmt.Println(k, v) })
func ForEachMap[K comparable, V any](collection map[K]V, iteratee func(V, K)) {
	for k, v := range collection {
		iteratee(v, k)
	}
}

// MapMap creates an array of values by running each element in a map through iteratee.
// Example: MapMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) string { return k + strconv.Itoa(v) }) -> []string{"a1", "b2"}
func MapMap[K comparable, V any, R any](collection map[K]V, iteratee func(V, K) R) []R {
	result := make([]R, 0, len(collection))
	for k, v := range collection {
		result = append(result, iteratee(v, k))
	}
	return result
}

// FilterMap filters elements of a map that satisfy the predicate.
// Example: FilterMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) bool { return v > 1 }) -> map[string]int{"b": 2}
func FilterMap[K comparable, V any](collection map[K]V, predicate func(V, K) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range collection {
		if predicate(v, k) {
			result[k] = v
		}
	}
	return result
}

// ReduceMap reduces a map to a value by iterating through the map and applying an accumulator function.
// Example: ReduceMap(map[string]int{"a": 1, "b": 2}, func(sum int, v int, k string) int { return sum + v }, 0) -> 3
func ReduceMap[K comparable, V any, R any](collection map[K]V, iteratee func(R, V, K) R, accumulator R) R {
	result := accumulator
	for k, v := range collection {
		result = iteratee(result, v, k)
	}
	return result
}
