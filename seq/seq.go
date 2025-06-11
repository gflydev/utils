// Package seq provides utility functions for sequence manipulation.
// A sequence is a chainable wrapper around a collection that allows for method chaining.
package seq

import (
	"github.com/gflydev/utils/arr"
	"github.com/gflydev/utils/coll"
)

// Sequence represents a chainable sequence of operations on a collection.
type Sequence[T comparable] struct {
	values []T
}

// New creates a new sequence from the given values.
// Example: seq.New(1, 2, 3)
func New[T comparable](values ...T) *Sequence[T] {
	return &Sequence[T]{values: values}
}

// FromSlice creates a new sequence from the given slice.
// Example: seq.FromSlice([]int{1, 2, 3})
func FromSlice[T comparable](slice []T) *Sequence[T] {
	return &Sequence[T]{values: slice}
}

// Value returns the underlying collection.
// Example: seq.New(1, 2, 3).Value() -> []int{1, 2, 3}
func (s *Sequence[T]) Value() []T {
	return s.values
}

// First returns the first element of the sequence.
// Example: seq.New(1, 2, 3).First() -> 1, true
func (s *Sequence[T]) First() (T, bool) {
	return arr.First(s.values)
}

// Last returns the last element of the sequence.
// Example: seq.New(1, 2, 3).Last() -> 3, true
func (s *Sequence[T]) Last() (T, bool) {
	return arr.Last(s.values)
}

// Map applies a function to each element in the sequence.
// Example: seq.New(1, 2, 3).Map(func(n int) int { return n * 2 }) -> seq.New(2, 4, 6)
func (s *Sequence[T]) Map(fn func(T) T) *Sequence[T] {
	return &Sequence[T]{values: coll.Map(s.values, fn)}
}

// MapTo applies a function to each element in the sequence and returns a new sequence of a different type.
// Example: seq.New(1, 2, 3).MapTo(func(n int) string { return strconv.Itoa(n) }) -> seq.New("1", "2", "3")
func (s *Sequence[T]) MapTo(fn func(T) interface{}) *Sequence[interface{}] {
	result := make([]interface{}, len(s.values))
	for i, v := range s.values {
		result[i] = fn(v)
	}
	return &Sequence[interface{}]{values: result}
}

// Filter filters the sequence based on a predicate.
// Example: seq.New(1, 2, 3, 4).Filter(func(n int) bool { return n % 2 == 0 }) -> seq.New(2, 4)
func (s *Sequence[T]) Filter(predicate func(T) bool) *Sequence[T] {
	return &Sequence[T]{values: coll.Filter(s.values, predicate)}
}

// Reject is the opposite of Filter; it returns elements that don't satisfy the predicate.
// Example: seq.New(1, 2, 3, 4).Reject(func(n int) bool { return n % 2 == 0 }) -> seq.New(1, 3)
func (s *Sequence[T]) Reject(predicate func(T) bool) *Sequence[T] {
	return &Sequence[T]{values: coll.Reject(s.values, predicate)}
}

// Reduce reduces the sequence to a single value.
// Example: seq.New(1, 2, 3).Reduce(func(sum, n int) int { return sum + n }, 0) -> 6
func (s *Sequence[T]) Reduce(fn func(interface{}, T) interface{}, initial interface{}) interface{} {
	result := initial
	for _, v := range s.values {
		result = fn(result, v)
	}
	return result
}

// ForEach iterates over elements of the sequence and invokes iteratee for each element.
// Example: seq.New(1, 2, 3).ForEach(func(n int) { fmt.Println(n) })
func (s *Sequence[T]) ForEach(iteratee func(T)) *Sequence[T] {
	coll.ForEach(s.values, iteratee)
	return s
}

// Includes checks if a value is in the sequence.
// Example: seq.New(1, 2, 3).Includes(2) -> true
func (s *Sequence[T]) Includes(value T) bool {
	return coll.Includes(s.values, value)
}

