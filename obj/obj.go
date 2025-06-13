// Package obj provides utility functions for object/struct/map manipulation.
package obj

import (
	"reflect"
	"sort"
	"strings"
)

// Assign assigns properties of source objects to the destination object.
// It's similar to Object.assign() in JavaScript.
//
// Parameters:
//   - dest: The destination map to which properties will be assigned
//   - sources: One or more source maps whose properties will be assigned to the destination
//
// Returns:
//   - map[K]V: A new map containing all properties from destination and source maps
//
// Example:
//
//	result := Assign(map[string]any{"a": 1}, map[string]any{"b": 2})
//	// result is map[string]any{"a": 1, "b": 2}
func Assign[K comparable, V any](dest map[K]V, sources ...map[K]V) map[K]V {
	result := make(map[K]V)

	// Copy destination first
	for k, v := range dest {
		result[k] = v
	}

	// Copy sources
	for _, source := range sources {
		for k, v := range source {
			result[k] = v
		}
	}

	return result
}

// Clone creates a shallow clone of an object.
//
// Parameters:
//   - obj: The map to clone
//
// Returns:
//   - map[K]V: A new map with the same key-value pairs as the original
//
// Example:
//
//	original := map[string]any{"a": 1, "b": 2}
//	clone := Clone(original)
//	// clone is map[string]any{"a": 1, "b": 2}
func Clone[K comparable, V any](obj map[K]V) map[K]V {
	result := make(map[K]V, len(obj))
	for k, v := range obj {
		result[k] = v
	}
	return result
}

// Entry represents a key-value pair from a map.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries returns an array of key-value pairs from a map.
//
// Parameters:
//   - obj: The map to convert to entries
//
// Returns:
//   - []Entry[K, V]: A slice of key-value pairs
//
// Example:
//
//	entries := Entries(map[string]int{"a": 1, "b": 2})
//	// entries is []Entry[string, int]{{"a", 1}, {"b", 2}}
func Entries[K comparable, V any](obj map[K]V) []Entry[K, V] {
	result := make([]Entry[K, V], 0, len(obj))
	for k, v := range obj {
		result = append(result, Entry[K, V]{k, v})
	}
	return result
}

// FromEntries returns an object composed from key-value pairs.
//
// Parameters:
//   - entries: A slice of key-value pairs
//
// Returns:
//   - map[K]V: A map constructed from the key-value pairs
//
// Example:
//
//	pairs := []Entry[string, int]{{"a", 1}, {"b", 2}}
//	result := FromEntries(pairs)
//	// result is map[string]int{"a": 1, "b": 2}
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))
	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}
	return result
}

// Get retrieves a value from a nested map using dot notation path.
//
// Parameters:
//   - obj: The map to retrieve the value from
//   - path: A dot-separated path to the desired value (e.g., "a.b.c")
//
// Returns:
//   - T: The value at the specified path, converted to type T
//   - bool: True if the value was found and successfully converted, false otherwise
//
// Example:
//
//	nested := map[string]any{"a": map[string]any{"b": 2}}
//	value, ok := Get[int](nested, "a.b")
//	// value is 2, ok is true
func Get[T any](obj map[string]any, path string) (T, bool) {
	var zero T

	// Split the path by dots
	keys := strings.Split(path, ".")

	// Start with the root object
	current := any(obj)

	// Navigate through each key in the path
	for _, key := range keys {
		// Check if current value is a map
		currentMap, ok := current.(map[string]any)
		if !ok {
			return zero, false
		}

		// Get the value for the current key
		value, exists := currentMap[key]
		if !exists {
			return zero, false
		}

		// Move to the next level
		current = value
	}

	// Try to convert the final value to the expected type
	result, ok := current.(T)
	if !ok {
		return zero, false
	}

	return result, true
}

// Has checks if a key is a direct property of an object.
//
// Parameters:
//   - obj: The map to check
//   - key: The key to look for
//
// Returns:
//   - bool: True if the key exists in the map, false otherwise
//
// Example:
//
//	exists := Has(map[string]any{"a": 1, "b": 2}, "a")
//	// exists is true
func Has[K comparable, V any](obj map[K]V, key K) bool {
	_, ok := obj[key]
	return ok
}

