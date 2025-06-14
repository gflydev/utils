# arr - Array and Slice Utility Functions for Go

The `arr` package provides a comprehensive set of utility functions for working with arrays, slices, maps, and sets in Go. It's inspired by libraries like Lodash for JavaScript and offers a wide range of functions to make working with collections easier and more expressive.

## Installation

```bash
go get github.com/gflydev/utils/arr
```

## Usage

```go
import "github.com/gflydev/utils/arr"
```

## Functions

### Basic Array Operations

#### Chunk

Splits an array into chunks of the specified size. If the array can't be divided evenly, the last chunk will contain the remaining elements.

Parameters:
- `array`: The array to split into chunks
- `size`: The size of each chunk

Returns:
- A new array containing chunks of the original array

```go
result := arr.Chunk([]int{1, 2, 3, 4}, 2)
// result: [][]int{{1, 2}, {3, 4}}

result := arr.Chunk([]int{1, 2, 3, 4, 5}, 2)
// result: [][]int{{1, 2}, {3, 4}, {5}}
```

Note: If `size` is less than or equal to 0, an empty array will be returned.

#### Compact

Removes zero values from an array. In Go, we consider zero values (like 0 for integers, "" for strings) as falsey.

Parameters:
- `array`: The array to compact

Returns:
- A new array with all zero values removed

```go
result := arr.Compact([]int{0, 1, 2, 0, 3})
// result: []int{1, 2, 3}

result := arr.Compact([]string{"", "a", "", "b"})
// result: []string{"a", "b"}
```

#### Contains

Checks if a slice contains a specific element.
It returns true if the element is found, false otherwise.

```go
// Check if a number exists in a slice
result := arr.Contains([]int{1, 2, 3, 4}, 3)
// result: true
result := arr.Contains([]int{1, 2, 3, 4}, 5)
// result: false

// Check if a string exists in a slice
result := arr.Contains([]string{"apple", "banana", "orange"}, "banana")
// result: true
result := arr.Contains([]string{"apple", "banana", "orange"}, "grape")
// result: false
```

#### Concat

Concatenates multiple arrays into a single array.

Parameters:
- `arrays`: Variable number of arrays to concatenate

Returns:
- A new array containing all elements from the input arrays

```go
result := arr.Concat([]int{1, 2}, []int{3, 4})
// result: []int{1, 2, 3, 4}

result := arr.Concat([]string{"a", "b"}, []string{"c"}, []string{"d", "e"})
// result: []string{"a", "b", "c", "d", "e"}
```

Note: This function works with any type and preserves the order of elements from the input arrays.

#### Difference

Returns the elements in the first array that are not in the other arrays.

Parameters:
- `array`: The base array to compare against
- `others`: Variable number of arrays to compare with the base array

Returns:
- A new array containing elements that are in the base array but not in any of the other arrays

```go
result := arr.Difference([]int{1, 2, 3}, []int{2, 3, 4})
// result: []int{1}

result := arr.Difference([]string{"a", "b", "c"}, []string{"b"}, []string{"c", "d"})
// result: []string{"a"}
```

Note: This function requires elements to be comparable. It efficiently uses a map to track elements in the other arrays.

#### Drop

Creates a slice with n elements dropped from the beginning.

Parameters:
- `array`: The input array
- `n`: Number of elements to drop from the beginning

Returns:
- A new array with the first n elements removed

```go
result := arr.Drop([]int{1, 2, 3, 4}, 2)
// result: []int{3, 4}

result := arr.Drop([]string{"a", "b", "c"}, 1)
// result: []string{"b", "c"}
```

Note: If n is less than or equal to 0, the original array is returned. If n is greater than or equal to the length of the array, an empty array is returned.

#### DropRight

Creates a slice with n elements dropped from the end.

Parameters:
- `array`: The input array
- `n`: Number of elements to drop from the end

Returns:
- A new array with the last n elements removed

```go
result := arr.DropRight([]int{1, 2, 3, 4}, 2)
// result: []int{1, 2}

result := arr.DropRight([]string{"a", "b", "c"}, 1)
// result: []string{"a", "b"}
```

Note: If n is less than or equal to 0, the original array is returned. If n is greater than or equal to the length of the array, an empty array is returned.

#### Fill

Fills elements of an array with a value from start up to, but not including, end.

Parameters:
- `array`: The input array
- `value`: The value to fill the array with
- `start`: The starting index (inclusive)
- `end`: The ending index (exclusive)

Returns:
- A new array with elements filled with the specified value

```go
result := arr.Fill([]int{1, 2, 3, 4}, 0, 1, 3)
// result: []int{1, 0, 0, 4}

result := arr.Fill([]string{"a", "b", "c", "d"}, "x", 0, 2)
// result: []string{"x", "x", "c", "d"}
```

Note: If `start` is less than 0, it will be set to 0. If `end` is greater than the length of the array, it will be set to the length of the array. If `start` is greater than or equal to `end`, the original array is returned.

#### FindIndex

Returns the index of the first element that satisfies the provided testing function.

Parameters:
- `array`: The input array
- `predicate`: A function that returns true for elements that satisfy the condition

Returns:
- The index of the first element that satisfies the predicate, or -1 if none found

```go
result := arr.FindIndex([]int{1, 2, 3, 4}, func(n int) bool { return n > 2 })
// result: 2 (index of 3)

result := arr.FindIndex([]string{"a", "b", "c"}, func(s string) bool { return s == "b" })
// result: 1
```

#### FindLastIndex

Returns the index of the last element that satisfies the provided testing function.

Parameters:
- `array`: The input array
- `predicate`: A function that returns true for elements that satisfy the condition

Returns:
- The index of the last element that satisfies the predicate, or -1 if none found

```go
result := arr.FindLastIndex([]int{1, 3, 2, 3}, func(n int) bool { return n > 2 })
// result: 3 (index of the last 3)

result := arr.FindLastIndex([]string{"a", "b", "c", "b"}, func(s string) bool { return s == "b" })
// result: 3
```

#### First

Returns the first element of an array.

Parameters:
- `array`: The input array

Returns:
- The first element of the array
- A boolean indicating whether the array is not empty (true if the array is not empty, false otherwise)

```go
result, ok := arr.First([]int{1, 2, 3})
// result: 1, ok: true

result, ok := arr.First([]int{})
// result: 0, ok: false

result, ok := arr.First([]string{"a", "b"})
// result: "a", ok: true
```

Note: If the array is empty, the zero value of the element type and false are returned.

#### FirstOrDefault

Returns the first element of an array or a default value if the array is empty.

```go
result := arr.FirstOrDefault([]int{1, 2, 3}, 0)
// result: 1

result := arr.FirstOrDefault([]int{}, 42)
// result: 42

// Works with any type
result := arr.FirstOrDefault([]string{"apple", "banana"}, "default")
// result: "apple"
result := arr.FirstOrDefault([]string{}, "default")
// result: "default"

// Works with structs
type User struct {
    Name string
}
users := []User{{Name: "Alice"}, {Name: "Bob"}}
defaultUser := User{Name: "Unknown"}
result := arr.FirstOrDefault(users, defaultUser)
// result: {Name: "Alice"}
result := arr.FirstOrDefault([]User{}, defaultUser)
// result: {Name: "Unknown"}
```

