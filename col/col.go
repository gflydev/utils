// Package col provides utility functions for collection manipulation.
// In Go, collections can be slices or maps.
package col

import (
	"github.com/gflydev/utils/arr"
	"github.com/gflydev/utils/num"
	"math/rand/v2"
	"sort"
)

// CountBy counts the occurrences of elements in a collection based on the iteratee function.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function that returns the key to group by
//
// Returns:
//   - map[K]int: A map where keys are the values returned by iteratee and values are counts
//
// Example:
//
//	CountBy([]int{1, 2, 3, 4}, func(n int) string {
//	    if n % 2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	})
//	// Returns: map[string]int{"odd": 2, "even": 2}
func CountBy[T any, K comparable](collection []T, iteratee func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range collection {
		key := iteratee(item)
		result[key]++
	}
	return result
}

// Every checks if all elements in the collection satisfy the predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for elements to include
//
// Returns:
//   - bool: True if all elements satisfy the predicate, false otherwise
//
// Example:
//
//	Every([]int{2, 4, 6}, func(n int) bool { return n % 2 == 0 })
//	// Returns: true
func Every[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Filter filters elements of a collection that satisfy the predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for elements to include
//
// Returns:
//   - []T: A new slice containing only the elements that satisfy the predicate
//
// Example:
//
//	Filter([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 })
//	// Returns: []int{2, 4}
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
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for the element to find
//
// Returns:
//   - T: The first element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
//	// Returns: 3, true
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
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for the element to find
//
// Returns:
//   - T: The last element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	FindLast([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
//	// Returns: 4, true
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
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to invoke for each element
//
// Example:
//
//	ForEach([]int{1, 2, 3}, func(n int) {
//	    fmt.Println(n)
//	})
//	// Prints: 1, 2, 3
func ForEach[T any](collection []T, iteratee func(T)) {
	for _, item := range collection {
		iteratee(item)
	}
}

// ForEachWithIndex iterates over elements of a collection and invokes iteratee for each element with its index.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to invoke for each element with its index
//
// Example:
//
//	ForEachWithIndex([]int{1, 2, 3}, func(n int, i int) {
//	    fmt.Println(i, n)
//	})
//	// Prints: 0 1, 1 2, 2 3
func ForEachWithIndex[T any](collection []T, iteratee func(T, int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// GroupBy creates an object composed of keys generated from the results of running each element of collection through iteratee.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function that returns the key to group by
//
// Returns:
//   - map[K][]T: A map where keys are the values returned by iteratee and values are slices of elements
//
// Example:
//
//	GroupBy([]int{1, 2, 3, 4}, func(n int) string {
//	    if n % 2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	})
//	// Returns: map[string][]int{"odd": {1, 3}, "even": {2, 4}}
func GroupBy[T any, K comparable](collection []T, iteratee func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range collection {
		key := iteratee(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Includes checks if a value is in the collection.
//
// Parameters:
//   - collection: The slice to process
//   - value: The value to check for
//
// Returns:
//   - bool: True if the value is in the collection, false otherwise
//
// Example:
//
//	Includes([]int{1, 2, 3}, 2)
//	// Returns: true
func Includes[T comparable](collection []T, value T) bool {
	return arr.Includes(collection, value)
}

// KeyBy creates an object composed of keys generated from the results of running each element of collection through iteratee.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function that returns the key for each element
//
// Returns:
//   - map[K]T: A map where keys are the values returned by iteratee and values are the original elements
//
// Example:
//
//	type User struct {
//	    ID int
//	    Name string
//	}
//	users := []User{{1, "Alice"}, {2, "Bob"}}
//	KeyBy(users, func(u User) int { return u.ID })
//	// Returns: map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
func KeyBy[T any, K comparable](collection []T, iteratee func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range collection {
		key := iteratee(item)
		result[key] = item
	}
	return result
}

// Map creates an array of values by running each element in collection through iteratee.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to transform each element
//
// Returns:
//   - []R: A new slice containing the transformed elements
//
// Example:
//
//	Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
//	// Returns: []int{2, 4, 6}
func Map[T any, R any](collection []T, iteratee func(T) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item)
	}
	return result
}

// MapWithIndex creates an array of values by running each element in collection through iteratee with its index.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to transform each element with its index
//
// Returns:
//   - []R: A new slice containing the transformed elements
//
// Example:
//
//	MapWithIndex([]int{1, 2, 3}, func(n, i int) int { return n * i })
//	// Returns: []int{0, 2, 6}
func MapWithIndex[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Partition splits a collection into two groups, the first of which contains elements that satisfy the predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for elements to include in the first group
//
// Returns:
//   - [][]T: A slice containing two slices: the first with elements that satisfy the predicate,
//     the second with elements that don't satisfy the predicate
//
// Example:
//
//	Partition([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 })
//	// Returns: [][]int{{2, 4}, {1, 3}}
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
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to apply to each element with the accumulator
//   - accumulator: The initial value of the accumulator
//
// Returns:
//   - R: The final accumulated value
//
// Example:
//
//	Reduce([]int{1, 2, 3}, func(sum, n int) int { return sum + n }, 0)
//	// Returns: 6
func Reduce[T any, R any](collection []T, iteratee func(R, T) R, accumulator R) R {
	result := accumulator
	for _, item := range collection {
		result = iteratee(result, item)
	}
	return result
}

// ReduceRight reduces a collection to a value by iterating through the collection from right to left.
//
// Parameters:
//   - collection: The slice to process
//   - iteratee: The function to apply to each element with the accumulator
//   - accumulator: The initial value of the accumulator
//
// Returns:
//   - R: The final accumulated value
//
// Example:
//
//	ReduceRight([]int{1, 2, 3}, func(result, n int) int { return result * 10 + n }, 0)
//	// Returns: 321
func ReduceRight[T any, R any](collection []T, iteratee func(R, T) R, accumulator R) R {
	result := accumulator
	for i := len(collection) - 1; i >= 0; i-- {
		result = iteratee(result, collection[i])
	}
	return result
}

// Reject is the opposite of Filter; it returns elements that don't satisfy the predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for elements to exclude
//
// Returns:
//   - []T: A new slice containing only the elements that don't satisfy the predicate
//
// Example:
//
//	Reject([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 })
//	// Returns: []int{1, 3}
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
//
// Parameters:
//   - collection: The slice to process
//
// Returns:
//   - T: A random element from the collection
//   - bool: True if an element was returned, false if the collection was empty
//
// Example:
//
//	element, found := Sample([]int{1, 2, 3, 4})
//	// Returns: a random element from the collection and true
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
//
// Parameters:
//   - collection: The slice to process
//   - n: The number of random elements to return
//
// Returns:
//   - []T: A slice containing n random elements from the collection
//
// Example:
//
//	SampleSize([]int{1, 2, 3, 4}, 2)
//	// Returns: []int{3, 1} (random elements)
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
//
// Parameters:
//   - collection: The slice to measure
//
// Returns:
//   - int: The number of elements in the collection
//
// Example:
//
//	Size([]int{1, 2, 3})
//	// Returns: 3
func Size[T any](collection []T) int {
	return len(collection)
}

// Some checks if any element in the collection satisfies the predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for elements to check
//
// Returns:
//   - bool: True if any element satisfies the predicate, false otherwise
//
// Example:
//
//	Some([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 })
//	// Returns: true
func Some[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// SortBy sorts a collection by the results of running each element through iteratee.
//
// Parameters:
//   - collection: The slice to sort
//   - iteratee: The function that returns the value to sort by
//
// Returns:
//   - []T: A new sorted slice
//
// Example:
//
//	SortBy([]int{1, 3, 2}, func(n int) int { return n })
//	// Returns: []int{1, 2, 3}
func SortBy[T any, U int | int8 | int16 | int32 | int64 | float32 | float64 | string](collection []T, iteratee func(T) U) []T {
	return arr.SortBy(collection, iteratee)
}

// OrderBy sorts a collection by a single iteratee with specified sort direction.
//
// Parameters:
//   - collection: The slice to sort
//   - iteratee: The function that returns the value to sort by
//   - ascending: Whether to sort in ascending (true) or descending (false) order
//
// Returns:
//   - []T: A new sorted slice
//
// Example:
//
//	type User struct {
//	    Name string
//	    Age int
//	}
//	users := []User{{Name: "fred", Age: 48}, {Name: "barney", Age: 34}}
//	OrderBy(users, func(u User) int { return u.Age }, true)
//	// Returns: []User{{Name: "barney", Age: 34}, {Name: "fred", Age: 48}}
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
//
// Parameters:
//   - collection: The map to process
//   - iteratee: The function to invoke for each element with its key
//
// Example:
//
//	ForEachMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) {
//	    fmt.Println(k, v)
//	})
//	// Prints: a 1, b 2
func ForEachMap[K comparable, V any](collection map[K]V, iteratee func(V, K)) {
	for k, v := range collection {
		iteratee(v, k)
	}
}

// MapMap creates an array of values by running each element in a map through iteratee.
//
// Parameters:
//   - collection: The map to process
//   - iteratee: The function to transform each element with its key
//
// Returns:
//   - []R: A slice containing the transformed elements
//
// Example:
//
//	MapMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) string {
//	    return k + strconv.Itoa(v)
//	})
//	// Returns: []string{"a1", "b2"}
func MapMap[K comparable, V any, R any](collection map[K]V, iteratee func(V, K) R) []R {
	result := make([]R, 0, len(collection))
	for k, v := range collection {
		result = append(result, iteratee(v, k))
	}
	return result
}

// FilterMap filters elements of a map that satisfy the predicate.
//
// Parameters:
//   - collection: The map to process
//   - predicate: The function that returns true for elements to include
//
// Returns:
//   - map[K]V: A new map containing only the elements that satisfy the predicate
//
// Example:
//
//	FilterMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) bool { return v > 1 })
//	// Returns: map[string]int{"b": 2}
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
//
// Parameters:
//   - collection: The map to process
//   - iteratee: The function to apply to each element with the accumulator and key
//   - accumulator: The initial value of the accumulator
//
// Returns:
//   - R: The final accumulated value
//
// Example:
//
//	ReduceMap(map[string]int{"a": 1, "b": 2}, func(sum int, v int, k string) int {
//	    return sum + v
//	}, 0)
//	// Returns: 3
func ReduceMap[K comparable, V any, R any](collection map[K]V, iteratee func(R, V, K) R, accumulator R) R {
	result := accumulator
	for k, v := range collection {
		result = iteratee(result, v, k)
	}
	return result
}

// Avg returns the average value of a collection using the provided value function.
//
// Parameters:
//   - collection: The slice to process
//   - valueFunc: The function that returns the numeric value for each element
//
// Returns:
//   - float64: The average value of the collection
//
// Example:
//
//	Avg([]int{1, 2, 3, 4}, func(n int) float64 { return float64(n) })
//	// Returns: 2.5
func Avg[T any](collection []T, valueFunc func(T) float64) float64 {
	if len(collection) == 0 {
		return 0
	}

	sum := 0.0
	for _, item := range collection {
		sum += valueFunc(item)
	}

	return sum / float64(len(collection))
}

// Chunk breaks the collection into multiple, smaller collections of a given size.
//
// Parameters:
//   - collection: The slice to chunk
//   - size: The size of each chunk
//
// Returns:
//   - [][]T: A slice of slices, each of the specified size (except possibly the last one)
//
// Example:
//
//	Chunk([]int{1, 2, 3, 4, 5}, 2)
//	// Returns: [][]int{{1, 2}, {3, 4}, {5}}
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	chunks := make([][]T, 0, (len(collection)+size-1)/size)

	for i := 0; i < len(collection); i += size {
		end := i + size
		if end > len(collection) {
			end = len(collection)
		}
		chunks = append(chunks, collection[i:end])
	}

	return chunks
}

// Collapse collapses a collection of arrays into a single, flat collection.
//
// Parameters:
//   - collection: The slice of slices to collapse
//
// Returns:
//   - []T: A single flattened slice containing all elements from the input slices
//
// Example:
//
//	Collapse([][]int{{1, 2}, {3, 4}})
//	// Returns: []int{1, 2, 3, 4}
func Collapse[T any](collection [][]T) []T {
	totalLen := 0
	for _, slice := range collection {
		totalLen += len(slice)
	}

	result := make([]T, 0, totalLen)
	for _, slice := range collection {
		result = append(result, slice...)
	}
	return result
}

// Contains determines whether the collection contains a given item.
//
// Parameters:
//   - collection: The slice to search
//   - item: The item to search for
//
// Returns:
//   - bool: True if the item is in the collection, false otherwise
//
// Example:
//
//	Contains([]int{1, 2, 3}, 2)
//	// Returns: true
func Contains[T comparable](collection []T, item T) bool {
	for _, value := range collection {
		if value == item {
			return true
		}
	}
	return false
}

// ContainsFn determines whether the collection contains an item that satisfies the given predicate.
//
// Parameters:
//   - collection: The slice to process
//   - predicate: The function that returns true for the element to check
//
// Returns:
//   - bool: True if any element satisfies the predicate, false otherwise
//
// Example:
//
//	ContainsFn([]int{1, 2, 3}, func(n int) bool { return n > 2 })
//	// Returns: true
func ContainsFn[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Count returns the total number of items in the collection.
//
// Parameters:
//   - collection: The slice to count
//
// Returns:
//   - int: The number of elements in the collection
//
// Example:
//
//	Count([]int{1, 2, 3})
//	// Returns: 3
func Count[T any](collection []T) int {
	return len(collection)
}

// CrossJoin cross joins the collection with the given arrays or collections.
//
// Parameters:
//   - collection: The base slice to cross join
//   - arrays: Variable number of slices to cross join with the collection
//
// Returns:
//   - [][]T: A new slice containing all possible combinations of items from the collection and arrays
//
// Example:
//
//	CrossJoin([]int{1, 2}, []int{3, 4})
//	// Returns: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}}
func CrossJoin[T any](collection []T, arrays ...[]T) [][]T {
	if len(collection) == 0 {
		return [][]T{}
	}

	// Start with the original collection as single-item arrays
	result := make([][]T, len(collection))
	for i, item := range collection {
		result[i] = []T{item}
	}

	// Cross join with each additional array
	for _, array := range arrays {
		newResult := make([][]T, 0, len(result)*len(array))
		for _, item := range result {
			for _, value := range array {
				newItem := make([]T, len(item)+1)
				copy(newItem, item)
				newItem[len(item)] = value
				newResult = append(newResult, newItem)
			}
		}
		result = newResult
	}

	return result
}

// Diff compares the collection against another collection or array and returns the values in the collection
// that are not present in the given items.
//
// Parameters:
//   - collection: The base slice to compare
//   - items: The slice to compare against
//
// Returns:
//   - []T: A new slice containing the values from collection that are not present in items
//
// Example:
//
//	Diff([]int{1, 2, 3}, []int{2, 3, 4})
//	// Returns: []int{1}
func Diff[T comparable](collection, items []T) []T {
	// Create a map for faster lookup
	itemMap := make(map[T]struct{})
	for _, item := range items {
		itemMap[item] = struct{}{}
	}

	// Keep elements from the collection that are not in items
	result := make([]T, 0)
	for _, item := range collection {
		if _, exists := itemMap[item]; !exists {
			result = append(result, item)
		}
	}

	return result
}

// DiffAssoc compares the collection against another collection or array based on its keys and values.
// Returns the key-value pairs in the collection that are not present in the given items or have different values.
//
// Parameters:
//   - collection: The base map to compare
//   - items: The map to compare against
//
// Returns:
//   - map[K]V: A new map containing the key-value pairs from collection that are not present in items or have different values
//
// Example:
//
//	DiffAssoc(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 3})
//	// Returns: map[string]int{"b": 2}
func DiffAssoc[K comparable, V comparable](collection, items map[K]V) map[K]V {
	result := make(map[K]V)

	for key, value := range collection {
		if itemValue, exists := items[key]; !exists || itemValue != value {
			result[key] = value
		}
	}

	return result
}

// DiffKeys compares the collection against another collection or array based on its keys.
// Returns the key-value pairs in the collection where the keys are not present in the given items.
//
// Parameters:
//   - collection: The base map to compare
//   - items: The map to compare against
//
// Returns:
//   - map[K]V: A new map containing the key-value pairs from collection where the keys are not present in items
//
// Example:
//
//	DiffKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 3})
//	// Returns: map[string]int{"b": 2}
func DiffKeys[K comparable, V any](collection, items map[K]V) map[K]V {
	result := make(map[K]V)

	for key, value := range collection {
		if _, exists := items[key]; !exists {
			result[key] = value
		}
	}

	return result
}

// Each iterates over the collection and passes each item to the given callback.
// The iteration stops if the callback returns false.
//
// Parameters:
//   - collection: The slice to iterate over
//   - callback: The function to call for each element, receives the element and its index.
//     Return false to stop iteration, true to continue.
//
// Example:
//
//	Each([]int{1, 2, 3}, func(n int, i int) bool {
//	    fmt.Println(n)
//	    return true // continue iteration
//	})
//	// Prints: 1, 2, 3
func Each[T any](collection []T, callback func(T, int) bool) {
	for i, item := range collection {
		if !callback(item, i) {
			break
		}
	}
}

// Except returns all items in the collection except for those with the specified keys.
//
// Parameters:
//   - collection: The map to filter
//   - keys: The keys to exclude from the result
//
// Returns:
//   - map[K]V: A new map containing all key-value pairs from collection except those with keys in the keys slice
//
// Example:
//
//	Except(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"})
//	// Returns: map[string]int{"b": 2}
func Except[K comparable, V any](collection map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	// Create a map for faster lookup
	keysMap := make(map[K]struct{})
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}

	for key, value := range collection {
		if _, exists := keysMap[key]; !exists {
			result[key] = value
		}
	}

	return result
}

// First returns the first element in the collection that passes a given truth test.
//
// Parameters:
//   - collection: The slice to search
//   - predicate: The function that returns true for the element to find
//
// Returns:
//   - T: The first element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	First([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
//	// Returns: 3, true
func First[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// FirstOrDefault returns the first element in the collection or a default value if the collection is empty.
//
// Parameters:
//   - collection: The slice to get the first element from
//   - defaultValue: The value to return if the collection is empty
//
// Returns:
//   - T: The first element in the collection or the default value if the collection is empty
//
// Example:
//
//	FirstOrDefault([]int{}, 0)
//	// Returns: 0
//
//	FirstOrDefault([]int{1, 2, 3}, 0)
//	// Returns: 1
func FirstOrDefault[T any](collection []T, defaultValue T) T {
	if len(collection) > 0 {
		return collection[0]
	}
	return defaultValue
}

// FlatMap iterates through the collection and passes each value to the given callback.
// The callback should return a slice, and all slices are flattened into a single result slice.
//
// Parameters:
//   - collection: The slice to process
//   - callback: The function that maps each element to a slice of elements
//
// Returns:
//   - []R: A new slice containing all elements from the slices returned by the callback
//
// Example:
//
//	FlatMap([]int{1, 2}, func(n int) []int { return []int{n, n * 2} })
//	// Returns: []int{1, 2, 2, 4}
func FlatMap[T any, R any](collection []T, callback func(T) []R) []R {
	result := make([]R, 0)
	for _, item := range collection {
		result = append(result, callback(item)...)
	}
	return result
}

// Flatten flattens a multi-dimensional collection into a single dimension.
//
// Parameters:
//   - collection: The slice of slices to flatten
//
// Returns:
//   - []T: A new slice containing all elements from all slices in the collection
//
// Example:
//
//	Flatten([][]int{{1, 2}, {3, 4}})
//	// Returns: []int{1, 2, 3, 4}
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for _, slice := range collection {
		totalLen += len(slice)
	}

	result := make([]T, 0, totalLen)
	for _, slice := range collection {
		result = append(result, slice...)
	}
	return result
}

// Flip swaps the collection's keys with their corresponding values.
//
// Parameters:
//   - collection: The map whose keys and values will be swapped
//
// Returns:
//   - map[V]K: A new map where the keys are the values from the original map and the values are the keys from the original map
//
// Example:
//
//	Flip(map[string]int{"a": 1, "b": 2})
//	// Returns: map[int]string{1: "a", 2: "b"}
func Flip[K comparable, V comparable](collection map[K]V) map[V]K {
	result := make(map[V]K)
	for key, value := range collection {
		result[value] = key
	}
	return result
}

// Forget removes items from the collection by their keys.
//
// Parameters:
//   - collection: The map to remove items from
//   - keys: The keys of the items to remove
//
// Returns:
//   - map[K]V: A new map containing all key-value pairs from collection except those with keys in the keys parameter
//
// Example:
//
//	Forget(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c")
//	// Returns: map[string]int{"b": 2}
func Forget[K comparable, V any](collection map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)
	for k, v := range collection {
		result[k] = v
	}

	for _, key := range keys {
		delete(result, key)
	}

	return result
}

// Get retrieves an item from the collection by its key.
//
// Parameters:
//   - collection: The map to get the value from
//   - key: The key to look up
//   - defaultValue: The value to return if the key doesn't exist in the collection
//
// Returns:
//   - V: The value associated with the key, or the default value if the key doesn't exist
//
// Example:
//
//	Get(map[string]int{"a": 1, "b": 2}, "a", 0)
//	// Returns: 1
//
//	Get(map[string]int{"a": 1}, "c", 0)
//	// Returns: 0
func Get[K comparable, V any](collection map[K]V, key K, defaultValue V) V {
	if value, exists := collection[key]; exists {
		return value
	}
	return defaultValue
}

// Has determines if a given key exists in the collection.
//
// Parameters:
//   - collection: The map to check
//   - key: The key to look for
//
// Returns:
//   - bool: True if the key exists in the collection, false otherwise
//
// Example:
//
//	Has(map[string]int{"a": 1, "b": 2}, "a")
//	// Returns: true
//
//	Has(map[string]int{"a": 1, "b": 2}, "c")
//	// Returns: false
func Has[K comparable, V any](collection map[K]V, key K) bool {
	_, exists := collection[key]
	return exists
}

// Implode joins the items in a collection into a single string.
//
// Parameters:
//   - collection: The slice to join
//   - separator: The string to place between elements
//   - toString: A function that converts each element to a string
//
// Returns:
//   - string: A string containing all elements joined by the separator
//
// Example:
//
//	Implode([]int{1, 2, 3}, ", ", func(n int) string { return strconv.Itoa(n) })
//	// Returns: "1, 2, 3"
func Implode[T any](collection []T, separator string, toString func(T) string) string {
	if len(collection) == 0 {
		return ""
	}

	result := toString(collection[0])
	for i := 1; i < len(collection); i++ {
		result += separator + toString(collection[i])
	}
	return result
}

// Intersect removes any values from the original collection that are not present in the given array or collection.
//
// Parameters:
//   - collection: The base slice to filter
//   - items: The slice to compare against
//
// Returns:
//   - []T: A new slice containing only the elements that are present in both collection and items
//
// Example:
//
//	Intersect([]int{1, 2, 3}, []int{2, 3, 4})
//	// Returns: []int{2, 3}
func Intersect[T comparable](collection, items []T) []T {
	// Create a map for faster lookup
	itemMap := make(map[T]struct{})
	for _, item := range items {
		itemMap[item] = struct{}{}
	}

	// Keep elements from the collection that are in items
	result := make([]T, 0)
	for _, item := range collection {
		if _, exists := itemMap[item]; exists {
			result = append(result, item)
		}
	}

	return result
}

// IntersectByKeys removes any keys from the original collection that are not present in the given array or collection.
//
// Parameters:
//   - collection: The base map to filter
//   - keys: The slice of keys to keep
//
// Returns:
//   - map[K]V: A new map containing only the key-value pairs from collection where the key is in the keys slice
//
// Example:
//
//	IntersectByKeys(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"})
//	// Returns: map[string]int{"a": 1, "c": 3}
func IntersectByKeys[K comparable, V any](collection map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	// Create a map for faster lookup
	keysMap := make(map[K]struct{})
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}

	for key, value := range collection {
		if _, exists := keysMap[key]; exists {
			result[key] = value
		}
	}

	return result
}

// IsEmpty determines if the collection is empty.
//
// Parameters:
//   - collection: The slice to check
//
// Returns:
//   - bool: True if the collection is empty, false otherwise
//
// Example:
//
//	IsEmpty([]int{})
//	// Returns: true
//
//	IsEmpty([]int{1, 2})
//	// Returns: false
func IsEmpty[T any](collection []T) bool {
	return len(collection) == 0
}

// IsNotEmpty determines if the collection is not empty.
//
// Parameters:
//   - collection: The slice to check
//
// Returns:
//   - bool: True if the collection is not empty, false otherwise
//
// Example:
//
//	IsNotEmpty([]int{1, 2})
//	// Returns: true
//
//	IsNotEmpty([]int{})
//	// Returns: false
func IsNotEmpty[T any](collection []T) bool {
	return len(collection) > 0
}

// Keys returns all of the collection's keys.
//
// Parameters:
//   - collection: The map to extract keys from
//
// Returns:
//   - []K: A slice containing all keys from the collection
//
// Example:
//
//	Keys(map[string]int{"a": 1, "b": 2})
//	// Returns: []string{"a", "b"}
func Keys[K comparable, V any](collection map[K]V) []K {
	keys := make([]K, 0, len(collection))
	for key := range collection {
		keys = append(keys, key)
	}
	return keys
}

// Last returns the last element in the collection that passes a given truth test.
//
// Parameters:
//   - collection: The slice to search
//   - predicate: The function that returns true for the element to find
//
// Returns:
//   - T: The last element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	Last([]int{1, 2, 3, 4}, func(n int) bool { return n < 3 })
//	// Returns: 2, true
func Last[T any](collection []T, predicate func(T) bool) (T, bool) {
	for i := len(collection) - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], true
		}
	}
	var zero T
	return zero, false
}

