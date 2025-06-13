package obj

import (
	"reflect"
	"testing"
)

func TestAssign(t *testing.T) {
	tests := []struct {
		dest     map[string]int
		sources  []map[string]int
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2},
			[]map[string]int{{"b": 3, "c": 4}},
			map[string]int{"a": 1, "b": 3, "c": 4},
		},
		{
			map[string]int{},
			[]map[string]int{{"a": 1}, {"b": 2}},
			map[string]int{"a": 1, "b": 2},
		},
		{
			map[string]int{"a": 1},
			[]map[string]int{},
			map[string]int{"a": 1},
		},
	}

	for _, test := range tests {
		// Create a copy of dest to avoid modifying the test case
		dest := make(map[string]int)
		for k, v := range test.dest {
			dest[k] = v
		}

		result := Assign(dest, test.sources...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Assign(%v, %v) = %v, expected %v", test.dest, test.sources, result, test.expected)
		}
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected map[string]int
	}{
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}},
		{map[string]int{}, map[string]int{}},
	}

	for _, test := range tests {
		result := Clone(test.obj)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Clone(%v) = %v, expected %v", test.obj, result, test.expected)
		}

		// Ensure the result is a different map
		if len(test.obj) > 0 {
			for k := range result {
				result[k] = 999
				break
			}
			if reflect.DeepEqual(result, test.obj) {
				t.Errorf("Clone(%v) returned the same map, not a clone", test.obj)
			}
		}
	}
}

func TestEntries(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected []Entry[string, int]
	}{
		{
			map[string]int{"a": 1, "b": 2},
			[]Entry[string, int]{{"a", 1}, {"b", 2}},
		},
		{
			map[string]int{},
			[]Entry[string, int]{},
		},
	}

	for _, test := range tests {
		result := Entries(test.obj)

		// Since map iteration order is not guaranteed, we need to check if all entries are present
		if len(result) != len(test.expected) {
			t.Errorf("Entries(%v) returned %d entries, expected %d", test.obj, len(result), len(test.expected))
			continue
		}

		// Check that all expected entries are in the result
		for _, entry := range test.expected {
			found := false
			for _, resultEntry := range result {
				if resultEntry.Key == entry.Key && resultEntry.Value == entry.Value {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Entries(%v) = %v, missing entry %v", test.obj, result, entry)
			}
		}
	}
}

func TestFromEntries(t *testing.T) {
	tests := []struct {
		entries  []Entry[string, int]
		expected map[string]int
	}{
		{
			[]Entry[string, int]{{"a", 1}, {"b", 2}},
			map[string]int{"a": 1, "b": 2},
		},
		{
			[]Entry[string, int]{},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := FromEntries(test.entries)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FromEntries(%v) = %v, expected %v", test.entries, result, test.expected)
		}
	}
}

func TestGet(t *testing.T) {
	obj := map[string]any{
		"a": 1,
		"b": map[string]any{
			"c": 2,
			"d": map[string]any{
				"e": 3,
			},
		},
	}

	tests := []struct {
		path     string
		expected int
		ok       bool
	}{
		{"a", 1, true},
		{"b.c", 2, true},
		{"b.d.e", 3, true},
		{"b.d.f", 0, false},
		{"x", 0, false},
	}

	for _, test := range tests {
		result, ok := Get[int](obj, test.path)
		if ok != test.ok || (ok && result != test.expected) {
			t.Errorf("Get(%v, %q) = (%v, %v), expected (%v, %v)", obj, test.path, result, ok, test.expected, test.ok)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		key      string
		expected bool
	}{
		{map[string]int{"a": 1, "b": 2}, "a", true},
		{map[string]int{"a": 1, "b": 2}, "c", false},
		{map[string]int{}, "a", false},
	}

	for _, test := range tests {
		result := Has(test.obj, test.key)
		if result != test.expected {
			t.Errorf("Has(%v, %q) = %v, expected %v", test.obj, test.key, result, test.expected)
		}
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected []string
	}{
		{map[string]int{"a": 1, "b": 2}, []string{"a", "b"}},
		{map[string]int{}, []string{}},
	}

	for _, test := range tests {
		result := Keys(test.obj)

		// Since map iteration order is not guaranteed, we need to check if all keys are present
		if len(result) != len(test.expected) {
			t.Errorf("Keys(%v) returned %d keys, expected %d", test.obj, len(result), len(test.expected))
			continue
		}

		// Check that all expected keys are in the result
		for _, key := range test.expected {
			found := false
			for _, resultKey := range result {
				if resultKey == key {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Keys(%v) = %v, missing key %q", test.obj, result, key)
			}
		}
	}
}

func TestKeysSorted(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected []string
	}{
		{map[string]int{"c": 1, "a": 2, "b": 3}, []string{"a", "b", "c"}},
		{map[string]int{}, []string{}},
	}

	for _, test := range tests {
		result := KeysSorted(test.obj)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("KeysSorted(%v) = %v, expected %v", test.obj, result, test.expected)
		}
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		iteratee func(int) string
		expected map[string]string
	}{
		{
			map[string]int{"a": 1, "b": 2},
			func(n int) string { return string(rune('0' + n)) },
			map[string]string{"a": "1", "b": "2"},
		},
		{
			map[string]int{},
			func(n int) string { return string(rune('0' + n)) },
			map[string]string{},
		},
	}

	for _, test := range tests {
		result := MapValues(test.obj, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapValues(%v, func) = %v, expected %v", test.obj, result, test.expected)
		}
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		iteratee func(string) int
		expected map[int]int
	}{
		{
			map[string]int{"a": 1, "b": 2},
			func(s string) int { return int(s[0]) },
			map[int]int{97: 1, 98: 2}, // ASCII values for 'a' and 'b'
		},
		{
			map[string]int{},
			func(s string) int { return int(s[0]) },
			map[int]int{},
		},
	}

	for _, test := range tests {
		result := MapKeys(test.obj, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapKeys(%v, func) = %v, expected %v", test.obj, result, test.expected)
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		dest     map[string]int
		sources  []map[string]int
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2},
			[]map[string]int{{"b": 3, "c": 4}, {"d": 5}},
			map[string]int{"a": 1, "b": 3, "c": 4, "d": 5},
		},
		{
			map[string]int{},
			[]map[string]int{{"a": 1}, {"b": 2}},
			map[string]int{"a": 1, "b": 2},
		},
		{
			map[string]int{"a": 1},
			[]map[string]int{},
			map[string]int{"a": 1},
		},
	}

	for _, test := range tests {
		// Create a copy of dest to avoid modifying the test case
		dest := make(map[string]int)
		for k, v := range test.dest {
			dest[k] = v
		}

		result := Merge(dest, test.sources...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Merge(%v, %v) = %v, expected %v", test.dest, test.sources, result, test.expected)
		}
	}
}