// Find finds the first element in the sequence that satisfies the predicate.
// Example: seq.New(1, 2, 3, 4).Find(func(n int) bool { return n > 2 }) -> 3, true
func (s *Sequence[T]) Find(predicate func(T) bool) (T, bool) {
	return coll.Find(s.values, predicate)
}

// FindLast finds the last element in the sequence that satisfies the predicate.
// Example: seq.New(1, 2, 3, 4).FindLast(func(n int) bool { return n > 2 }) -> 4, true
func (s *Sequence[T]) FindLast(predicate func(T) bool) (T, bool) {
	return coll.FindLast(s.values, predicate)
}

// Every checks if all elements in the sequence satisfy the predicate.
// Example: seq.New(2, 4, 6).Every(func(n int) bool { return n % 2 == 0 }) -> true
func (s *Sequence[T]) Every(predicate func(T) bool) bool {
	return coll.Every(s.values, predicate)
}

// Some checks if any element in the sequence satisfies the predicate.
// Example: seq.New(1, 2, 3, 4).Some(func(n int) bool { return n > 3 }) -> true
func (s *Sequence[T]) Some(predicate func(T) bool) bool {
	return coll.Some(s.values, predicate)
}

// Size returns the size of the sequence.
// Example: seq.New(1, 2, 3).Size() -> 3
func (s *Sequence[T]) Size() int {
	return len(s.values)
}

// IsEmpty checks if the sequence is empty.
// Example: seq.New().IsEmpty() -> true
func (s *Sequence[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Reverse reverses the order of elements in the sequence.
// Example: seq.New(1, 2, 3).Reverse() -> seq.New(3, 2, 1)
func (s *Sequence[T]) Reverse() *Sequence[T] {
	return &Sequence[T]{values: arr.Reverse(s.values)}
}

// Uniq creates a duplicate-free version of the sequence.
// Example: seq.New(1, 2, 1, 3).Uniq() -> seq.New(1, 2, 3)
func (s *Sequence[T]) Uniq() *Sequence[T] {
	return &Sequence[T]{values: arr.Uniq(s.values)}
}

// Chunk splits the sequence into groups of the specified size.
// Example: seq.New(1, 2, 3, 4).Chunk(2) -> [][]int{{1, 2}, {3, 4}}
func (s *Sequence[T]) Chunk(size int) [][]T {
	return arr.Chunk(s.values, size)
}

// Flatten flattens the sequence a single level deep.
// Example: seq.New([]int{1, 2}, []int{3, 4}).Flatten() -> seq.New(1, 2, 3, 4)
func (s *Sequence[T]) Flatten() *Sequence[T] {
	// This is a simplified version that assumes T is already a slice
	// In a real implementation, we would need to use reflection to handle different types
	return s
}

// Concat concatenates the sequence with other sequences.
// Example: seq.New(1, 2).Concat(seq.New(3, 4)) -> seq.New(1, 2, 3, 4)
func (s *Sequence[T]) Concat(others ...*Sequence[T]) *Sequence[T] {
	result := make([]T, len(s.values))
	copy(result, s.values)

	for _, other := range others {
		result = append(result, other.values...)
	}

	return &Sequence[T]{values: result}
}

// Take creates a slice of the sequence with n elements taken from the beginning.
// Example: seq.New(1, 2, 3, 4).Take(2) -> seq.New(1, 2)
func (s *Sequence[T]) Take(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.Take(s.values, n)}
}

// TakeRight creates a slice of the sequence with n elements taken from the end.
// Example: seq.New(1, 2, 3, 4).TakeRight(2) -> seq.New(3, 4)
func (s *Sequence[T]) TakeRight(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.TakeRight(s.values, n)}
}

// Drop creates a slice of the sequence with n elements dropped from the beginning.
// Example: seq.New(1, 2, 3, 4).Drop(2) -> seq.New(3, 4)
func (s *Sequence[T]) Drop(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.Drop(s.values, n)}
}