// LastOrDefault returns the last element in the collection or a default value if the collection is empty.
//
// Parameters:
//   - collection: The slice to get the last element from
//   - defaultValue: The value to return if the collection is empty
//
// Returns:
//   - T: The last element in the collection or the default value if the collection is empty
//
// Example:
//
//	LastOrDefault([]int{1, 2, 3}, 0)
//	// Returns: 3
//
//	LastOrDefault([]int{}, 0)
//	// Returns: 0
func LastOrDefault[T any](collection []T, defaultValue T) T {
	if len(collection) > 0 {
		return collection[len(collection)-1]
	}
	return defaultValue
}

// Max returns the maximum value of a given key.
//
// Parameters:
//   - collection: The slice to process
//   - valueFunc: The function that extracts a numeric value from each element
//
// Returns:
//   - V: The maximum value found, or zero if the collection is empty
//
// Example:
//
//	Max([]int{1, 2, 3, 4}, func(n int) int { return n })
//	// Returns: 4
//
//	Max([]struct{Age int}{{Age: 25}, {Age: 30}, {Age: 20}}, func(p struct{Age int}) int { return p.Age })
//	// Returns: 30
func Max[T any, V float64 | int | int64 | float32 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](collection []T, valueFunc func(T) V) V {
	if len(collection) == 0 {
		var zero V
		return zero
	}

	max := valueFunc(collection[0])
	for i := 1; i < len(collection); i++ {
		value := valueFunc(collection[i])
		if value > max {
			max = value
		}
	}

	return max
}

