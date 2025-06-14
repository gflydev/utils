# gFlyDev Utils (gflydev/utils)

A modern Go utility library delivering modularity, performance, and extras. Inspired by [Lodash](https://lodash.com/) and [Laravel](https://laravel.com) ([Helpers](https://laravel.com/docs/12.x/helpers), [String](https://laravel.com/docs/12.x/strings), [Collection](https://laravel.com/docs/12.x/collections)).

## Overview

This library provides a comprehensive set of utility functions for Go, organized by type:

- **String utilities** (`str`): Functions for string manipulation
- **Number utilities** (`num`): Functions for number manipulation
- **Array utilities** (`arr`): Functions for array/slice manipulation
- **Object utilities** (`obj`): Functions for object/map manipulation
- **Collection utilities** (`col`): Functions for collection manipulation
- **Function utilities** (`fn`): Functions for function manipulation
- **Sequence utilities** (`seq`): Functions for sequence manipulation
- **Network utilities** (`net`): Functions for HTTP and network operations

## Installation

```bash
go get github.com/gflydev/utils
```

## Usage

### String Utilities [Full document](str/README.md)

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

// Slugify - Convert a string to a URL-friendly slug
result := str.Slugify("Hello, World!") // "hello-world"

// ContainsAny - Check if a string contains any of the given substrings
result := str.ContainsAny("hello world", []string{"hello", "goodbye"}) // true

// ToTitleCase - Convert a string to title case
result := str.ToTitleCase("hello world") // "Hello World"

// OnlyAlphanumeric - Remove all non-alphanumeric characters from a string
result := str.OnlyAlphanumeric("Hello, World!") // "HelloWorld"

// Mask - Mask a portion of a string with a character
result := str.Mask("1234567890", '*', 4, 4) // "1234****90"

// PadLeft - Pad a string on the left to a specified length
result := str.PadLeft("hello", 10, ' ') // "     hello"

// PadRight - Pad a string on the right to a specified length
result := str.PadRight("hello", 10, ' ') // "hello     "

// Reverse - Reverse a string
result := str.Reverse("hello") // "olleh"

// CountWords - Count the number of words in a string
result := str.CountWords("hello world") // 2

// TruncateWords - Truncate a string to a specified number of words
result := str.TruncateWords("hello world goodbye", 2) // "hello world..."

// FormatWithCommas - Format a number as a string with commas
result := str.FormatWithCommas(1234567) // "1,234,567"

// After - Get the substring after the first occurrence of a delimiter
result := str.After("hello-world", "-") // "world"

// AfterLast - Get the substring after the last occurrence of a delimiter
result := str.AfterLast("hello-world-goodbye", "-") // "goodbye"

// Before - Get the substring before the first occurrence of a delimiter
result := str.Before("hello-world", "-") // "hello"

// BeforeLast - Get the substring before the last occurrence of a delimiter
result := str.BeforeLast("hello-world-goodbye", "-") // "hello-world"

// Between - Get the substring between two delimiters
result := str.Between("hello [world] goodbye", "[", "]") // "world"

// ContainsAll - Check if a string contains all of the given substrings
result := str.ContainsAll("hello world", []string{"hello", "world"}) // true

// Finish - Ensure a string ends with a specific suffix
result := str.Finish("hello", "!") // "hello!"

// Is - Check if a string matches a pattern
result := str.Is("foo*", "foobar") // true

// IsAscii - Check if a string contains only ASCII characters
result := str.IsAscii("hello") // true

// Limit - Limit the number of characters in a string
result := str.Limit("hello world", 5) // "hello..."

// Random - Generate a random string of a specified length
result := str.Random(10) // e.g., "aB3cD7eF9g"

// ReplaceArray - Replace multiple occurrences of a placeholder with different values
result := str.ReplaceArray("?", []string{"hello", "world"}, "? ?") // "hello world"

// ReplaceFirst - Replace the first occurrence of a substring
result := str.ReplaceFirst("hello hello", "hello", "hi") // "hi hello"

// ReplaceLast - Replace the last occurrence of a substring
result := str.ReplaceLast("hello hello", "hello", "hi") // "hello hi"

// Start - Ensure a string starts with a specific prefix
result := str.Start("world", "hello ") // "hello world"

// Studly - Convert a string to StudlyCase
result := str.Studly("hello_world") // "HelloWorld"

// Substr - Get a substring of a string
result := str.Substr("hello world", 6, 5) // "world"

// Ucfirst - Capitalize the first character of a string
result := str.Ucfirst("hello") // "Hello"

// Lcfirst - Lowercase the first character of a string
result := str.Lcfirst("Hello") // "hello"

// Plural - Convert a singular word to its plural form
result := str.Plural("apple") // "apples"

// Singular - Convert a plural word to its singular form
result := str.Singular("apples") // "apple"

// Wordwrap - Wrap a string to a given number of characters
result := str.Wordwrap("hello world", 5, "\n", true) // "hello\nworld"
```

### Number Utilities [Full document](num/README.md)

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

// Ceiling - Alias for Ceil
result := num.Ceiling(4.3) // 5

// Format - Format a number with a specified number of decimal places
result := num.Format(123.456, 2) // "123.46"

// FormatCompact - Format a number in a compact form
result := num.FormatCompact(1234567) // "1.2M"

// FormatPercentage - Format a number as a percentage
result := num.FormatPercentage(0.1234, 2) // "12.34%"

// Percent - Calculate the percentage of a value relative to a total
result := num.Percent(25, 100) // 25.0
```

### Array Utilities [Full document](arr/README.md)

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

// Contains - Check if an array contains a specific element
result := arr.Contains([]int{1, 2, 3, 4}, 3) // true

// Filter - Filter an array based on a predicate function
result := arr.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{2, 4}

// Map - Map an array using a transformation function
result := arr.Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // []int{2, 4, 6}

// Find - Find the first element that satisfies a predicate
result, ok := arr.Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) // 3, true

