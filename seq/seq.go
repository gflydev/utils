// Package seq provides utility functions for sequence manipulation.
// A sequence is a chainable wrapper around a collection that allows for method chaining.
// It enables functional programming style operations on collections with a fluent interface.
package seq

import (
	"github.com/gflydev/utils/arr"
	"github.com/gflydev/utils/col"
)

// Sequence represents a chainable sequence of operations on a collection.
// It wraps a slice of comparable values and provides methods for manipulating the collection
// in a functional programming style with method chaining.
type Sequence[T comparable] struct {
	values []T
}

// New creates a new sequence from the given values.
//
// Parameters:
//   - values: The values to include in the sequence
//
// Returns:
//   - *Sequence[T]: A new sequence containing the provided values
//
// Example:
//
//	seq.New(1, 2, 3) // Creates a sequence with values [1, 2, 3]
func New[T comparable](values ...T) *Sequence[T] {
	return &Sequence[T]{values: values}
}

// FromSlice creates a new sequence from the given slice.
//
// Parameters:
//   - slice: The slice to convert into a sequence
//
// Returns:
//   - *Sequence[T]: A new sequence containing the elements from the slice
//
// Example:
//
//	seq.FromSlice([]int{1, 2, 3}) // Creates a sequence with values [1, 2, 3]
func FromSlice[T comparable](slice []T) *Sequence[T] {
	return &Sequence[T]{values: slice}
}

// Value returns the underlying collection as a slice.
//
// Returns:
//   - []T: The slice containing all elements in the sequence
//
// Example:
//
//	seq.New(1, 2, 3).Value() // Returns []int{1, 2, 3}
func (s *Sequence[T]) Value() []T {
	return s.values
}

// First returns the first element of the sequence.
//
// Returns:
//   - T: The first element in the sequence
//   - bool: True if the sequence is not empty, false otherwise
//
// Example:
//
//	value, exists := seq.New(1, 2, 3).First() // Returns 1, true
//	value, exists := seq.New().First() // Returns zero value, false
func (s *Sequence[T]) First() (T, bool) {
	return arr.First(s.values)
}

// Last returns the last element of the sequence.
//
// Returns:
//   - T: The last element in the sequence
//   - bool: True if the sequence is not empty, false otherwise
//
// Example:
//
//	value, exists := seq.New(1, 2, 3).Last() // Returns 3, true
//	value, exists := seq.New().Last() // Returns zero value, false
func (s *Sequence[T]) Last() (T, bool) {
	return arr.Last(s.values)
}

// Map applies a function to each element in the sequence and returns a new sequence with the results.
//
// Parameters:
//   - fn: The function to apply to each element
//
// Returns:
//   - *Sequence[T]: A new sequence containing the transformed elements
//
// Example:
//
//	seq.New(1, 2, 3).Map(func(n int) int { return n * 2 }) // Returns sequence with [2, 4, 6]
func (s *Sequence[T]) Map(fn func(T) T) *Sequence[T] {
	return &Sequence[T]{values: col.Map(s.values, fn)}
}

// MapTo applies a function to each element in the sequence and returns a new sequence of a different type.
//
// Parameters:
//   - fn: The function to apply to each element, converting it to a different type
//
// Returns:
//   - *Sequence[any]: A new sequence containing the transformed elements of the new type
//
// Example:
//
//	// Convert integers to strings
//	seq.New(1, 2, 3).MapTo(func(n int) any {
//	    return strconv.Itoa(n)
//	}) // Returns sequence with ["1", "2", "3"]
func (s *Sequence[T]) MapTo(fn func(T) any) *Sequence[any] {
	result := make([]any, len(s.values))
	for i, v := range s.values {
		result[i] = fn(v)
	}
	return &Sequence[any]{values: result}
}

// Filter creates a new sequence with all elements that pass the test implemented by the provided function.
//
// Parameters:
//   - predicate: Function that tests each element; return true to keep the element, false otherwise
//
// Returns:
//   - *Sequence[T]: A new sequence containing only the elements that passed the test
//
// Example:
//
//	// Get only even numbers
//	seq.New(1, 2, 3, 4).Filter(func(n int) bool {
//	    return n % 2 == 0
//	}) // Returns sequence with [2, 4]
func (s *Sequence[T]) Filter(predicate func(T) bool) *Sequence[T] {
	return &Sequence[T]{values: col.Filter(s.values, predicate)}
}