// Merge merges the given array or collection with the original collection.
// If a key exists in both collections, the value from items will be used.
//
// Parameters:
//   - collection: The base map to merge into
//   - items: The map to merge from
//
// Returns:
//   - map[K]V: A new map containing all key-value pairs from both maps, with items taking precedence for duplicate keys
//
// Example:
//
//	Merge(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
//	// Returns: map[string]int{"a": 1, "b": 3, "c": 4}
func Merge[K comparable, V any](collection, items map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range collection {
		result[k] = v
	}
	for k, v := range items {
		result[k] = v
	}
	return result
}

// Min returns the minimum value of a given key.
//
// Parameters:
//   - collection: The slice to process
//   - valueFunc: The function that extracts a numeric value from each element
//
// Returns:
//   - V: The minimum value found, or zero if the collection is empty
//
// Example:
//
//	Min([]int{1, 2, 3, 4}, func(n int) int { return n })
//	// Returns: 1
//
//	Min([]struct{Age int}{{Age: 25}, {Age: 30}, {Age: 20}}, func(p struct{Age int}) int { return p.Age })
//	// Returns: 20
func Min[T any, V float64 | int | int64 | float32 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](collection []T, valueFunc func(T) V) V {
	if len(collection) == 0 {
		var zero V
		return zero
	}

	min := valueFunc(collection[0])
	for i := 1; i < len(collection); i++ {
		value := valueFunc(collection[i])
		if value < min {
			min = value
		}
	}

	return min
}