// FirstOrDefault - Get the first element or a default value if the array is empty
result := arr.FirstOrDefault([]int{1, 2, 3}, 0) // 1
result := arr.FirstOrDefault([]int{}, 42) // 42

// LastOrDefault - Get the last element or a default value if the array is empty
result := arr.LastOrDefault([]int{1, 2, 3}, 0) // 3
result := arr.LastOrDefault([]int{}, 42) // 42

// Prepend - Add elements to the beginning of an array
result := arr.Prepend([]int{3, 4}, 1, 2) // []int{1, 2, 3, 4}

// Unique - Remove duplicates from an array
result := arr.Unique([]int{1, 2, 2, 3, 3, 3}) // []int{1, 2, 3}

// SortedCopy - Create a sorted copy of an array
result := arr.SortedCopy([]int{3, 1, 2}, func(a, b int) bool { return a < b }) // []int{1, 2, 3}

// Reduce - Reduce an array to a single value
result := arr.Reduce([]int{1, 2, 3}, 0, func(acc, item int) int { return acc + item }) // 6

// GroupBy - Group array elements by a key function
result := arr.GroupBy([]string{"one", "two", "three"}, func(s string) int { return len(s) })
// map[int][]string{3: {"one", "two"}, 5: {"three"}}

// MapMerge - Merge multiple maps into one
result := arr.MapMerge(map[string]int{"a": 1}, map[string]int{"b": 2}) // map[string]int{"a": 1, "b": 2}

// MapKeys - Get the keys of a map
result := arr.MapKeys(map[string]int{"a": 1, "b": 2}) // []string{"a", "b"}

// MapValues - Get the values of a map
result := arr.MapValues(map[string]int{"a": 1, "b": 2}) // []int{1, 2}

// MapValuesFn - Transform the values of a map
result := arr.MapValuesFn(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 })
// map[string]int{"a": 2, "b": 4}