// Reject is the opposite of Filter; it creates a new sequence with all elements that do not pass the test.
//
// Parameters:
//   - predicate: Function that tests each element; return true to remove the element, false to keep it
//
// Returns:
//   - *Sequence[T]: A new sequence containing only the elements that did not pass the test
//
// Example:
//
//	// Get only odd numbers by rejecting even numbers
//	seq.New(1, 2, 3, 4).Reject(func(n int) bool {
//	    return n % 2 == 0
//	}) // Returns sequence with [1, 3]
func (s *Sequence[T]) Reject(predicate func(T) bool) *Sequence[T] {
	return &Sequence[T]{values: col.Reject(s.values, predicate)}
}

// Reduce applies a function against an accumulator and each element in the sequence
// to reduce it to a single value.
//
// Parameters:
//   - fn: Function to execute on each element, taking the accumulator and current value as arguments
//   - initial: The initial value of the accumulator
//
// Returns:
//   - any: The final accumulated value
//
// Example:
//
//	// Sum all numbers in the sequence
//	seq.New(1, 2, 3).Reduce(func(sum any, n int) any {
//	    return sum.(int) + n
//	}, 0) // Returns 6
func (s *Sequence[T]) Reduce(fn func(any, T) any, initial any) any {
	result := initial
	for _, v := range s.values {
		result = fn(result, v)
	}
	return result
}

// ForEach executes a provided function once for each sequence element.
//
// Parameters:
//   - iteratee: Function to execute on each element
//
// Returns:
//   - *Sequence[T]: The original sequence (for chaining)
//
// Example:
//
//	// Print each number in the sequence
//	seq.New(1, 2, 3).ForEach(func(n int) {
//	    fmt.Println(n)
//	}) // Prints 1, 2, 3 and returns the original sequence
func (s *Sequence[T]) ForEach(iteratee func(T)) *Sequence[T] {
	col.ForEach(s.values, iteratee)
	return s
}

// Includes determines whether the sequence includes a certain value.
//
// Parameters:
//   - value: The value to search for
//
// Returns:
//   - bool: True if the value is found, false otherwise
//
// Example:
//
//	seq.New(1, 2, 3).Includes(2) // Returns true
//	seq.New(1, 2, 3).Includes(4) // Returns false
func (s *Sequence[T]) Includes(value T) bool {
	return col.Includes(s.values, value)
}

// Find returns the first element in the sequence that satisfies the provided testing function.
//
// Parameters:
//   - predicate: Function to test each element; return true to indicate a match
//
// Returns:
//   - T: The first element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	// Find the first number greater than 2
//	value, found := seq.New(1, 2, 3, 4).Find(func(n int) bool {
//	    return n > 2
//	}) // Returns 3, true
func (s *Sequence[T]) Find(predicate func(T) bool) (T, bool) {
	return col.Find(s.values, predicate)
}

// FindLast returns the last element in the sequence that satisfies the provided testing function.
//
// Parameters:
//   - predicate: Function to test each element; return true to indicate a match
//
// Returns:
//   - T: The last element that satisfies the predicate
//   - bool: True if an element was found, false otherwise
//
// Example:
//
//	// Find the last number greater than 2
//	value, found := seq.New(1, 2, 3, 4).FindLast(func(n int) bool {
//	    return n > 2
//	}) // Returns 4, true
func (s *Sequence[T]) FindLast(predicate func(T) bool) (T, bool) {
	return col.FindLast(s.values, predicate)
}

// Every tests whether all elements in the sequence pass the test implemented by the provided function.
//
// Parameters:
//   - predicate: Function to test each element; should return a boolean
//
// Returns:
//   - bool: True if all elements pass the test, false otherwise
//
// Example:
//
//	// Check if all numbers are even
//	seq.New(2, 4, 6).Every(func(n int) bool {
//	    return n % 2 == 0
//	}) // Returns true
//
//	seq.New(1, 2, 3).Every(func(n int) bool {
//	    return n % 2 == 0
//	}) // Returns false
func (s *Sequence[T]) Every(predicate func(T) bool) bool {
	return col.Every(s.values, predicate)
}