// Only returns the items in the collection with the specified keys.
//
// Parameters:
//   - collection: The map to filter
//   - keys: The keys to keep in the result
//
// Returns:
//   - map[K]V: A new map containing only the key-value pairs from collection where the key is in the keys slice
//
// Example:
//
//	Only(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"})
//	// Returns: map[string]int{"a": 1, "c": 3}
func Only[K comparable, V any](collection map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	// Create a map for faster lookup
	keysMap := make(map[K]struct{})
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}

	for key, value := range collection {
		if _, exists := keysMap[key]; exists {
			result[key] = value
		}
	}

	return result
}

// Pad fills the array to the specified size with a value.
//
// Parameters:
//   - collection: The slice to pad
//   - size: The target size of the resulting slice
//   - value: The value to use for padding
//
// Returns:
//   - []T: A new slice padded to the specified size with the given value,
//     or the original slice if size is less than or equal to the length of the collection
//
// Example:
//
//	Pad([]int{1, 2}, 4, 0)
//	// Returns: []int{1, 2, 0, 0}
func Pad[T any](collection []T, size int, value T) []T {
	if size <= len(collection) {
		return collection
	}

	result := make([]T, size)
	copy(result, collection)
	for i := len(collection); i < size; i++ {
		result[i] = value
	}
	return result
}