// MapGetOrDefault - Get a value from a map or a default if the key doesn't exist
result := arr.MapGetOrDefault(map[string]int{"a": 1}, "a", 0) // 1
result := arr.MapGetOrDefault(map[string]int{"a": 1}, "b", 0) // 0

// SetContains - Check if a set contains an element
result := arr.SetContains(map[string]struct{}{"a": {}, "b": {}}, "a") // true

// SetToSlice - Convert a set to a slice
result := arr.SetToSlice(map[string]struct{}{"a": {}, "b": {}}) // []string{"a", "b"}

// SliceToSet - Convert a slice to a set
result := arr.SliceToSet([]string{"a", "b", "a"}) // map[string]struct{}{"a": {}, "b": {}}
```

### Object Utilities [Full document](obj/README.md)

```go
import "github.com/gflydev/utils/obj"

// Assign
result := obj.Assign(map[string]any{"a": 1}, map[string]any{"b": 2}) // map[string]any{"a": 1, "b": 2}

// Clone
result := obj.Clone(map[string]any{"a": 1, "b": 2}) // map[string]any{"a": 1, "b": 2}

// Entries
result := obj.Entries(map[string]int{"a": 1, "b": 2}) // []obj.Entry[string, int]{{"a", 1}, {"b", 2}}

// FromEntries
result := obj.FromEntries([]obj.Entry[string, int]{{"a", 1}, {"b", 2}}) // map[string]int{"a": 1, "b": 2}

// Get - Get a value from an object by path
result, ok := obj.Get[int](map[string]any{"a": map[string]any{"b": 1}}, "a.b") // 1, true

// Has
result := obj.Has(map[string]any{"a": 1, "b": 2}, "a") // true

// Keys
result := obj.Keys(map[string]any{"a": 1, "b": 2}) // []string{"a", "b"}

// KeysSorted - Get sorted keys of an object
result := obj.KeysSorted(map[string]any{"b": 2, "a": 1}) // []string{"a", "b"}

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

### Collection Utilities [Full document](col/README.md)

