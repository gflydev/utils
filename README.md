# Go Lodash (gflydev/utils)

A modern Go utility library delivering modularity, performance, and extras. Inspired by [Lodash](https://lodash.com/) and [lo](https://github.com/samber/lo).

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

// Camelcase
result := str.Camelcase("foo bar") // "fooBar"

// Capitalize
result := str.Capitalize("fred") // "Fred"

// EndsWith
result := str.EndsWith("abc", "c") // true

// StartsWith
result := str.StartsWith("abc", "a") // true

// Trim
result := str.Trim("  abc  ") // "abc"
result := str.Trim("-_-abc-_-", "-_") // "abc"

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

// RandomInt
result := num.RandomInt(1, 10) // a random integer between 1 and 10

// Round
result := num.Round(4.7) // 5

// Floor
result := num.Floor(4.7) // 4

// Ceil
result := num.Ceil(4.3) // 5

// Max
result := num.Max(1, 2, 3) // 3

// Min
result := num.Min(1, 2, 3) // 1

// Sum
result := num.Sum(1, 2, 3) // 6

// Mean
result := num.Mean(1, 2, 3) // 2
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

// First
result, ok := arr.First([]int{1, 2, 3}) // 1, true

// Includes
result := arr.Includes([]int{1, 2, 3}, 2) // true

// IndexOf
result := arr.IndexOf([]int{1, 2, 3, 2}, 2) // 1

// Initial
result := arr.Initial([]int{1, 2, 3}) // []int{1, 2}

// Intersection
result := arr.Intersection([]int{1, 2, 3}, []int{2, 3, 4}) // []int{2, 3}

// Last
result, ok := arr.Last([]int{1, 2, 3}) // 3, true

// Pull
result := arr.Pull([]int{1, 2, 3, 1, 2, 3}, 2, 3) // []int{1, 1}

// Reverse
result := arr.Reverse([]int{1, 2, 3}) // []int{3, 2, 1}

// Shuffle
result := arr.Shuffle([]int{1, 2, 3, 4}) // []int{3, 1, 4, 2} (random order)

// Slice
result := arr.Slice([]int{1, 2, 3, 4}, 1, 3) // []int{2, 3}

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

// Has
result := obj.Has(map[string]interface{}{"a": 1, "b": 2}, "a") // true

// Keys
result := obj.Keys(map[string]interface{}{"a": 1, "b": 2}) // []string{"a", "b"}

// MapValues
result := obj.MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 }) // map[string]int{"a": 2, "b": 4}

// Omit
result := obj.Omit(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c") // map[string]int{"b": 2}

// Pick
result := obj.Pick(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c") // map[string]int{"a": 1, "c": 3}

// Values
result := obj.Values(map[string]int{"a": 1, "b": 2}) // []int{1, 2}
```

### Collection Utilities

```go
import "github.com/gflydev/utils/coll"

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

// ForEach
coll.ForEach([]int{1, 2, 3}, func(n int) { fmt.Println(n) })

// GroupBy
result := coll.GroupBy([]int{1, 2, 3, 4}, func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}) // map[string][]int{"odd": {1, 3}, "even": {2, 4}}

// Map
result := coll.Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // []int{2, 4, 6}

// Partition
result := coll.Partition([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // [][]int{{2, 4}, {1, 3}}

// Reduce
result := coll.Reduce([]int{1, 2, 3}, func(sum, n int) int { return sum + n }, 0) // 6

// Reject
result := coll.Reject([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{1, 3}

// Sample
result, ok := coll.Sample([]int{1, 2, 3, 4}) // a random element from the collection

// Some
result := coll.Some([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 }) // true
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

// Throttle
f := fn.Throttle(func() { fmt.Println("called") }, 100*time.Millisecond)
f() // Calls the function immediately
f() // Ignores this call (less than 100ms since last call)
// Wait 100ms
f() // Calls the function again
```

### Sequence Utilities

```go
import "github.com/gflydev/utils/seq"

// Create a new sequence
s := seq.New(1, 2, 3, 4)

// Chain operations
result := s.
    Filter(func(n int) bool { return n % 2 == 0 }).
    Map(func(n int) int { return n * 2 }).
    Value() // []int{4, 8}

// More examples
s.First() // 1, true
s.Last() // 4, true
s.Includes(2) // true
s.Reverse().Value() // []int{4, 3, 2, 1}
s.Take(2).Value() // []int{1, 2}
s.Drop(2).Value() // []int{3, 4}
```

## License

MIT