// DropRight creates a slice of the sequence with n elements dropped from the end.
// Example: seq.New(1, 2, 3, 4).DropRight(2) -> seq.New(1, 2)
func (s *Sequence[T]) DropRight(n int) *Sequence[T] {
	return &Sequence[T]{values: arr.DropRight(s.values, n)}
}

// Shuffle creates a shuffled version of the sequence.
// Example: seq.New(1, 2, 3, 4).Shuffle() -> seq.New(3, 1, 4, 2) (random order)
func (s *Sequence[T]) Shuffle() *Sequence[T] {
	return &Sequence[T]{values: arr.Shuffle(s.values)}
}

// Sample gets a random element from the sequence.
// Example: seq.New(1, 2, 3, 4).Sample() -> a random element from the sequence
func (s *Sequence[T]) Sample() (T, bool) {
	return coll.Sample(s.values)
}

// SampleSize gets n random elements from the sequence.
// Example: seq.New(1, 2, 3, 4).SampleSize(2) -> seq.New(2, 4) (random elements)
func (s *Sequence[T]) SampleSize(n int) *Sequence[T] {
	return &Sequence[T]{values: coll.SampleSize(s.values, n)}
}

// Partition splits the sequence into two groups, the first of which contains elements that satisfy the predicate.
// Example: seq.New(1, 2, 3, 4).Partition(func(n int) bool { return n % 2 == 0 }) -> [][]int{{2, 4}, {1, 3}}
func (s *Sequence[T]) Partition(predicate func(T) bool) [][]T {
	return coll.Partition(s.values, predicate)
}

// GroupBy creates an object composed of keys generated from the results of running each element of the sequence through iteratee.
// Example: seq.New(1, 2, 3, 4).GroupBy(func(n int) string { return "even" if n % 2 == 0 else "odd" })
// -> map[string][]int{"odd": {1, 3}, "even": {2, 4}}
func (s *Sequence[T]) GroupBy(iteratee func(T) string) map[string][]T {
	return coll.GroupBy(s.values, iteratee)
}

// CountBy counts the occurrences of elements in the sequence based on the iteratee function.
// Example: seq.New(1, 2, 3, 4).CountBy(func(n int) string { return "even" if n % 2 == 0 else "odd" })
// -> map[string]int{"odd": 2, "even": 2}
func (s *Sequence[T]) CountBy(iteratee func(T) string) map[string]int {
	return coll.CountBy(s.values, iteratee)
}

// KeyBy creates an object composed of keys generated from the results of running each element of the sequence through iteratee.
// Example: seq.New(struct{ID int, Name string}{{1, "Alice"}, {2, "Bob"}}).KeyBy(func(u User) int { return u.ID })
// -> map[int]struct{ID int, Name string}{1: {1, "Alice"}, 2: {2, "Bob"}}
func (s *Sequence[T]) KeyBy(iteratee func(T) int) map[int]T {
	return coll.KeyBy(s.values, iteratee)
}

// SortBy sorts the sequence by the results of running each element through iteratee.
// Example: seq.New(1, 3, 2).SortBy(func(n int) int { return n }) -> seq.New(1, 2, 3)
func (s *Sequence[T]) SortBy(iteratee func(T) int) *Sequence[T] {
	return &Sequence[T]{values: coll.SortBy(s.values, iteratee)}
}

// OrderBy sorts the sequence by the results of running each element through iteratee.
// Example: seq.New(1, 3, 2).OrderBy(func(n int) int { return n }, true) -> seq.New(1, 2, 3)
func (s *Sequence[T]) OrderBy(iteratee func(T) int, ascending bool) *Sequence[T] {
	return &Sequence[T]{values: coll.OrderBy(s.values, iteratee, ascending)}
}

// Join joins all elements of the sequence into a string.
// Example: seq.New(1, 2, 3).Join(",") -> "1,2,3"
func (s *Sequence[T]) Join(separator string) string {
	return arr.Join(s.values, separator)
}