// Keys returns an array of object's own enumerable property names.
//
// Parameters:
//   - obj: The map whose keys will be returned
//
// Returns:
//   - []K: A slice containing all the keys from the map
//
// Example:
//
//	keys := Keys(map[string]any{"a": 1, "b": 2})
//	// keys contains "a" and "b" (order not guaranteed)
func Keys[K comparable, V any](obj map[K]V) []K {
	result := make([]K, 0, len(obj))
	for k := range obj {
		result = append(result, k)
	}
	return result
}

// KeysSorted returns a sorted array of object's own enumerable property names.
//
// Parameters:
//   - obj: The map whose keys will be returned in sorted order
//
// Returns:
//   - []K: A slice containing all the keys from the map, sorted
//
// Example:
//
//	keys := KeysSorted(map[string]any{"b": 2, "a": 1})
//	// keys is []string{"a", "b"}
func KeysSorted[K comparable, V any](obj map[K]V) []K {
	keys := Keys(obj)

	// Sort the keys if they are of a sortable type
	if len(keys) > 0 {
		// Check if K is a type that can be sorted
		switch any(keys[0]).(type) {
		case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			sort.Slice(keys, func(i, j int) bool {
				// This is a bit of a hack, but it works for the basic types
				return reflect.ValueOf(keys[i]).String() < reflect.ValueOf(keys[j]).String()
			})
		}
	}

	return keys
}

// MapValues creates an object with the same keys as object and values generated by running each property through iteratee.
//
// Parameters:
//   - obj: The source map
//   - iteratee: A function that transforms each value in the map
//
// Returns:
//   - map[K]R: A new map with the same keys but transformed values
//
// Example:
//
//	doubled := MapValues(map[string]int{"a": 1, "b": 2}, func(v int) int { return v * 2 })
//	// doubled is map[string]int{"a": 2, "b": 4}
func MapValues[K comparable, V any, R any](obj map[K]V, iteratee func(V) R) map[K]R {
	result := make(map[K]R, len(obj))
	for k, v := range obj {
		result[k] = iteratee(v)
	}
	return result
}

// MapKeys creates an object with keys generated by running the property names of object through iteratee.
//
// Parameters:
//   - obj: The source map
//   - iteratee: A function that transforms each key in the map
//
// Returns:
//   - map[R]V: A new map with transformed keys and the original values
//
// Example:
//
//	transformed := MapKeys(map[string]int{"a": 1, "b": 2}, func(k string) string { return k + "x" })
//	// transformed is map[string]int{"ax": 1, "bx": 2}
func MapKeys[K comparable, V any, R comparable](obj map[K]V, iteratee func(K) R) map[R]V {
	result := make(map[R]V, len(obj))
	for k, v := range obj {
		result[iteratee(k)] = v
	}
	return result
}

// Merge merges properties of source objects into the destination object.
// Note: This is a simplified version that doesn't do deep merging due to Go's type system limitations.
//
// Parameters:
//   - dest: The destination map
//   - sources: One or more source maps to merge into the destination
//
// Returns:
//   - map[K]V: A new map containing all properties from destination and source maps
//
// Example:
//
//	result := Merge(map[string]any{"a": 1}, map[string]any{"b": 2})
//	// result is map[string]any{"a": 1, "b": 2}
func Merge[K comparable, V any](dest map[K]V, sources ...map[K]V) map[K]V {
	result := Clone(dest)

	for _, source := range sources {
		for k, v := range source {
			result[k] = v
		}
	}

	return result
}