// Pluck retrieves all of the values for a given key.
//
// Parameters:
//   - collection: The slice to process
//   - key: The function that extracts a value from each element
//
// Returns:
//   - []V: A new slice containing the values extracted from each element in the collection
//
// Example:
//
//	Pluck([]struct{Name string}{{"Alice"}, {"Bob"}}, func(p struct{Name string}) string { return p.Name })
//	// Returns: []string{"Alice", "Bob"}
func Pluck[T any, V any](collection []T, key func(T) V) []V {
	result := make([]V, len(collection))
	for i, item := range collection {
		result[i] = key(item)
	}
	return result
}

// Prepend adds items to the beginning of the collection.
//
// Parameters:
//   - collection: The slice to prepend to
//   - values: The values to add to the beginning of the collection
//
// Returns:
//   - []T: A new slice with the values prepended to the collection
//
// Example:
//
//	Prepend([]int{3, 4}, 1, 2)
//	// Returns: []int{1, 2, 3, 4}
func Prepend[T any](collection []T, values ...T) []T {
	result := make([]T, len(values)+len(collection))
	copy(result, values)
	copy(result[len(values):], collection)
	return result
}

// Pull removes and returns an item from the collection by key.
//
// Parameters:
//   - collection: The slice to remove an item from
//   - index: The index of the item to remove
//
// Returns:
//   - T: The removed item
//   - []T: A new slice with the item removed
//
// Example:
//
//	Pull([]int{1, 2, 3}, 1)
//	// Returns: 2, []int{1, 3}
func Pull[T any](collection []T, index int) (T, []T) {
	if index < 0 || index >= len(collection) {
		var zero T
		return zero, collection
	}

	item := collection[index]
	return item, append(collection[:index], collection[index+1:]...)
}

// Push adds items to the end of the collection.
//
// Parameters:
//   - collection: The slice to append to
//   - values: The values to add to the end of the collection
//
// Returns:
//   - []T: A new slice with the values appended to the collection
//
// Example:
//
//	Push([]int{1, 2}, 3, 4)
//	// Returns: []int{1, 2, 3, 4}
func Push[T any](collection []T, values ...T) []T {
	return append(collection, values...)
}

// Put sets the given key and value in the collection.
//
// Parameters:
//   - collection: The map to add the key-value pair to
//   - key: The key to set
//   - value: The value to set
//
// Returns:
//   - map[K]V: A new map with the key-value pair added or updated
//
// Example:
//
//	Put(map[string]int{"a": 1}, "b", 2)
//	// Returns: map[string]int{"a": 1, "b": 2}
func Put[K comparable, V any](collection map[K]V, key K, value V) map[K]V {
	result := make(map[K]V)
	for k, v := range collection {
		result[k] = v
	}
	result[key] = value
	return result
}

// Random retrieves a random item from the collection.
//
// Parameters:
//   - collection: The slice to get a random item from
//
// Returns:
//   - T: A random item from the collection
//   - bool: True if an item was returned, false if the collection is empty
//
// Example:
//
//	Random([]int{1, 2, 3})
//	// Returns: a random element from the slice, true
func Random[T any](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}

	return collection[rand.IntN(len(collection))], true
}

// RandomOrDefault retrieves a random item from the collection or a default value if the collection is empty.
//
// Parameters:
//   - collection: The slice to get a random element from
//   - defaultValue: The value to return if the collection is empty
//
// Returns:
//   - T: A random element from the collection or the default value if the collection is empty
//
// Example:
//
//	RandomOrDefault([]int{1, 2, 3}, 0)
//	// Returns: a random element from the slice (e.g., 2)
//
//	RandomOrDefault([]int{}, 0)
//	// Returns: 0 (the default value)
func RandomOrDefault[T any](collection []T, defaultValue T) T {
	if len(collection) == 0 {
		return defaultValue
	}

	return collection[rand.IntN(len(collection))]
}