#### Flatten

Flattens a nested array into a single-level array.

Parameters:
- `array`: The nested array to flatten

Returns:
- A new array with all nested elements combined into a single level

```go
// Flatten an array of integer arrays
result := arr.Flatten([][]int{{1, 2}, {3, 4}})
// result: []int{1, 2, 3, 4}

// Flatten an array of string arrays
result := arr.Flatten([][]string{{"a", "b"}, {"c"}})
// result: []string{"a", "b", "c"}

// Flatten an empty array
result := arr.Flatten([][]int{})
// result: []int{}

// Flatten an array with empty nested arrays
result := arr.Flatten([][]int{{}, {}, {}})
// result: []int{}
```

Note: This function efficiently pre-allocates memory for the result based on the total length of all nested arrays.

#### Includes

Checks if an array includes a certain value.

Parameters:
- `array`: The array to search in
- `value`: The value to search for

Returns:
- `bool`: True if the value is found in the array, false otherwise

```go
// Check if an integer exists in an array
result := arr.Includes([]int{1, 2, 3}, 2)
// result: true

// Check if a value is not in an array
result := arr.Includes([]int{1, 2, 3}, 4)
// result: false

// Works with strings
result := arr.Includes([]string{"a", "b", "c"}, "b")
// result: true

// Works with any comparable type
result := arr.Includes([]float64{1.1, 2.2, 3.3}, 2.2)
// result: true
```

Note: This function requires elements to be comparable. It performs a linear search through the array.

#### IndexOf

Returns the first index at which a given element can be found in the array.

Parameters:
- `array`: The array to search in
- `value`: The value to search for

Returns:
- `int`: The index of the first occurrence of the value, or -1 if not found

```go
// Find the index of a value in an integer array
result := arr.IndexOf([]int{1, 2, 3, 2}, 2)
// result: 1

// Find the index of a string in a string array
result := arr.IndexOf([]string{"a", "b", "c"}, "c")
// result: 2

// Return -1 when the value is not found
result := arr.IndexOf([]int{1, 2, 3}, 4)
// result: -1
```

Note: This function requires elements to be comparable. It performs a linear search through the array and returns the first matching index.

#### Initial

Gets all but the last element of an array.

Parameters:
- `array`: The input array

Returns:
- `[]T`: A new array containing all elements except the last one

```go
// Basic usage
result := arr.Initial([]int{1, 2, 3})
// result: []int{1, 2}

// With string array
result := arr.Initial([]string{"a", "b", "c"})
// result: []string{"a", "b"}

// With single element array
result := arr.Initial([]int{1})
// result: []int{} (empty array)

// With empty array
result := arr.Initial([]int{})
// result: []int{} (empty array)
```

Note: This function works with any type. If the input array has one or zero elements, an empty array is returned.

#### Intersection

Returns an array of unique values that are included in all given arrays.

Parameters:
- `arrays`: Variable number of arrays to find common elements from

Returns:
- `[]T`: A new array containing elements that exist in all input arrays

```go
// Basic usage with two arrays
result := arr.Intersection([]int{1, 2, 3}, []int{2, 3, 4})
// result: []int{2, 3}

// With multiple arrays
result := arr.Intersection([]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "e"})
// result: []string{"b"}

// With arrays that have no common elements
result := arr.Intersection([]int{1, 2}, []int{3, 4})
// result: []int{} (empty array)

// With a single array (returns unique values)
result := arr.Intersection([]int{1, 2, 2, 3})
// result: []int{1, 2, 3}

// With empty arrays
result := arr.Intersection([]int{}, []int{1, 2})
// result: []int{} (empty array)
```

Note: This function requires elements to be comparable. It efficiently handles multiple arrays and returns unique values that appear in all input arrays. If only one array is provided, it returns the unique values from that array.

#### Join

Joins all elements of an array into a string using a specified separator.

Parameters:
- `array`: The array whose elements will be joined
- `separator`: The string to insert between elements

Returns:
- `string`: A string containing all array elements joined with the separator

```go
// Join an array of integers
result := arr.Join([]int{1, 2, 3}, ", ")
// result: "1, 2, 3"

// Join an array of strings
result := arr.Join([]string{"a", "b", "c"}, "-")
// result: "a-b-c"

// Join an array of booleans
result := arr.Join([]bool{true, false}, " and ")
// result: "true and false"

// Join an empty array
result := arr.Join([]int{}, ",")
// result: ""
```

Note: This function works with any type. Elements are converted to strings using the str.ToString function, which provides proper string representation for various types.

#### Last

Returns the last element of an array.

Parameters:
- `array`: The input array

Returns:
- `T`: The last element of the array
- `bool`: True if the array is not empty, false otherwise

```go
// Basic usage
result, ok := arr.Last([]int{1, 2, 3})
// result: 3, ok: true

// With string array
result, ok := arr.Last([]string{"a", "b", "c"})
// result: "c", ok: true

// With empty array
result, ok := arr.Last([]int{})
// result: 0, ok: false (returns zero value of type T)

// With single element array
result, ok := arr.Last([]int{42})
// result: 42, ok: true
```

Note: This function works with any type. If the array is empty, it returns the zero value of type T and false. Otherwise, it returns the last element and true.

#### LastOrDefault

Returns the last element of an array or a default value if the array is empty.
It safely handles empty arrays by returning the provided default value.

```go
// Get the last element from a non-empty array
result := arr.LastOrDefault([]int{1, 2, 3}, 0)
// result: 3

// Get the default value from an empty array
result := arr.LastOrDefault([]int{}, 42)
// result: 42

// Works with any type
result := arr.LastOrDefault([]string{"apple", "banana"}, "default")
// result: "banana"
result := arr.LastOrDefault([]string{}, "default")
// result: "default"

// Works with structs
type User struct {
    Name string
}
users := []User{{Name: "Alice"}, {Name: "Bob"}}
defaultUser := User{Name: "Unknown"}
result := arr.LastOrDefault(users, defaultUser)
// result: {Name: "Bob"}
result := arr.LastOrDefault([]User{}, defaultUser)
// result: {Name: "Unknown"}
```

#### LastIndexOf

Returns the last index at which a given element can be found in the array.

Parameters:
- `array`: The input array
- `value`: The value to search for

Returns:
- `int`: The index of the last occurrence of the value, or -1 if not found

```go
// Basic usage
result := arr.LastIndexOf([]int{1, 2, 3, 2}, 2)
// result: 3 (index of the last occurrence of 2)

// With string array
result := arr.LastIndexOf([]string{"a", "b", "c", "b"}, "b")
// result: 3 (index of the last occurrence of "b")

// When value is not found
result := arr.LastIndexOf([]int{1, 2, 3}, 4)
// result: -1

// With empty array
result := arr.LastIndexOf([]int{}, 1)
// result: -1
```

