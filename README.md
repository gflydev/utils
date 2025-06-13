# gFlyDev Utils (gflydev/utils)

A modern Go utility library delivering modularity, performance, and extras. Inspired by [Lodash](https://lodash.com/) and [Laravel Helpers](https://laravel.com/docs/12.x/helpers), [Laravel String](https://laravel.com/docs/12.x/strings), [Laravel Collection](https://laravel.com/docs/12.x/collections).

## Overview

This library provides a comprehensive set of utility functions for Go, organized by type:

- **String utilities** (`str`): Functions for string manipulation
- **Number utilities** (`num`): Functions for number manipulation
- **Array utilities** (`arr`): Functions for array/slice manipulation
- **Object utilities** (`obj`): Functions for object/map manipulation
- **Collection utilities** (`coll`): Functions for collection manipulation
- **Function utilities** (`fn`): Functions for function manipulation
- **Sequence utilities** (`seq`): Functions for sequence manipulation

## Installation

```bash
go get github.com/gflydev/utils
```

## Usage

### String Utilities

```go
import "github.com/gflydev/utils/str"

// ToString - Convert any value to string
result := str.ToString(123) // "123"

// RuneLength - Get the number of runes in a string
result := str.RuneLength("hello") // 5

// Words - Split a string into words
result := str.Words("fred, barney, & pebbles") // []string{"fred", "barney", "pebbles"}

// WordsPattern - Split a string into words using a custom pattern
result := str.WordsPattern("fred, barney, & pebbles", "[^,]+") // []string{"fred", " barney", " & pebbles"}

// Camelcase - Convert string to camelCase
result := str.Camelcase("foo bar") // "fooBar"

// KebabCase - Convert string to kebab-case
result := str.KebabCase("fooBar") // "foo-bar"

// SnakeCase - Convert string to snake_case
result := str.SnakeCase("fooBar") // "foo_bar"

// PascalCase - Convert string to PascalCase
result := str.PascalCase("foo bar") // "FooBar"

// Capitalize
result := str.Capitalize("fred") // "Fred"

// EndsWith
result := str.EndsWith("abc", "c") // true

// StartsWith
result := str.StartsWith("abc", "a") // true

// Trim
result := str.Trim("  abc  ") // "abc"
result := str.Trim("-_-abc-_-", "-_") // "abc"

// TrimStart - Trim characters from the start of a string
result := str.TrimStart("  abc  ") // "abc  "

// TrimEnd - Trim characters from the end of a string
result := str.TrimEnd("  abc  ") // "  abc"

// ToLower
result := str.ToLower("FRED") // "fred"

// ToUpper
result := str.ToUpper("fred") // "FRED"

// Split
result := str.Split("a-b-c", "-") // []string{"a", "b", "c"}

// Join
result := str.Join([]string{"a", "b", "c"}, "-") // "a-b-c"

// Repeat
result := str.Repeat("abc", 2) // "abcabc"

// Replace
result := str.Replace("Hi Fred", "Fred", "Barney") // "Hi Barney"

// Contains
result := str.Contains("abc", "b") // true

// Count - Count occurrences of a substring
result := str.Count("abcabc", "ab") // 2

// Index - Find the index of a substring
result := str.Index("abc", "b") // 1

// LastIndex - Find the last index of a substring
result := str.LastIndex("abcabc", "ab") // 3

// Ellipsis - Truncate a string with ellipsis
result := str.Ellipsis("This is a long text", 10) // "This is..."
```

### Number Utilities

```go
import "github.com/gflydev/utils/num"

// Clamp
result := num.Clamp(10, 0, 5) // 5

// InRange
result := num.InRange(3, 2, 4) // true

// Random
result := num.Random(1, 10) // a random number between 1 and 10

// Round
result := num.Round(4.7) // 5
result := num.Round(4.7234, 2) // 4.72

// Floor
result := num.Floor(4.7) // 4
result := num.Floor(4.7234, 2) // 4.72

// Ceil
result := num.Ceil(4.3) // 5
result := num.Ceil(4.3234, 2) // 4.33

// Max
result := num.Max(1, 2, 3) // 3

// MaxBy - Find the maximum value using an iteratee function
result := num.MaxBy([]string{"abc", "a", "ab"}, func(s string) float64 { 
    return float64(len(s)) 
}) // "abc"

// Min
result := num.Min(1, 2, 3) // 1

// MinBy - Find the minimum value using an iteratee function
result := num.MinBy([]string{"abc", "a", "ab"}, func(s string) float64 { 
    return float64(len(s)) 
}) // "a"

// Sum
result := num.Sum(1, 2, 3) // 6

// SumBy - Sum values using an iteratee function
result := num.SumBy([]string{"a", "ab", "abc"}, func(s string) float64 { 
    return float64(len(s)) 
}) // 6.0

// Mean
result := num.Mean(1, 2, 3) // 2

// MeanBy - Calculate mean using an iteratee function
result := num.MeanBy([]string{"a", "ab", "abc"}, func(s string) float64 { 
    return float64(len(s)) 
}) // 2.0

// Abs - Get absolute value
result := num.Abs(-5) // 5

// Pow - Calculate power
result := num.Pow(2, 3) // 8

// Sqrt - Calculate square root
result := num.Sqrt(9) // 3
```

### Array Utilities

```go
import "github.com/gflydev/utils/arr"

// Chunk
result := arr.Chunk([]int{1, 2, 3, 4}, 2) // [][]int{{1, 2}, {3, 4}}

// Compact
result := arr.Compact([]int{0, 1, 2, 0, 3}) // []int{1, 2, 3}

// Concat
result := arr.Concat([]int{1, 2}, []int{3, 4}) // []int{1, 2, 3, 4}

// Difference
result := arr.Difference([]int{1, 2, 3}, []int{2, 3, 4}) // []int{1}

// Drop
result := arr.Drop([]int{1, 2, 3, 4}, 2) // []int{3, 4}

// DropRight
result := arr.DropRight([]int{1, 2, 3, 4}, 2) // []int{1, 2}

// Fill
result := arr.Fill([]int{1, 2, 3, 4}, 0, 1, 3) // []int{1, 0, 0, 4}

// FindIndex
result := arr.FindIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) // 2

// FindLastIndex - Find the last index where predicate returns true
result := arr.FindLastIndex([]int{1, 2, 3, 2}, func(n int) bool { return n == 2 }) // 3

// First
result, ok := arr.First([]int{1, 2, 3}) // 1, true

// Flatten - Flatten a nested array
result := arr.Flatten([][]int{{1, 2}, {3, 4}}) // []int{1, 2, 3, 4}

// Includes
result := arr.Includes([]int{1, 2, 3}, 2) // true

// IndexOf
result := arr.IndexOf([]int{1, 2, 3, 2}, 2) // 1

// LastIndexOf - Find the last index of a value
result := arr.LastIndexOf([]int{1, 2, 3, 2}, 2) // 3

// Initial
result := arr.Initial([]int{1, 2, 3}) // []int{1, 2}

// Intersection
result := arr.Intersection([]int{1, 2, 3}, []int{2, 3, 4}) // []int{2, 3}

// Join - Join array elements with a separator
result := arr.Join([]int{1, 2, 3}, "-") // "1-2-3"

// Last
result, ok := arr.Last([]int{1, 2, 3}) // 3, true

// Nth - Get the nth element of an array
result, ok := arr.Nth([]int{1, 2, 3, 4}, 2) // 3, true

// Pull
result := arr.Pull([]int{1, 2, 3, 1, 2, 3}, 2, 3) // []int{1, 1}

// Random - Get n random elements from an array
result := arr.Random([]int{1, 2, 3, 4, 5}, 2) // e.g., []int{3, 1}

// RandomChoice - Get a random element from an array
result, ok := arr.RandomChoice([]int{1, 2, 3, 4, 5}) // e.g., 3, true

// Reverse
result := arr.Reverse([]int{1, 2, 3}) // []int{3, 2, 1}

// Shuffle
result := arr.Shuffle([]int{1, 2, 3, 4}) // []int{3, 1, 4, 2} (random order)

// Slice
result := arr.Slice([]int{1, 2, 3, 4}, 1, 3) // []int{2, 3}

// SortedIndex - Get the index at which value should be inserted
result := arr.SortedIndex([]int{1, 3, 5, 7}, 4) // 2

// Tail
result := arr.Tail([]int{1, 2, 3}) // []int{2, 3}

// Take
result := arr.Take([]int{1, 2, 3, 4}, 2) // []int{1, 2}

// TakeRight
result := arr.TakeRight([]int{1, 2, 3, 4}, 2) // []int{3, 4}

// Union
result := arr.Union([]int{1, 2}, []int{2, 3}) // []int{1, 2, 3}

// Uniq
result := arr.Uniq([]int{1, 2, 1, 3}) // []int{1, 2, 3}

// Without
result := arr.Without([]int{1, 2, 3, 4}, 2, 4) // []int{1, 3}

// Zip - Create an array of grouped elements
result := arr.Zip([]int{1, 2}, []int{3, 4}) // [][]int{{1, 3}, {2, 4}}

// SortBy - Sort an array using an iteratee function
result := arr.SortBy([]string{"abc", "a", "ab"}, func(s string) int { 
    return len(s) 
}) // []string{"a", "ab", "abc"}
```

### Object Utilities

```go
import "github.com/gflydev/utils/obj"

// Assign
result := obj.Assign(map[string]interface{}{"a": 1}, map[string]interface{}{"b": 2}) // map[string]interface{}{"a": 1, "b": 2}

// Clone
result := obj.Clone(map[string]interface{}{"a": 1, "b": 2}) // map[string]interface{}{"a": 1, "b": 2}

// Entries
result := obj.Entries(map[string]int{"a": 1, "b": 2}) // []obj.Entry[string, int]{{"a", 1}, {"b", 2}}

// FromEntries
result := obj.FromEntries([]obj.Entry[string, int]{{"a", 1}, {"b", 2}}) // map[string]int{"a": 1, "b": 2}

// Get - Get a value from an object by path
result, ok := obj.Get[int](map[string]any{"a": map[string]any{"b": 1}}, "a.b") // 1, true

// Has
result := obj.Has(map[string]interface{}{"a": 1, "b": 2}, "a") // true

// Keys
result := obj.Keys(map[string]interface{}{"a": 1, "b": 2}) // []string{"a", "b"}

// KeysSorted - Get sorted keys of an object
result := obj.KeysSorted(map[string]interface{}{"b": 2, "a": 1}) // []string{"a", "b"}

// MapValues
result := obj.MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 }) // map[string]int{"a": 2, "b": 4}

// MapKeys - Transform keys of an object
result := obj.MapKeys(map[string]int{"a": 1, "b": 2}, func(k string) string { return k + "x" }) // map[string]int{"ax": 1, "bx": 2}

// Merge - Merge objects
result := obj.Merge(map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"c": 3}) // map[string]int{"a": 1, "b": 2, "c": 3}

// Omit
result := obj.Omit(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c") // map[string]int{"b": 2}

// OmitBy - Omit properties that satisfy the predicate
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { return v > 1 }) // map[string]int{"a": 1}

// Pick
result := obj.Pick(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c") // map[string]int{"a": 1, "c": 3}

// PickBy - Pick properties that satisfy the predicate
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { return v > 1 }) // map[string]int{"b": 2, "c": 3}

// Values
result := obj.Values(map[string]int{"a": 1, "b": 2}) // []int{1, 2}

// Size - Get the size of an object
result := obj.Size(map[string]int{"a": 1, "b": 2}) // 2

// IsEmpty - Check if an object is empty
result := obj.IsEmpty(map[string]int{}) // true

// IsEqual - Check if two objects are equal
result := obj.IsEqual(map[string]int{"a": 1}, map[string]int{"a": 1}) // true
```

### Collection Utilities

```go
import (
    "fmt"
    "github.com/gflydev/utils/coll"
)

// CountBy
result := coll.CountBy([]int{1, 2, 3, 4}, func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}) // map[string]int{"odd": 2, "even": 2}

// Every
result := coll.Every([]int{2, 4, 6}, func(n int) bool { return n % 2 == 0 }) // true

// Filter
result := coll.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{2, 4}

// Find
result, ok := coll.Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) // 3, true

// FindLast - Find the last element that satisfies the predicate
result, ok := coll.FindLast([]int{1, 2, 3, 4, 3}, func(n int) bool { return n > 2 }) // 3, true

// ForEach
coll.ForEach([]int{1, 2, 3}, func(n int) { fmt.Println(n) })

// ForEachWithIndex - Iterate with index
coll.ForEachWithIndex([]int{1, 2, 3}, func(n, i int) { fmt.Printf("%d: %d\n", i, n) })

// GroupBy
result := coll.GroupBy([]int{1, 2, 3, 4}, func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}) // map[string][]int{"odd": {1, 3}, "even": {2, 4}}

// KeyBy - Create an object from an array using a key function
result := coll.KeyBy([]string{"a", "ab", "abc"}, func(s string) int { return len(s) }) // map[int]string{1: "a", 2: "ab", 3: "abc"}

// Map
result := coll.Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // []int{2, 4, 6}

// MapWithIndex - Map with index
result := coll.MapWithIndex([]int{1, 2, 3}, func(n, i int) int { return n * i }) // []int{0, 2, 6}

// Partition
result := coll.Partition([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // [][]int{{2, 4}, {1, 3}}

// Reduce
result := coll.Reduce([]int{1, 2, 3}, func(sum, n int) int { return sum + n }, 0) // 6

// ReduceRight - Reduce from right to left
result := coll.ReduceRight([]int{1, 2, 3}, func(sum, n int) int { return sum - n }, 0) // -6

// Reject
result := coll.Reject([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{1, 3}

// Sample
result, ok := coll.Sample([]int{1, 2, 3, 4}) // a random element from the collection

// SampleSize - Get n random elements
result := coll.SampleSize([]int{1, 2, 3, 4, 5}, 2) // e.g., []int{3, 1}

// Size - Get the size of a collection
result := coll.Size([]int{1, 2, 3}) // 3

// Some
result := coll.Some([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 }) // true

// SortBy - Sort a collection using an iteratee function
result := coll.SortBy([]string{"abc", "a", "ab"}, func(s string) int { return len(s) }) // []string{"a", "ab", "abc"}

// OrderBy - Sort a collection with direction control
result := coll.OrderBy([]int{3, 1, 2}, func(n int) int { return n }, false) // []int{3, 2, 1}

// ForEachMap - Iterate over a map
coll.ForEachMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) { fmt.Printf("%s: %d\n", k, v) })

// MapMap - Map over a map's values
result := coll.MapMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) string { return fmt.Sprintf("%s-%d", k, v) }) // []string{"a-1", "b-2"}

// FilterMap - Filter a map
result := coll.FilterMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int, k string) bool { return v > 1 }) // map[string]int{"b": 2, "c": 3}

// ReduceMap - Reduce a map
result := coll.ReduceMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(acc int, v int, k string) int { return acc + v }, 0) // 6
```

### Function Utilities

```go
import (
    "fmt"
    "github.com/gflydev/utils/fn"
    "time"
)

// After
f := fn.After(3, func() string { return "done" })
f() // ""
f() // ""
f() // "done"

// Before
f := fn.Before(3, func() string { return "called" })
f() // "called"
f() // "called"
f() // "called"
f() // "called" (returns the result of the last successful invocation)

// Curry
add := func(a, b int) int { return a + b }
addCurried := fn.Curry(add, 2)
add1 := addCurried(1)
add1(2) // 3

// Debounce
f := fn.Debounce(func() { fmt.Println("called") }, 100*time.Millisecond)
f() // Schedules the function to be called after 100ms
f() // Cancels the previous schedule and reschedules
f() // Cancels the previous schedule and reschedules

// Delay - Delay execution of a function
fn.Delay(func() { fmt.Println("delayed") }, 100*time.Millisecond)

// Memoize
fibonacci := fn.Memoize(func(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
})
fibonacci(10) // 55 (calculated efficiently with memoization)

// Once
initialize := fn.Once(func() string { return "initialized" })
initialize() // "initialized"
initialize() // "initialized" (doesn't call the function again)

// Partial - Partially apply a function
add := func(a, b int) int { return a + b }
add5 := fn.Partial(add, 5)
add5(10) // 15

// Rearg - Rearrange arguments
swap := fn.Rearg(func(a, b int) string { return fmt.Sprintf("%d,%d", a, b) })
swap(1, 2) // "2,1"

// Throttle
f := fn.Throttle(func() { fmt.Println("called") }, 100*time.Millisecond)
f() // Calls the function immediately
f() // Ignores this call (less than 100ms since last call)
// Wait 100ms
f() // Calls the function again

// Wrap - Wrap a function with another function
hello := func(name string) string { return "hello " + name }
wrapped := fn.Wrap(hello, func(f func(string) string, name string) string {
    return f(name) + "!"
})
wrapped("world") // "hello world!"

// Retry - Retry a function until it succeeds or max retries is reached
result, err := fn.Retry(func() (string, error) {
    // Some operation that might fail
    return "success", nil
}, 3, 100*time.Millisecond)()

// Compose - Compose functions from right to left
addOne := func(x int) int { return x + 1 }
double := func(x int) int { return x * 2 }
composed := fn.Compose(double, addOne)
composed(3) // 8 (double(addOne(3)))

// Pipe - Compose functions from left to right
piped := fn.Pipe(addOne, double)
piped(3) // 8 (double(addOne(3)))

// Negate - Create a function that negates the result of the predicate
isEven := func(n int) bool { return n % 2 == 0 }
isOdd := fn.Negate(isEven)
isOdd(3) // true

// Spread - Convert a function that takes multiple arguments to one that takes an array
add := func(a, b int) int { return a + b }
addSpread := fn.Spread(add)
addSpread([]int{1, 2}) // 3

// TransformList - Transform a list using a transformer function
numbers := []int{1, 2, 3}
squares := fn.TransformList(numbers, func(x int) int { return x * x }) // []int{1, 4, 9}

// TransformMap - Transform a map using a transformer function
ages := map[string]int{"John": 30, "Jane": 25}
doubled := fn.TransformMap(ages, func(age int) int { return age * 2 }) // map[string]int{"John": 60, "Jane": 50}

// TransformListWithError - Transform a list and collect errors
results, errors := fn.TransformListWithError([]string{"1", "2", "x"}, func(s string) (int, error) {
    return strconv.Atoi(s)
}) // results = []int{1, 2}, errors contains the error for "x"

// TransformConcurrent - Transform a list concurrently with multiple workers
numbers := []int{1, 2, 3, 4, 5}
squares := fn.TransformConcurrent(numbers, func(x int) int { return x * x }, 2) // []int{1, 4, 9, 16, 25}

// TransformBatch - Transform a list in batches
numbers := []int{1, 2, 3, 4, 5}
doubled := fn.TransformBatch(numbers, func(batch []int) []int {
    var result []int
    for _, n := range batch {
        result = append(result, n*2)
    }
    return result
}, 2) // []int{2, 4, 6, 8, 10}

// TransformBatch - Transform a list in batches
input := []int{1, 2, 3, 4, 5, 6, 7}
batchSize := 3
var batches [][]int

transformerFn := func(batch []int) []string {
    batches = append(batches, batch)
    result := make([]string, len(batch))
    for i, v := range batch {
        result[i] = strconv.Itoa(v)
    }
    return result
}

result := TransformBatch(input, transformerFn, batchSize) 
// result is []string{"1", "2", "3", "4", "5", "6", "7"}
// batches is [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
```

### Sequence Utilities

```go
import (
    "fmt"
    "github.com/gflydev/utils/seq"
)

// Create a new sequence
s := seq.New(1, 2, 3, 4)

// FromSlice - Create a sequence from a slice
s := seq.FromSlice([]int{1, 2, 3, 4})

// Chain operations
result := s.
    Filter(func(n int) bool { return n % 2 == 0 }).
    Map(func(n int) int { return n * 2 }).
    Value() // []int{4, 8}

// More examples
s.First() // 1, true
s.Last() // 4, true
s.Includes(2) // true
s.Find(func(n int) bool { return n > 2 }) // 3, true
s.FindLast(func(n int) bool { return n < 4 }) // 3, true
s.Every(func(n int) bool { return n > 0 }) // true
s.Some(func(n int) bool { return n > 3 }) // true
s.Size() // 4
s.IsEmpty() // false
s.Reverse().Value() // []int{4, 3, 2, 1}
s.Uniq().Value() // []int{1, 2, 3, 4}
s.Chunk(2) // [][]int{{1, 2}, {3, 4}}
s.Take(2).Value() // []int{1, 2}
s.TakeRight(2).Value() // []int{3, 4}
s.Drop(2).Value() // []int{3, 4}
s.DropRight(2).Value() // []int{1, 2}
s.Shuffle().Value() // random order
s.Sample() // random element
s.SampleSize(2).Value() // 2 random elements
s.Partition(func(n int) bool { return n % 2 == 0 }) // [][]int{{2, 4}, {1, 3}}
s.GroupBy(func(n int) string { return fmt.Sprintf("%d", n % 2) }) // map[string][]int{"0": {2, 4}, "1": {1, 3}}
s.CountBy(func(n int) string { return fmt.Sprintf("%d", n % 2) }) // map[string]int{"0": 2, "1": 2}
s.KeyBy(func(n int) int { return n % 2 }) // map[int]int{0: 4, 1: 3}
s.SortBy(func(n int) int { return -n }).Value() // []int{4, 3, 2, 1}
s.OrderBy(func(n int) int { return n }, false).Value() // []int{4, 3, 2, 1}
s.Join("-") // "1-2-3-4"

// MapTo - Map to a different type
result := s.MapTo(func(n int) string { return fmt.Sprintf("%d", n) }).Value() // []interface{}{"1", "2", "3", "4"}

// Reject - Filter out elements that satisfy the predicate
result := s.Reject(func(n int) bool { return n % 2 == 0 }).Value() // []int{1, 3}

// Reduce - Reduce the sequence to a single value
result := s.Reduce(func(acc interface{}, n int) interface{} { 
    return acc.(int) + n 
}, 0) // 10

// ForEach - Iterate over the sequence
s.ForEach(func(n int) { fmt.Println(n) }) // Prints 1, 2, 3, 4

// Flatten - Flatten a sequence of sequences
nestedSeq := seq.New(seq.New(1, 2), seq.New(3, 4))
result := nestedSeq.Flatten().Value() // []interface{}{1, 2, 3, 4}

// Concat - Concatenate sequences
result := s.Concat(seq.New(5, 6)).Value() // []int{1, 2, 3, 4, 5, 6}
```

## License

MIT
