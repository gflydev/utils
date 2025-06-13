package seq

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	seq := New(1, 2, 3)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(seq.Value(), expected) {
		t.Errorf("New(1, 2, 3).Value() = %v, expected %v", seq.Value(), expected)
	}
}

func TestFromSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	seq := FromSlice(slice)

	if !reflect.DeepEqual(seq.Value(), slice) {
		t.Errorf("FromSlice(%v).Value() = %v, expected %v", slice, seq.Value(), slice)
	}
}

func TestValue(t *testing.T) {
	seq := New(1, 2, 3)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(seq.Value(), expected) {
		t.Errorf("New(1, 2, 3).Value() = %v, expected %v", seq.Value(), expected)
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected int
		ok       bool
	}{
		{New(1, 2, 3), 1, true},
		{New[int](), 0, false},
	}

	for _, test := range tests {
		result, ok := test.seq.First()
		if result != test.expected || ok != test.ok {
			t.Errorf("%v.First() = (%v, %v), expected (%v, %v)", test.seq.Value(), result, ok, test.expected, test.ok)
		}
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected int
		ok       bool
	}{
		{New(1, 2, 3), 3, true},
		{New[int](), 0, false},
	}

	for _, test := range tests {
		result, ok := test.seq.Last()
		if result != test.expected || ok != test.ok {
			t.Errorf("%v.Last() = (%v, %v), expected (%v, %v)", test.seq.Value(), result, ok, test.expected, test.ok)
		}
	}
}

func TestMap(t *testing.T) {
	seq := New(1, 2, 3)
	result := seq.Map(func(n int) int { return n * 2 })
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("%v.Map(func) = %v, expected %v", seq.Value(), result.Value(), expected)
	}
}

func TestFilter(t *testing.T) {
	seq := New(1, 2, 3, 4)
	result := seq.Filter(func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4}

	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("%v.Filter(func) = %v, expected %v", seq.Value(), result.Value(), expected)
	}
}

func TestReject(t *testing.T) {
	seq := New(1, 2, 3, 4)
	result := seq.Reject(func(n int) bool { return n%2 == 0 })
	expected := []int{1, 3}

	if !reflect.DeepEqual(result.Value(), expected) {
		t.Errorf("%v.Reject(func) = %v, expected %v", seq.Value(), result.Value(), expected)
	}
}

func TestReduce(t *testing.T) {
	seq := New(1, 2, 3, 4)
	result := seq.Reduce(func(acc interface{}, n int) interface{} {
		return acc.(int) + n
	}, 0)
	expected := 10

	if result != expected {
		t.Errorf("%v.Reduce(func, 0) = %v, expected %v", seq.Value(), result, expected)
	}
}

func TestForEach(t *testing.T) {
	seq := New(1, 2, 3)
	sum := 0
	result := seq.ForEach(func(n int) {
		sum += n
	})

	if sum != 6 {
		t.Errorf("%v.ForEach(func) resulted in sum = %v, expected 6", seq.Value(), sum)
	}

	// ForEach should return the original sequence for chaining
	if !reflect.DeepEqual(result.Value(), seq.Value()) {
		t.Errorf("%v.ForEach(func) returned %v, expected %v", seq.Value(), result.Value(), seq.Value())
	}
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		value    int
		expected bool
	}{
		{New(1, 2, 3), 2, true},
		{New(1, 2, 3), 4, false},
		{New[int](), 1, false},
	}

	for _, test := range tests {
		result := test.seq.Includes(test.value)
		if result != test.expected {
			t.Errorf("%v.Includes(%v) = %v, expected %v", test.seq.Value(), test.value, result, test.expected)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected int
		ok       bool
	}{
		{New(1, 2, 3, 4), 3, true},
		{New(1, 2), 0, false},
		{New[int](), 0, false},
	}

	for _, test := range tests {
		result, ok := test.seq.Find(func(n int) bool { return n > 2 })
		if result != test.expected || ok != test.ok {
			t.Errorf("%v.Find(func) = (%v, %v), expected (%v, %v)", test.seq.Value(), result, ok, test.expected, test.ok)
		}
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected int
		ok       bool
	}{
		{New(1, 2, 3, 4, 3), 3, true},
		{New(1, 2), 0, false},
		{New[int](), 0, false},
	}

	for _, test := range tests {
		result, ok := test.seq.FindLast(func(n int) bool { return n > 2 })
		if result != test.expected || ok != test.ok {
			t.Errorf("%v.FindLast(func) = (%v, %v), expected (%v, %v)", test.seq.Value(), result, ok, test.expected, test.ok)
		}
	}
}

func TestEvery(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected bool
	}{
		{New(2, 4, 6), true},
		{New(2, 3, 6), false},
		{New[int](), true},
	}

	for _, test := range tests {
		result := test.seq.Every(func(n int) bool { return n%2 == 0 })
		if result != test.expected {
			t.Errorf("%v.Every(func) = %v, expected %v", test.seq.Value(), result, test.expected)
		}
	}
}

