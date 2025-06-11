package num

import (
	"math"
	"reflect"
	"testing"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		n        float64
		lower    float64
		upper    float64
		expected float64
	}{
		{5, 0, 10, 5},
		{-5, 0, 10, 0},
		{15, 0, 10, 10},
		{5, 10, 0, 5}, // lower > upper, should swap
	}

	for _, test := range tests {
		result := Clamp(test.n, test.lower, test.upper)
		if result != test.expected {
			t.Errorf("Clamp(%f, %f, %f) = %f, expected %f", test.n, test.lower, test.upper, result, test.expected)
		}
	}
}

func TestInRange(t *testing.T) {
	tests := []struct {
		n        float64
		start    float64
		end      float64
		expected bool
	}{
		{3, 2, 4, true},
		{2, 2, 4, true},
		{4, 2, 4, true},
		{1, 2, 4, false},
		{5, 2, 4, false},
		{3, 4, 2, true}, // start > end, should swap
	}

	for _, test := range tests {
		result := InRange(test.n, test.start, test.end)
		if result != test.expected {
			t.Errorf("InRange(%f, %f, %f) = %v, expected %v", test.n, test.start, test.end, result, test.expected)
		}
	}
}

func TestRandom(t *testing.T) {
	minV, maxV := 1, 10
	for i := 0; i < 100; i++ {
		result := Random(minV, maxV)
		if result < minV || result > maxV {
			t.Errorf("Random(%d, %d) = %d, which is outside the range [%d, %d]", minV, maxV, result, minV, maxV)
		}
	}
}

func TestRound(t *testing.T) {
	tests := []struct {
		n        float64
		expected float64
	}{
		{4.7, 5},
		{4.3, 4},
		{4.5, 5},
		{-4.7, -5},
		{-4.3, -4},
	}

	for _, test := range tests {
		result := Round(test.n)
		if result != test.expected {
			t.Errorf("Round(%f) = %f, expected %f", test.n, result, test.expected)
		}
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		n        float64
		expected float64
	}{
		{4.7, 4},
		{4.3, 4},
		{4.0, 4},
		{-4.7, -5},
		{-4.3, -5},
	}

	for _, test := range tests {
		result := Floor(test.n)
		if result != test.expected {
			t.Errorf("Floor(%f) = %f, expected %f", test.n, result, test.expected)
		}
	}
}

func TestCeil(t *testing.T) {
	tests := []struct {
		n        float64
		expected float64
	}{
		{4.7, 5},
		{4.3, 5},
		{4.0, 4},
		{-4.7, -4},
		{-4.3, -4},
	}

	for _, test := range tests {
		result := Ceil(test.n)
		if result != test.expected {
			t.Errorf("Ceil(%f) = %f, expected %f", test.n, result, test.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3}, 3},
		{[]float64{3, 2, 1}, 3},
		{[]float64{-1, -2, -3}, -1},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Max(test.numbers...)
		if result != test.expected {
			t.Errorf("Max(%v) = %f, expected %f", test.numbers, result, test.expected)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3}, 1},
		{[]float64{3, 2, 1}, 1},
		{[]float64{-1, -2, -3}, -3},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Min(test.numbers...)
		if result != test.expected {
			t.Errorf("Min(%v) = %f, expected %f", test.numbers, result, test.expected)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3}, 6},
		{[]float64{-1, -2, -3}, -6},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Sum(test.numbers...)
		if result != test.expected {
			t.Errorf("Sum(%v) = %f, expected %f", test.numbers, result, test.expected)
		}
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 3, 5, 7}, 4},
		{[]float64{}, 0},
	}

	for _, test := range tests {
		result := Mean(test.numbers...)
		if result != test.expected {
			t.Errorf("Mean(%v) = %f, expected %f", test.numbers, result, test.expected)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		n        float64
		expected float64
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
	}

	for _, test := range tests {
		result := Abs(test.n)
		if result != test.expected {
			t.Errorf("Abs(%f) = %f, expected %f", test.n, result, test.expected)
		}
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		base     float64
		exponent float64
		expected float64
	}{
		{2, 3, 8},
		{3, 2, 9},
		{2, 0, 1},
		{0, 2, 0},
		{-2, 2, 4},
		{-2, 3, -8},
	}

	for _, test := range tests {
		result := Pow(test.base, test.exponent)
		if result != test.expected {
			t.Errorf("Pow(%f, %f) = %f, expected %f", test.base, test.exponent, result, test.expected)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		n        float64
		expected float64
	}{
		{9, 3},
		{4, 2},
		{2, math.Sqrt(2)},
		{0, 0},
	}

	for _, test := range tests {
		result := Sqrt(test.n)
		if result != test.expected {
			t.Errorf("Sqrt(%f) = %f, expected %f", test.n, result, test.expected)
		}
	}
}

func TestRoundWithPrecision(t *testing.T) {
	tests := []struct {
		n         float64
		precision int
		expected  float64
	}{
		{4.7, 0, 5},
		{4.7, 1, 4.7},
		{4.75, 1, 4.8},
		{4.749, 2, 4.75},
		{-4.7, 0, -5},
		{-4.7, 1, -4.7},
	}

	for _, test := range tests {
		result := Round(test.n, test.precision)
		if result != test.expected {
			t.Errorf("Round(%f, %d) = %f, expected %f", test.n, test.precision, result, test.expected)
		}
	}
}

func TestFloorWithPrecision(t *testing.T) {
	tests := []struct {
		n         float64
		precision int
		expected  float64
	}{
		{4.7, 0, 4},
		{4.78, 1, 4.7},
		{4.753, 2, 4.75},
		{-4.7, 0, -5},
		{-4.78, 1, -4.8},
	}

	for _, test := range tests {
		result := Floor(test.n, test.precision)
		if result != test.expected {
			t.Errorf("Floor(%f, %d) = %f, expected %f", test.n, test.precision, result, test.expected)
		}
	}
}

func TestCeilWithPrecision(t *testing.T) {
	tests := []struct {
		n         float64
		precision int
		expected  float64
	}{
		{4.3, 0, 5},
		{4.78, 1, 4.8},
		{4.753, 2, 4.76},
		{-4.3, 0, -4},
		{-4.78, 1, -4.7},
	}

	for _, test := range tests {
		result := Ceil(test.n, test.precision)
		if result != test.expected {
			t.Errorf("Ceil(%f, %d) = %f, expected %f", test.n, test.precision, result, test.expected)
		}
	}
}

func TestMaxBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		collection []person
		iteratee   func(person) float64
		expected   person
	}{
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(p.age) },
			person{"Bob", 30},
		},
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(len(p.name)) },
			person{"Charlie", 20},
		},
		{
			[]person{},
			func(p person) float64 { return float64(p.age) },
			person{},
		},
	}

	for _, test := range tests {
		result := MaxBy(test.collection, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MaxBy(%v, func) = %v, expected %v", test.collection, result, test.expected)
		}
	}

	// Test with numbers
	numbers := []int{1, 2, 3, 4, 5}
	result := MaxBy(numbers, func(n int) float64 { return float64(n * n) })
	if result != 5 {
		t.Errorf("MaxBy(%v, func) = %v, expected %v", numbers, result, 5)
	}
}

func TestMinBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		collection []person
		iteratee   func(person) float64
		expected   person
	}{
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(p.age) },
			person{"Charlie", 20},
		},
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(len(p.name)) },
			person{"Bob", 30},
		},
		{
			[]person{},
			func(p person) float64 { return float64(p.age) },
			person{},
		},
	}

	for _, test := range tests {
		result := MinBy(test.collection, test.iteratee)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("MinBy(%v, func) = %v, expected %v", test.collection, result, test.expected)
		}
	}

	// Test with numbers
	numbers := []int{1, 2, 3, 4, 5}
	result := MinBy(numbers, func(n int) float64 { return float64(n * n) })
	if result != 1 {
		t.Errorf("MinBy(%v, func) = %v, expected %v", numbers, result, 1)
	}
}

func TestSumBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		collection []person
		iteratee   func(person) float64
		expected   float64
	}{
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(p.age) },
			75,
		},
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(len(p.name)) },
			16,
		},
		{
			[]person{},
			func(p person) float64 { return float64(p.age) },
			0,
		},
	}

	for _, test := range tests {
		result := SumBy(test.collection, test.iteratee)
		if result != test.expected {
			t.Errorf("SumBy(%v, func) = %f, expected %f", test.collection, result, test.expected)
		}
	}

	// Test with numbers
	numbers := []int{1, 2, 3, 4, 5}
	result := SumBy(numbers, func(n int) float64 { return float64(n * 2) })
	if result != 30 {
		t.Errorf("SumBy(%v, func) = %f, expected %f", numbers, result, 30.0)
	}
}

func TestMeanBy(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	tests := []struct {
		collection []person
		iteratee   func(person) float64
		expected   float64
	}{
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(p.age) },
			25,
		},
		{
			[]person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}},
			func(p person) float64 { return float64(len(p.name)) },
			5.333333333333333,
		},
		{
			[]person{},
			func(p person) float64 { return float64(p.age) },
			0,
		},
	}

	for _, test := range tests {
		result := MeanBy(test.collection, test.iteratee)
		if math.Abs(result-test.expected) > 0.000001 {
			t.Errorf("MeanBy(%v, func) = %f, expected %f", test.collection, result, test.expected)
		}
	}

	// Test with numbers
	numbers := []int{1, 2, 3, 4, 5}
	result := MeanBy(numbers, func(n int) float64 { return float64(n * 2) })
	if result != 6 {
		t.Errorf("MeanBy(%v, func) = %f, expected %f", numbers, result, 6.0)
	}
}