// Some tests whether at least one element in the sequence passes the test implemented by the provided function.
//
// Parameters:
//   - predicate: Function to test each element; should return a boolean
//
// Returns:
//   - bool: True if at least one element passes the test, false otherwise
//
// Example:
//
//	// Check if any number is greater than 3
//	seq.New(1, 2, 3, 4).Some(func(n int) bool {
//	    return n > 3
//	}) // Returns true
//
//	seq.New(1, 2, 3).Some(func(n int) bool {
//	    return n > 3
//	}) // Returns false
func (s *Sequence[T]) Some(predicate func(T) bool) bool {
	return col.Some(s.values, predicate)
}

// Size returns the number of elements in the sequence.
//
// Returns:
//   - int: The number of elements in the sequence
//
// Example:
//
//	seq.New(1, 2, 3).Size() // Returns 3
//	seq.New().Size() // Returns 0
func (s *Sequence[T]) Size() int {
	return len(s.values)
}

// IsEmpty checks if the sequence contains no elements.
//
// Returns:
//   - bool: True if the sequence is empty, false otherwise
//
// Example:
//
//	seq.New().IsEmpty() // Returns true
//	seq.New(1, 2, 3).IsEmpty() // Returns false
func (s *Sequence[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Reverse creates a new sequence with the elements in reverse order.
//
// Returns:
//   - *Sequence[T]: A new sequence with elements in reverse order
//
// Example:
//
//	seq.New(1, 2, 3).Reverse() // Returns sequence with [3, 2, 1]
func (s *Sequence[T]) Reverse() *Sequence[T] {
	return &Sequence[T]{values: arr.Reverse(s.values)}
}

// Uniq creates a new sequence with all duplicate elements removed.
//
// Returns:
//   - *Sequence[T]: A new sequence with only unique elements
//
// Example:
//
//	seq.New(1, 2, 1, 3).Uniq() // Returns sequence with [1, 2, 3]
func (s *Sequence[T]) Uniq() *Sequence[T] {
	return &Sequence[T]{values: arr.Uniq(s.values)}
}

// Chunk splits the sequence into groups of the specified size.
//
// Parameters:
//   - size: The size of each chunk
//
// Returns:
//   - [][]T: A slice of slices, where each inner slice is a chunk of the original sequence
//
// Example:
//
//	seq.New(1, 2, 3, 4).Chunk(2) // Returns [][]int{{1, 2}, {3, 4}}
//	seq.New(1, 2, 3, 4, 5).Chunk(2) // Returns [][]int{{1, 2}, {3, 4}, {5}}
func (s *Sequence[T]) Chunk(size int) [][]T {
	return arr.Chunk(s.values, size)
}

// Flatten flattens the sequence a single level deep.
//
// Note: This is a simplified version that assumes T is already a slice.
// In a real implementation, reflection would be needed to handle different types.
//
// Returns:
//   - *Sequence[T]: A new sequence with elements flattened one level
//
// Example:
//
//	// Note: This example is conceptual as the current implementation is simplified
//	seq.New([]int{1, 2}, []int{3, 4}).Flatten() // Would return sequence with [1, 2, 3, 4]
func (s *Sequence[T]) Flatten() *Sequence[T] {
	// This is a simplified version that assumes T is already a slice
	// In a real implementation, we would need to use reflection to handle different types
	return s
}

// Concat creates a new sequence by concatenating the current sequence with other sequences.
//
// Parameters:
//   - others: One or more sequences to concatenate with the current sequence
//
// Returns:
//   - *Sequence[T]: A new sequence containing all elements from the current sequence followed by elements from the other sequences
//
// Example:
//
//	seq1 := seq.New(1, 2)
//	seq2 := seq.New(3, 4)
//	seq1.Concat(seq2) // Returns sequence with [1, 2, 3, 4]
//
//	seq.New(1, 2).Concat(seq.New(3), seq.New(4, 5)) // Returns sequence with [1, 2, 3, 4, 5]
func (s *Sequence[T]) Concat(others ...*Sequence[T]) *Sequence[T] {
	result := make([]T, len(s.values))
	copy(result, s.values)

	for _, other := range others {
		result = append(result, other.values...)
	}

	return &Sequence[T]{values: result}
}

// Take creates a new sequence with n elements taken from the beginning of the current sequence.
//
// Parameters:
//   - n: The number of elements to take
//
// Returns:
//   - *Sequence[T]: A new sequence with the first n elements
//
// Example:
//
//	seq.New(1, 2, 3, 4).Take(2) // Returns sequence with [1, 2]
//	seq.New(1, 2).Take(3) // Returns sequence with [1, 2] (takes all available elements)
func (s *Sequence[T]) Take(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.Take(s.values, n)}
}

// TakeRight creates a new sequence with n elements taken from the end of the current sequence.
//
// Parameters:
//   - n: The number of elements to take from the end
//
// Returns:
//   - *Sequence[T]: A new sequence with the last n elements
//
// Example:
//
//	seq.New(1, 2, 3, 4).TakeRight(2) // Returns sequence with [3, 4]
//	seq.New(1, 2).TakeRight(3) // Returns sequence with [1, 2] (takes all available elements)
func (s *Sequence[T]) TakeRight(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.TakeRight(s.values, n)}
}