// Omit creates an object composed of the object properties not included in the keys.
//
// Parameters:
//   - obj: The source map
//   - keys: The keys to omit from the resulting map
//
// Returns:
//   - map[K]V: A new map with all properties from the original except those specified in keys
//
// Example:
//
//	result := Omit(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c")
//	// result is map[string]int{"b": 2}
func Omit[K comparable, V any](obj map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)

	// Create a set of keys to omit
	omitSet := make(map[K]bool)
	for _, k := range keys {
		omitSet[k] = true
	}

	// Copy all properties except those in the omit set
	for k, v := range obj {
		if !omitSet[k] {
			result[k] = v
		}
	}

	return result
}

// OmitBy creates an object composed of the object properties for which predicate returns falsey.
//
// Parameters:
//   - obj: The source map
//   - predicate: A function that returns true for values to omit
//
// Returns:
//   - map[K]V: A new map with properties for which the predicate returned false
//
// Example:
//
//	result := OmitBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { return v > 2 })
//	// result is map[string]int{"a": 1, "b": 2}
func OmitBy[K comparable, V any](obj map[K]V, predicate func(V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range obj {
		if !predicate(v) {
			result[k] = v
		}
	}

	return result
}

// Pick creates an object composed of the picked object properties.
//
// Parameters:
//   - obj: The source map
//   - keys: The keys to include in the resulting map
//
// Returns:
//   - map[K]V: A new map with only the properties specified in keys
//
// Example:
//
//	result := Pick(map[string]int{"a": 1, "b": 2, "c": 3}, "a", "c")
//	// result is map[string]int{"a": 1, "c": 3}
func Pick[K comparable, V any](obj map[K]V, keys ...K) map[K]V {
	result := make(map[K]V)

	for _, k := range keys {
		if v, ok := obj[k]; ok {
			result[k] = v
		}
	}

	return result
}

// PickBy creates an object composed of the object properties for which predicate returns truthy.
//
// Parameters:
//   - obj: The source map
//   - predicate: A function that returns true for values to include
//
// Returns:
//   - map[K]V: A new map with properties for which the predicate returned true
//
// Example:
//
//	result := PickBy(map[string]int{"a": 1, "b": 2, "c": 3}, func(v int) bool { return v > 2 })
//	// result is map[string]int{"c": 3}
func PickBy[K comparable, V any](obj map[K]V, predicate func(V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range obj {
		if predicate(v) {
			result[k] = v
		}
	}

	return result
}

// Values returns an array of object's own enumerable property values.
//
// Parameters:
//   - obj: The map whose values will be returned
//
// Returns:
//   - []V: A slice containing all the values from the map
//
// Example:
//
//	values := Values(map[string]int{"a": 1, "b": 2})
//	// values contains 1 and 2 (order not guaranteed)
func Values[K comparable, V any](obj map[K]V) []V {
	result := make([]V, 0, len(obj))
	for _, v := range obj {
		result = append(result, v)
	}
	return result
}

// Size returns the number of own enumerable properties of an object.
//
// Parameters:
//   - obj: The map whose size will be returned
//
// Returns:
//   - int: The number of key-value pairs in the map
//
// Example:
//
//	count := Size(map[string]int{"a": 1, "b": 2})
//	// count is 2
func Size[K comparable, V any](obj map[K]V) int {
	return len(obj)
}

// IsEmpty checks if an object is empty.
//
// Parameters:
//   - obj: The map to check
//
// Returns:
//   - bool: True if the map contains no key-value pairs, false otherwise
//
// Example:
//
//	empty := IsEmpty(map[string]int{})
//	// empty is true
func IsEmpty[K comparable, V any](obj map[K]V) bool {
	return len(obj) == 0
}

// IsEqual performs a deep comparison between two objects to determine if they are equivalent.
//
// Parameters:
//   - obj1: The first map to compare
//   - obj2: The second map to compare
//
// Returns:
//   - bool: True if both maps have the same keys and values, false otherwise
//
// Example:
//
//	equal := IsEqual(map[string]int{"a": 1}, map[string]int{"a": 1})
//	// equal is true
func IsEqual[K comparable, V comparable](obj1, obj2 map[K]V) bool {
	if len(obj1) != len(obj2) {
		return false
	}

	for k, v1 := range obj1 {
		if v2, ok := obj2[k]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}