Note: This function requires elements to be comparable. It performs a linear search through the array from the end to the beginning and returns the index of the first matching element found (which is the last occurrence in the original array).

#### Nth

Gets the element at index n of an array. If n is negative, it gets the element from the end.

Parameters:
- `array`: The input array
- `n`: The index of the element to retrieve

Returns:
- The element at the specified index
- A boolean indicating if a valid element was found

```go
result, ok := arr.Nth([]int{1, 2, 3}, 1)
// result: 2, ok: true

result, ok := arr.Nth([]int{1, 2, 3}, -1)
// result: 3, ok: true (negative indices count from the end)

result, ok := arr.Nth([]int{1, 2, 3}, 5)
// result: 0, ok: false (index out of bounds)

result, ok := arr.Nth([]int{}, 0)
// result: 0, ok: false (empty array)
```

#### Pull

Removes all given values from an array.

Parameters:
- `array`: The input array
- `values`: Variable number of values to remove from the array

Returns:
- A new array with all specified values removed

```go
result := arr.Pull([]int{1, 2, 3, 1, 2, 3}, 2, 3)
// result: []int{1, 1}

result := arr.Pull([]string{"a", "b", "c", "a"}, "a")
// result: []string{"b", "c"}

result := arr.Pull([]int{1, 2, 3}, 4)
// result: []int{1, 2, 3} (no values removed)
```

Note: If the array or values are empty, the original array is returned.

#### Reverse

Reverses the order of elements in array.

Parameters:
- `slice`: The input array to reverse

Returns:
- A new array with elements in reverse order

```go
result := arr.Reverse([]int{1, 2, 3})
// result: []int{3, 2, 1}

result := arr.Reverse([]string{"a", "b", "c"})
// result: []string{"c", "b", "a"}

result := arr.Reverse([]int{1})
// result: []int{1}
```

#### Shuffle

Returns a new slice with elements in random order.

Parameters:
- `slice`: The input array to shuffle

Returns:
- A new array with elements randomly reordered

```go
result := arr.Shuffle([]int{1, 2, 3, 4, 5})
// result: [3, 1, 5, 2, 4] (elements will vary)

result := arr.Shuffle([]string{"a", "b", "c"})
// result: ["c", "a", "b"] (elements will vary)
```

Note: This function uses the Fisher-Yates shuffle algorithm to randomly reorder elements. It creates a copy of the original array, so the original array is not modified.

#### Random

Returns n random elements from the given slice without replacement.

- If n <= 0 or the slice is empty, an empty slice is returned
- If n >= len(slice), a shuffled copy of the entire slice is returned
- The original slice is not modified
- Uses the Shuffle function internally to randomize the elements

```go
result := arr.Random([]int{1, 2, 3, 4, 5}, 3)
// result: []int{2, 4, 1} (elements will vary)

result := arr.Random([]string{"a", "b", "c", "d"}, 2)
// result: []string{"c", "a"} (elements will vary)

// If n >= len(slice), returns all elements in random order
result := arr.Random([]int{1, 2}, 3)
// result: []int{2, 1} (order will vary)
```

#### RandomChoice

Returns a randomly selected element from an array along with a boolean indicating success.

- If the array is empty, returns the zero value of the type and false
- If the array has only one element, returns that element and true
- Otherwise, returns a randomly selected element and true

```go
result, ok := arr.RandomChoice([]string{"a", "b", "c"})
// result: "b", ok: true (element will vary)

result, ok := arr.RandomChoice([]int{1, 2, 3, 4})
// result: 3, ok: true (element will vary)

result, ok := arr.RandomChoice([]int{})
// result: 0, ok: false
```

#### Slice

Creates a slice of array from start up to, but not including, end.

Parameters:
- `array`: The input array
- `start`: The starting index (inclusive)
- `end`: The ending index (exclusive)

Returns:
- A new array containing elements from start index up to but not including end index

```go
result := arr.Slice([]int{1, 2, 3, 4}, 1, 3)
// result: []int{2, 3}

result := arr.Slice([]string{"a", "b", "c", "d"}, 0, 2)
// result: []string{"a", "b"}

result := arr.Slice([]int{1, 2, 3}, 2, 2)
// result: []int{}
```

Note: The function handles edge cases gracefully. If `start` is negative, it's set to 0. If `end` is greater than the array length, it's set to the array length. If `start` is greater than or equal to `end`, an empty array is returned.

#### SortedIndex

Returns the index at which a value should be inserted into a sorted array to maintain sort order.

```go
result := arr.SortedIndex([]int{1, 3, 5, 7}, 4)
// result: 2

result := arr.SortedIndex([]int{10, 20, 30, 40}, 25)
// result: 2

result := arr.SortedIndex([]float64{1.5, 3.5, 5.5}, 0.5)
// result: 0
```

#### Tail

Returns all but the first element of array.

Parameters:
- `array`: The input array

Returns:
- A new array containing all elements except the first one

```go
result := arr.Tail([]int{1, 2, 3})
// result: []int{2, 3}

result := arr.Tail([]string{"a", "b", "c"})
// result: []string{"b", "c"}

result := arr.Tail([]int{1})
// result: []int{}
```

Note: If the array has 0 or 1 elements, an empty array is returned.

#### Take

Creates a slice of array with n elements taken from the beginning.

Parameters:
- `array`: The input array
- `n`: Number of elements to take from the beginning

Returns:
- A new array with the first n elements

```go
result := arr.Take([]int{1, 2, 3, 4}, 2)
// result: []int{1, 2}

result := arr.Take([]string{"a", "b", "c"}, 1)
// result: []string{"a"}

result := arr.Take([]int{1, 2}, 3)
// result: []int{1, 2}
```

Note: If n is less than or equal to 0, an empty array is returned. If n is greater than or equal to the length of the array, the entire array is returned.

#### TakeRight

Creates a slice of array with n elements taken from the end.

Parameters:
- `array`: The input array
- `n`: Number of elements to take from the end

Returns:
- A new array with the last n elements

```go
result := arr.TakeRight([]int{1, 2, 3, 4}, 2)
// result: []int{3, 4}

result := arr.TakeRight([]string{"a", "b", "c"}, 1)
// result: []string{"c"}

result := arr.TakeRight([]int{1, 2}, 3)
// result: []int{1, 2}
```

Note: If n is less than or equal to 0, an empty array is returned. If n is greater than or equal to the length of the array, the entire array is returned.

#### Union

Creates an array of unique values from all given arrays.

Parameters:
- `arrays`: Variable number of arrays to combine

Returns:
- A new array containing all unique elements from the input arrays

```go
result := arr.Union([]int{1, 2}, []int{2, 3})
// result: []int{1, 2, 3}

result := arr.Union([]string{"a", "b"}, []string{"b", "c"}, []string{"c", "d"})
// result: []string{"a", "b", "c", "d"}

result := arr.Union([]int{1, 1, 2}, []int{2, 2, 3})
// result: []int{1, 2, 3}
```

#### Uniq

Creates an array of unique values.

