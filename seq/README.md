# seq - Sequence Utility Functions for Go

The `seq` package provides utility functions for sequence manipulation in Go. A sequence is a chainable wrapper around a collection that allows for method chaining, enabling functional programming style operations on collections with a fluent interface.

## Installation

```bash
go get github.com/gflydev/utils/seq
```

## Usage

```go
import "github.com/gflydev/utils/seq"
```

## Functions

### Creating Sequences

#### New

Creates a new sequence from the given values.

```go
seq := seq.New(1, 2, 3)
// seq.Value() returns []int{1, 2, 3}
```

#### FromSlice

Creates a new sequence from the given slice.

```go
slice := []int{1, 2, 3}
seq := seq.FromSlice(slice)
// seq.Value() returns []int{1, 2, 3}
```

### Basic Operations

#### Value

Returns the underlying collection as a slice.

```go
seq := seq.New(1, 2, 3)
values := seq.Value()
// values is []int{1, 2, 3}
```

#### First

Returns the first element of the sequence.

```go
seq := seq.New(1, 2, 3)
first, exists := seq.First()
// first is 1, exists is true

emptySeq := seq.New[int]()
first, exists := emptySeq.First()
// first is 0, exists is false
```

#### Last

Returns the last element of the sequence.

```go
seq := seq.New(1, 2, 3)
last, exists := seq.Last()
// last is 3, exists is true

emptySeq := seq.New[int]()
last, exists := emptySeq.Last()
// last is 0, exists is false
```

### Transformation

#### Map

Applies a function to each element in the sequence and returns a new sequence with the results.

Parameters:
- `fn`: The function to apply to each element

Returns:
- `*Sequence[T]`: A new sequence containing the transformed elements

```go
seq := seq.New(1, 2, 3)
result := seq.Map(func(n int) int { return n * 2 })
// result.Value() returns []int{2, 4, 6}
```

#### MapTo

Applies a function to each element in the sequence and returns a new sequence of a different type.

Parameters:
- `fn`: The function to apply to each element, converting it to a different type

Returns:
- `*Sequence[any]`: A new sequence containing the transformed elements of the new type

```go
// Convert integers to strings
seq := seq.New(1, 2, 3)
result := seq.MapTo(func(n int) any {
    return strconv.Itoa(n)
}) 
// result.Value() returns []any{"1", "2", "3"}
```

#### Filter

Creates a new sequence with all elements that pass the test implemented by the provided function.

Parameters:
- `predicate`: Function that tests each element; return true to keep the element, false otherwise

Returns:
- `*Sequence[T]`: A new sequence containing only the elements that passed the test

```go
// Get only even numbers
seq := seq.New(1, 2, 3, 4)
result := seq.Filter(func(n int) bool { return n%2 == 0 })
// result.Value() returns []int{2, 4}
```

#### Reject

Creates a new sequence with all elements that do not pass the test implemented by the provided function.

Parameters:
- `predicate`: Function that tests each element; return true to remove the element, false to keep it

Returns:
- `*Sequence[T]`: A new sequence containing only the elements that did not pass the test

```go
// Get only odd numbers by rejecting even numbers
seq := seq.New(1, 2, 3, 4)
result := seq.Reject(func(n int) bool { return n%2 == 0 })
// result.Value() returns []int{1, 3}
```

#### Reduce

Applies a function against an accumulator and each element in the sequence to reduce it to a single value.

Parameters:
- `fn`: Function to execute on each element, taking the accumulator and current value as arguments
- `initial`: The initial value of the accumulator

Returns:
- `any`: The final accumulated value

```go
// Sum all numbers in the sequence
seq := seq.New(1, 2, 3, 4)
result := seq.Reduce(func(acc any, n int) any {
    return acc.(int) + n
}, 0)
// result is 10
```

#### ForEach

Executes a provided function once for each sequence element.

Parameters:
- `iteratee`: Function to execute on each element

Returns:
- `*Sequence[T]`: The original sequence (for chaining)

```go
seq := seq.New(1, 2, 3)
sum := 0
result := seq.ForEach(func(n int) {
    sum += n
})
// sum is 6
// result is the original sequence (for chaining)
```

### Searching

#### Includes

Determines whether the sequence includes a certain value.

Parameters:
- `value`: The value to search for

Returns:
- `bool`: True if the value is found, false otherwise