func TestOmit(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		keys     []string
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			[]string{"a", "c"},
			map[string]int{"b": 2},
		},
		{
			map[string]int{"a": 1, "b": 2},
			[]string{"c"},
			map[string]int{"a": 1, "b": 2},
		},
		{
			map[string]int{},
			[]string{"a"},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := Omit(test.obj, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Omit(%v, %v) = %v, expected %v", test.obj, test.keys, result, test.expected)
		}
	}
}

func TestPick(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		keys     []string
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			[]string{"a", "c"},
			map[string]int{"a": 1, "c": 3},
		},
		{
			map[string]int{"a": 1, "b": 2},
			[]string{"c"},
			map[string]int{},
		},
		{
			map[string]int{},
			[]string{"a"},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := Pick(test.obj, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Pick(%v, %v) = %v, expected %v", test.obj, test.keys, result, test.expected)
		}
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected []int
	}{
		{map[string]int{"a": 1, "b": 2}, []int{1, 2}},
		{map[string]int{}, []int{}},
	}

	for _, test := range tests {
		result := Values(test.obj)

		// Since map iteration order is not guaranteed, we need to check if all values are present
		if len(result) != len(test.expected) {
			t.Errorf("Values(%v) returned %d values, expected %d", test.obj, len(result), len(test.expected))
			continue
		}

		// Check that all expected values are in the result
		for _, value := range test.expected {
			found := false
			for _, resultValue := range result {
				if resultValue == value {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Values(%v) = %v, missing value %d", test.obj, result, value)
			}
		}
	}
}

func TestValuesSorted(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected []int
	}{
		{map[string]int{"c": 3, "a": 1, "b": 2}, []int{1, 2, 3}},
		{map[string]int{}, []int{}},
	}

	for _, test := range tests {
		result := ValuesSorted(test.obj)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ValuesSorted(%v) = %v, expected %v", test.obj, result, test.expected)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected int
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, 3},
		{map[string]int{"a": 1}, 1},
		{map[string]int{}, 0},
	}

	for _, test := range tests {
		result := Size(test.obj)
		if result != test.expected {
			t.Errorf("Size(%v) = %d, expected %d", test.obj, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		obj      map[string]int
		expected bool
	}{
		{map[string]int{"a": 1, "b": 2}, false},
		{map[string]int{}, true},
	}

	for _, test := range tests {
		result := IsEmpty(test.obj)
		if result != test.expected {
			t.Errorf("IsEmpty(%v) = %v, expected %v", test.obj, result, test.expected)
		}
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		obj1     map[string]string
		obj2     map[string]string
		expected bool
	}{
		{
			map[string]string{"a": "1", "b": "2"},
			map[string]string{"a": "1", "b": "2"},
			true,
		},
		{
			map[string]string{"a": "1", "b": "2"},
			map[string]string{"a": "1", "b": "3"},
			false,
		},
		{
			map[string]string{"a": "1", "b": "2"},
			map[string]string{"a": "1"},
			false,
		},
		{
			map[string]string{},
			map[string]string{},
			true,
		},
	}

	for _, test := range tests {
		result := IsEqual(test.obj1, test.obj2)
		if result != test.expected {
			t.Errorf("IsEqual(%v, %v) = %v, expected %v", test.obj1, test.obj2, result, test.expected)
		}
	}
}