Parameters:
- `array`: The input array

Returns:
- A new array with duplicate elements removed

```go
result := arr.Uniq([]int{1, 2, 2, 3})
// result: []int{1, 2, 3}

result := arr.Uniq([]string{"a", "b", "a", "c", "b"})
// result: []string{"a", "b", "c"}

result := arr.Uniq([]int{1, 1, 1})
// result: []int{1}
```

#### Unique

Removes duplicate elements from a slice, returning a new slice with only unique elements.

```go
result := arr.Unique([]int{1, 2, 2, 3})
// result: []int{1, 2, 3}

// Works with any comparable type
result := arr.Unique([]string{"apple", "banana", "apple", "orange"})
// result: []string{"apple", "banana", "orange"}

// For structs, all fields must match for it to be considered a duplicate
type User struct {
    ID int
    Name string
}
users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 1, Name: "Alice"}, // Duplicate
    {ID: 3, Name: "Charlie"},
    {ID: 2, Name: "Bob"},   // Duplicate
}
uniqueUsers := arr.Unique(users)
// Returns [{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}, {ID: 3, Name: "Charlie"}]
```

#### Without

Returns an array excluding all given values.

Parameters:
- `array`: The input array
- `values`: Variable number of values to exclude from the array

Returns:
- A new array with all specified values excluded

```go
result := arr.Without([]int{1, 2, 3, 4}, 2, 3)
// result: []int{1, 4}

result := arr.Without([]string{"a", "b", "c"}, "a", "c")
// result: []string{"b"}

result := arr.Without([]int{1, 2, 3}, 4)
// result: []int{1, 2, 3}
```

Note: This function is implemented using the Pull function.

#### Collapse

Collapses a slice of slices into a single slice.
It flattens a two-dimensional slice into a one-dimensional slice.

```go
// Collapse multiple slices into one
Collapse([][]any{
    {1, 2, 3},
    {4, 5},
    {6, 7, 8, 9},
}) // Returns [1, 2, 3, 4, 5, 6, 7, 8, 9]

// Collapse with mixed types
Collapse([][]any{
    {"a", "b"},
    {1, 2},
    {true, false},
}) // Returns ["a", "b", 1, 2, true, false]

// Collapse with empty slices
Collapse([][]any{
    {1, 2},
    {},
    {3, 4},
}) // Returns [1, 2, 3, 4]

// Collapse an empty slice
Collapse([][]any{}) // Returns []
```

#### CrossJoin

Cross joins the given arrays, returning a cartesian product with all possible permutations.
It generates all possible combinations by taking one element from each input array.

```go
// Cross join two arrays
CrossJoin([]int{1, 2}, []int{3, 4})
// Returns [[1, 3], [1, 4], [2, 3], [2, 4]]

// Cross join three arrays
CrossJoin([]string{"a", "b"}, []string{"c", "d"}, []string{"e", "f"})
// Returns [
//   ["a", "c", "e"], ["a", "c", "f"],
//   ["a", "d", "e"], ["a", "d", "f"],
//   ["b", "c", "e"], ["b", "c", "f"],
//   ["b", "d", "e"], ["b", "d", "f"]
// ]

// Cross join with a single array
CrossJoin([]int{1, 2, 3})
// Returns [[1], [2], [3]]

// Cross join with an empty array
CrossJoin([]int{}, []int{1, 2})
// Returns [] (empty result because one array is empty)

// Cross join with no arrays
CrossJoin[int]()
// Returns [] (empty result)
```

#### Zip

Zips multiple arrays together, creating a new array where each element is an array containing elements from the input arrays at the same index.

```go
result := arr.Zip([]int{1, 2}, []int{3, 4})
// result: [][]int{{1, 3}, {2, 4}}

result := arr.Zip([]string{"a", "b"}, []string{"c", "d"}, []string{"e", "f"})
// result: [][]string{{"a", "c", "e"}, {"b", "d", "f"}}

result := arr.Zip([]int{1, 2, 3}, []int{4, 5})
// result: [][]int{{1, 4}, {2, 5}}
```

#### Accessible

Checks if a value is directly accessible as an array, slice, or map.
Returns true if the value is an array, slice, or map; false otherwise.

```go
// Check arrays, slices, and maps
result := arr.Accessible([]int{1, 2, 3}) // Returns true
result := arr.Accessible(map[string]int{"a": 1}) // Returns true
result := arr.Accessible([3]int{1, 2, 3}) // Returns true

// Check other types
result := arr.Accessible(42) // Returns false
result := arr.Accessible("hello") // Returns false
result := arr.Accessible(struct{}{}) // Returns false

// Check nil
result := arr.Accessible(nil) // Returns false

// Check pointers to arrays/slices/maps
arr := []int{1, 2, 3}
result := arr.Accessible(&arr) // Returns false (it's a pointer, not directly accessible)
```

#### Wrap

Ensures a value is contained in a slice. If the value is already a slice or array, it converts it to []any. Otherwise, it creates a new slice containing the value.

```go
// Wrapping a single value
result := arr.Wrap(42)
// result: []any{42}

// Wrapping an existing slice
nums := []int{1, 2, 3}
result := arr.Wrap(nums)
// result: []any{1, 2, 3}

// Handling nil
result := arr.Wrap(nil)
// result: []any{} (empty slice)

// Wrapping a struct
type User struct {
    Name string
}
user := User{Name: "John"}
result := arr.Wrap(user)
// result: []any{User{Name: "John"}}
```

### Functional Programming

#### Filter

Returns a new array with all elements that pass the test implemented by the provided function.

```go
result := arr.Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
// result: []int{2, 4}
```

#### WhereNotNull

Filters an array by removing nil values.

```go
// Filter nil pointers from a slice
type User struct {
    Name string
}
var u1 = &User{Name: "Alice"}
var u2 *User = nil
var u3 = &User{Name: "Bob"}
users := []*User{u1, u2, u3}
result := arr.WhereNotNull(users)
// result: [&User{Name: "Alice"}, &User{Name: "Bob"}]

// Works with any type of nil values
values := []any{1, nil, "hello", nil, true}
result := arr.WhereNotNull(values)
// result: []any{1, "hello", true}
```

#### Find

Returns the first element in the array that satisfies the provided testing function and a boolean indicating whether such an element was found.

Parameters:
- `slice`: The input slice to search in
- `predicate`: A function that returns true for the element to find

Returns:
- The first element for which the predicate returns true
- A boolean indicating whether such an element was found (true if found, false otherwise)

```go
// Find the first even number
result, ok := arr.Find([]int{1, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
// result: 4, ok: true

// Find a string with a specific prefix
result, ok := arr.Find([]string{"apple", "banana", "cherry"}, func(s string) bool {
    return strings.HasPrefix(s, "b")
})
// result: "banana", ok: true

// When no element is found
result, ok := arr.Find([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 })
// result: 0 (zero value), ok: false
```

#### Map

Creates a new array with the results of calling a provided function on every element in the calling array.

Parameters:
- `slice`: The input slice to transform
- `mapFunc`: A function that transforms each element from type T to type R