```go
seq := seq.New(1, 2, 3)
result := seq.Includes(2)
// result is true

result = seq.Includes(4)
// result is false

emptySeq := seq.New[int]()
result = emptySeq.Includes(1)
// result is false
```

#### Find

Returns the first element in the sequence that satisfies the provided testing function.

Parameters:
- `predicate`: Function to test each element; return true to indicate a match

Returns:
- `T`: The first element that satisfies the predicate
- `bool`: True if an element was found, false otherwise

```go
seq := seq.New(1, 2, 3, 4)
value, found := seq.Find(func(n int) bool { return n > 2 })
// value is 3, found is true

seq = seq.New(1, 2)
value, found = seq.Find(func(n int) bool { return n > 2 })
// value is 0, found is false

emptySeq := seq.New[int]()
value, found = emptySeq.Find(func(n int) bool { return n > 2 })
// value is 0, found is false
```

#### FindLast

Returns the last element in the sequence that satisfies the provided testing function.

Parameters:
- `predicate`: Function to test each element; return true to indicate a match

Returns:
- `T`: The last element that satisfies the predicate
- `bool`: True if an element was found, false otherwise

```go
seq := seq.New(1, 2, 3, 4, 3)
value, found := seq.FindLast(func(n int) bool { return n > 2 })
// value is 3, found is true

seq = seq.New(1, 2)
value, found = seq.FindLast(func(n int) bool { return n > 2 })
// value is 0, found is false

emptySeq := seq.New[int]()
value, found = emptySeq.FindLast(func(n int) bool { return n > 2 })
// value is 0, found is false
```

### Testing

#### Every

Tests whether all elements in the sequence pass the test implemented by the provided function.

Parameters:
- `predicate`: Function to test each element; should return a boolean

Returns:
- `bool`: True if all elements pass the test, false otherwise

```go
seq := seq.New(2, 4, 6)
result := seq.Every(func(n int) bool { return n%2 == 0 })
// result is true

seq = seq.New(2, 3, 6)
result = seq.Every(func(n int) bool { return n%2 == 0 })
// result is false

emptySeq := seq.New[int]()
result = emptySeq.Every(func(n int) bool { return n%2 == 0 })
// result is true (vacuously true for empty sequences)
```

#### Some

Tests whether at least one element in the sequence passes the test implemented by the provided function.

Parameters:
- `predicate`: Function to test each element; should return a boolean

Returns:
- `bool`: True if at least one element passes the test, false otherwise

```go
seq := seq.New(1, 2, 3)
result := seq.Some(func(n int) bool { return n%2 == 0 })
// result is true

seq = seq.New(1, 3, 5)
result = seq.Some(func(n int) bool { return n%2 == 0 })
// result is false

emptySeq := seq.New[int]()
result = emptySeq.Some(func(n int) bool { return n%2 == 0 })
// result is false
```

### Information

#### Size

Returns the number of elements in the sequence.

Returns:
- `int`: The number of elements in the sequence

```go
seq := seq.New(1, 2, 3, 4)
size := seq.Size()
// size is 4

seq = seq.New(1)
size = seq.Size()
// size is 1

emptySeq := seq.New[int]()
size = emptySeq.Size()
// size is 0
```

#### IsEmpty

Checks if the sequence contains no elements.

Returns:
- `bool`: True if the sequence is empty, false otherwise

```go
seq := seq.New(1, 2, 3)
isEmpty := seq.IsEmpty()
// isEmpty is false

emptySeq := seq.New[int]()
isEmpty = emptySeq.IsEmpty()
// isEmpty is true
```

### Manipulation

#### Reverse

Creates a new sequence with the elements in reverse order.

Returns:
- `*Sequence[T]`: A new sequence with elements in reverse order

```go
seq := seq.New(1, 2, 3)
result := seq.Reverse()
// result.Value() returns []int{3, 2, 1}

seq = seq.New(1)
result = seq.Reverse()
// result.Value() returns []int{1}
```

#### Uniq

Creates a new sequence with all duplicate elements removed.

Returns:
- `*Sequence[T]`: A new sequence with only unique elements

```go
seq := seq.New(1, 2, 1, 3, 2)
result := seq.Uniq()
// result.Value() returns []int{1, 2, 3}

seq = seq.New(1, 2, 3)
result = seq.Uniq()
// result.Value() returns []int{1, 2, 3}
```

#### Chunk

Splits the sequence into groups of the specified size.

Parameters:
- `size`: The size of each chunk

Returns:
- `[][]T`: A slice of slices, where each inner slice is a chunk of the original sequence

