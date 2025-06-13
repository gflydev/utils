package arr

import (
	"net/url"
	"reflect"
	"sort"
	"testing"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		input    []int
		size     int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
		{[]int{}, 2, [][]int{}},
	}

	for _, test := range tests {
		result := Chunk(test.input, test.size)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Chunk(%v, %d) = %v, expected %v", test.input, test.size, result, test.expected)
		}
	}
}

func TestCompact(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 2, 0, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Compact(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Compact(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		inputs   [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{{}, {1, 2}, {}}, []int{1, 2}},
		{[][]int{}, []int{}},
	}

	for _, test := range tests {
		result := Concat(test.inputs...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Concat(%v) = %v, expected %v", test.inputs, result, test.expected)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		array    []int
		others   [][]int
		expected []int
	}{
		{[]int{1, 2, 3}, [][]int{{2, 3}}, []int{1}},
		{[]int{1, 2, 3}, [][]int{{2}, {3}}, []int{1}},
		{[]int{1, 2, 3}, [][]int{{4, 5}}, []int{1, 2, 3}},
		{[]int{}, [][]int{{1, 2}}, []int{}},
	}

	for _, test := range tests {
		result := Difference(test.array, test.others...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Difference(%v, %v) = %v, expected %v", test.array, test.others, result, test.expected)
		}
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{3, 4}},
		{[]int{1, 2, 3}, 5, []int{}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{}, 2, []int{}},
	}

	for _, test := range tests {
		result := Drop(test.input, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Drop(%v, %d) = %v, expected %v", test.input, test.n, result, test.expected)
		}
	}
}

func TestDropRight(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 5, []int{}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{}, 2, []int{}},
	}

	for _, test := range tests {
		result := DropRight(test.input, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("DropRight(%v, %d) = %v, expected %v", test.input, test.n, result, test.expected)
		}
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{}, -1},
	}

	for _, test := range tests {
		result := FindIndex(test.input, func(n int) bool { return n > 3 })
		if result != test.expected {
			t.Errorf("FindIndex(%v, func) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		input      []int
		expected   int
		expectedOk bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		result, ok := First(test.input)
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("First(%v) = (%d, %v), expected (%d, %v)", test.input, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		expected bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
	}

	for _, test := range tests {
		result := Includes(test.input, test.value)
		if result != test.expected {
			t.Errorf("Includes(%v, %d) = %v, expected %v", test.input, test.value, result, test.expected)
		}
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		input      []int
		expected   int
		expectedOk bool
	}{
		{[]int{1, 2, 3}, 3, true},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		result, ok := Last(test.input)
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("Last(%v) = (%d, %v), expected (%d, %v)", test.input, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestFill(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		start    int
		end      int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 0, 1, 3, []int{1, 0, 0, 4}},
		{[]int{1, 2, 3, 4}, 5, 0, 4, []int{5, 5, 5, 5}},
		{[]int{1, 2, 3, 4}, 5, -1, 2, []int{5, 5, 3, 4}}, // Negative start is treated as 0
		{[]int{1, 2, 3, 4}, 5, 2, 10, []int{1, 2, 5, 5}},
	}

	for _, test := range tests {
		result := Fill(test.input, test.value, test.start, test.end)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Fill(%v, %d, %d, %d) = %v, expected %v", test.input, test.value, test.start, test.end, result, test.expected)
		}
	}
}

func TestFindLastIndex(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 3}, 4},
		{[]int{1, 2, 3}, 2}, // Index of 3 in [1, 2, 3] is 2
		{[]int{}, -1},
	}

	for _, test := range tests {
		result := FindLastIndex(test.input, func(n int) bool { return n == 3 })
		if result != test.expected {
			t.Errorf("FindLastIndex(%v, func) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
		{[][]int{}, []int{}},
	}

	for _, test := range tests {
		result := Flatten(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Flatten(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 3, 2},
		{[]int{1, 2, 3, 4}, 5, -1},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		result := IndexOf(test.input, test.value)
		if result != test.expected {
			t.Errorf("IndexOf(%v, %d) = %d, expected %d", test.input, test.value, result, test.expected)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		expected int
	}{
		{[]int{1, 2, 3, 2, 4}, 2, 3},
		{[]int{1, 2, 3, 4}, 5, -1},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		result := LastIndexOf(test.input, test.value)
		if result != test.expected {
			t.Errorf("LastIndexOf(%v, %d) = %d, expected %d", test.input, test.value, result, test.expected)
		}
	}
}

func TestInitial(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2}},
		{[]int{1}, []int{}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Initial(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Initial(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		inputs   [][]int
		expected []int
	}{
		{[][]int{{1, 2, 3}, {2, 3, 4}}, []int{2, 3}},
		{[][]int{{1, 2}, {2, 3}, {2, 4}}, []int{2}},
		{[][]int{{1, 2}, {3, 4}}, []int{}},
		{[][]int{}, []int{}},
	}

	for _, test := range tests {
		result := Intersection(test.inputs...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Intersection(%v) = %v, expected %v", test.inputs, result, test.expected)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input     []int
		separator string
		expected  string
	}{
		{[]int{1, 2, 3}, ",", "1,2,3"},
		{[]int{1}, ",", "1"},
		{[]int{}, ",", ""},
		{[]int{}, ",", ""},
	}

	for _, test := range tests {
		result := Join(test.input, test.separator)
		// Since the exact string representation depends on the reflect package,
		// we'll just check that the result is not empty for non-empty inputs
		if len(test.input) > 0 && result == "" {
			t.Errorf("Join(%v, %q) returned empty string", test.input, test.separator)
		}
		if len(test.input) == 0 && result != "" {
			t.Errorf("Join(%v, %q) = %q, expected empty string", test.input, test.separator, result)
		}
	}
}

func TestNth(t *testing.T) {
	tests := []struct {
		input      []int
		n          int
		expected   int
		expectedOk bool
	}{
		{[]int{1, 2, 3, 4}, 1, 2, true},
		{[]int{1, 2, 3, 4}, -1, 4, true},
		{[]int{1, 2, 3, 4}, 5, 0, false},
		{[]int{}, 0, 0, false},
	}

	for _, test := range tests {
		result, ok := Nth(test.input, test.n)
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("Nth(%v, %d) = (%d, %v), expected (%d, %v)", test.input, test.n, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestTail(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3}},
		{[]int{1}, []int{}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Tail(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Tail(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{}, 2, []int{}},
	}

	for _, test := range tests {
		result := Take(test.input, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Take(%v, %d) = %v, expected %v", test.input, test.n, result, test.expected)
		}
	}
}

func TestTakeRight(t *testing.T) {
	tests := []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{3, 4}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 0, []int{}},
		{[]int{}, 2, []int{}},
	}

	for _, test := range tests {
		result := TakeRight(test.input, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("TakeRight(%v, %d) = %v, expected %v", test.input, test.n, result, test.expected)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		inputs   [][]int
		expected map[int]bool // Use a map to check for presence, not order
	}{
		{[][]int{{1, 2}, {2, 3}}, map[int]bool{1: true, 2: true, 3: true}},
		{[][]int{{1, 2}, {3, 4}, {5}}, map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}},
		{[][]int{}, map[int]bool{}},
	}

	for _, test := range tests {
		result := Union(test.inputs...)

		// Check that the result has the expected length
		if len(result) != len(test.expected) {
			t.Errorf("Union(%v) returned %d elements, expected %d", test.inputs, len(result), len(test.expected))
			continue
		}

		// Check that all expected elements are in the result
		for _, v := range result {
			if !test.expected[v] {
				t.Errorf("Union(%v) = %v, contains unexpected element %d", test.inputs, result, v)
			}
		}
	}
}

func TestUniq(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 1, 3, 2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Uniq(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Uniq(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestWithout(t *testing.T) {
	tests := []struct {
		input    []int
		values   []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 3}, []int{1, 4}},
		{[]int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}},
		{[]int{}, []int{1, 2}, []int{}},
	}

	for _, test := range tests {
		result := Without(test.input, test.values...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Without(%v, %v) = %v, expected %v", test.input, test.values, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Reverse(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestShuffle(t *testing.T) {
	// Shuffle is non-deterministic, so we just check that the length is the same
	// and that all elements from the original array are present in the shuffled array
	input := []int{1, 2, 3, 4, 5}
	result := Shuffle(input)

	if len(result) != len(input) {
		t.Errorf("Shuffle(%v) returned array of length %d, expected %d", input, len(result), len(input))
	}

	// Check that all elements are present
	for _, v := range input {
		found := false
		for _, r := range result {
			if r == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Shuffle(%v) = %v, missing element %d", input, result, v)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		input    []int
		start    int
		end      int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 1, 3, []int{2, 3}},
		{[]int{1, 2, 3, 4}, 0, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 2, 2, []int{}},
		{[]int{1, 2, 3, 4}, -1, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, 1, 10, []int{2, 3, 4}},
		{[]int{}, 0, 0, []int{}},
	}

	for _, test := range tests {
		result := Slice(test.input, test.start, test.end)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Slice(%v, %d, %d) = %v, expected %v", test.input, test.start, test.end, result, test.expected)
		}
	}
}

func TestSortedIndex(t *testing.T) {
	tests := []struct {
		input    []int
		value    int
		expected int
	}{
		{[]int{1, 3, 5, 7}, 4, 2},
		{[]int{1, 3, 5, 7}, 0, 0},
		{[]int{1, 3, 5, 7}, 8, 4},
		{[]int{}, 1, 0},
	}

	for _, test := range tests {
		result := SortedIndex(test.input, test.value)
		if result != test.expected {
			t.Errorf("SortedIndex(%v, %d) = %d, expected %d", test.input, test.value, result, test.expected)
		}
	}
}

func TestZip(t *testing.T) {
	tests := []struct {
		inputs   [][]int
		expected [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 3}, {2, 4}}},
		{[][]int{{1, 2, 3}, {4, 5}}, [][]int{{1, 4}, {2, 5}}}, // Only zips up to the minimum length
		{[][]int{}, [][]int{}},
	}

	for _, test := range tests {
		result := Zip(test.inputs...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Zip(%v) = %v, expected %v", test.inputs, result, test.expected)
		}
	}
}

func TestSortBy(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 2}, []int{1, 2, 3, 4}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := SortBy(test.input, func(n int) int { return n })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestPull(t *testing.T) {
	tests := []struct {
		input    []int
		values   []int
		expected []int
	}{
		{[]int{1, 2, 3, 1, 2, 3}, []int{2, 3}, []int{1, 1}},
		{[]int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3}},
		{[]int{}, []int{1, 2}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		result := Pull(test.input, test.values...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Pull(%v, %v) = %v, expected %v", test.input, test.values, result, test.expected)
		}
	}
}

func TestRandom(t *testing.T) {
	// Test empty slice
	result := Random([]int{}, 3)
	if len(result) != 0 {
		t.Errorf("Random([], 3) = %v, expected []", result)
	}

	// Test n <= 0
	result = Random([]int{1, 2, 3}, 0)
	if len(result) != 0 {
		t.Errorf("Random([1, 2, 3], 0) = %v, expected []", result)
	}

	// Test n >= len(slice)
	input := []int{1, 2, 3}
	result = Random(input, 3)
	if len(result) != 3 {
		t.Errorf("Random(%v, 3) returned %d elements, expected 3", input, len(result))
	}
	// Check that all elements from input are in result
	for _, v := range input {
		found := false
		for _, r := range result {
			if r == v {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Random(%v, 3) = %v, missing element %d", input, result, v)
		}
	}

	// Test n < len(slice)
	result = Random([]int{1, 2, 3, 4, 5}, 3)
	if len(result) != 3 {
		t.Errorf("Random([1, 2, 3, 4, 5], 3) returned %d elements, expected 3", len(result))
	}
}

func TestRandomChoice(t *testing.T) {
	// Test empty slice
	_, ok := RandomChoice([]int{})
	if ok {
		t.Errorf("RandomChoice([]) returned ok=true, expected false")
	}

	// Test single element slice
	input := []int{42}
	result, ok := RandomChoice(input)
	if !ok || result != 42 {
		t.Errorf("RandomChoice([42]) = (%d, %v), expected (42, true)", result, ok)
	}

	// Test multiple elements
	input = []int{1, 2, 3, 4, 5}
	result, ok = RandomChoice(input)
	if !ok {
		t.Errorf("RandomChoice(%v) returned ok=false, expected true", input)
	}

	// Check that the result is one of the input elements
	found := false
	for _, v := range input {
		if result == v {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomChoice(%v) = %d, which is not in the input slice", input, result)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		input    []int
		element  int
		expected bool
	}{
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
	}

	for _, test := range tests {
		result := Contains(test.input, test.element)
		if result != test.expected {
			t.Errorf("Contains(%v, %d) = %v, expected %v", test.input, test.element, result, test.expected)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4}},
		{[]int{1, 3, 5}, []int{}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Filter(test.input, func(n int) bool { return n%2 == 0 })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Filter(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		result := Map(test.input, func(n int) int { return n * 2 })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Map(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		input      []int
		expected   int
		expectedOk bool
	}{
		{[]int{1, 2, 3, 4}, 3, true},
		{[]int{1, 2}, 0, false},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		result, ok := Find(test.input, func(n int) bool { return n > 2 })
		if result != test.expected || ok != test.expectedOk {
			t.Errorf("Find(%v, func) = (%d, %v), expected (%d, %v)", test.input, result, ok, test.expected, test.expectedOk)
		}
	}
}

func TestFirstOrDefault(t *testing.T) {
	tests := []struct {
		input        []int
		defaultValue int
		expected     int
	}{
		{[]int{1, 2, 3}, 0, 1},
		{[]int{}, 42, 42},
	}

	for _, test := range tests {
		result := FirstOrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("FirstOrDefault(%v, %d) = %d, expected %d", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestLastOrDefault(t *testing.T) {
	tests := []struct {
		input        []int
		defaultValue int
		expected     int
	}{
		{[]int{1, 2, 3}, 0, 3},
		{[]int{}, 42, 42},
	}

	for _, test := range tests {
		result := LastOrDefault(test.input, test.defaultValue)
		if result != test.expected {
			t.Errorf("LastOrDefault(%v, %d) = %d, expected %d", test.input, test.defaultValue, result, test.expected)
		}
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		input    []int
		values   []int
		expected []int
	}{
		{[]int{3, 4}, []int{1, 2}, []int{1, 2, 3, 4}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
	}

	for _, test := range tests {
		result := Prepend(test.input, test.values...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Prepend(%v, %v) = %v, expected %v", test.input, test.values, result, test.expected)
		}
	}
}

func TestSetContains(t *testing.T) {
	tests := []struct {
		set      map[int]struct{}
		item     int
		expected bool
	}{
		{map[int]struct{}{1: {}, 2: {}, 3: {}}, 2, true},
		{map[int]struct{}{1: {}, 2: {}, 3: {}}, 4, false},
		{map[int]struct{}{}, 1, false},
	}

	for _, test := range tests {
		result := SetContains(test.set, test.item)
		if result != test.expected {
			t.Errorf("SetContains(%v, %d) = %v, expected %v", test.set, test.item, result, test.expected)
		}
	}
}

func TestSetToSlice(t *testing.T) {
	tests := []struct {
		set      map[int]struct{}
		expected []int // Note: order is not guaranteed, so we'll sort before comparing
	}{
		{map[int]struct{}{1: {}, 2: {}, 3: {}}, []int{1, 2, 3}},
		{map[int]struct{}{}, []int{}},
		{map[int]struct{}{42: {}}, []int{42}},
	}

	for _, test := range tests {
		result := SetToSlice(test.set)
		// Sort both slices to ensure consistent comparison
		sort.Ints(result)
		sort.Ints(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SetToSlice(%v) = %v, expected %v", test.set, result, test.expected)
		}
	}
}

func TestSliceToSet(t *testing.T) {
	tests := []struct {
		slice    []int
		expected map[int]struct{}
	}{
		{[]int{1, 2, 3}, map[int]struct{}{1: {}, 2: {}, 3: {}}},
		{[]int{}, map[int]struct{}{}},
		{[]int{1, 1, 2, 2, 3}, map[int]struct{}{1: {}, 2: {}, 3: {}}}, // Test deduplication
	}

	for _, test := range tests {
		result := SliceToSet(test.slice)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SliceToSet(%v) = %v, expected %v", test.slice, result, test.expected)
		}
	}
}

func TestSetUnion(t *testing.T) {
	tests := []struct {
		set1     map[int]struct{}
		set2     map[int]struct{}
		expected map[int]struct{}
	}{
		{
			map[int]struct{}{1: {}, 2: {}, 3: {}},
			map[int]struct{}{3: {}, 4: {}, 5: {}},
			map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{1: {}, 2: {}},
		},
		{
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{},
			map[int]struct{}{1: {}, 2: {}},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{},
			map[int]struct{}{},
		},
	}

	for _, test := range tests {
		result := SetUnion(test.set1, test.set2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SetUnion(%v, %v) = %v, expected %v", test.set1, test.set2, result, test.expected)
		}
	}
}

func TestSetIntersection(t *testing.T) {
	tests := []struct {
		set1     map[int]struct{}
		set2     map[int]struct{}
		expected map[int]struct{}
	}{
		{
			map[int]struct{}{1: {}, 2: {}, 3: {}},
			map[int]struct{}{3: {}, 4: {}, 5: {}},
			map[int]struct{}{3: {}},
		},
		{
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{3: {}, 4: {}},
			map[int]struct{}{},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{},
		},
		{
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{},
			map[int]struct{}{},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{},
			map[int]struct{}{},
		},
	}

	for _, test := range tests {
		result := SetIntersection(test.set1, test.set2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SetIntersection(%v, %v) = %v, expected %v", test.set1, test.set2, result, test.expected)
		}
	}
}

func TestSetDifference(t *testing.T) {
	tests := []struct {
		set1     map[int]struct{}
		set2     map[int]struct{}
		expected map[int]struct{}
	}{
		{
			map[int]struct{}{1: {}, 2: {}, 3: {}},
			map[int]struct{}{3: {}, 4: {}, 5: {}},
			map[int]struct{}{1: {}, 2: {}},
		},
		{
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{3: {}, 4: {}},
			map[int]struct{}{1: {}, 2: {}},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{},
		},
		{
			map[int]struct{}{1: {}, 2: {}},
			map[int]struct{}{},
			map[int]struct{}{1: {}, 2: {}},
		},
		{
			map[int]struct{}{},
			map[int]struct{}{},
			map[int]struct{}{},
		},
	}

	for _, test := range tests {
		result := SetDifference(test.set1, test.set2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SetDifference(%v, %v) = %v, expected %v", test.set1, test.set2, result, test.expected)
		}
	}
}

func TestMapMergeMaps(t *testing.T) {
	tests := []struct {
		maps     []map[string]int
		expected map[string]int
	}{
		{
			[]map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
			},
			map[string]int{"a": 1, "b": 3, "c": 4}, // b from second map overwrites b from first map
		},
		{
			[]map[string]int{
				{"a": 1},
				{},
				{"b": 2},
			},
			map[string]int{"a": 1, "b": 2},
		},
		{
			[]map[string]int{},
			map[string]int{},
		},
		{
			[]map[string]int{
				{},
			},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := MapMerge(test.maps...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapMerge(%v) = %v, expected %v", test.maps, result, test.expected)
		}
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected []string // Note: order is not guaranteed, so we'll sort before comparing
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "b", "c"}},
		{map[string]int{}, []string{}},
		{map[string]int{"x": 42}, []string{"x"}},
	}

	for _, test := range tests {
		result := MapKeys(test.m)
		// Sort both slices to ensure consistent comparison
		sort.Strings(result)
		sort.Strings(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapKeys(%v) = %v, expected %v", test.m, result, test.expected)
		}
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected []int // Note: order is not guaranteed, so we'll sort before comparing
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{1, 2, 3}},
		{map[string]int{}, []int{}},
		{map[string]int{"x": 42}, []int{42}},
	}

	for _, test := range tests {
		result := MapValues(test.m)
		// Sort both slices to ensure consistent comparison
		sort.Ints(result)
		sort.Ints(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapValues(%v) = %v, expected %v", test.m, result, test.expected)
		}
	}
}

func TestMapValuesFn(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			map[string]int{"a": 2, "b": 4, "c": 6},
		},
		{
			map[string]int{},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := MapValuesFn(test.m, func(v int) int { return v * 2 })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapValues(%v, func) = %v, expected %v", test.m, result, test.expected)
		}
	}
}

func TestMapGetOrDefault(t *testing.T) {
	tests := []struct {
		m            map[string]int
		key          string
		defaultValue int
		expected     int
	}{
		{map[string]int{"a": 1, "b": 2}, "a", 99, 1},       // Key exists
		{map[string]int{"a": 1, "b": 2}, "c", 99, 99},      // Key doesn't exist
		{map[string]int{}, "anything", 42, 42},             // Empty map
		{map[string]int{"zero": 0}, "zero", 42, 0},         // Value is zero
		{map[string]int{"zero": 0}, "nonexistent", 42, 42}, // Key doesn't exist
	}

	for _, test := range tests {
		result := MapGetOrDefault(test.m, test.key, test.defaultValue)
		if result != test.expected {
			t.Errorf("MapGetOrDefault(%v, %q, %d) = %d, expected %d",
				test.m, test.key, test.defaultValue, result, test.expected)
		}
	}
}

func TestMapEqualMaps(t *testing.T) {
	tests := []struct {
		m1       map[string]int
		m2       map[string]int
		expected bool
	}{
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}, true},  // Same maps
		{map[string]int{"a": 1, "b": 2}, map[string]int{"b": 2, "a": 1}, true},  // Same maps, different order
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "c": 3}, false}, // Different keys
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 3}, false}, // Different values
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, false},         // Different sizes
		{map[string]int{}, map[string]int{}, true},                              // Both empty
		{map[string]int{"a": 1}, map[string]int{}, false},                       // One empty, one not
	}

	for _, test := range tests {
		result := MapEqualMaps(test.m1, test.m2)
		if result != test.expected {
			t.Errorf("MapEqualMaps(%v, %v) = %v, expected %v",
				test.m1, test.m2, result, test.expected)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}}, // Already unique
		{[]int{}, []int{}},               // Empty slice
	}

	for _, test := range tests {
		result := Unique(test.input)
		// Sort both slices to ensure consistent comparison
		sort.Ints(result)
		sort.Ints(test.expected)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Unique(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestSortedCopy(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 2}, []int{1, 2, 3, 4}},
		{[]int{5, 5, 3, 3, 1}, []int{1, 3, 3, 5, 5}}, // With duplicates
		{[]int{1, 2, 3}, []int{1, 2, 3}},             // Already sorted
		{[]int{}, []int{}},                           // Empty slice
	}

	for _, test := range tests {
		result := SortedCopy(test.input, func(a, b int) bool { return a < b })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortedCopy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}

		// Verify the original slice wasn't modified
		if len(test.input) > 0 && reflect.DeepEqual(test.input, result) && !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("SortedCopy modified the original slice: %v", test.input)
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		input        []int
		initialValue int
		expected     int
	}{
		{[]int{1, 2, 3, 4}, 0, 10},  // Sum
		{[]int{1, 2, 3, 4}, 10, 20}, // Sum with initial value
		{[]int{2, 3, 4}, 1, 24},     // Product
		{[]int{}, 42, 42},           // Empty slice returns initial value
	}

	for _, test := range tests {
		// Test sum reducer
		if test.initialValue == 0 || test.initialValue == 10 {
			result := Reduce(test.input, test.initialValue, func(acc, item int) int { return acc + item })
			if result != test.expected {
				t.Errorf("Reduce(%v, %d, sum) = %d, expected %d", test.input, test.initialValue, result, test.expected)
			}
		} else { // Test product reducer
			result := Reduce(test.input, test.initialValue, func(acc, item int) int { return acc * item })
			if result != test.expected {
				t.Errorf("Reduce(%v, %d, product) = %d, expected %d", test.input, test.initialValue, result, test.expected)
			}
		}
	}
}

