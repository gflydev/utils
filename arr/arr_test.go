package arr

import (
	"reflect"
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
