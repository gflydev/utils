# obj - Object and Map Utility Functions for Go

The `obj` package provides a comprehensive set of utility functions for object/struct/map manipulation in Go. It's inspired by libraries like Lodash for JavaScript and offers a wide range of functions to make working with maps and objects easier and more expressive.

## Installation

```bash
go get github.com/gflydev/utils/obj
```

## Usage

```
import "github.com/gflydev/utils/obj"
```

## Functions

### Assign

Assigns properties of source objects to the destination object.
It's similar to Object.assign() in JavaScript. Properties in the source objects
will overwrite properties with the same key in the destination object.

```
result := obj.Assign(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4})
// result: map[string]int{"a": 1, "b": 3, "c": 4}

result := obj.Assign(map[string]int{}, map[string]int{"a": 1}, map[string]int{"b": 2})
// result: map[string]int{"a": 1, "b": 2}

result := obj.Assign(map[string]int{"a": 1}, []map[string]int{}...)
// result: map[string]int{"a": 1}
```

### Clone

Creates a shallow clone of an object. A shallow clone means that only the top-level
key-value pairs are copied. If the values are reference types (like maps, slices, or pointers),
the clone will reference the same underlying data as the original.

```
original := map[string]int{"a": 1, "b": 2}
clone := obj.Clone(original)
// clone: map[string]int{"a": 1, "b": 2}

// Modifying the clone doesn't affect the original
clone["c"] = 3
// original is still map[string]int{"a": 1, "b": 2}
// clone is now map[string]int{"a": 1, "b": 2, "c": 3}

empty := map[string]int{}
clone := obj.Clone(empty)
// clone: map[string]int{}
```

### Entries

Returns an array of key-value pairs from a map. This function is useful
when you need to iterate over map entries in a specific order or when you need
to serialize the map data.

```
entries := obj.Entries(map[string]int{"a": 1, "b": 2})
// entries: []obj.Entry[string, int]{{"a", 1}, {"b", 2}}

// You can iterate over entries in a specific order
sort.Slice(entries, func(i, j int) bool {
    return entries[i].Key < entries[j].Key
})
for _, entry := range entries {
    fmt.Printf("%s: %d\n", entry.Key, entry.Value)
}

entries := obj.Entries(map[string]int{})
// entries: []obj.Entry[string, int]{}
```

### FromEntries

Returns an object composed from key-value pairs. This function
is the inverse of Entries() and is useful for converting a slice of key-value
pairs back into a map.

```
pairs := []obj.Entry[string, int]{{"a", 1}, {"b", 2}}
result := obj.FromEntries(pairs)
// result: map[string]int{"a": 1, "b": 2}

// Common use case: transform a map, then convert back
original := map[string]int{"a": 1, "b": 2}
entries := obj.Entries(original)
// Transform entries
for i := range entries {
    entries[i].Value *= 2
}
transformed := obj.FromEntries(entries)
// transformed is map[string]int{"a": 2, "b": 4}

result := obj.FromEntries([]obj.Entry[string, int]{})
// result: map[string]int{}
```

### Get

Retrieves a value from a nested map using dot notation path. This function
allows you to safely access deeply nested values in a map structure without
having to check for the existence of each level manually.

Notes:
- If any part of the path doesn't exist, the function returns the zero value of T and false
- If the value exists but can't be converted to type T, the function returns the zero value of T and false
- The function only works with map[string]any as the input map type

```
nested := map[string]any{
    "a": 1,
    "b": map[string]any{
        "c": 2,
        "d": map[string]any{
            "e": 3,
        },
        "f": "hello",
    },
}

// Get an integer value
value, ok := obj.Get[int](nested, "a")
// value: 1, ok: true

// Get a nested integer value
value, ok := obj.Get[int](nested, "b.c")
// value: 2, ok: true

// Get a deeply nested integer value
value, ok := obj.Get[int](nested, "b.d.e")
// value: 3, ok: true

// Get a string value
strValue, ok := obj.Get[string](nested, "b.f")
// strValue: "hello", ok: true

// Path doesn't exist
value, ok := obj.Get[int](nested, "b.d.f")
// value: 0, ok: false

// Type mismatch
value, ok := obj.Get[int](nested, "b.f")
// value: 0, ok: false

// Root path doesn't exist
value, ok := obj.Get[int](nested, "x")
// value: 0, ok: false
```

### Has

Checks if a key is a direct property of an object. This function provides
a type-safe way to check for the existence of a key in a map without having to
use the comma-ok idiom directly.

Notes:
- This function only checks for direct properties, not nested ones
- The function returns false if the map is nil
- The function works with any map type where the key is comparable