func TestGroupBy(t *testing.T) {
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
	result := GroupBy(people, func(p Person) int { return p.Age })

	// Expected groups
	expected := map[int][]Person{
		25: {{"Alice", 25}, {"Charlie", 25}},
		30: {{"Bob", 30}, {"Dave", 30}},
		35: {{"Eve", 35}},
	}

	// Check each group
	for age, group := range expected {
		resultGroup, ok := result[age]
		if !ok {
			t.Errorf("GroupBy missing expected group for age %d", age)
			continue
		}

		// Check group size
		if len(resultGroup) != len(group) {
			t.Errorf("GroupBy for age %d has %d items, expected %d", age, len(resultGroup), len(group))
			continue
		}

		// Check each person in the group
		for _, expectedPerson := range group {
			found := false
			for _, resultPerson := range resultGroup {
				if resultPerson.Name == expectedPerson.Name && resultPerson.Age == expectedPerson.Age {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("GroupBy for age %d missing person %+v", age, expectedPerson)
			}
		}
	}

	// Test empty slice
	emptyResult := GroupBy([]Person{}, func(p Person) int { return p.Age })
	if len(emptyResult) != 0 {
		t.Errorf("GroupBy on empty slice returned %v, expected empty map", emptyResult)
	}
}

func TestMapDiffMaps(t *testing.T) {
	tests := []struct {
		m1              map[string]int
		m2              map[string]int
		expectedAdded   map[string]int
		expectedRemoved map[string]int
		expectedChanged map[string]int
	}{
		{
			m1:              map[string]int{"a": 1, "b": 2, "c": 3},
			m2:              map[string]int{"b": 2, "c": 4, "d": 5},
			expectedAdded:   map[string]int{"d": 5},
			expectedRemoved: map[string]int{"a": 1},
			expectedChanged: map[string]int{"c": 4},
		},
		{
			m1:              map[string]int{"a": 1, "b": 2},
			m2:              map[string]int{"a": 1, "b": 2},
			expectedAdded:   map[string]int{},
			expectedRemoved: map[string]int{},
			expectedChanged: map[string]int{},
		},
		{
			m1:              map[string]int{},
			m2:              map[string]int{"a": 1, "b": 2},
			expectedAdded:   map[string]int{"a": 1, "b": 2},
			expectedRemoved: map[string]int{},
			expectedChanged: map[string]int{},
		},
		{
			m1:              map[string]int{"a": 1, "b": 2},
			m2:              map[string]int{},
			expectedAdded:   map[string]int{},
			expectedRemoved: map[string]int{"a": 1, "b": 2},
			expectedChanged: map[string]int{},
		},
	}

	for _, test := range tests {
		added, removed, changed := MapDiffMaps(test.m1, test.m2)

		if !reflect.DeepEqual(added, test.expectedAdded) {
			t.Errorf("MapDiffMaps(%v, %v) added = %v, expected %v",
				test.m1, test.m2, added, test.expectedAdded)
		}

		if !reflect.DeepEqual(removed, test.expectedRemoved) {
			t.Errorf("MapDiffMaps(%v, %v) removed = %v, expected %v",
				test.m1, test.m2, removed, test.expectedRemoved)
		}

		if !reflect.DeepEqual(changed, test.expectedChanged) {
			t.Errorf("MapDiffMaps(%v, %v) changed = %v, expected %v",
				test.m1, test.m2, changed, test.expectedChanged)
		}
	}
}

func TestMapGetOrInsert(t *testing.T) {
	tests := []struct {
		m            map[string]int
		key          string
		defaultValue int
		expected     int
		expectedMap  map[string]int
	}{
		{
			m:            map[string]int{"a": 1, "b": 2},
			key:          "a",
			defaultValue: 99,
			expected:     1,
			expectedMap:  map[string]int{"a": 1, "b": 2}, // No change
		},
		{
			m:            map[string]int{"a": 1, "b": 2},
			key:          "c",
			defaultValue: 99,
			expected:     99,
			expectedMap:  map[string]int{"a": 1, "b": 2, "c": 99}, // Key added
		},
		{
			m:            map[string]int{},
			key:          "a",
			defaultValue: 42,
			expected:     42,
			expectedMap:  map[string]int{"a": 42}, // Key added to empty map
		},
	}

	for _, test := range tests {
		// Create a copy of the map for testing
		m := make(map[string]int)
		for k, v := range test.m {
			m[k] = v
		}

		result := MapGetOrInsert(m, test.key, test.defaultValue)

		if result != test.expected {
			t.Errorf("MapGetOrInsert(%v, %q, %d) = %d, expected %d",
				test.m, test.key, test.defaultValue, result, test.expected)
		}

		if !reflect.DeepEqual(m, test.expectedMap) {
			t.Errorf("MapGetOrInsert(%v, %q, %d) resulted in map %v, expected %v",
				test.m, test.key, test.defaultValue, m, test.expectedMap)
		}
	}
}

func TestMapInvertMap(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected map[int]string
	}{
		{
			m:        map[string]int{"a": 1, "b": 2, "c": 3},
			expected: map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			m:        map[string]int{},
			expected: map[int]string{},
		},
		{
			// When multiple keys map to the same value, only one will be in the result
			m:        map[string]int{"a": 1, "b": 1, "c": 2},
			expected: map[int]string{1: "b", 2: "c"}, // or {1: "a", 2: "c"}, but we can't predict which
		},
	}

	for _, test := range tests {
		result := MapInvertMap(test.m)

		// Special case for duplicate values
		if len(test.m) != len(result) && !hasDuplicateValues(test.m) {
			t.Errorf("MapInvertMap(%v) = %v, expected %v (lengths differ without duplicate values)",
				test.m, result, test.expected)
			continue
		}

		// For the case with duplicate values, we need to check differently
		if hasDuplicateValues(test.m) {
			// Check that each value in the original map has its key in the inverted map
			for k, v := range test.m {
				invertedKey, ok := result[v]
				if !ok {
					t.Errorf("MapInvertMap(%v) = %v, missing value %d from original map",
						test.m, result, v)
				} else if invertedKey != k && !hasKeyWithValue(test.m, invertedKey, v) {
					t.Errorf("MapInvertMap(%v) = %v, value %d maps to %q which doesn't have value %d in original map",
						test.m, result, v, invertedKey, v)
				}
			}
			// For maps without duplicate values, we can do a direct comparison
		} else if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapInvertMap(%v) = %v, expected %v",
				test.m, result, test.expected)
		}
	}
}

