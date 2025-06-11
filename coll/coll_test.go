package coll

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestCountBy(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int) string
		expected map[string]int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) string {
				if n%2 == 0 {
					return "even"
				}
				return "odd"
			},
			map[string]int{"odd": 2, "even": 2},
		},
		{
			[]int{1, 3, 5},
			func(n int) string {
				if n%2 == 0 {
					return "even"
				}
				return "odd"
			},
			map[string]int{"odd": 3},
		},
		{
			[]int{},
			func(n int) string { return "any" },
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := CountBy(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CountBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			[]int{2, 4, 6},
			func(n int) bool { return n%2 == 0 },
			true,
		},
		{
			[]int{2, 3, 6},
			func(n int) bool { return n%2 == 0 },
			false,
		},
		{
			[]int{},
			func(n int) bool { return n%2 == 0 },
			true,
		},
	}

	for _, test := range tests {
		result := Every(test.input, test.predicate)
		if result != test.expected {
			t.Errorf("Every(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) bool { return n%2 == 0 },
			[]int{2, 4},
		},
		{
			[]int{1, 3, 5},
			func(n int) bool { return n%2 == 0 },
			[]int{},
		},
		{
			[]int{},
			func(n int) bool { return n%2 == 0 },
			[]int{},
		},
	}

	for _, test := range tests {
		result := Filter(test.input, test.predicate)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Filter(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		input      []int
		predicate  func(int) bool
		expected   int
		expectedOk bool
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) bool { return n > 2 },
			3,
			true,
		},
		{
			[]int{1, 2},
			func(n int) bool { return n > 2 },
			0,
			false,
		},
		{
			[]int{},
			func(n int) bool { return n > 2 },
			0,
			false,
		},
	}

	for _, test := range tests {
		result, ok := Find(test.input, test.predicate)
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("Find(%v, func) = (%v, %v), expected (%v, %v)", test.input, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		input      []int
		predicate  func(int) bool
		expected   int
		expectedOk bool
	}{
		{
			[]int{1, 2, 3, 4, 3},
			func(n int) bool { return n > 2 },
			3,
			true,
		},
		{
			[]int{1, 2},
			func(n int) bool { return n > 2 },
			0,
			false,
		},
		{
			[]int{},
			func(n int) bool { return n > 2 },
			0,
			false,
		},
	}

	for _, test := range tests {
		result, ok := FindLast(test.input, test.predicate)
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("FindLast(%v, func) = (%v, %v), expected (%v, %v)", test.input, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, 3, true},
		{[]int{1, 2, 3, 4}, 5, false},
		{[]int{}, 1, false},
	}

	for _, test := range tests {
		result := Includes(test.input, test.value)
		if result != test.expected {
			t.Errorf("Includes(%v, %d) = %v, expected %v", test.input, test.value, result, test.expected)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		result := Size(test.input)
		if result != test.expected {
			t.Errorf("Size(%v) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  bool
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) bool { return n > 3 },
			true,
		},
		{
			[]int{1, 2, 3},
			func(n int) bool { return n > 3 },
			false,
		},
		{
			[]int{},
			func(n int) bool { return n > 3 },
			false,
		},
	}

	for _, test := range tests {
		result := Some(test.input, test.predicate)
		if result != test.expected {
			t.Errorf("Some(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestSortBy(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int) int
		expected []int
	}{
		{
			[]int{3, 1, 4, 2},
			func(n int) int { return n },
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3},
			func(n int) int { return -n },
			[]int{3, 2, 1},
		},
		{
			[]int{},
			func(n int) int { return n },
			[]int{},
		},
	}

	for _, test := range tests {
		result := SortBy(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int) int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			func(n int) int { return n * 2 },
			[]int{2, 4, 6},
		},
		{
			[]int{},
			func(n int) int { return n * 2 },
			[]int{},
		},
	}

	for _, test := range tests {
		result := Map(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Map(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestReject(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) bool { return n%2 == 0 },
			[]int{1, 3},
		},
		{
			[]int{2, 4, 6},
			func(n int) bool { return n%2 == 0 },
			[]int{},
		},
		{
			[]int{},
			func(n int) bool { return n%2 == 0 },
			[]int{},
		},
	}

	for _, test := range tests {
		result := Reject(test.input, test.predicate)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Reject(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		input       []int
		iteratee    func(int, int) int
		accumulator int
		expected    int
	}{
		{
			[]int{1, 2, 3, 4},
			func(acc, n int) int { return acc + n },
			0,
			10,
		},
		{
			[]int{1, 2, 3},
			func(acc, n int) int { return acc * n },
			1,
			6,
		},
		{
			[]int{},
			func(acc, n int) int { return acc + n },
			5,
			5,
		},
	}

	for _, test := range tests {
		result := Reduce(test.input, test.iteratee, test.accumulator)
		if result != test.expected {
			t.Errorf("Reduce(%v, func, %d) = %d, expected %d", test.input, test.accumulator, result, test.expected)
		}
	}
}

func TestGroupBy(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int) string
		expected map[string][]int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) string {
				if n%2 == 0 {
					return "even"
				}
				return "odd"
			},
			map[string][]int{"odd": {1, 3}, "even": {2, 4}},
		},
		{
			[]int{},
			func(n int) string { return "any" },
			map[string][]int{},
		},
	}

	for _, test := range tests {
		result := GroupBy(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("GroupBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestKeyBy(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int) string
		expected map[string]int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) string {
				if n%2 == 0 {
					return "even"
				}
				return "odd"
			},
			map[string]int{"odd": 3, "even": 4},
		},
		{
			[]int{},
			func(n int) string { return "any" },
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := KeyBy(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("KeyBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		input     []int
		predicate func(int) bool
		expected  [][]int
	}{
		{
			[]int{1, 2, 3, 4},
			func(n int) bool { return n%2 == 0 },
			[][]int{{2, 4}, {1, 3}},
		},
		{
			[]int{1, 3, 5},
			func(n int) bool { return n%2 == 0 },
			[][]int{{}, {1, 3, 5}},
		},
		{
			[]int{},
			func(n int) bool { return n%2 == 0 },
			[][]int{{}, {}},
		},
	}

	for _, test := range tests {
		result := Partition(test.input, test.predicate)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Partition(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestReduceRight(t *testing.T) {
	tests := []struct {
		input       []int
		iteratee    func(int, int) int
		accumulator int
		expected    int
	}{
		{
			[]int{1, 2, 3, 4},
			func(acc, n int) int { return acc + n },
			0,
			10,
		},
		{
			[]int{1, 2, 3},
			func(acc, n int) int { return acc * n },
			1,
			6,
		},
		{
			[]int{},
			func(acc, n int) int { return acc + n },
			5,
			5,
		},
	}

	for _, test := range tests {
		result := ReduceRight(test.input, test.iteratee, test.accumulator)
		if result != test.expected {
			t.Errorf("ReduceRight(%v, func, %d) = %d, expected %d", test.input, test.accumulator, result, test.expected)
		}
	}
}

func TestOrderBy(t *testing.T) {
	tests := []struct {
		input     []int
		iteratee  func(int) int
		ascending bool
		expected  []int
	}{
		{
			[]int{3, 1, 4, 2},
			func(n int) int { return n },
			true,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{3, 1, 4, 2},
			func(n int) int { return n },
			false,
			[]int{4, 3, 2, 1},
		},
		{
			[]int{},
			func(n int) int { return n },
			true,
			[]int{},
		},
	}

	for _, test := range tests {
		result := OrderBy(test.input, test.iteratee, test.ascending)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("OrderBy(%v, func, %v) = %v, expected %v", test.input, test.ascending, result, test.expected)
		}
	}
}

func TestForEach(t *testing.T) {
	// Test that ForEach calls the iteratee function for each element
	input := []int{1, 2, 3}
	count := 0
	ForEach(input, func(n int) {
		count++
	})
	if count != len(input) {
		t.Errorf("ForEach(%v, func) called iteratee %d times, expected %d", input, count, len(input))
	}
}

func TestMapWithIndex(t *testing.T) {
	tests := []struct {
		input    []int
		iteratee func(int, int) int
		expected []int
	}{
		{
			[]int{1, 2, 3},
			func(n, i int) int { return n * i },
			[]int{0, 2, 6},
		},
		{
			[]int{},
			func(n, i int) int { return n * i },
			[]int{},
		},
	}

	for _, test := range tests {
		result := MapWithIndex(test.input, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapWithIndex(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestForEachWithIndex(t *testing.T) {
	// Test that ForEachWithIndex calls the iteratee function for each element with the correct index
	input := []int{10, 20, 30}
	indices := []int{}
	values := []int{}

	ForEachWithIndex(input, func(n, i int) {
		indices = append(indices, i)
		values = append(values, n)
	})

	if !reflect.DeepEqual(indices, []int{0, 1, 2}) {
		t.Errorf("ForEachWithIndex(%v, func) called iteratee with indices %v, expected %v", input, indices, []int{0, 1, 2})
	}

	if !reflect.DeepEqual(values, input) {
		t.Errorf("ForEachWithIndex(%v, func) called iteratee with values %v, expected %v", input, values, input)
	}
}

func TestSample(t *testing.T) {
	// Test that Sample returns an element from the collection
	input := []int{1, 2, 3, 4, 5}
	result, ok := Sample(input)

	if !ok {
		t.Errorf("Sample(%v) returned ok=false, expected true", input)
	}

	found := false
	for _, v := range input {
		if v == result {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Sample(%v) = %d, which is not in the input array", input, result)
	}

	// Test with empty collection
	_, ok = Sample([]int{})
	if ok {
		t.Errorf("Sample([]) returned ok=true, expected false")
	}
}

func TestSampleSize(t *testing.T) {
	// Test that SampleSize returns n random elements from the collection
	input := []int{1, 2, 3, 4, 5}
	n := 3
	result := SampleSize(input, n)

	if len(result) != n {
		t.Errorf("SampleSize(%v, %d) returned array of length %d, expected %d", input, n, len(result), n)
	}

	// Check that all elements in result are from the input array
	for _, r := range result {
		found := false
		for _, v := range input {
			if r == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("SampleSize(%v, %d) = %v, contains element %d which is not in the input array", input, n, result, r)
		}
	}

	// Test with n > len(input)
	result = SampleSize(input, 10)
	if len(result) != len(input) {
		t.Errorf("SampleSize(%v, %d) returned array of length %d, expected %d", input, 10, len(result), len(input))
	}

	// Test with empty collection
	result = SampleSize([]int{}, 3)
	if len(result) != 0 {
		t.Errorf("SampleSize([], %d) returned array of length %d, expected 0", 3, len(result))
	}
}

func TestForEachMap(t *testing.T) {
	// Test that ForEachMap calls the iteratee function for each key-value pair
	input := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := []string{}
	values := []int{}

	ForEachMap(input, func(v int, k string) {
		keys = append(keys, k)
		values = append(values, v)
	})

	// Sort keys and values for deterministic comparison
	sort.Strings(keys)
	sort.Ints(values)

	if !reflect.DeepEqual(keys, []string{"a", "b", "c"}) {
		t.Errorf("ForEachMap(%v, func) called iteratee with keys %v, expected %v", input, keys, []string{"a", "b", "c"})
	}

	if !reflect.DeepEqual(values, []int{1, 2, 3}) {
		t.Errorf("ForEachMap(%v, func) called iteratee with values %v, expected %v", input, values, []int{1, 2, 3})
	}
}

func TestMapMap(t *testing.T) {
	tests := []struct {
		input    map[string]int
		iteratee func(int, string) string
		expected []string
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			func(v int, k string) string { return k + ":" + string(rune('0'+v)) },
			[]string{"a:1", "b:2", "c:3"},
		},
		{
			map[string]int{},
			func(v int, k string) string { return k + ":" + string(rune('0'+v)) },
			[]string{},
		},
	}

	for _, test := range tests {
		result := MapMap(test.input, test.iteratee)
		// Sort for deterministic comparison
		sort.Strings(result)
		expected := test.expected
		sort.Strings(expected)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("MapMap(%v, func) = %v, expected %v", test.input, result, expected)
		}
	}
}

func TestFilterMap(t *testing.T) {
	tests := []struct {
		input     map[string]int
		predicate func(int, string) bool
		expected  map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			func(v int, k string) bool { return v%2 == 0 },
			map[string]int{"b": 2, "d": 4},
		},
		{
			map[string]int{"a": 1, "c": 3},
			func(v int, k string) bool { return v%2 == 0 },
			map[string]int{},
		},
		{
			map[string]int{},
			func(v int, k string) bool { return v%2 == 0 },
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := FilterMap(test.input, test.predicate)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FilterMap(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestReduceMap(t *testing.T) {
	tests := []struct {
		input       map[string]int
		iteratee    func(string, int, string) string
		accumulator string
		expected    string
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			func(acc string, v int, k string) string { return acc + k },
			"",
			"abc",
		},
		{
			map[string]int{},
			func(acc string, v int, k string) string { return acc + k },
			"initial",
			"initial",
		},
	}

	for _, test := range tests {
		result := ReduceMap(test.input, test.iteratee, test.accumulator)
		// Since map iteration order is not guaranteed, we need to check if the result contains all the expected characters
		if len(result) != len(test.expected) {
			t.Errorf("ReduceMap(%v, func, %q) = %q, expected string of length %d", test.input, test.accumulator, result, len(test.expected))
		}
		for _, ch := range test.expected {
			if !strings.Contains(result, string(ch)) {
				t.Errorf("ReduceMap(%v, func, %q) = %q, expected to contain %q", test.input, test.accumulator, result, string(ch))
			}
		}
	}
}