```
exists := obj.Has(map[string]int{"a": 1, "b": 2}, "a")
// exists: true

exists := obj.Has(map[string]int{"a": 1, "b": 2}, "c")
// exists: false

exists := obj.Has(map[string]int{}, "a")
// exists: false

// Works with any map type
exists := obj.Has(map[int]string{1: "one", 2: "two"}, 1)
// exists: true

// Safe with nil maps
var nilMap map[string]int
exists := obj.Has(nilMap, "a")
// exists: false
```

### Keys

Returns an array of object's own enumerable property names. This function extracts all the keys from a map and returns them as a slice. The order of the keys in the resulting slice is not guaranteed due to the nature of Go maps.

Notes:
- The order of keys is not guaranteed to be consistent between different runs
- Returns an empty slice for empty maps
- Safe to use with nil maps (returns an empty slice)

```
keys := obj.Keys(map[string]int{"a": 1, "b": 2})
// keys: []string{"a", "b"} (order not guaranteed)

keys := obj.Keys(map[string]int{})
// keys: []string{}

// Works with any map type
keys := obj.Keys(map[int]string{1: "one", 2: "two"})
// keys: []int{1, 2} (order not guaranteed)

// Safe with nil maps
var nilMap map[string]int
keys := obj.Keys(nilMap)
// keys: []string{}
```

### KeysSorted

Returns a sorted array of object's own enumerable property names. This function extracts all the keys from a map, sorts them, and returns them as a slice. The keys are sorted in ascending order based on their string representation.

Notes:
- Only works with keys that can be sorted (strings, numbers, etc.)
- The sorting is based on the string representation of the keys
- Returns an empty slice for empty maps
- Safe to use with nil maps (returns an empty slice)

```
keys := obj.KeysSorted(map[string]int{"c": 1, "a": 2, "b": 3})
// keys: []string{"a", "b", "c"}

keys := obj.KeysSorted(map[string]int{})
// keys: []string{}

// Works with numeric keys
keys := obj.KeysSorted(map[int]string{3: "three", 1: "one", 2: "two"})
// keys: []int{1, 2, 3}

// Safe with nil maps
var nilMap map[string]int
keys := obj.KeysSorted(nilMap)
// keys: []string{}
```

### MapValues

Creates an object with the same keys as object and values generated by running each property through iteratee. This function transforms the values in a map by applying a transformation function to each value, while keeping the keys unchanged.

Notes:
- The resulting map has the same keys as the original map
- The values in the resulting map can be of a different type than the original values
- The transformation function is applied to each value in the map
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
// Transform integers to strings
result := obj.MapValues(map[string]int{"a": 1, "b": 2}, func(v int) string { 
    return string(rune('0' + v)) 
})
// result: map[string]string{"a": "1", "b": "2"}

// Double each value
result := obj.MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { 
    return v * 2 
})
// result: map[string]int{"a": 2, "b": 4}

result := obj.MapValues(map[string]int{}, func(v int) string { 
    return string(rune('0' + v)) 
})
// result: map[string]string{}

// Safe with nil maps
var nilMap map[string]int
result := obj.MapValues(nilMap, func(v int) string { 
    return string(rune('0' + v)) 
})
// result: map[string]string{}
```

### MapKeys

Creates an object with keys generated by running the property names of object through iteratee. This function transforms the keys in a map by applying a transformation function to each key, while keeping the values unchanged.

Notes:
- The resulting map has transformed keys but the same values as the original map
- The keys in the resulting map can be of a different type than the original keys
- The transformation function is applied to each key in the map
- If the transformation function produces duplicate keys, later values will overwrite earlier ones
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
// Transform string keys to their ASCII values
result := obj.MapKeys(map[string]int{"a": 1, "b": 2}, func(s string) int { 
    return int(s[0]) 
})
// result: map[int]int{97: 1, 98: 2} // ASCII values for 'a' and 'b'

// Add a prefix to each key
result := obj.MapKeys(map[string]int{"a": 1, "b": 2}, func(s string) string { 
    return "key_" + s 
})
// result: map[string]int{"key_a": 1, "key_b": 2}

result := obj.MapKeys(map[string]int{}, func(s string) int { 
    return int(s[0]) 
})
// result: map[int]int{}

// Safe with nil maps
var nilMap map[string]int
result := obj.MapKeys(nilMap, func(s string) int { 
    return int(s[0]) 
})
// result: map[int]int{}

// Be careful with duplicate keys
result := obj.MapKeys(map[string]int{"a": 1, "b": 2, "c": 3}, func(s string) string { 
    if s == "a" || s == "b" {
        return "x"
    }
    return s
})
// result: map[string]int{"x": 2, "c": 3} // "a" and "b" map to the same key, "b" overwrites "a"
```