Returns:
- A new slice containing the transformed elements

```go
// Double each number
result := arr.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
// result: []int{2, 4, 6}

// Convert integers to strings
result := arr.Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("Number %d", n) })
// result: []string{"Number 1", "Number 2", "Number 3"}

// Extract a field from structs
type User struct {
    ID int
    Name string
}
users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 3, Name: "Charlie"},
}
names := arr.Map(users, func(u User) string { return u.Name })
// Returns ["Alice", "Bob", "Charlie"]
```

#### Reduce

Applies a function against an accumulator and each element in the array to reduce it to a single value.

```go
// Sum reducer
result := arr.Reduce([]int{1, 2, 3}, 0, func(acc, item int) int { return acc + item })
// result: 6

// Product reducer
result := arr.Reduce([]int{1, 2, 3}, 1, func(acc, item int) int { return acc * item })
// result: 6
```

#### SortBy

Sorts an array by the results of running each element through the iteratee function.
It returns a new sorted array without modifying the original.

```go
// Sort numbers in ascending order
result := arr.SortBy([]int{1, 3, 2}, func(n int) int { return n })
// result: []int{1, 2, 3}

// Sort strings by length
result := arr.SortBy([]string{"apple", "banana", "kiwi"}, func(s string) int { return len(s) })
// result: []string{"kiwi", "apple", "banana"}

// Sort structs by a specific field
type Person struct { Age int }
people := []Person{{Age: 30}, {Age: 25}, {Age: 40}}
result := arr.SortBy(people, func(p Person) int { return p.Age })
// result: []Person{{Age: 25}, {Age: 30}, {Age: 40}}
```

#### SortedCopy

Returns a sorted copy of the slice without modifying the original.
It uses the provided less function to determine the order.

```go
// Sort integers in ascending order
result := arr.SortedCopy([]int{3, 1, 4, 2}, func(i, j int) bool { return i < j })
// result: []int{1, 2, 3, 4}

// Sort integers in descending order
result := arr.SortedCopy([]int{3, 1, 4, 2}, func(i, j int) bool { return i > j })
// result: []int{4, 3, 2, 1}

// Sort strings by length
result := arr.SortedCopy([]string{"apple", "banana", "kiwi", "orange"}, func(i, j string) bool {
    return len(i) < len(j)
})
// result: []string{"kiwi", "apple", "orange", "banana"}

// Sort structs by a specific field
type Person struct {
    Name string
    Age int
}
people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
    {Name: "Charlie", Age: 35},
}
// Sort by age
result := arr.SortedCopy(people, func(i, j Person) bool { return i.Age < j.Age })
// result: [{Name: "Bob", Age: 25}, {Name: "Alice", Age: 30}, {Name: "Charlie", Age: 35}]
```

#### GroupBy

Groups the elements of an array according to the result of calling the provided function on each element.

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob", 30},
    {"Charlie", 25},
    {"Dave", 30},
    {"Eve", 35},
}

// Group by age
result := arr.GroupBy(people, func(p Person) int { return p.Age })
// result: map[int][]Person{
//    25: {{"Alice", 25}, {"Charlie", 25}},
//    30: {{"Bob", 30}, {"Dave", 30}},
//    35: {{"Eve", 35}},
// }
```

#### KeyBy

Creates a map from an array, using the result of the key function as the map key.
It transforms a slice into a map where each element is indexed by a key derived from the element itself.

```go
// Key numbers by themselves
result := arr.KeyBy([]int{1, 2, 3}, func(n int) int { return n })
// Returns map[1:1 2:2 3:3]

// Key strings by their length
result := arr.KeyBy([]string{"apple", "banana", "kiwi"}, func(s string) int {
    return len(s)
})
// Returns map[5:"apple" 6:"banana" 4:"kiwi"]

// Key structs by a field
type User struct {
    ID int
    Name string
}
users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
    {ID: 3, Name: "Charlie"},
}
userMap := arr.KeyBy(users, func(u User) int { return u.ID })
// Returns map[1:{ID:1 Name:"Alice"} 2:{ID:2 Name:"Bob"} 3:{ID:3 Name:"Charlie"}]

// Note: If multiple elements produce the same key, later elements will overwrite earlier ones
```

### Map Operations

#### MapMerge

Merges multiple maps into a single map.

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
result := arr.MapMerge(m1, m2)
// result: map[string]int{"a": 1, "b": 3, "c": 4}
```

#### Add

Adds a key/value pair to a map if the key doesn't already exist.
It returns a new map without modifying the original.

```go
// Add a new key/value pair
original := map[string]any{"name": "John", "age": 30}
result := arr.Add(original, "city", "New York")
// result = {"name": "John", "age": 30, "city": "New York"}
// original remains unchanged

// Try to add an existing key
original := map[string]any{"name": "John", "age": 30}
result := arr.Add(original, "name", "Jane")
// result = {"name": "John", "age": 30}
// The key "name" already exists, so the value is not changed

// Add to an empty map
empty := map[string]any{}
result := arr.Add(empty, "status", "active")
// result = {"status": "active"}
```

#### MapKeys

Extracts all keys from a map into a slice.

Parameters:
- m: The source map

Returns:
- A slice containing all keys from the map

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
keys := arr.MapKeys(data)
// keys: []string{"a", "b", "c"} (order may vary)
```

#### MapValues

Extracts all values from a map into a slice.

Parameters:
- m: The source map

Returns:
- A slice containing all values from the map

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
values := arr.MapValues(data)
// values: []int{1, 2, 3} (order may vary)
```

#### MapValuesFn

Transforms all values in a map using a mapping function.

Parameters:
- m: The source map
- mapFunc: A function that transforms values of type V to type R

Returns:
- A new map with the same keys but transformed values

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
doubled := arr.MapValuesFn(data, func(v int) int {
    return v * 2
})
// doubled: {"a": 2, "b": 4, "c": 6}

// Converting types
toString := arr.MapValuesFn(data, func(v int) string {
    return fmt.Sprintf("value-%d", v)
})
// toString: {"a": "value-1", "b": "value-2", "c": "value-3"}
```

#### MapFindKey

Finds the first key in a map that corresponds to a specific value.

Parameters:
- m: The source map to search in
- value: The value to search for

Returns:
- The first key that maps to the specified value
- A boolean indicating whether such a key was found

```go
data := map[string]int{"a": 1, "b": 2, "c": 1}
key, found := arr.MapFindKey(data, 1)
// key: "a" (or "c" depending on map iteration order), found: true

key, found = arr.MapFindKey(data, 5)
// key: "" (zero value for string), found: false
```

#### MapToSlice

Converts a map to a slice of key-value pair structs.

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
pairs := arr.MapToSlice(data)
// pairs is a slice of struct{Key string; Value int} with entries like:
// [{Key: "a", Value: 1}, {Key: "b", Value: 2}, {Key: "c", Value: 3}]
// (order may vary due to map iteration)

// You can iterate over the pairs:
for _, pair := range pairs {
    fmt.Printf("Key: %s, Value: %d\n", pair.Key, pair.Value)
}
```