```go
import (
    "fmt"
    "github.com/gflydev/utils/col"
)

// CountBy
result := col.CountBy([]int{1, 2, 3, 4}, func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}) // map[string]int{"odd": 2, "even": 2}

// Every
result := col.Every([]int{2, 4, 6}, func(n int) bool { return n % 2 == 0 }) // true

// Filter
result := col.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{2, 4}

// Find
result, ok := col.Find([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 }) // 3, true

// FindLast - Find the last element that satisfies the predicate
result, ok := col.FindLast([]int{1, 2, 3, 4, 3}, func(n int) bool { return n > 2 }) // 3, true

// ForEach
col.ForEach([]int{1, 2, 3}, func(n int) { fmt.Println(n) })

// ForEachWithIndex - Iterate with index
col.ForEachWithIndex([]int{1, 2, 3}, func(n, i int) { fmt.Printf("%d: %d\n", i, n) })

// GroupBy
result := col.GroupBy([]int{1, 2, 3, 4}, func(n int) string {
    if n % 2 == 0 {
        return "even"
    }
    return "odd"
}) // map[string][]int{"odd": {1, 3}, "even": {2, 4}}

// KeyBy - Create an object from an array using a key function
result := col.KeyBy([]string{"a", "ab", "abc"}, func(s string) int { return len(s) }) // map[int]string{1: "a", 2: "ab", 3: "abc"}

// Map
result := col.Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // []int{2, 4, 6}

// MapWithIndex - Map with index
result := col.MapWithIndex([]int{1, 2, 3}, func(n, i int) int { return n * i }) // []int{0, 2, 6}

// Partition
result := col.Partition([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // [][]int{{2, 4}, {1, 3}}

// Reduce
result := col.Reduce([]int{1, 2, 3}, func(sum, n int) int { return sum + n }, 0) // 6

// ReduceRight - Reduce from right to left
result := col.ReduceRight([]int{1, 2, 3}, func(sum, n int) int { return sum - n }, 0) // -6

// Reject
result := col.Reject([]int{1, 2, 3, 4}, func(n int) bool { return n % 2 == 0 }) // []int{1, 3}

// Sample
result, ok := col.Sample([]int{1, 2, 3, 4}) // a random element from the collection

// SampleSize - Get n random elements
result := col.SampleSize([]int{1, 2, 3, 4, 5}, 2) // e.g., []int{3, 1}

// Size - Get the size of a collection
result := col.Size([]int{1, 2, 3}) // 3

// Some
result := col.Some([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 }) // true

// SortBy - Sort a collection using an iteratee function
result := col.SortBy([]string{"abc", "a", "ab"}, func(s string) int { return len(s) }) // []string{"a", "ab", "abc"}

// OrderBy - Sort a collection with direction control
result := col.OrderBy([]int{3, 1, 2}, func(n int) int { return n }, false) // []int{3, 2, 1}

// ForEachMap - Iterate over a map
col.ForEachMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) { fmt.Printf("%s: %d\n", k, v) })

// MapMap - Map over a map's values
result := col.MapMap(map[string]int{"a": 1, "b": 2}, func(v int, k string) string { return fmt.Sprintf("%s-%d", k, v) }) // []string{"a-1", "b-2"}

// FilterMap - Filter a map
result := col.FilterMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int, k string) bool { return v > 1 }) // map[string]int{"b": 2, "c": 3}

// ReduceMap - Reduce a map
result := col.ReduceMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(acc int, v int, k string) int { return acc + v }, 0) // 6

// Avg - Calculate the average of a collection
result := col.Avg([]int{1, 2, 3, 4, 5}) // 3.0

// Chunk - Split a collection into chunks of a given size
result := col.Chunk([]int{1, 2, 3, 4, 5}, 2) // [][]int{{1, 2}, {3, 4}, {5}}

// Contains - Check if a collection contains a specific element
result := col.Contains([]int{1, 2, 3, 4}, 3) // true

// Reverse - Reverse the order of elements in a collection
result := col.Reverse([]int{1, 2, 3, 4}) // []int{4, 3, 2, 1}

// Slice - Get a slice of a collection
result := col.Slice([]int{1, 2, 3, 4, 5}, 1, 3) // []int{2, 3, 4}

// SliceWithLength - Get a slice of a collection with a specified length
result := col.SliceWithLength([]int{1, 2, 3, 4, 5}, 1, 3) // []int{2, 3, 4}

// Shuffle - Randomize the order of elements in a collection
result := col.Shuffle([]int{1, 2, 3, 4, 5}) // e.g., []int{3, 1, 5, 2, 4}

// Collapse - Flatten a collection of arrays into a single array
result := col.Collapse([][]int{{1, 2}, {3, 4}}) // []int{1, 2, 3, 4}

// CrossJoin - Cross join multiple collections
result := col.CrossJoin([]int{1, 2}, []string{"a", "b"}) // [][]any{{1, "a"}, {1, "b"}, {2, "a"}, {2, "b"}}

// Diff - Get the difference between two collections
result := col.Diff([]int{1, 2, 3, 4}, []int{2, 4, 5, 6}) // []int{1, 3}

// DiffAssoc - Get the difference between two collections with keys
result := col.DiffAssoc(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2, "c": 3}) // map[string]int{"a": 1}

// DiffKeys - Get the difference between two collections by keys
result := col.DiffKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}) // map[string]int{"a": 1}

// Except - Get all elements except those with specified keys
result := col.Except(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"}) // map[string]int{"b": 2}

// First - Get the first element of a collection
result, ok := col.First([]int{1, 2, 3}) // 1, true

// FirstOrDefault - Get the first element or a default value if the collection is empty
result := col.FirstOrDefault([]int{1, 2, 3}, 0) // 1
result := col.FirstOrDefault([]int{}, 42) // 42

// FlatMap - Map a collection and flatten the result
result := col.FlatMap([]string{"a,b", "c,d"}, func(s string) []string { return strings.Split(s, ",") }) // []string{"a", "b", "c", "d"}

// Flatten - Flatten a nested collection
result := col.Flatten([][]int{{1, 2}, {3, 4}}) // []int{1, 2, 3, 4}

// Flip - Swap the keys and values of a map
result := col.Flip(map[string]int{"a": 1, "b": 2}) // map[int]string{1: "a", 2: "b"}

// Forget - Remove an element from a map by key
result := col.Forget(map[string]int{"a": 1, "b": 2}, "a") // map[string]int{"b": 2}

// Get - Get a value from a map by key with a default value
result := col.Get(map[string]int{"a": 1, "b": 2}, "c", 3) // 3

// Has - Check if a map has a specific key
result := col.Has(map[string]int{"a": 1, "b": 2}, "a") // true

// Implode - Join the elements of a collection with a string
result := col.Implode([]int{1, 2, 3}, "-") // "1-2-3"

// Intersect - Get the intersection of two collections
result := col.Intersect([]int{1, 2, 3}, []int{2, 3, 4}) // []int{2, 3}

// IntersectByKeys - Get the intersection of two maps by keys
result := col.IntersectByKeys(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}) // map[string]int{"b": 2}

// IsEmpty - Check if a collection is empty
result := col.IsEmpty([]int{}) // true

// IsNotEmpty - Check if a collection is not empty
result := col.IsNotEmpty([]int{1, 2, 3}) // true

// Keys - Get the keys of a map
result := col.Keys(map[string]int{"a": 1, "b": 2}) // []string{"a", "b"}

// Last - Get the last element of a collection
result, ok := col.Last([]int{1, 2, 3}) // 3, true

// LastOrDefault - Get the last element or a default value if the collection is empty
result := col.LastOrDefault([]int{1, 2, 3}, 0) // 3
result := col.LastOrDefault([]int{}, 42) // 42

// Max - Get the maximum value in a collection
result := col.Max([]int{1, 5, 3, 2, 4}) // 5

// Merge - Merge two maps
result := col.Merge(map[string]int{"a": 1}, map[string]int{"b": 2}) // map[string]int{"a": 1, "b": 2}

// Min - Get the minimum value in a collection
result := col.Min([]int{5, 3, 1, 4, 2}) // 1

// Only - Get a subset of a map with only the specified keys
result := col.Only(map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "c"}) // map[string]int{"a": 1, "c": 3}

// Pad - Pad a collection to a specified length with a value
result := col.Pad([]int{1, 2, 3}, 5, 0) // []int{1, 2, 3, 0, 0}

// Pluck - Extract a specific key from a collection of maps
type User struct { Name string; Age int }
users := []User{{"Alice", 25}, {"Bob", 30}}
result := col.Pluck(users, func(u User) string { return u.Name }) // []string{"Alice", "Bob"}

// Prepend - Add elements to the beginning of a collection
result := col.Prepend([]int{3, 4}, 1, 2) // []int{1, 2, 3, 4}

// Pull - Remove and return an element from a collection by key
result, ok := col.Pull([]int{1, 2, 3, 4}, 2) // 3, true

// Push - Add an element to the end of a collection
result := col.Push([]int{1, 2, 3}, 4) // []int{1, 2, 3, 4}

// Put - Set a value in a map
result := col.Put(map[string]int{"a": 1}, "b", 2) // map[string]int{"a": 1, "b": 2}

// Random - Get a random element from a collection
result, ok := col.Random([]int{1, 2, 3, 4, 5}) // e.g., 3, true

// RandomOrDefault - Get a random element or a default value if the collection is empty
result := col.RandomOrDefault([]int{1, 2, 3, 4, 5}, 0) // e.g., 3
result := col.RandomOrDefault([]int{}, 42) // 42

// Search - Search for a value in a collection and return its key
result, ok := col.Search([]int{1, 2, 3, 4}, 3) // 2, true

// Shift - Remove and return the first element of a collection
result, rest, ok := col.Shift([]int{1, 2, 3}) // 1, []int{2, 3}, true

// Sort - Sort a collection
result := col.Sort([]int{3, 1, 4, 2}) // []int{1, 2, 3, 4}

// SortByDesc - Sort a collection in descending order
result := col.SortByDesc([]int{1, 2, 3, 4}) // []int{4, 3, 2, 1}

// Splice - Remove and return a portion of a collection
result, rest := col.Splice([]int{1, 2, 3, 4, 5}, 1, 3) // []int{2, 3, 4}, []int{1, 5}

// Split - Split a collection into groups of a given size
result := col.Split([]int{1, 2, 3, 4, 5, 6}, 2) // [][]int{{1, 2, 3}, {4, 5, 6}}

// Sum - Calculate the sum of a collection
result := col.Sum([]int{1, 2, 3, 4, 5}) // 15

// Take - Take the first n elements of a collection
result := col.Take([]int{1, 2, 3, 4, 5}, 3) // []int{1, 2, 3}

// Tap - Pass a collection to a callback and return the collection
result := col.Tap([]int{1, 2, 3}, func(arr []int) { fmt.Println(arr) }) // []int{1, 2, 3}

// Unique - Remove duplicates from a collection
result := col.Unique([]int{1, 2, 2, 3, 3, 3}) // []int{1, 2, 3}

// UniqueBy - Remove duplicates from a collection using a key function
result := col.UniqueBy([]int{1, 2, 3, 4, 5}, func(n int) int { return n % 3 }) // []int{1, 2, 3}

// Values - Get the values of a map
result := col.Values(map[string]int{"a": 1, "b": 2}) // []int{1, 2}

// Zip - Combine multiple collections
result := col.Zip([]int{1, 2}, []string{"a", "b"}) // [][]any{{1, "a"}, {2, "b"}}

// Unless - Execute a callback unless a condition is true
result := col.Unless(true, []int{1, 2, 3}, func(arr []int) []int { return append(arr, 4) }) // []int{1, 2, 3}

// UnlessEmpty - Execute a callback unless a collection is empty
result := col.UnlessEmpty([]int{}, []int{1, 2, 3}, func(arr []int) []int { return append(arr, 4) }) // []int{1, 2, 3}

// UnlessNotEmpty - Execute a callback unless a collection is not empty
result := col.UnlessNotEmpty([]int{1, 2, 3}, []int{}, func(arr []int) []int { return append(arr, 4) }) // []int{}

// When - Execute a callback when a condition is true
result := col.When(true, []int{1, 2, 3}, func(arr []int) []int { return append(arr, 4) }) // []int{1, 2, 3, 4}

// WhenEmpty - Execute a callback when a collection is empty
result := col.WhenEmpty([]int{}, []int{}, func(arr []int) []int { return append(arr, 1) }) // []int{1}

// WhenNotEmpty - Execute a callback when a collection is not empty
result := col.WhenNotEmpty([]int{1, 2, 3}, []int{1, 2, 3}, func(arr []int) []int { return append(arr, 4) }) // []int{1, 2, 3, 4}

// Where - Filter a collection by a key/value pair
result := col.Where([]map[string]int{{"id": 1, "value": 10}, {"id": 2, "value": 20}}, "id", 1) // []map[string]int{{"id": 1, "value": 10}}

// WhereIn - Filter a collection by a key/value pair with multiple values
result := col.WhereIn([]map[string]int{{"id": 1}, {"id": 2}, {"id": 3}}, "id", []int{1, 3}) // []map[string]int{{"id": 1}, {"id": 3}}

// WhereNotIn - Filter a collection by a key/value pair with excluded values
result := col.WhereNotIn([]map[string]int{{"id": 1}, {"id": 2}, {"id": 3}}, "id", []int{1, 3}) // []map[string]int{{"id": 2}}

// ContainsFn - Check if a collection contains an element that satisfies a predicate
result := col.ContainsFn([]int{1, 2, 3, 4}, func(n int) bool { return n > 3 }) // true

// Count - Count the number of elements in a collection
result := col.Count([]int{1, 2, 3, 4, 5}) // 5

// Each - Iterate over a collection
result := col.Each([]int{1, 2, 3}, func(n int) bool { fmt.Println(n); return true }) // []int{1, 2, 3}
```