// Drop creates a new sequence with n elements removed from the beginning of the current sequence.
//
// Parameters:
//   - n: The number of elements to exclude from the beginning
//
// Returns:
//   - *Sequence[T]: A new sequence with the first n elements removed
//
// Example:
//
//	seq.New(1, 2, 3, 4).Drop(2) // Returns sequence with [3, 4]
//	seq.New(1, 2).Drop(3) // Returns empty sequence (drops all elements)
func (s *Sequence[T]) Drop(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.Drop(s.values, n)}
}

// DropRight creates a new sequence with n elements removed from the end of the current sequence.
//
// Parameters:
//   - n: The number of elements to exclude from the end
//
// Returns:
//   - *Sequence[T]: A new sequence with the last n elements removed
//
// Example:
//
//	seq.New(1, 2, 3, 4).DropRight(2) // Returns sequence with [1, 2]
//	seq.New(1, 2).DropRight(3) // Returns empty sequence (drops all elements)
func (s *Sequence[T]) DropRight(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.DropRight(s.values, n)}
}

// Shuffle creates a new sequence with elements randomly reordered.
//
// Returns:
//   - *Sequence[T]: A new sequence with the same elements in random order
//
// Example:
//
//	// Result will have the same elements in random order
//	seq.New(1, 2, 3, 4).Shuffle() // Might return sequence with [3, 1, 4, 2]
func (s *Sequence[T]) Shuffle() *Sequence[T] {
	return &Sequence[T]{values: arr.Shuffle(s.values)}
}

// Sample returns a randomly selected element from the sequence.
//
// Returns:
//   - T: A randomly selected element from the sequence
//   - bool: True if the sequence is not empty and an element was selected, false otherwise
//
// Example:
//
//	// Get a random element
//	value, exists := seq.New(1, 2, 3, 4).Sample() // Might return 2, true
//	value, exists := seq.New().Sample() // Returns zero value, false
func (s *Sequence[T]) Sample() (T, bool) {
	return col.Sample(s.values)
}

// SampleSize returns a new sequence with n randomly selected elements from the current sequence.
//
// Parameters:
//   - n: The number of elements to randomly select
//
// Returns:
//   - *Sequence[T]: A new sequence containing n randomly selected elements
//
// Example:
//
//	// Get 2 random elements
//	seq.New(1, 2, 3, 4).SampleSize(2) // Might return sequence with [2, 4]
//
//	// If n is greater than the sequence size, returns all elements in random order
//	seq.New(1, 2).SampleSize(3) // Returns all elements in random order
func (s *Sequence[T]) SampleSize(n int) *Sequence[T] {
	return &Sequence[T]{values: col.SampleSize(s.values, n)}
}

// Partition divides the sequence into two groups: elements that satisfy the predicate and elements that don't.
//
// Parameters:
//   - predicate: Function to test each element; return true to include in first group, false for second group
//
// Returns:
//   - [][]T: A slice containing two slices: the first with elements that passed the test, the second with elements that failed
//
// Example:
//
//	// Partition into even and odd numbers
//	seq.New(1, 2, 3, 4).Partition(func(n int) bool {
//	    return n % 2 == 0
//	}) // Returns [][]int{{2, 4}, {1, 3}}
func (s *Sequence[T]) Partition(predicate func(T) bool) [][]T {
	return col.Partition(s.values, predicate)
}