#### MapSliceToMap

Converts a slice of key-value pair structs to a map. This is the inverse operation of MapToSlice.

```go
pairs := []struct {
    Key   string
    Value int
}{
    {Key: "a", Value: 1},
    {Key: "b", Value: 2},
    {Key: "c", Value: 3},
}

data := arr.MapSliceToMap(pairs)
// data: map[string]int{"a": 1, "b": 2, "c": 3}

// If there are duplicate keys, the last one wins:
pairs2 := []struct {
    Key   string
    Value int
}{
    {Key: "a", Value: 1},
    {Key: "a", Value: 10},
}
data2 := arr.MapSliceToMap(pairs2)
// data2: map[string]int{"a": 10}
```

#### MapEqualMaps

Checks if two maps contain exactly the same key-value pairs.

```go
map1 := map[string]int{"a": 1, "b": 2, "c": 3}
map2 := map[string]int{"c": 3, "b": 2, "a": 1} // Same content, different order
map3 := map[string]int{"a": 1, "b": 2}         // Missing a key
map4 := map[string]int{"a": 1, "b": 2, "c": 4} // Different value

arr.MapEqualMaps(map1, map2) // Returns: true
arr.MapEqualMaps(map1, map3) // Returns: false
arr.MapEqualMaps(map1, map4) // Returns: false
```

#### MapDiffMaps

Compares two maps and returns three maps containing the added, removed, and changed entries.

```go
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 2, "c": 4, "d": 5}
added, removed, changed := arr.MapDiffMaps(m1, m2)
// added: map[string]int{"d": 5}
// removed: map[string]int{"a": 1}
// changed: map[string]int{"c": 4}
```

#### Divide

Returns two slices, one containing the keys, and the other containing the values of the original map.
It separates a map into its keys and values while preserving the corresponding order.

```go
// Divide a map into keys and values
keys, values := arr.Divide(map[string]any{
    "name": "John",
    "age": 30,
    "city": "New York",
})
// keys could be ["name", "age", "city"] (order may vary)
// values could be ["John", 30, "New York"] (in the same order as keys)

// Divide an empty map
keys, values := arr.Divide(map[string]any{})
// keys = [] (empty slice)
// values = [] (empty slice)

// Note: The order of keys and values is not guaranteed to be the same across different runs
// due to the non-deterministic iteration order of Go maps
```

#### Dot

Flattens a multi-dimensional map into a single level map with "dot" notation.
It converts nested maps into a flat map where keys are paths to values using dot separators.

```go
// Flatten a nested map
nested := map[string]any{
    "user": map[string]any{
        "name": "John",
        "address": map[string]any{
            "city": "New York",
            "zip": 10001,
        },
    },
    "status": "active",
}

flat := arr.Dot(nested)
// Returns:
// {
//    "user.name": "John",
//    "user.address.city": "New York",
//    "user.address.zip": 10001,
//    "status": "active"
// }

// Flatten an empty map
arr.Dot(map[string]any{}) // Returns an empty map

// Flatten a map with no nested structures
arr.Dot(map[string]any{"a": 1, "b": 2}) // Returns the same map {"a": 1, "b": 2}
```

#### Undot

Converts a flattened map with dot notation keys into a nested map structure.
It's the opposite operation of Dot.

```go
// Convert a flat map to a nested structure
data := map[string]any{
    "user.name": "John",
    "user.address.city": "New York",
    "user.address.zip": "10001",
    "status": "active",
}

result := arr.Undot(data)
// Returns:
// {
//   "user": {
//     "name": "John",
//     "address": {
//       "city": "New York",
//       "zip": "10001"
//     }
//   },
//   "status": "active"
// }

// Handle an empty map
result := arr.Undot(map[string]any{})
// Returns an empty map

// Handle a map with no dot notation
data := map[string]any{"name": "John", "age": 30}
result := arr.Undot(data)
// Returns the same structure: {"name": "John", "age": 30}
```

#### Except

Creates a new map excluding the specified keys from the original map.
It returns a copy of the map without the specified keys.

```go
// Remove specific keys
original := map[string]any{"name": "John", "age": 30, "city": "New York"}
result := arr.Except(original, "age", "city")
// Returns {"name": "John"}

// Remove all keys
original := map[string]any{"a": 1, "b": 2}
result := arr.Except(original, "a", "b")
// Returns {} (empty map)

// No keys to remove
original := map[string]any{"a": 1, "b": 2}
result := arr.Except(original)
// Returns {"a": 1, "b": 2} (unchanged)
```

#### Exists

Checks if a given key exists in a map.
It returns true if the key exists, false otherwise.

```go
// Check if a key exists
user := map[string]any{
    "name": "John",
    "age": 30,
}
result := arr.Exists(user, "name") 
// Returns true
result := arr.Exists(user, "email") 
// Returns false

// Check in an empty map
result := arr.Exists(map[string]any{}, "key") 
// Returns false

// Check with a nil map
var nilMap map[string]any
result := arr.Exists(nilMap, "key") 
// Returns false (safe to use with nil maps)
```

#### Get

Retrieves a value from a map using dot notation for nested keys, with a default value if the key doesn't exist.

```go
// Get values from a nested map
nested := map[string]any{
    "user": map[string]any{
        "name": "John",
        "address": map[string]any{
            "city": "New York",
        },
    },
    "status": "active",
}

result := arr.Get(nested, "user.name", "Unknown")
// Returns "John"

result := arr.Get(nested, "user.address.city", "Unknown")
// Returns "New York"

result := arr.Get(nested, "user.address.country", "USA")
// Returns "USA" (key doesn't exist, so default value is returned)

result := arr.Get(nested, "user.email", nil)
// Returns nil (key doesn't exist, so default value is returned)

// Empty key returns the entire map
result := arr.Get(nested, "", nil)
// Returns the entire nested map
```

#### Has

Determines if all of the specified keys exist in the map using "dot" notation.
It checks if every key in the provided list exists in the map, including nested keys.
Returns true only if all keys exist.

```go
// Check simple keys
data := map[string]any{
    "name": "John",
    "age": 30,
}
result := arr.Has(data, "name")
// Returns true

result := arr.Has(data, "email")
// Returns false

result := arr.Has(data, "name", "age")
// Returns true (both keys exist)

result := arr.Has(data, "name", "email")
// Returns false (not all keys exist)

// Check nested keys
nested := map[string]any{
    "user": map[string]any{
        "name": "John",
        "address": map[string]any{
            "city": "New York",
        },
    },
}

result := arr.Has(nested, "user.name")
// Returns true

result := arr.Has(nested, "user.address.city")
// Returns true

result := arr.Has(nested, "user.address.country")
// Returns false

result := arr.Has(nested, "user.name", "user.address.city")
// Returns true (both keys exist)

// Empty keys list
result := arr.Has(data)
// Returns false (no keys to check)
```

#### HasAny

Determines if any of the specified keys exist in the map using "dot" notation.
It checks if at least one key in the provided list exists in the map, including nested keys.
Returns true if any key exists.