### Merge

Merges properties of source objects into the destination object. This function combines multiple maps into a single map by copying all key-value pairs from the source maps into a clone of the destination map.

Notes:
- This is a simplified version that doesn't do deep merging due to Go's type system limitations
- If a key exists in multiple source maps, the value from the last source map will be used
- The original destination map is not modified; a new map is returned
- Source maps are processed in the order they are provided
- Returns a clone of the destination map if no source maps are provided
- Safe to use with nil maps (treats them as empty maps)

```
// Merge multiple maps
result := obj.Merge(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "c": 4}, map[string]int{"d": 5})
// result: map[string]int{"a": 1, "b": 3, "c": 4, "d": 5}

// Merge into an empty map
result := obj.Merge(map[string]int{}, map[string]int{"a": 1}, map[string]int{"b": 2})
// result: map[string]int{"a": 1, "b": 2}

// Merge with no source maps (returns a clone of destination)
result := obj.Merge(map[string]int{"a": 1}, []map[string]int{}...)
// result: map[string]int{"a": 1}

// Later values overwrite earlier ones for the same key
result := obj.Merge(map[string]int{"a": 1}, map[string]int{"a": 2}, map[string]int{"a": 3})
// result: map[string]int{"a": 3}

// Safe with nil maps
var nilMap map[string]int
result := obj.Merge(nilMap, map[string]int{"a": 1})
// result: map[string]int{"a": 1}

result := obj.Merge(map[string]int{"a": 1}, nilMap)
// result: map[string]int{"a": 1}
```

### Omit

Creates an object composed of the object properties not included in the keys. This function filters a map by excluding specific keys, returning a new map with only the key-value pairs that weren't excluded.

Notes:
- The original map is not modified; a new map is returned
- Keys that don't exist in the original map are ignored
- If all keys are omitted, an empty map is returned
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
result := obj.Omit(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c")
// result: map[string]int{"b": 2}

result := obj.Omit(map[string]int{"a": 1, "b": 2}, "c")
// result: map[string]int{"a": 1, "b": 2}

result := obj.Omit(map[string]int{}, "a")
// result: map[string]int{}

// Omit all keys
result := obj.Omit(map[string]int{"a": 1, "b": 2}, "a", "b")
// result: map[string]int{}

// Safe with nil maps
var nilMap map[string]int
result := obj.Omit(nilMap, "a")
// result: map[string]int{}
```

### OmitBy

Creates an object composed of the object properties for which predicate returns falsey. This function filters a map by excluding key-value pairs based on a predicate function applied to each value, returning a new map with only the pairs where the predicate returns false.

Notes:
- The original map is not modified; a new map is returned
- The predicate function is called once for each value in the map
- Only values for which the predicate returns false are included in the result
- If the predicate returns true for all values, an empty map is returned
- If the predicate returns false for all values, a copy of the original map is returned
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
// Omit values greater than 2
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v > 2 
})
// result: map[string]int{"a": 1, "b": 2}

// Omit values less than 2
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v < 2 
})
// result: map[string]int{"b": 2, "c": 3}

// Omit values equal to 2
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v == 2 
})
// result: map[string]int{"a": 1, "c": 3}

// Omit all values (predicate always returns true)
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return true 
})
// result: map[string]int{}

// Keep all values (predicate always returns false)
result := obj.OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return false 
})
// result: map[string]int{"a": 1, "b": 2, "c": 3}

// Safe with nil maps
var nilMap map[string]int
result := obj.OmitBy(nilMap, func(v int) bool { 
    return v > 0 
})
// result: map[string]int{}
```

### Pick

Creates an object composed of the picked object properties. This function is the opposite of `Omit` - it filters a map by including only specific keys, returning a new map with only the key-value pairs that were explicitly selected.

Notes:
- The original map is not modified; a new map is returned
- Keys that don't exist in the original map are ignored
- If none of the specified keys exist in the original map, an empty map is returned
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
// Pick specific keys
result := obj.Pick(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c")
// result: map[string]int{"a": 1, "c": 3}

// Pick a key that doesn't exist
result := obj.Pick(map[string]int{"a": 1, "b": 2}, "c")
// result: map[string]int{}

// Pick from an empty map
result := obj.Pick(map[string]int{}, "a")
// result: map[string]int{}

// Pick all keys
result := obj.Pick(map[string]int{"a": 1, "b": 2}, "a", "b")
// result: map[string]int{"a": 1, "b": 2}

// Safe with nil maps
var nilMap map[string]int
result := obj.Pick(nilMap, "a")
// result: map[string]int{}
```

### PickBy

Creates an object composed of the object properties for which predicate returns truthy. This function is the opposite of `OmitBy` - it filters a map by including key-value pairs based on a predicate function applied to each value, returning a new map with only the pairs where the predicate returns true.

Notes:
- The original map is not modified; a new map is returned
- The predicate function is called once for each value in the map
- Only values for which the predicate returns true are included in the result
- If the predicate returns false for all values, an empty map is returned
- If the predicate returns true for all values, a copy of the original map is returned
- Returns an empty map for empty maps
- Safe to use with nil maps (returns an empty map)

```
// Pick values greater than 2
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v > 2 
})
// result: map[string]int{"c": 3}