```go
seq := seq.New(1, 2, 3, 4)
chunks := seq.Chunk(2)
// chunks is [][]int{{1, 2}, {3, 4}}

seq = seq.New(1, 2, 3, 4, 5)
chunks = seq.Chunk(2)
// chunks is [][]int{{1, 2}, {3, 4}, {5}}

seq = seq.New(1, 2, 3)
chunks = seq.Chunk(5)
// chunks is [][]int{{1, 2, 3}}

emptySeq := seq.New[int]()
chunks = emptySeq.Chunk(2)
// chunks is [][]int{}
```

#### Flatten

Flattens the sequence a single level deep.

Note: This is a simplified version that assumes T is already a slice. In a real implementation, reflection would be needed to handle different types.

Returns:
- `*Sequence[T]`: A new sequence with elements flattened one level

```go
// Note: This example is conceptual as the current implementation is simplified
seq := seq.New([]int{1, 2}, []int{3, 4})
result := seq.Flatten()
// Would return sequence with [1, 2, 3, 4]
```

#### Concat

Concatenates the current sequence with one or more other sequences.

Parameters:
- `others`: One or more sequences to concatenate with the current sequence

Returns:
- `*Sequence[T]`: A new sequence containing all elements from the current sequence followed by elements from the other sequences

```go
seq1 := seq.New(1, 2)
seq2 := seq.New(3, 4)
result := seq1.Concat(seq2)
// result.Value() returns []int{1, 2, 3, 4}

result = seq.New(1, 2).Concat(seq.New(3), seq.New(4, 5))
// result.Value() returns []int{1, 2, 3, 4, 5}
```

#### Take

Creates a new sequence with n elements taken from the beginning of the current sequence.

```go
seq := seq.New(1, 2, 3, 4)
result := seq.Take(2)
// result.Value() returns []int{1, 2}

seq = seq.New(1, 2, 3)
result = seq.Take(5)
// result.Value() returns []int{1, 2, 3}

seq = seq.New(1, 2, 3)
result = seq.Take(0)
// result.Value() returns []int{}
```

#### TakeRight

Creates a new sequence with n elements taken from the end of the current sequence.

```go
seq := seq.New(1, 2, 3, 4)
result := seq.TakeRight(2)
// result.Value() returns []int{3, 4}

seq = seq.New(1, 2, 3)
result = seq.TakeRight(5)
// result.Value() returns []int{1, 2, 3}

seq = seq.New(1, 2, 3)
result = seq.TakeRight(0)
// result.Value() returns []int{}
```

#### Drop

Creates a new sequence with n elements removed from the beginning of the current sequence.

```go
seq := seq.New(1, 2, 3, 4)
result := seq.Drop(2)
// result.Value() returns []int{3, 4}

seq = seq.New(1, 2, 3)
result = seq.Drop(5)
// result.Value() returns []int{}

seq = seq.New(1, 2, 3)
result = seq.Drop(0)
// result.Value() returns []int{1, 2, 3}

emptySeq := seq.New[int]()
result = emptySeq.Drop(2)
// result.Value() returns []int{}
```

#### DropRight

Creates a new sequence with n elements removed from the end of the current sequence.

```go
seq := seq.New(1, 2, 3, 4)
result := seq.DropRight(2)
// result.Value() returns []int{1, 2}

seq = seq.New(1, 2, 3)
result = seq.DropRight(5)
// result.Value() returns []int{}

seq = seq.New(1, 2, 3)
result = seq.DropRight(0)
// result.Value() returns []int{1, 2, 3}

emptySeq := seq.New[int]()
result = emptySeq.DropRight(2)
// result.Value() returns []int{}
```

#### Shuffle

Creates a new sequence with elements randomly reordered.

Returns:
- `*Sequence[T]`: A new sequence with the same elements in random order

```go
seq := seq.New(1, 2, 3, 4)
result := seq.Shuffle()
// result.Value() might return []int{3, 1, 4, 2}
```

#### Sample

Returns a randomly selected element from the sequence.

Returns:
- `T`: A randomly selected element from the sequence
- `bool`: True if the sequence is not empty and an element was selected, false otherwise

```go
seq := seq.New(1, 2, 3, 4)
value, exists := seq.Sample()
// value might be 2, exists is true

emptySeq := seq.New[int]()
value, exists = emptySeq.Sample()
// value is 0, exists is false
```

#### SampleSize