```go
// Check simple keys
data := map[string]any{
    "name": "John",
    "age": 30,
}
result := arr.HasAny(data, "name")
// Returns true

result := arr.HasAny(data, "email")
// Returns false

result := arr.HasAny(data, "name", "email")
// Returns true (at least one key exists)

result := arr.HasAny(data, "email", "phone")
// Returns false (none of the keys exist)

// Check nested keys
nested := map[string]any{
    "user": map[string]any{
        "name": "John",
        "address": map[string]any{
            "city": "New York",
        },
    },
}

result := arr.HasAny(nested, "user.name")
// Returns true

result := arr.HasAny(nested, "user.address.city")
// Returns true

result := arr.HasAny(nested, "user.address.country")
// Returns false

result := arr.HasAny(nested, "user.address.country", "user.name")
// Returns true (at least one key exists)

// Empty keys list
result := arr.HasAny(data)
// Returns false (no keys to check)
```

#### IsAssoc

Determines if a value is an associative array/map (has string keys).
It checks if the value is a map with string keys.

```go
// Check associative arrays (maps with string keys)
result := arr.IsAssoc(map[string]any{"name": "John", "age": 30})
// Returns true

result := arr.IsAssoc(map[string]int{"a": 1, "b": 2})
// Returns true

// Check non-associative arrays
result := arr.IsAssoc([]int{1, 2, 3})
// Returns false (slice, not a map)

result := arr.IsAssoc(map[int]string{1: "a", 2: "b"})
// Returns false (map with non-string keys)

// Check other types
result := arr.IsAssoc(42)
// Returns false

result := arr.IsAssoc("hello")
// Returns false

// Check nil
result := arr.IsAssoc(nil)
// Returns false
```

#### IsList

Determines if a value is a list (array or slice).
It checks if the value is an array or slice type.

```go
// Check list types
result := arr.IsList([]int{1, 2, 3})
// Returns true (slice)

result := arr.IsList([3]string{"a", "b", "c"})
// Returns true (array)

// Check non-list types
result := arr.IsList(map[string]int{"a": 1, "b": 2})
// Returns false (map, not a slice/array)

result := arr.IsList(42)
// Returns false

result := arr.IsList("hello")
// Returns false

// Check nil
result := arr.IsList(nil)
// Returns false

// Check empty slice
result := arr.IsList([]int{})
// Returns true (empty slice is still a list)
```

#### Forget

Removes the given key/value pairs from a map.
It returns a new map with the specified keys removed, without modifying the original map.

```go
// Remove a single key
original := map[string]any{
    "name": "John",
    "age": 30,
    "city": "New York",
}
result := arr.Forget(original, "age")
// Returns {"name": "John", "city": "New York"}

// Remove multiple keys
original := map[string]any{
    "name": "John",
    "age": 30,
    "city": "New York",
    "country": "USA",
}
result := arr.Forget(original, "age", "country")
// Returns {"name": "John", "city": "New York"}

// Remove keys that don't exist
original := map[string]any{"a": 1, "b": 2}
result := arr.Forget(original, "c")
// Returns {"a": 1, "b": 2} (unchanged since key doesn't exist)

// Note: Forget is similar to Except, but with a different parameter order
```

#### Only

Returns a new map containing only the specified keys from the original map.
It creates a filtered copy of the input map with just the requested keys.

```go
// Keep only specific keys
original := map[string]any{
    "name": "John",
    "age": 30,
    "city": "New York",
    "country": "USA",
}
result := arr.Only(original, "name", "city")
// Returns {"name": "John", "city": "New York"}

// Request keys that don't exist
original := map[string]any{"a": 1, "b": 2}
result := arr.Only(original, "a", "c")
// Returns {"a": 1} (only existing keys are included)

// Request no keys
original := map[string]any{"a": 1, "b": 2}
result := arr.Only(original)
// Returns {} (empty map)

// Note: Only is the opposite of Except - it keeps only the specified keys
```

#### Set

Sets a value at a specified key path in a map, supporting nested paths with dot notation.

```go
// Set a value at a specific key
data := map[string]any{"user": map[string]any{"name": "John"}}
result := arr.Set(data, "user.age", 30)
// result: {"user": {"name": "John", "age": 30}}

// Set a value at a non-existent path (creates intermediate maps)
data := map[string]any{}
result := arr.Set(data, "user.profile.verified", true)
// result: {"user": {"profile": {"verified": true}}}

// Set a value at the root level
data := map[string]any{"status": "pending"}
result := arr.Set(data, "status", "active")
// result: {"status": "active"}

// Empty key returns the original map
data := map[string]any{"a": 1}
result := arr.Set(data, "", "value")
// result: {"a": 1}
```

#### SortByKey

Sorts a map by keys in ascending alphabetical order.

```go
// Sort a map by keys
data := map[string]any{"c": 3, "a": 1, "b": 2}
result := arr.SortByKey(data)
// result: {"a": 1, "b": 2, "c": 3}

// Sort an empty map
result := arr.SortByKey(map[string]any{})
// result: {} (empty map)
```

#### SortByKeyDesc

Sorts a map by keys in descending alphabetical order.

```go
// Sort a map by keys in descending order
data := map[string]any{"a": 1, "b": 2, "c": 3}
result := arr.SortByKeyDesc(data)
// result: {"c": 3, "b": 2, "a": 1}

// Sort an empty map
result := arr.SortByKeyDesc(map[string]any{})
// result: {} (empty map)
```

#### SortRecursive

Recursively sorts maps by keys and nested arrays/maps.

```go
// Sort a nested map structure
data := map[string]any{
    "c": 3,
    "a": map[string]any{"z": 26, "x": 24},
    "b": []any{2, 1, 3}
}
result := arr.SortRecursive(data)
// result: {
//   "a": {"x": 24, "z": 26},
//   "b": [2, 1, 3], // Note: array order is preserved
//   "c": 3
// }

// Sort a simple map
data := map[string]any{"c": 3, "a": 1, "b": 2}
result := arr.SortRecursive(data)
// result: {"a": 1, "b": 2, "c": 3}

// Sort a non-map value
result := arr.SortRecursive(42)
// result: 42 (non-map values are returned as is)
```

#### MapFilterMap

Creates a new map by filtering entries from the original map based on a predicate function.

Parameters:
- m: The source map
- predicate: A function that takes a key and value and returns a boolean

Returns:
- A new map containing only the entries for which the predicate returns true

```go
// Filter a map to keep only entries with even values
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
evens := arr.MapFilterMap(data, func(k string, v int) bool {
    return v%2 == 0
})
// evens: {"b": 2, "d": 4}

// Keep only entries where key is "a" or "c"
filtered := arr.MapFilterMap(data, func(k string, v int) bool {
    return k == "a" || k == "c"
})
// filtered: {"a": 1, "c": 3}
```

#### MapInvertMap

Creates a new map by swapping the keys and values of the original map.

Parameters:
- m: The source map to invert

Returns:
- A new map with the keys and values swapped