// Pick values less than 2
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v < 2 
})
// result: map[string]int{"a": 1}

// Pick values equal to 2
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return v == 2 
})
// result: map[string]int{"b": 2}

// Pick all values (predicate always returns true)
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return true 
})
// result: map[string]int{"a": 1, "b": 2, "c": 3}

// Pick no values (predicate always returns false)
result := obj.PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { 
    return false 
})
// result: map[string]int{}

// Safe with nil maps
var nilMap map[string]int
result := obj.PickBy(nilMap, func(v int) bool { 
    return v > 0 
})
// result: map[string]int{}
```

### Values

Returns an array of object's own enumerable property values. This function extracts all the values from a map and returns them as a slice.

Notes:
- The order of values is not guaranteed to be consistent between different runs
- Returns an empty slice for empty maps
- Safe to use with nil maps (returns an empty slice)
- Works with any map type

```
values := obj.Values(map[string]int{"a": 1, "b": 2})
// values: []int{1, 2} (order not guaranteed)

values := obj.Values(map[string]int{})
// values: []int{}

// Works with any map type
values := obj.Values(map[int]string{1: "one", 2: "two"})
// values: []string{"one", "two"} (order not guaranteed)

// Safe with nil maps
var nilMap map[string]int
values := obj.Values(nilMap)
// values: []int{}
```

### Size

Returns the number of own enumerable properties of an object. This function returns the count of key-value pairs in a map.

Notes:
- Returns 0 for empty maps
- Safe to use with nil maps (returns 0)
- Works with any map type
- This is equivalent to using the built-in `len()` function but provides a consistent interface with other utility functions

```
count := obj.Size(map[string]int{"a": 1, "b": 2, "c": 3})
// count: 3

count := obj.Size(map[string]int{"a": 1})
// count: 1

count := obj.Size(map[string]int{})
// count: 0

// Works with any map type
count := obj.Size(map[int]string{1: "one", 2: "two"})
// count: 2

// Safe with nil maps
var nilMap map[string]int
count := obj.Size(nilMap)
// count: 0
```

### IsEmpty

Checks if an object is empty. This function determines whether a map contains any key-value pairs.

Notes:
- Returns true for maps with no key-value pairs
- Returns true for nil maps
- Returns false for maps with at least one key-value pair
- Works with any map type
- This is equivalent to checking if `Size(map) == 0` or `len(map) == 0`

```
empty := obj.IsEmpty(map[string]int{})
// empty: true

empty := obj.IsEmpty(map[string]int{"a": 1, "b": 2})
// empty: false

// Works with any map type
empty := obj.IsEmpty(map[int]string{})
// empty: true

// Safe with nil maps
var nilMap map[string]int
empty := obj.IsEmpty(nilMap)
// empty: true
```

### IsEqual

Performs a deep comparison between two objects to determine if they are equivalent. This function checks if two maps have the same keys and values.

Notes:
- Returns true if both maps have the same keys and values
- Returns false if the maps have different sizes
- Returns false if any key in the first map doesn't exist in the second map
- Returns false if any key exists in both maps but has different values
- Returns true if both maps are empty
- Safe to use with nil maps (a nil map is considered equal to an empty map)
- Only works with maps where the value type is comparable
- The comparison is not order-dependent since maps don't guarantee order

```
equal := obj.IsEqual(map[string]string{"a": "1", "b": "2"}, map[string]string{"a": "1", "b": "2"})
// equal: true

equal := obj.IsEqual(map[string]string{"a": "1", "b": "2"}, map[string]string{"a": "1", "b": "3"})
// equal: false

equal := obj.IsEqual(map[string]string{"a": "1", "b": "2"}, map[string]string{"a": "1"})
// equal: false

equal := obj.IsEqual(map[string]string{}, map[string]string{})
// equal: true

// Order doesn't matter
equal := obj.IsEqual(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2, "a": 1})
// equal: true

// Safe with nil maps
var nilMap map[string]int
equal := obj.IsEqual(nilMap, map[string]int{})
// equal: true
```

## License

This package is licensed under the MIT License - see the LICENSE file for details.