// GroupBy creates a map that groups sequence elements by keys generated from the iteratee function.
//
// Parameters:
//   - iteratee: Function that returns a string key for each element
//
// Returns:
//   - map[string][]T: A map where keys are the strings returned by iteratee and values are slices of elements that produced each key
//
// Example:
//
//	// Group numbers by even/odd
//	seq.New(1, 2, 3, 4).GroupBy(func(n int) string {
//	    if n % 2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	}) // Returns map[string][]int{"odd": {1, 3}, "even": {2, 4}}
func (s *Sequence[T]) GroupBy(iteratee func(T) string) map[string][]T {
	return col.GroupBy(s.values, iteratee)
}

// CountBy creates a map that counts elements by keys generated from the iteratee function.
//
// Parameters:
//   - iteratee: Function that returns a string key for each element
//
// Returns:
//   - map[string]int: A map where keys are the strings returned by iteratee and values are counts of elements that produced each key
//
// Example:
//
//	// Count numbers by even/odd
//	seq.New(1, 2, 3, 4).CountBy(func(n int) string {
//	    if n % 2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	}) // Returns map[string]int{"odd": 2, "even": 2}
func (s *Sequence[T]) CountBy(iteratee func(T) string) map[string]int {
	return col.CountBy(s.values, iteratee)
}

// KeyBy creates a map with keys generated by applying the iteratee function to each element.
//
// Parameters:
//   - iteratee: Function that returns an integer key for each element
//
// Returns:
//   - map[int]T: A map where keys are integers returned by iteratee and values are the elements that produced each key
//
// Example:
//
//	// Create a map of users keyed by their ID
//	type User struct {
//	    ID   int
//	    Name string
//	}
//	users := seq.New(User{1, "Alice"}, User{2, "Bob"})
//	users.KeyBy(func(u User) int {
//	    return u.ID
//	}) // Returns map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
func (s *Sequence[T]) KeyBy(iteratee func(T) int) map[int]T {
	return col.KeyBy(s.values, iteratee)
}

// SortBy creates a new sequence sorted by the values returned by the iteratee function.
//
// Parameters:
//   - iteratee: Function that returns a comparable value used for sorting
//
// Returns:
//   - *Sequence[T]: A new sequence with elements sorted by the iteratee function results
//
// Example:
//
//	// Sort numbers in ascending order
//	seq.New(1, 3, 2).SortBy(func(n int) int {
//	    return n
//	}) // Returns sequence with [1, 2, 3]
//
//	// Sort users by ID
//	type User struct {
//	    ID   int
//	    Name string
//	}
//	seq.New(User{3, "Charlie"}, User{1, "Alice"}, User{2, "Bob"}).SortBy(func(u User) int {
//	    return u.ID
//	}) // Returns sequence with users sorted by ID
func (s *Sequence[T]) SortBy(iteratee func(T) int) *Sequence[T] {
	return &Sequence[T]{values: col.SortBy(s.values, iteratee)}
}

// OrderBy sorts the sequence by the results of running each element through iteratee.
//
// Parameters:
//   - iteratee: Function that returns an integer value used for sorting comparison
//   - ascending: If true sort in ascending order, if false sort in descending order
//
// Returns:
//   - *Sequence[T]: A new sequence with elements sorted based on iteratee results
//
// Example:
//
//	seq.New(1, 3, 2).OrderBy(func(n int) int {
//		return n
//	}, true) // Returns sequence with [1, 2, 3]
//
//	// Sort users by age in descending order
//	type User struct {
//		Name string
//		Age  int
//	}
//	seq.New(
//		User{"Alice", 25},
//		User{"Bob", 30},
//		User{"Charlie", 20},
//	).OrderBy(func(u User) int {
//		return u.Age
//	}, false) // Returns sequence with [Bob(30), Alice(25), Charlie(20)]
func (s *Sequence[T]) OrderBy(iteratee func(T) int, ascending bool) *Sequence[T] {
	return &Sequence[T]{values: col.OrderBy(s.values, iteratee, ascending)}
}

// Join joins all elements of the sequence into a string.
//
// Parameters:
//   - separator: The string to insert between joined elements
//
// Returns:
//   - string: A string containing all elements joined together with the separator between them
//
// Example:
//
//	seq.New(1, 2, 3).Join(",") // Returns "1,2,3"
//	seq.New("a", "b", "c").Join("-") // Returns "a-b-c"
//	seq.New(true, false, true).Join(" and ") // Returns "true and false and true"
//	seq.New[int]().Join(",") // Returns "" (empty string for empty sequence)
func (s *Sequence[T]) Join(separator string) string {
	return arr.Join(s.values, separator)
}