Notes:
- If multiple keys map to the same value in the original map, only one key-value pair will be in the result
- The last key-value pair processed will be the one that appears in the result

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}
inverted := arr.MapInvertMap(data)
// inverted: {1: "a", 2: "b", 3: "c"}

// With duplicate values
data2 := map[string]int{"a": 1, "b": 2, "c": 1}
inverted2 := arr.MapInvertMap(data2)
// inverted2 might be {1: "c", 2: "b"} or {1: "a", 2: "b"} depending on map iteration order
```

#### MapGetOrDefault

Safely retrieves a value from a map, returning a default value if the key doesn't exist.

Parameters:
- m: The source map
- key: The key to look up
- defaultValue: The value to return if the key doesn't exist

Returns:
- The value associated with the key, or the default value if the key doesn't exist

```go
data := map[string]int{"a": 1, "b": 2, "c": 3}

value := arr.MapGetOrDefault(data, "b", 0)
// value: 2

value = arr.MapGetOrDefault(data, "d", 0)
// value: 0 (default value)
```

#### MapGetOrInsert

Retrieves a value from a map, or inserts a default value if the key doesn't exist.

Parameters:
- m: The source map (will be modified if the key doesn't exist)
- key: The key to look up
- defaultValue: The value to insert and return if the key doesn't exist

Returns:
- The value associated with the key, or the default value if the key didn't exist

```go
data := map[string]int{"a": 1, "b": 2}

// Key exists
value := arr.MapGetOrInsert(data, "b", 0)
// value: 2, data unchanged: {"a": 1, "b": 2}

// Key doesn't exist
value = arr.MapGetOrInsert(data, "c", 3)
// value: 3, data modified: {"a": 1, "b": 2, "c": 3}
```

### Set Operations

#### SetContains

Checks if a set (implemented as map[T]struct{}) contains a specific element.

```go
// Create a set
set := map[string]struct{}{
    "apple":  {},
    "banana": {},
    "cherry": {},
}

result := arr.SetContains(set, "banana")
// result: true

result = arr.SetContains(set, "orange")
// result: false
```

#### SetToSlice

Converts a set (implemented as map[T]struct{}) to a slice.

```go
// Create a set
set := map[string]struct{}{
    "apple":  {},
    "banana": {},
    "cherry": {},
}

slice := arr.SetToSlice(set)
// slice: []string{"apple", "banana", "cherry"} (order may vary)
```

#### SliceToSet

Converts a slice to a set (implemented as map[T]struct{}). This is the inverse operation of SetToSlice.

```go
// Create a slice with duplicate elements
slice := []string{"apple", "banana", "apple", "cherry", "banana"}

set := arr.SliceToSet(slice)
// set: map[string]struct{}{"apple": {}, "banana": {}, "cherry": {}}
// Note: duplicates are automatically removed
```

#### SetUnion

Creates a new set containing all elements from both input sets.

```go
set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}

union := arr.SetUnion(set1, set2)
// union: {"a": {}, "b": {}, "c": {}, "d": {}}
```

#### SetIntersection

Creates a new set containing only elements that exist in both input sets.

Parameters:
- `set1`: The first set
- `set2`: The second set

Returns:
- A new set containing only elements that appear in both set1 and set2

```go
set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}

intersection := arr.SetIntersection(set1, set2)
// intersection: {"b": {}, "c": {}}
```

Note:
- The function optimizes performance by iterating over the smaller set

#### SetDifference

Creates a new set containing elements that are in the first set but not in the second set.

Parameters:
- `set1`: The first set (source set)
- `set2`: The second set (elements to exclude)

Returns:
- A new set containing elements that are in set1 but not in set2

```go
set1 := map[string]struct{}{"a": {}, "b": {}, "c": {}}
set2 := map[string]struct{}{"b": {}, "c": {}, "d": {}}

difference := arr.SetDifference(set1, set2)
// difference: {"a": {}}

// Note that the difference is not symmetric:
difference2 := arr.SetDifference(set2, set1)
// difference2: {"d": {}}
```

#### Pluck

Extracts a specific property from each element in an array and returns an array of those properties.

```go
// Extract a property from a slice of structs
type Person struct {
    Name string
    Age  int
}
people := []Person{
    {Name: "Alice", Age: 25},
    {Name: "Bob", Age: 30},
    {Name: "Charlie", Age: 35},
}
ages := arr.Pluck(people, func(p Person) int {
    return p.Age
})
// Returns []int{25, 30, 35}

// Works with maps too
people := []map[string]any{
    {"name": "Alice", "age": 25},
    {"name": "Bob", "age": 30},
    {"name": "Charlie", "age": 35},
}
names := arr.Pluck(people, func(p map[string]any) string {
    return p["name"].(string)
})
// Returns []string{"Alice", "Bob", "Charlie"}
```

#### Prepend

Adds one or more items to the beginning of a slice.
It returns a new slice with the values added at the beginning, without modifying the original slice.

```go
// Prepend a single value
result := arr.Prepend([]int{2, 3, 4}, 1)
// Returns [1, 2, 3, 4]

// Prepend multiple values
result := arr.Prepend([]int{3, 4, 5}, 1, 2)
// Returns [1, 2, 3, 4, 5]

// Prepend to an empty slice
result := arr.Prepend([]string{}, "hello")
// Returns ["hello"]

// Works with any type
result := arr.Prepend([]string{"world"}, "hello")
// Returns ["hello", "world"]

// Prepend no values (returns a copy of the original)
result := arr.Prepend([]int{1, 2, 3})
// Returns [1, 2, 3]
```

#### Query

Builds a URL query string from a map.
It converts a map into a URL-encoded query string suitable for HTTP requests.

```go
// Simple key-value pairs
result := arr.Query(map[string]any{
    "name": "John Doe",
    "age": 30,
})
// Returns "age=30&name=John+Doe" (order may vary)

// With array values
result := arr.Query(map[string]any{
    "colors": []string{"red", "blue", "green"},
    "id": 123,
})
// Returns "colors%5B%5D=red&colors%5B%5D=blue&colors%5B%5D=green&id=123" (order may vary)
// Decoded: "colors[]=red&colors[]=blue&colors[]=green&id=123"

// With special characters
result := arr.Query(map[string]any{
    "search": "hello world",
    "filter": "price>100",
})
// Returns "filter=price%3E100&search=hello+world" (order may vary)
// Decoded: "filter=price>100&search=hello world"

// Empty map
result := arr.Query(map[string]any{})
// Returns "" (empty string)
```

#### RandomOrDefault

Returns a random value from a slice or a default value if the slice is empty.
It safely handles empty slices by returning the provided default value.

```go
// Get a random element from a non-empty slice
// Note: The actual returned value will vary due to randomness
result := arr.RandomOrDefault([]int{1, 2, 3, 4, 5}, 0)
// Returns one of 1, 2, 3, 4, or 5

// Get the default value from an empty slice
result := arr.RandomOrDefault([]int{}, 0)
// Returns 0
```

## License

This package is licensed under the MIT License - see the LICENSE file for details.