// Reverse reverses the order of the collection's items.
//
// Parameters:
//   - collection: The slice to reverse
//
// Returns:
//   - []T: A new slice with the elements in reverse order
//
// Example:
//
//	Reverse([]int{1, 2, 3})
//	// Returns: []int{3, 2, 1}
func Reverse[T any](collection []T) []T {
	result := make([]T, len(collection))
	for i, j := 0, len(collection)-1; j >= 0; i, j = i+1, j-1 {
		result[i] = collection[j]
	}
	return result
}

// Search searches the collection for a given value and returns the corresponding index if successful.
//
// Parameters:
//   - collection: The slice to search in
//   - value: The value to search for
//
// Returns:
//   - int: The index of the found element, or -1 if not found
//   - bool: True if the element was found, false otherwise
//
// Example:
//
//	Search([]int{1, 2, 3, 4}, 3)
//	// Returns: 2, true
//
//	Search([]int{1, 2, 3, 4}, 5)
//	// Returns: -1, false
func Search[T comparable](collection []T, value T) (int, bool) {
	for i, item := range collection {
		if item == value {
			return i, true
		}
	}
	return -1, false
}

// SearchFunc searches the collection using the given predicate function and returns the index of the first matching element.
//
// Parameters:
//   - collection: The slice to search in
//   - predicate: The function that returns true for the element to find
//
// Returns:
//   - int: The index of the first element that satisfies the predicate, or -1 if not found
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	SearchFunc([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
//	// Returns: 2, true
//
//	SearchFunc([]int{1, 2}, func(n int) bool { return n > 2 })
//	// Returns: -1, false
func SearchFunc[T any](collection []T, predicate func(T) bool) (int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return i, true
		}
	}
	return -1, false
}

// Shift removes and returns the first item from the collection.
//
// Parameters:
//   - collection: The slice to remove the first element from
//
// Returns:
//   - T: The first element of the collection, or zero value if the collection is empty
//   - []T: The remaining elements of the collection
//
// Example:
//
//	Shift([]int{1, 2, 3})
//	// Returns: 1, []int{2, 3}
//
//	Shift([]int{})
//	// Returns: 0, []int{}
func Shift[T any](collection []T) (T, []T) {
	if len(collection) == 0 {
		var zero T
		return zero, collection
	}

	item := collection[0]
	result := collection[1:]
	return item, result
}

// Shuffle randomly shuffles the items in the collection.
//
// Parameters:
//   - collection: The slice to shuffle
//
// Returns:
//   - []T: A new slice with the elements in random order
//
// Example:
//
//	Shuffle([]int{1, 2, 3})
//	// Returns: []int{2, 3, 1} (random order)
//
//	Shuffle([]int{})
//	// Returns: []int{}
func Shuffle[T any](collection []T) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result
}

// Slice returns a slice of the collection starting at the given index.
//
// Parameters:
//   - collection: The slice to extract a portion from
//   - start: The starting index (can be negative to count from the end)
//
// Returns:
//   - []T: A new slice containing elements from the start index to the end of the collection
//
// Example:
//
//	Slice([]int{1, 2, 3, 4}, 1)
//	// Returns: []int{2, 3, 4}
//
//	Slice([]int{1, 2, 3, 4}, -2)
//	// Returns: []int{3, 4}
//
//	Slice([]int{1, 2, 3, 4}, 5)
//	// Returns: []int{}
func Slice[T any](collection []T, start int) []T {
	if start >= len(collection) {
		return []T{}
	}

	if start < 0 {
		start = len(collection) + start
		if start < 0 {
			start = 0
		}
	}

	return collection[start:]
}

// SliceWithLength returns a slice of the collection starting at the given index with the specified length.
//
// Parameters:
//   - collection: The slice to extract a portion from
//   - start: The starting index (can be negative to count from the end)
//   - length: The number of elements to include in the result
//
// Returns:
//   - []T: A new slice containing the specified number of elements from the start index
//
// Example:
//
//	SliceWithLength([]int{1, 2, 3, 4}, 1, 2)
//	// Returns: []int{2, 3}
//
//	SliceWithLength([]int{1, 2, 3, 4}, -2, 2)
//	// Returns: []int{3, 4}
//
//	SliceWithLength([]int{1, 2, 3, 4}, 0, 10)
//	// Returns: []int{1, 2, 3, 4}
func SliceWithLength[T any](collection []T, start, length int) []T {
	if start >= len(collection) || length <= 0 {
		return []T{}
	}

	if start < 0 {
		start = len(collection) + start
		if start < 0 {
			start = 0
		}
	}

	end := start + length
	if end > len(collection) {
		end = len(collection)
	}

	return collection[start:end]
}

// Sort sorts the collection according to the given comparison function.
//
// Parameters:
//   - collection: The slice to sort
//   - less: The comparison function that defines the sort order.
//     Should return true if the first argument should be ordered before the second.
//
// Returns:
//   - []T: A new sorted slice
//
// Example:
//
//	Sort([]int{3, 1, 4, 2}, func(i, j int) bool { return i < j })
//	// Returns: []int{1, 2, 3, 4}
//
//	Sort([]int{3, 1, 4, 2}, func(i, j int) bool { return i > j })
//	// Returns: []int{4, 3, 2, 1}
func Sort[T any](collection []T, less func(i, j T) bool) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})

	return result
}

// SortByDesc sorts the collection by the given key in descending order.
//
// Parameters:
//   - collection: The slice to sort
//   - keyFunc: The function that extracts the key to sort by from each element
//   - less: The comparison function that defines the sort order.
//     Should return true if the first argument should be ordered before the second.
//
// Returns:
//   - []T: A new slice sorted in descending order by the extracted keys
//
// Example:
//
//	SortByDesc([]int{3, 1, 4, 2}, func(n int) string { return strconv.Itoa(n) }, func(i, j string) bool { return i < j })
//	// Returns: []int{4, 3, 2, 1}
//
//	SortByDesc([]string{"a", "c", "b"}, func(s string) string { return s }, func(i, j string) bool { return i < j })
//	// Returns: []string{"c", "b", "a"}
func SortByDesc[T any, K comparable](collection []T, keyFunc func(T) K, less func(i, j K) bool) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	sort.Slice(result, func(i, j int) bool {
		return less(keyFunc(result[j]), keyFunc(result[i]))
	})

	return result
}

// Splice removes and returns a slice of elements from the collection starting at the given index.
//
// Parameters:
//   - collection: The slice to modify
//   - start: The starting index (can be negative to count from the end)
//   - length: The number of elements to remove
//
// Returns:
//   - []T: A slice containing the removed elements
//   - []T: The modified collection with elements removed
//
// Example:
//
//	Splice([]int{1, 2, 3, 4}, 1, 2)
//	// Returns: []int{2, 3}, []int{1, 4}
//
//	Splice([]int{1, 2, 3}, -1, 1)
//	// Returns: []int{3}, []int{1, 2}
//
//	Splice([]int{1, 2, 3}, 5, 1)
//	// Returns: []int{}, []int{1, 2, 3}
func Splice[T any](collection []T, start, length int) ([]T, []T) {
	if start >= len(collection) || length <= 0 {
		return []T{}, collection
	}

	if start < 0 {
		start = len(collection) + start
		if start < 0 {
			start = 0
		}
	}

	end := start + length
	if end > len(collection) {
		end = len(collection)
	}

	removed := collection[start:end]
	return removed, append(collection[:start], collection[end:]...)
}