Returns a new sequence with n randomly selected elements from the current sequence.

Parameters:
- `n`: The number of elements to randomly select

Returns:
- `*Sequence[T]`: A new sequence containing n randomly selected elements

```go
seq := seq.New(1, 2, 3, 4)
result := seq.SampleSize(2)
// result.Value() might return []int{2, 4}

// If n is greater than the sequence size, returns all elements in random order
seq = seq.New(1, 2)
result = seq.SampleSize(3)
// result.Value() returns all elements in random order
```

#### Partition

Divides the sequence into two groups: elements that satisfy the predicate and elements that don't.

Parameters:
- `predicate`: Function to test each element; return true to include in first group, false for second group

Returns:
- `[][]T`: A slice containing two slices: the first with elements that passed the test, the second with elements that failed

```go
seq := seq.New(1, 2, 3, 4)
result := seq.Partition(func(n int) bool {
    return n % 2 == 0
})
// result is [][]int{{2, 4}, {1, 3}}
```

#### GroupBy

Creates a map that groups elements by keys generated from the iteratee function.

Parameters:
- `iteratee`: Function that returns a string key for each element

Returns:
- `map[string][]T`: A map where keys are the strings returned by iteratee and values are slices of elements that produced each key

```go
// Group numbers by even/odd
seq := seq.New(1, 2, 3, 4)
result := seq.GroupBy(func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
})
// result is map[string][]int{"odd": {1, 3}, "even": {2, 4}}
```

#### CountBy

Creates a map that counts elements by keys generated from the iteratee function.

Parameters:
- `iteratee`: Function that returns a string key for each element

Returns:
- `map[string]int`: A map where keys are the strings returned by iteratee and values are counts of elements that produced each key

```go
// Count numbers by even/odd
seq := seq.New(1, 2, 3, 4)
result := seq.CountBy(func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
})
// result is map[string]int{"odd": 2, "even": 2}
```

#### KeyBy

Creates a map with keys generated by applying the iteratee function to each element.

Parameters:
- `iteratee`: Function that returns an integer key for each element

Returns:
- `map[int]T`: A map where keys are integers returned by iteratee and values are the elements that produced each key

```go
// Create a map of users keyed by their ID
type User struct {
    ID   int
    Name string
}
users := seq.New(User{1, "Alice"}, User{2, "Bob"})
result := users.KeyBy(func(u User) int {
    return u.ID
})
// result is map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
```

#### SortBy

Creates a new sequence sorted by the values returned by the iteratee function.

Parameters:
- `iteratee`: Function that returns a comparable value used for sorting

Returns:
- `*Sequence[T]`: A new sequence with elements sorted by the iteratee function results

```go
// Sort numbers in ascending order
seq := seq.New(1, 3, 2)
result := seq.SortBy(func(n int) int {
    return n
})
// result.Value() returns []int{1, 2, 3}

// Sort users by ID
type User struct {
    ID   int
    Name string
}
users := seq.New(User{3, "Charlie"}, User{1, "Alice"}, User{2, "Bob"})
result := users.SortBy(func(u User) int {
    return u.ID
})
// result.Value() returns users sorted by ID: [{1 Alice} {2 Bob} {3 Charlie}]
```

#### OrderBy

Sorts the sequence by the results of running each element through iteratee.

Parameters:
- `iteratee`: Function that returns an integer value used for sorting comparison
- `ascending`: If true sort in ascending order, if false sort in descending order

Returns:
- `*Sequence[T]`: A new sequence with elements sorted based on iteratee results

```go
// Sort numbers in ascending order
seq := seq.New(1, 3, 2)
result := seq.OrderBy(func(n int) int {
    return n
}, true)
// result.Value() returns []int{1, 2, 3}

// Sort users by age in descending order
type User struct {
    Name string
    Age  int
}
users := seq.New(
    User{"Alice", 25},
    User{"Bob", 30},
    User{"Charlie", 20},
)
result := users.OrderBy(func(u User) int {
    return u.Age
}, false)
// result.Value() returns users sorted by age in descending order: 
// [{Bob 30} {Alice 25} {Charlie 20}]
```

#### Join

Joins all elements of the sequence into a string.

```go
seq := seq.New(1, 2, 3)
result := seq.Join(",")
// result is "1,2,3"

seq = seq.New(1)
result = seq.Join(",")
// result is "1"

emptySeq := seq.New[int]()
result = emptySeq.Join(",")
// result is ""
```