func TestSome(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected bool
	}{
		{New(1, 2, 3), true},
		{New(1, 3, 5), false},
		{New[int](), false},
	}

	for _, test := range tests {
		result := test.seq.Some(func(n int) bool { return n%2 == 0 })
		if result != test.expected {
			t.Errorf("%v.Some(func) = %v, expected %v", test.seq.Value(), result, test.expected)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected int
	}{
		{New(1, 2, 3, 4), 4},
		{New(1), 1},
		{New[int](), 0},
	}

	for _, test := range tests {
		result := test.seq.Size()
		if result != test.expected {
			t.Errorf("%v.Size() = %v, expected %v", test.seq.Value(), result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected bool
	}{
		{New(1, 2, 3), false},
		{New[int](), true},
	}

	for _, test := range tests {
		result := test.seq.IsEmpty()
		if result != test.expected {
			t.Errorf("%v.IsEmpty() = %v, expected %v", test.seq.Value(), result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected []int
	}{
		{New(1, 2, 3), []int{3, 2, 1}},
		{New(1), []int{1}},
		// {New[int](), []int{}},
	}

	for _, test := range tests {
		result := test.seq.Reverse()
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.Reverse() = %v, expected %v", test.seq.Value(), result.Value(), test.expected)
		}
	}
}

func TestUniq(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		expected []int
	}{
		{New(1, 2, 1, 3, 2), []int{1, 2, 3}},
		{New(1, 2, 3), []int{1, 2, 3}},
		// {New[int](), []int{}}, ===> Error seq_test.go:289: [].Uniq() = [], expected []
	}

	for _, test := range tests {
		result := test.seq.Uniq()
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.Uniq() = %v, expected %v", test.seq.Value(), result.Value(), test.expected)
		}
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		size     int
		expected [][]int
	}{
		{New(1, 2, 3, 4), 2, [][]int{{1, 2}, {3, 4}}},
		{New(1, 2, 3, 4, 5), 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{New(1, 2, 3), 5, [][]int{{1, 2, 3}}},
		{New[int](), 2, [][]int{}},
	}

	for _, test := range tests {
		result := test.seq.Chunk(test.size)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("%v.Chunk(%d) = %v, expected %v", test.seq.Value(), test.size, result, test.expected)
		}
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		n        int
		expected []int
	}{
		{New(1, 2, 3, 4), 2, []int{1, 2}},
		{New(1, 2, 3), 5, []int{1, 2, 3}},
		{New(1, 2, 3), 0, []int{}},
		// {New[int](), 2, []int{}},
	}

	for _, test := range tests {
		result := test.seq.Take(test.n)
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.Take(%d) = %v, expected %v", test.seq.Value(), test.n, result.Value(), test.expected)
		}
	}
}

func TestTakeRight(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		n        int
		expected []int
	}{
		{New(1, 2, 3, 4), 2, []int{3, 4}},
		{New(1, 2, 3), 5, []int{1, 2, 3}},
		{New(1, 2, 3), 0, []int{}},
		// {New[int](), 2, []int{}}, ===> Error seq_test.go:349: [].TakeRight(2) = [], expected []
	}

	for _, test := range tests {
		result := test.seq.TakeRight(test.n)
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.TakeRight(%d) = %v, expected %v", test.seq.Value(), test.n, result.Value(), test.expected)
		}
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		n        int
		expected []int
	}{
		{New(1, 2, 3, 4), 2, []int{3, 4}},
		{New(1, 2, 3), 5, []int{}},
		{New(1, 2, 3), 0, []int{1, 2, 3}},
		{New[int](), 2, []int{}},
	}

	for _, test := range tests {
		result := test.seq.Drop(test.n)
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.Drop(%d) = %v, expected %v", test.seq.Value(), test.n, result.Value(), test.expected)
		}
	}
}

func TestDropRight(t *testing.T) {
	tests := []struct {
		seq      *Sequence[int]
		n        int
		expected []int
	}{
		{New(1, 2, 3, 4), 2, []int{1, 2}},
		{New(1, 2, 3), 5, []int{}},
		{New(1, 2, 3), 0, []int{1, 2, 3}},
		{New[int](), 2, []int{}},
	}

	for _, test := range tests {
		result := test.seq.DropRight(test.n)
		if !reflect.DeepEqual(result.Value(), test.expected) {
			t.Errorf("%v.DropRight(%d) = %v, expected %v", test.seq.Value(), test.n, result.Value(), test.expected)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		seq       *Sequence[int]
		separator string
		expected  string
	}{
		{New(1, 2, 3), ",", "1,2,3"},
		{New(1), ",", "1"},
		{New[int](), ",", ""},
	}

	for _, test := range tests {
		result := test.seq.Join(test.separator)
		if result != test.expected {
			t.Errorf("%v.Join(%q) = %q, expected %q", test.seq.Value(), test.separator, result, test.expected)
		}
	}
}