### Function Utilities [Full document](fn/README.md)

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

### Sequence Utilities [Full document](seq/README.md)

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
result := s.MapTo(func(n int) string { return fmt.Sprintf("%d", n) }).Value() // []any{"1", "2", "3", "4"}

// Reject - Filter out elements that satisfy the predicate
result := s.Reject(func(n int) bool { return n % 2 == 0 }).Value() // []int{1, 3}

// Reduce - Reduce the sequence to a single value
result := s.Reduce(func(acc any, n int) any { 
    return acc.(int) + n 
}, 0) // 10

// ForEach - Iterate over the sequence
s.ForEach(func(n int) { fmt.Println(n) }) // Prints 1, 2, 3, 4

// Flatten - Flatten a sequence of sequences
nestedSeq := seq.New(seq.New(1, 2), seq.New(3, 4))
result := nestedSeq.Flatten().Value() // []any{1, 2, 3, 4}

// Concat - Concatenate sequences
result := s.Concat(seq.New(5, 6)).Value() // []int{1, 2, 3, 4, 5, 6}
```

### Network Utilities [Full document](net/README.md)

```go
import "github.com/gflydev/utils/net"

// BuildURL - Build a URL with query parameters
baseURL := "https://api.example.com/users"
params := map[string]string{"page": "1", "limit": "10"}
url, err := net.BuildURL(baseURL, params) // "https://api.example.com/users?limit=10&page=1"