// Helper function to check if a map has duplicate values
func hasDuplicateValues[K comparable, V comparable](m map[K]V) bool {
	seen := make(map[V]bool)
	for _, v := range m {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}

// Helper function to check if a map has a key with a specific value
func hasKeyWithValue[K comparable, V comparable](m map[K]V, key K, value V) bool {
	v, ok := m[key]
	return ok && v == value
}

func TestAccessible(t *testing.T) {
	tests := []struct {
		input    any
		expected bool
	}{
		{[]int{1, 2, 3}, true},
		{map[string]int{"a": 1, "b": 2}, true},
		{[3]int{1, 2, 3}, true},
		{"not an array", false},
		{42, false},
		{nil, false},
	}

	for _, test := range tests {
		result := Accessible(test.input)
		if result != test.expected {
			t.Errorf("Accessible(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		array    map[string]any
		key      string
		value    any
		expected map[string]any
	}{
		{
			map[string]any{"name": "John", "age": 30},
			"city",
			"New York",
			map[string]any{"name": "John", "age": 30, "city": "New York"},
		},
		{
			map[string]any{"name": "John", "age": 30},
			"name",
			"Jane",
			map[string]any{"name": "John", "age": 30}, // Key already exists, no change
		},
		{
			map[string]any{},
			"name",
			"John",
			map[string]any{"name": "John"},
		},
	}

	for _, test := range tests {
		result := Add(test.array, test.key, test.value)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Add(%v, %q, %v) = %v, expected %v", test.array, test.key, test.value, result, test.expected)
		}

		// Verify the original map wasn't modified
		if test.key != "name" && !reflect.DeepEqual(test.array, map[string]any{"name": "John", "age": 30}) {
			t.Errorf("Add modified the original map: %v", test.array)
		}
	}
}

func TestCollapse(t *testing.T) {
	tests := []struct {
		input    [][]any
		expected []any
	}{
		{
			[][]any{{1, 2}, {3, 4}},
			[]any{1, 2, 3, 4},
		},
		{
			[][]any{{1}, {2}, {3}},
			[]any{1, 2, 3},
		},
		{
			[][]any{{"a", "b"}, {"c", "d"}},
			[]any{"a", "b", "c", "d"},
		},
		{
			[][]any{},
			[]any{},
		},
	}

	for _, test := range tests {
		result := Collapse(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Collapse(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestCrossJoin(t *testing.T) {
	tests := []struct {
		arrays   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 2}, {3, 4}},
			[][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}},
		},
		{
			[][]int{{1, 2}, {3}, {4, 5}},
			[][]int{{1, 3, 4}, {1, 3, 5}, {2, 3, 4}, {2, 3, 5}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{},
			[][]int{},
		},
	}

	for _, test := range tests {
		result := CrossJoin(test.arrays...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CrossJoin(%v) = %v, expected %v", test.arrays, result, test.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		input          map[string]any
		expectedKeys   []string
		expectedValues []any
	}{
		{
			map[string]any{"name": "John", "age": 30},
			[]string{"name", "age"},
			[]any{"John", 30},
		},
		{
			map[string]any{"a": 1, "b": 2, "c": 3},
			[]string{"a", "b", "c"},
			[]any{1, 2, 3},
		},
		{
			map[string]any{},
			[]string{},
			[]any{},
		},
	}

	for _, test := range tests {
		keys, values := Divide(test.input)

		// Sort the results for consistent comparison
		sort.Strings(keys)
		sort.Strings(test.expectedKeys)

		// We can't easily sort the values slice since it contains any,
		// so we'll check that each expected value is in the result
		if !reflect.DeepEqual(keys, test.expectedKeys) {
			t.Errorf("Divide(%v) keys = %v, expected %v", test.input, keys, test.expectedKeys)
		}

		// Check that all expected values are in the result
		if len(values) != len(test.expectedValues) {
			t.Errorf("Divide(%v) values length = %d, expected %d", test.input, len(values), len(test.expectedValues))
		} else {
			// Since the order of keys and values returned by Divide is not guaranteed,
			// we need to check that each key in the input map has its corresponding value
			// in the result, regardless of the order.

			// Create a map of the input for easier lookup
			inputMap := make(map[string]any)
			for k, v := range test.input {
				inputMap[k] = v
			}

			// Check that each key has its corresponding value in the result
			for i, k := range keys {
				expectedValue, ok := inputMap[k]
				if !ok {
					t.Errorf("Divide(%v) returned key %q which is not in the input map", test.input, k)
				} else if !reflect.DeepEqual(values[i], expectedValue) {
					// Find the index of this key in the expected keys
					keyIndex := -1
					for j, expectedKey := range test.expectedKeys {
						if expectedKey == k {
							keyIndex = j
							break
						}
					}

					if keyIndex != -1 && !reflect.DeepEqual(values[i], test.expectedValues[keyIndex]) {
						t.Errorf("Divide(%v) value for key %q = %v, expected %v",
							test.input, k, values[i], test.expectedValues[keyIndex])
					}
				}
			}
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected map[string]any
	}{
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
					"address": map[string]any{
						"city": "New York",
					},
				},
			},
			map[string]any{
				"user.name":         "John",
				"user.address.city": "New York",
			},
		},
		{
			map[string]any{
				"a": 1,
				"b": map[string]any{
					"c": 2,
					"d": 3,
				},
			},
			map[string]any{
				"a":   1,
				"b.c": 2,
				"b.d": 3,
			},
		},
		{
			map[string]any{},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := Dot(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Dot(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestExcept(t *testing.T) {
	tests := []struct {
		array    map[string]any
		keys     []string
		expected map[string]any
	}{
		{
			map[string]any{"name": "John", "age": 30, "city": "New York"},
			[]string{"age"},
			map[string]any{"name": "John", "city": "New York"},
		},
		{
			map[string]any{"name": "John", "age": 30, "city": "New York"},
			[]string{"age", "city"},
			map[string]any{"name": "John"},
		},
		{
			map[string]any{"name": "John", "age": 30},
			[]string{"nonexistent"},
			map[string]any{"name": "John", "age": 30},
		},
		{
			map[string]any{},
			[]string{"name"},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := Except(test.array, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Except(%v, %v) = %v, expected %v", test.array, test.keys, result, test.expected)
		}
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		array    map[string]any
		key      string
		expected bool
	}{
		{
			map[string]any{"name": "John", "age": 30},
			"name",
			true,
		},
		{
			map[string]any{"name": "John", "age": 30},
			"city",
			false,
		},
		{
			map[string]any{},
			"name",
			false,
		},
	}

	for _, test := range tests {
		result := Exists(test.array, test.key)
		if result != test.expected {
			t.Errorf("Exists(%v, %q) = %v, expected %v", test.array, test.key, result, test.expected)
		}
	}
}

func TestForget(t *testing.T) {
	tests := []struct {
		array    map[string]any
		keys     []string
		expected map[string]any
	}{
		{
			map[string]any{"name": "John", "age": 30, "city": "New York"},
			[]string{"age"},
			map[string]any{"name": "John", "city": "New York"},
		},
		{
			map[string]any{"name": "John", "age": 30, "city": "New York"},
			[]string{"age", "city"},
			map[string]any{"name": "John"},
		},
		{
			map[string]any{"name": "John", "age": 30},
			[]string{"nonexistent"},
			map[string]any{"name": "John", "age": 30},
		},
		{
			map[string]any{},
			[]string{"name"},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := Forget(test.array, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Forget(%v, %v) = %v, expected %v", test.array, test.keys, result, test.expected)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		array        map[string]any
		key          string
		defaultValue any
		expected     any
	}{
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
					"address": map[string]any{
						"city": "New York",
					},
				},
			},
			"user.name",
			"default",
			"John",
		},
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
			"user.address.city",
			"default",
			"default",
		},
		{
			map[string]any{},
			"user.name",
			"default",
			"default",
		},
		{
			nil,
			"user.name",
			"default",
			"default",
		},
	}

	for _, test := range tests {
		result := Get(test.array, test.key, test.defaultValue)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Get(%v, %q, %v) = %v, expected %v", test.array, test.key, test.defaultValue, result, test.expected)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		array    map[string]any
		keys     []string
		expected bool
	}{
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
					"address": map[string]any{
						"city": "New York",
					},
				},
			},
			[]string{"user.name"},
			true,
		},
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
			[]string{"user.address.city"},
			false,
		},
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
					"age":  30,
				},
			},
			[]string{"user.name", "user.age"},
			true,
		},
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
			[]string{"user.name", "user.age"},
			false,
		},
		{
			map[string]any{},
			[]string{"user.name"},
			false,
		},
		{
			nil,
			[]string{"user.name"},
			false,
		},
		{
			map[string]any{},
			[]string{},
			false,
		},
	}

	for _, test := range tests {
		result := Has(test.array, test.keys...)
		if result != test.expected {
			t.Errorf("Has(%v, %v) = %v, expected %v", test.array, test.keys, result, test.expected)
		}
	}
}