// Split breaks a collection into the given number of groups.
//
// Parameters:
//   - collection: The slice to split
//   - numberOfGroups: The number of groups to split the collection into
//
// Returns:
//   - [][]T: A slice of slices, where each inner slice contains elements distributed evenly
//
// Example:
//
//	Split([]int{1, 2, 3, 4, 5, 6}, 3)
//	// Returns: [][]int{{1, 4}, {2, 5}, {3, 6}}
//
//	Split([]int{1, 2, 3, 4, 5}, 2)
//	// Returns: [][]int{{1, 3, 5}, {2, 4}}
//
//	Split([]int{}, 3)
//	// Returns: [][]int{}
func Split[T any](collection []T, numberOfGroups int) [][]T {
	if numberOfGroups <= 0 {
		return [][]T{}
	}

	if len(collection) == 0 {
		return [][]T{}
	}

	result := make([][]T, numberOfGroups)
	for i := 0; i < numberOfGroups; i++ {
		result[i] = make([]T, 0)
	}

	for i, item := range collection {
		result[i%numberOfGroups] = append(result[i%numberOfGroups], item)
	}

	return result
}

// Sum returns the sum of all items in the collection.
//
// Parameters:
//   - collection: The slice to sum
//   - valueFunc: The function that extracts a numeric value from each element
//
// Returns:
//   - V: The sum of all values extracted from the collection
//
// Example:
//
//	Sum([]int{1, 2, 3, 4}, func(n int) int { return n })
//	// Returns: 10
//
//	Sum([]int{1, 2, 3}, func(n int) int { return n * 2 })
//	// Returns: 12
//
//	Sum([]struct{Value int}{{1}, {2}}, func(x struct{Value int}) int { return x.Value })
//	// Returns: 3
func Sum[T any, V float64 | int | int64 | float32 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8](collection []T, valueFunc func(T) V) V {
	var sum V
	for _, item := range collection {
		sum += valueFunc(item)
	}
	return sum
}

// Take returns a new collection with the specified number of items.
//
// Parameters:
//   - collection: The slice to take elements from
//   - limit: The maximum number of elements to take
//
// Returns:
//   - []T: A new slice containing at most the specified number of elements from the beginning of the collection
//
// Example:
//
//	Take([]int{1, 2, 3, 4}, 2)
//	// Returns: []int{1, 2}
//
//	Take([]int{1, 2, 3}, 5)
//	// Returns: []int{1, 2, 3}
//
//	Take([]int{1, 2, 3}, 0)
//	// Returns: []int{}
func Take[T any](collection []T, limit int) []T {
	if limit <= 0 {
		return []T{}
	}

	if limit >= len(collection) {
		return collection
	}

	return collection[:limit]
}

// Tap passes the collection to the given callback then returns the collection.
//
// Parameters:
//   - collection: The slice to pass to the callback
//   - callback: The function to call with the collection
//
// Returns:
//   - []T: The original collection (which may be modified by the callback)
//
// Example:
//
//	Tap([]int{1, 2, 3}, func(x []int) { fmt.Println(x) })
//	// Returns: []int{1, 2, 3} (and prints [1 2 3])
//
//	nums := []int{1,2}
//	Tap(nums, func(x []int) { x[0] = 99 })
//	// Returns: []int{99, 2}
func Tap[T any](collection []T, callback func([]T)) []T {
	callback(collection)
	return collection
}