// IsSuccessStatusCode - Check if an HTTP status code indicates success
result := net.IsSuccessStatusCode(200) // true
result := net.IsSuccessStatusCode(404) // false

// ParseQueryParams - Parse query parameters from a query string
params, err := net.ParseQueryParams("page=1&limit=10") // map[string]string{"page": "1", "limit": "10"}

// CreateHTTPClient - Create a customized HTTP client
client := net.CreateHTTPClient(10*time.Second, 100, 10, 100)

// GetJSON - Make a GET request and parse the JSON response
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
var user User
err := net.GetJSON("https://api.example.com/users/1", &user, nil)

// PostJSON - Make a POST request with JSON body and parse the JSON response
payload := map[string]string{"name": "John Doe"}
var response User
err := net.PostJSON("https://api.example.com/users", payload, &response, nil)

// PutJSON - Make a PUT request with JSON body and parse the JSON response
payload := map[string]string{"name": "John Doe"}
var response User
err := net.PutJSON("https://api.example.com/users/1", payload, &response, nil)

// DeleteJSON - Make a DELETE request and parse the JSON response
var response map[string]string
err := net.DeleteJSON("https://api.example.com/users/1", &response, nil)

// DownloadFile - Download a file from a URL
data, err := net.DownloadFile("https://example.com/file.pdf", 30)

// UploadFile - Upload a file to a URL
filePath := "/path/to/file.jpg"
additionalFields := map[string]string{"description": "Profile picture"}
headers := map[string]string{"Authorization": "Bearer token"}
response, err := net.UploadFile("https://api.example.com/upload", "file", filePath, additionalFields, headers)
```

## License

MIT License

    gFly Dev
    https://github.com/gflydev
    Copyright © 2023, JiveCode

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