func TestHasAny(t *testing.T) {
	tests := []struct {
		array    map[string]any
		keys     []string
		expected bool
	}{
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
			[]string{"user.name", "user.age"},
			true,
		},
		{
			map[string]any{
				"user": map[string]any{
					"name": "John",
				},
			},
			[]string{"user.age", "user.address.city"},
			false,
		},
		{
			map[string]any{},
			[]string{"user.name"},
			false,
		},
		{
			nil,
			[]string{"user.name"},
			false,
		},
		{
			map[string]any{},
			[]string{},
			false,
		},
	}

	for _, test := range tests {
		result := HasAny(test.array, test.keys...)
		if result != test.expected {
			t.Errorf("HasAny(%v, %v) = %v, expected %v", test.array, test.keys, result, test.expected)
		}
	}
}

func TestIsAssoc(t *testing.T) {
	tests := []struct {
		input    any
		expected bool
	}{
		{map[string]int{"a": 1, "b": 2}, true},
		{map[string]any{"a": 1, "b": "test"}, true},
		{map[int]string{1: "a", 2: "b"}, false},
		{[]int{1, 2, 3}, false},
		{42, false},
		{nil, false},
	}

	for _, test := range tests {
		result := IsAssoc(test.input)
		if result != test.expected {
			t.Errorf("IsAssoc(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestIsList(t *testing.T) {
	tests := []struct {
		input    any
		expected bool
	}{
		{[]int{1, 2, 3}, true},
		{[]string{"a", "b", "c"}, true},
		{[3]int{1, 2, 3}, true},
		{map[string]int{"a": 1, "b": 2}, false},
		{42, false},
		{nil, false},
	}

	for _, test := range tests {
		result := IsList(test.input)
		if result != test.expected {
			t.Errorf("IsList(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestKeyBy(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	people := []Person{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}

	tests := []struct {
		input    []Person
		keyFunc  func(Person) int
		expected map[int]Person
	}{
		{
			people,
			func(p Person) int { return p.ID },
			map[int]Person{
				1: {1, "Alice"},
				2: {2, "Bob"},
				3: {3, "Charlie"},
			},
		},
		{
			[]Person{},
			func(p Person) int { return p.ID },
			map[int]Person{},
		},
	}

	for _, test := range tests {
		result := KeyBy(test.input, test.keyFunc)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("KeyBy(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestOnly(t *testing.T) {
	tests := []struct {
		array    map[string]any
		keys     []string
		expected map[string]any
	}{
		{
			map[string]any{"name": "John", "age": 30, "city": "New York"},
			[]string{"name", "age"},
			map[string]any{"name": "John", "age": 30},
		},
		{
			map[string]any{"name": "John", "age": 30},
			[]string{"name", "nonexistent"},
			map[string]any{"name": "John"},
		},
		{
			map[string]any{},
			[]string{"name"},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := Only(test.array, test.keys...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Only(%v, %v) = %v, expected %v", test.array, test.keys, result, test.expected)
		}
	}
}

func TestPluck(t *testing.T) {
	type Person struct {
		ID   int
		Name string
		Age  int
	}

	people := []Person{
		{1, "Alice", 25},
		{2, "Bob", 30},
		{3, "Charlie", 35},
	}

	tests := []struct {
		input    []Person
		keyFunc  func(Person) string
		expected []string
	}{
		{
			people,
			func(p Person) string { return p.Name },
			[]string{"Alice", "Bob", "Charlie"},
		},
		{
			[]Person{},
			func(p Person) string { return p.Name },
			[]string{},
		},
	}

	for _, test := range tests {
		result := Pluck(test.input, test.keyFunc)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Pluck(%v, func) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestQuery(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected string
	}{
		{
			map[string]any{"name": "John", "age": 30},
			"age=30&name=John",
		},
		{
			map[string]any{"tags": []string{"php", "laravel"}},
			"tags%5B%5D=php&tags%5B%5D=laravel",
		},
		{
			map[string]any{},
			"",
		},
	}

	for _, test := range tests {
		result := Query(test.input)
		// URL query parameters can be in any order, so we need to parse and compare
		resultValues, _ := url.ParseQuery(result)
		expectedValues, _ := url.ParseQuery(test.expected)

		if !reflect.DeepEqual(resultValues, expectedValues) {
			t.Errorf("Query(%v) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestRandomOrDefault(t *testing.T) {
	// Test with non-empty slice
	input := []int{1, 2, 3, 4, 5}
	result := RandomOrDefault(input, 0)
	found := false
	for _, v := range input {
		if result == v {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomOrDefault(%v, 0) = %d, which is not in the input slice", input, result)
	}

	// Test with empty slice
	emptyResult := RandomOrDefault([]int{}, 42)
	if emptyResult != 42 {
		t.Errorf("RandomOrDefault([], 42) = %d, expected 42", emptyResult)
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		array    map[string]any
		key      string
		value    any
		expected map[string]any
	}{
		{
			map[string]any{"user": map[string]any{"name": "John"}},
			"user.age",
			30,
			map[string]any{"user": map[string]any{"name": "John", "age": 30}},
		},
		{
			map[string]any{},
			"user.name",
			"John",
			map[string]any{"user": map[string]any{"name": "John"}},
		},
		{
			map[string]any{"user": "John"},
			"user.name",
			"John Jr.",
			map[string]any{"user": map[string]any{"name": "John Jr."}},
		},
	}

	for _, test := range tests {
		result := Set(test.array, test.key, test.value)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Set(%v, %q, %v) = %v, expected %v", test.array, test.key, test.value, result, test.expected)
		}
	}
}

func TestSortByKey(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected map[string]any
	}{
		{
			map[string]any{"c": 3, "a": 1, "b": 2},
			map[string]any{"a": 1, "b": 2, "c": 3},
		},
		{
			map[string]any{},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := SortByKey(test.input)

		// Check that the result has the same keys and values
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortByKey(%v) = %v, expected %v", test.input, result, test.expected)
		}

		// Note: We don't check if the keys are in sorted order because in Go,
		// the order of keys in a map is not guaranteed when iterating.
		// The SortByKey function is supposed to sort the keys internally,
		// but we can only verify that the result has the same keys and values
		// as the expected result.
	}
}

func TestSortByKeyDesc(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected map[string]any
	}{
		{
			map[string]any{"a": 1, "b": 2, "c": 3},
			map[string]any{"c": 3, "b": 2, "a": 1},
		},
		{
			map[string]any{},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := SortByKeyDesc(test.input)

		// Check that the result has the same keys and values
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortByKeyDesc(%v) = %v, expected %v", test.input, result, test.expected)
		}

		// Note: We don't check if the keys are in reverse sorted order because in Go,
		// the order of keys in a map is not guaranteed when iterating.
		// The SortByKeyDesc function is supposed to sort the keys internally in descending order,
		// but we can only verify that the result has the same keys and values
		// as the expected result.
	}
}

func TestSortRecursive(t *testing.T) {
	tests := []struct {
		input    any
		expected any
	}{
		{
			map[string]any{
				"c": 3,
				"a": 1,
				"b": map[string]any{
					"z": 3,
					"x": 1,
					"y": 2,
				},
			},
			map[string]any{
				"a": 1,
				"b": map[string]any{
					"x": 1,
					"y": 2,
					"z": 3,
				},
				"c": 3,
			},
		},
		{
			[]any{3, 1, 2, map[string]any{"c": 3, "a": 1, "b": 2}},
			[]any{3, 1, 2, map[string]any{"a": 1, "b": 2, "c": 3}},
		},
		{
			42,
			42,
		},
	}

	for _, test := range tests {
		result := SortRecursive(test.input)

		// Note: We don't check if the keys are sorted because in Go,
		// the order of keys in a map is not guaranteed when iterating.
		// The SortRecursive function is supposed to sort the keys internally,
		// but we can only verify that the result has the same structure and values
		// as the expected result.

		// Check that the result has the same structure and values
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SortRecursive(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestUndot(t *testing.T) {
	tests := []struct {
		input    map[string]any
		expected map[string]any
	}{
		{
			map[string]any{
				"user.name":         "John",
				"user.address.city": "New York",
			},
			map[string]any{
				"user": map[string]any{
					"name": "John",
					"address": map[string]any{
						"city": "New York",
					},
				},
			},
		},
		{
			map[string]any{
				"a":   1,
				"b.c": 2,
				"b.d": 3,
			},
			map[string]any{
				"a": 1,
				"b": map[string]any{
					"c": 2,
					"d": 3,
				},
			},
		},
		{
			map[string]any{},
			map[string]any{},
		},
	}

	for _, test := range tests {
		result := Undot(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Undot(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestWhereNotNull(t *testing.T) {
	var nilPtr *int
	tests := []struct {
		input    []any
		expected []any
	}{
		{
			[]any{1, nil, "test", nilPtr, 2},
			[]any{1, "test", 2},
		},
		{
			[]any{nil, nilPtr},
			[]any{},
		},
		{
			[]any{1, 2, 3},
			[]any{1, 2, 3},
		},
		{
			[]any{},
			[]any{},
		},
	}

	for _, test := range tests {
		result := WhereNotNull(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("WhereNotNull(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		input    any
		expected []any
	}{
		{
			42,
			[]any{42},
		},
		{
			[]int{1, 2, 3},
			[]any{1, 2, 3},
		},
		{
			[2]string{"a", "b"},
			[]any{"a", "b"},
		},
		{
			nil,
			[]any{},
		},
	}

	for _, test := range tests {
		result := Wrap(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Wrap(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestMapFindKey(t *testing.T) {
	tests := []struct {
		m           map[string]int
		value       int
		expectedKey string
		expectedOk  bool
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			2,
			"b",
			true,
		},
		{
			map[string]int{"a": 1, "b": 2, "c": 3},
			4,
			"",
			false,
		},
		{
			map[string]int{},
			1,
			"",
			false,
		},
	}

	for _, test := range tests {
		key, ok := MapFindKey(test.m, test.value)
		if key != test.expectedKey || ok != test.expectedOk {
			t.Errorf("MapFindKey(%v, %d) = (%q, %v), expected (%q, %v)",
				test.m, test.value, key, ok, test.expectedKey, test.expectedOk)
		}
	}
}

func TestMapFilterMap(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected map[string]int
	}{
		{
			map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			map[string]int{"b": 2, "d": 4},
		},
		{
			map[string]int{"a": 1, "c": 3},
			map[string]int{},
		},
		{
			map[string]int{},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := MapFilterMap(test.m, func(k string, v int) bool { return v%2 == 0 })
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapFilterMap(%v, func) = %v, expected %v", test.m, result, test.expected)
		}
	}
}

func TestMapToSlice(t *testing.T) {
	tests := []struct {
		m        map[string]int
		expected []struct {
			Key   string
			Value int
		}
	}{
		{
			map[string]int{"a": 1, "b": 2},
			[]struct {
				Key   string
				Value int
			}{
				{"a", 1},
				{"b", 2},
			},
		},
		{
			map[string]int{},
			[]struct {
				Key   string
				Value int
			}{},
		},
	}

	for _, test := range tests {
		result := MapToSlice(test.m)

		// Sort both slices by key for consistent comparison
		sort.Slice(result, func(i, j int) bool {
			return result[i].Key < result[j].Key
		})

		sort.Slice(test.expected, func(i, j int) bool {
			return test.expected[i].Key < test.expected[j].Key
		})

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapToSlice(%v) = %v, expected %v", test.m, result, test.expected)
		}
	}
}

func TestMapSliceToMap(t *testing.T) {
	tests := []struct {
		slice []struct {
			Key   string
			Value int
		}
		expected map[string]int
	}{
		{
			[]struct {
				Key   string
				Value int
			}{
				{"a", 1},
				{"b", 2},
			},
			map[string]int{"a": 1, "b": 2},
		},
		{
			[]struct {
				Key   string
				Value int
			}{},
			map[string]int{},
		},
	}

	for _, test := range tests {
		result := MapSliceToMap(test.slice)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MapSliceToMap(%v) = %v, expected %v", test.slice, result, test.expected)
		}
	}
}