// Unique returns all of the unique items in the collection.
//
// Parameters:
//   - collection: The slice to remove duplicates from
//
// Returns:
//   - []T: A new slice containing only unique elements, preserving the original order of first occurrence
//
// Example:
//
//	Unique([]int{1, 2, 2, 3, 3, 3})
//	// Returns: []int{1, 2, 3}
//
//	Unique([]string{"a", "a", "b", "c"})
//	// Returns: []string{"a", "b", "c"}
func Unique[T comparable](collection []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// UniqueBy returns all of the unique items in the collection using the given key function.
//
// Parameters:
//   - collection: The slice to remove duplicates from
//   - keyFunc: The function that extracts the key to determine uniqueness
//
// Returns:
//   - []T: A new slice containing only elements with unique keys, preserving the original order of first occurrence
//
// Example:
//
//	UniqueBy([]int{1, 2, 3, 4, 5, 6}, func(n int) int { return n % 3 })
//	// Returns: []int{1, 2, 3} (because 1%3=1, 2%3=2, 3%3=0, 4%3=1, 5%3=2, 6%3=0)
//
//	UniqueBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })
//	// Returns: []string{"one", "three"} (because len("one")=3, len("two")=3, len("three")=5)
func UniqueBy[T any, K comparable](collection []T, keyFunc func(T) K) []T {
	seen := make(map[K]struct{})
	result := make([]T, 0)

	for _, item := range collection {
		key := keyFunc(item)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// Unless executes the given callback when the condition is false.
//
// Parameters:
//   - condition: The boolean condition to check
//   - collection: The slice to pass to the callback if the condition is false
//   - callback: The function to call with the collection if the condition is false
//
// Returns:
//   - []T: The result of the callback if the condition is false, otherwise the original collection
//
// Example:
//
//	Unless(false, []int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3, 4} (because condition is false)
//
//	Unless(true, []int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3} (because condition is true)
func Unless[T any](condition bool, collection []T, callback func([]T) []T) []T {
	if !condition {
		return callback(collection)
	}
	return collection
}

// UnlessEmpty executes the given callback when the collection is not empty.
//
// Parameters:
//   - collection: The slice to check and pass to the callback if not empty
//   - callback: The function to call with the collection if it's not empty
//
// Returns:
//   - []T: The result of the callback if the collection is not empty, otherwise the original collection
//
// Example:
//
//	UnlessEmpty([]int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3, 4} (because collection is not empty)
//
//	UnlessEmpty([]int{}, func(x []int) []int { return append(x, 1) })
//	// Returns: []int{} (because collection is empty)
func UnlessEmpty[T any](collection []T, callback func([]T) []T) []T {
	if len(collection) > 0 {
		return callback(collection)
	}
	return collection
}

// UnlessNotEmpty executes the given callback when the collection is empty.
//
// Parameters:
//   - collection: The slice to check and pass to the callback if empty
//   - callback: The function to call with the collection if it's empty
//
// Returns:
//   - []T: The result of the callback if the collection is empty, otherwise the original collection
//
// Example:
//
//	UnlessNotEmpty([]int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3} (because collection is not empty)
//
//	UnlessNotEmpty([]int{}, func(x []int) []int { return append(x, 1) })
//	// Returns: []int{1} (because collection is empty)
func UnlessNotEmpty[T any](collection []T, callback func([]T) []T) []T {
	if len(collection) == 0 {
		return callback(collection)
	}
	return collection
}

// Values returns all of the values in the map collection.
//
// Parameters:
//   - collection: The map to extract values from
//
// Returns:
//   - []V: A slice containing all values from the map
//
// Example:
//
//	Values(map[string]int{"a": 1, "b": 2, "c": 3})
//	// Returns: []int{1, 2, 3}
//
//	Values(map[string]int{})
//	// Returns: []int{}
func Values[K comparable, V any](collection map[K]V) []V {
	values := make([]V, 0, len(collection))
	for _, value := range collection {
		values = append(values, value)
	}
	return values
}

// When executes the given callback when the condition is true.
//
// Parameters:
//   - condition: The boolean condition to check
//   - collection: The slice to pass to the callback if the condition is true
//   - callback: The function to call with the collection if the condition is true
//
// Returns:
//   - []T: The result of the callback if the condition is true, otherwise the original collection
//
// Example:
//
//	When(true, []int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3, 4} (because condition is true)
//
//	When(false, []int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3} (because condition is false)
func When[T any](condition bool, collection []T, callback func([]T) []T) []T {
	if condition {
		return callback(collection)
	}
	return collection
}

// WhenEmpty executes the given callback when the collection is empty.
//
// Parameters:
//   - collection: The slice to check and pass to the callback if empty
//   - callback: The function to call with the collection if it's empty
//
// Returns:
//   - []T: The result of the callback if the collection is empty, otherwise the original collection
//
// Example:
//
//	WhenEmpty([]int{}, func(x []int) []int { return append(x, 1) })
//	// Returns: []int{1} (because collection is empty)
//
//	WhenEmpty([]int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3} (because collection is not empty)
func WhenEmpty[T any](collection []T, callback func([]T) []T) []T {
	if len(collection) == 0 {
		return callback(collection)
	}
	return collection
}

// WhenNotEmpty executes the given callback when the collection is not empty.
//
// Parameters:
//   - collection: The slice to check and pass to the callback if not empty
//   - callback: The function to call with the collection if it's not empty
//
// Returns:
//   - []T: The result of the callback if the collection is not empty, otherwise the original collection
//
// Example:
//
//	WhenNotEmpty([]int{1, 2, 3}, func(x []int) []int { return append(x, 4) })
//	// Returns: []int{1, 2, 3, 4} (because collection is not empty)
//
//	WhenNotEmpty([]int{}, func(x []int) []int { return append(x, 1) })
//	// Returns: []int{} (because collection is empty)
func WhenNotEmpty[T any](collection []T, callback func([]T) []T) []T {
	if len(collection) > 0 {
		return callback(collection)
	}
	return collection
}

// Where filters the collection using the given predicate function.
//
// Parameters:
//   - collection: The slice to filter
//   - predicate: The function that determines whether an element should be included in the result
//
// Returns:
//   - []T: A new slice containing only the elements that satisfy the predicate
//
// Example:
//
//	Where([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 })
//	// Returns: []int{2, 4}
//
//	Where([]int{1, 3, 5}, func(n int) bool { return n % 2 == 0 })
//	// Returns: []int{}
func Where[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// WhereIn filters the collection by a given key/value contained within the given array.
//
// Parameters:
//   - collection: The slice to filter
//   - keyFunc: The function that extracts the key to check against the values
//   - values: The slice of values to check against
//
// Returns:
//   - []T: A new slice containing only the elements whose extracted keys are in the values slice
//
// Example:
//
//	WhereIn([]int{1, 2, 3, 4}, func(n int) int { return n }, []int{2, 4})
//	// Returns: []int{2, 4}
//
//	WhereIn([]int{1, 2, 3, 4, 5}, func(n int) int { return n % 3 }, []int{0, 1})
//	// Returns: []int{1, 3, 4} (because 1%3=1, 3%3=0, 4%3=1, which are in [0,1])
func WhereIn[T any, K comparable](collection []T, keyFunc func(T) K, values []K) []T {
	// Create a map for faster lookup
	valuesMap := make(map[K]struct{})
	for _, value := range values {
		valuesMap[value] = struct{}{}
	}

	result := make([]T, 0)
	for _, item := range collection {
		key := keyFunc(item)
		if _, exists := valuesMap[key]; exists {
			result = append(result, item)
		}
	}

	return result
}

// WhereNotIn filters the collection by a given key/value not contained within the given array.
//
// Parameters:
//   - collection: The slice to filter
//   - keyFunc: The function that extracts the key to check against the values
//   - values: The slice of values to check against
//
// Returns:
//   - []T: A new slice containing only the elements whose extracted keys are not in the values slice
//
// Example:
//
//	WhereNotIn([]int{1, 2, 3, 4}, func(n int) int { return n }, []int{2, 4})
//	// Returns: []int{1, 3}
//
//	WhereNotIn([]int{1, 2, 3, 4, 5}, func(n int) int { return n % 3 }, []int{0, 1})
//	// Returns: []int{2, 5} (because 2%3=2, 5%3=2, which are not in [0,1])
func WhereNotIn[T any, K comparable](collection []T, keyFunc func(T) K, values []K) []T {
	// Create a map for faster lookup
	valuesMap := make(map[K]struct{})
	for _, value := range values {
		valuesMap[value] = struct{}{}
	}

	result := make([]T, 0)
	for _, item := range collection {
		key := keyFunc(item)
		if _, exists := valuesMap[key]; !exists {
			result = append(result, item)
		}
	}

	return result
}

// Zip merges together the values of the given arrays with the values of the original collection.
//
// Parameters:
//   - collection: The base slice to merge with other arrays
//   - arrays: Variable number of slices to merge with the collection
//
// Returns:
//   - [][]T: A new slice of slices where each inner slice contains elements from the same position across all input slices
//
// Example:
//
//	Zip([]int{1, 2, 3}, [][]int{{4, 5, 6}})
//	// Returns: [][]int{{1, 4}, {2, 5}, {3, 6}}
//
//	Zip([]int{1, 2}, [][]int{{3, 4}, {5, 6}})
//	// Returns: [][]int{{1, 3, 5}, {2, 4, 6}}
func Zip[T any](collection []T, arrays ...[]T) [][]T {
	if len(collection) == 0 {
		return [][]T{}
	}

	// Determine the length of the result (minimum length of all arrays)
	minLength := len(collection)
	for _, array := range arrays {
		if len(array) < minLength {
			minLength = len(array)
		}
	}

	// Create the result array
	result := make([][]T, minLength)
	for i := 0; i < minLength; i++ {
		// Each inner array has a length of 1 (for the collection) + len(arrays)
		result[i] = make([]T, 1+len(arrays))
		result[i][0] = collection[i]

		// Add items from the other arrays
		for j, array := range arrays {
			result[i][j+1] = array[i]
		}
	}

	return result
}
