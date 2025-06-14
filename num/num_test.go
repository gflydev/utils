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
			15,
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
			5.00000,
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

func TestCeiling(t *testing.T) {
	tests := []struct {
		n         float64
		precision []int
		expected  float64
	}{
		{4.3, []int{}, 5},
		{4.7, []int{}, 5},
		{-4.3, []int{}, -4},
		{-4.7, []int{}, -4},
		{4.357, []int{2}, 4.36},
		{4.351, []int{2}, 4.36},
		{-4.357, []int{2}, -4.35},
		{-4.351, []int{2}, -4.35},
		{4.357, []int{1}, 4.4},
		{4.351, []int{1}, 4.4},
	}

	for _, test := range tests {
		result := Ceiling(test.n, test.precision...)
		if result != test.expected {
			t.Errorf("Ceiling(%f, %v) = %f, expected %f", test.n, test.precision, result, test.expected)
		}
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		number             float64
		decimals           int
		decimalSeparator   string
		thousandsSeparator string
		expected           string
	}{
		{1234.5678, 2, ".", ",", "1,234.57"},
		{1234567.89, 1, ",", " ", "1 234 567,9"},
		{1000000, 0, ".", ",", "1,000,000"},
		{-1234.5678, 2, ".", ",", "-1,234.57"},
		{0.5678, 3, ".", ",", "0.568"},
		{0, 2, ".", ",", "0.00"},
	}

	for _, test := range tests {
		result := Format(test.number, test.decimals, test.decimalSeparator, test.thousandsSeparator)
		if result != test.expected {
			t.Errorf("Format(%f, %d, %q, %q) = %q, expected %q",
				test.number, test.decimals, test.decimalSeparator, test.thousandsSeparator, result, test.expected)
		}
	}
}

func TestFormatPercentage(t *testing.T) {
	tests := []struct {
		number   float64
		decimals int
		expected string
	}{
		{0.156, 1, "15.6%"},
		{0.5, 0, "50%"},
		{1, 2, "100.00%"},
		{0, 1, "0.0%"},
		{-0.25, 0, "-25%"},
		{1.5, 1, "150.0%"},
	}

	for _, test := range tests {
		result := FormatPercentage(test.number, test.decimals)
		if result != test.expected {
			t.Errorf("FormatPercentage(%f, %d) = %q, expected %q", test.number, test.decimals, result, test.expected)
		}
	}
}

func TestPercent(t *testing.T) {
	tests := []struct {
		number   float64
		total    float64
		decimals []int
		expected float64
	}{
		{25, 100, []int{}, 25.0},
		{1, 3, []int{2}, 33.33},
		{1, 0, []int{}, 0},
		{0, 100, []int{}, 0},
		{50, 200, []int{1}, 25.0},
		{200, 50, []int{}, 400},
		{-25, 100, []int{}, -25},
		{25, -100, []int{}, -25},
	}

	for _, test := range tests {
		result := Percent(test.number, test.total, test.decimals...)
		if result != test.expected {
			t.Errorf("Percent(%f, %f, %v) = %f, expected %f", test.number, test.total, test.decimals, result, test.expected)
		}
	}
}

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		number    float64
		precision []int
		expected  string
	}{
		{1000, []int{}, "1K"},
		{489939, []int{}, "490K"},
		{1230000, []int{2}, "1.23M"},
		{1000000000, []int{}, "1B"},
		{1500000000000, []int{1}, "1.5T"},
		{999, []int{}, "999"},
		{-1234567, []int{1}, "-1.2M"},
		{0, []int{}, "0"},
		{1500, []int{2}, "1.50K"},
		{1500000, []int{3}, "1.500M"},
		{1500000000, []int{0}, "2B"},
		{1234567890123, []int{2}, "1.23T"},
		{1234567, []int{1}, "1.2M"},
		{1234, []int{2}, "1.23K"},
		{1000000000, []int{1}, "1.0B"},
		{999, []int{1}, "999.0"},
		{-1234567, []int{1}, "-1.2M"},
		{0, []int{1}, "0.0"},
		{1500, []int{0}, "2K"},
		{1500000, []int{0}, "2M"},
		{1500000000, []int{0}, "2B"},
	}

	for _, test := range tests {
		result := Abbreviate(test.number, test.precision...)
		if result != test.expected {
			t.Errorf("Abbreviate(%f, %v) = %q, expected %q", test.number, test.precision, result, test.expected)
		}
	}
}

func TestCurrencySymbol(t *testing.T) {
	tests := []struct {
		code     string
		expected string
	}{
		{"USD", "$"},
		{"EUR", "€"},
		{"GBP", "£"},
		{"JPY", "¥"},
		{"CNY", "¥"},
		{"INR", "₹"},
		{"RUB", "₽"},
		{"BRL", "R$"},
		{"KRW", "₩"},
		{"AUD", "A$"},
		{"CAD", "C$"},
		{"CHF", "CHF"},
		{"HKD", "HK$"},
		{"SGD", "S$"},
		{"SEK", "kr"},
		{"NOK", "kr"},
		{"DKK", "kr"},
		{"PLN", "zł"},
		{"THB", "฿"},
		{"MXN", "Mex$"},
		{"ZAR", "R"},
		{"XYZ", "XYZ"}, // Unknown currency code should return the code itself
	}

	for _, test := range tests {
		result := CurrencySymbol(test.code)
		if result != test.expected {
			t.Errorf("CurrencySymbol(%q) = %q, expected %q", test.code, result, test.expected)
		}
	}
}

func TestGetLocaleInfo(t *testing.T) {
	tests := []struct {
		locale   string
		expected LocaleInfo
	}{
		{"en", LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"}},
		{"de", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "suffix"}},
		{"fr", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: " ", SymbolPosition: "suffix"}},
		{"es", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "suffix"}},
		{"it", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "suffix"}},
		{"nl", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "prefix"}},
		{"pt", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "prefix"}},
		{"ru", LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: " ", SymbolPosition: "suffix"}},
		{"ja", LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"}},
		{"zh", LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"}},
		{"xx", LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"}}, // Unknown locale should default to English
	}

	for _, test := range tests {
		result := GetLocaleInfo(test.locale)
		if result != test.expected {
			t.Errorf("GetLocaleInfo(%q) = %+v, expected %+v", test.locale, result, test.expected)
		}
	}
}

func TestFileSize(t *testing.T) {
	tests := []struct {
		bytes     float64
		precision []int
		expected  string
	}{
		// Examples from the issue description
		{1024, []int{}, "1 KB"},
		{1024 * 1024, []int{}, "1 MB"},
		{1024, []int{2}, "1.00 KB"},

		// Additional test cases
		{0, []int{}, "0 B"},
		{500, []int{}, "500 B"},
		{1500, []int{}, "1 KB"},
		{1500, []int{2}, "1.46 KB"},
		{1500000, []int{}, "1 MB"},
		{1500000, []int{1}, "1.4 MB"},
		{1073741824, []int{}, "1 GB"},       // 1 GB (1024^3)
		{1099511627776, []int{}, "1 TB"},    // 1 TB (1024^4)
		{1125899906842624, []int{}, "1 PB"}, // 1 PB (1024^5)
		{2000000, []int{2}, "1.91 MB"},
		{-1024, []int{}, "-1 KB"},
		{-1500, []int{2}, "-1.46 KB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 2, []int{1}, "2.0 PB"}, // 2 PB
		{1023, []int{}, "1023 B"},                                  // Just below 1 KB
		{1048575, []int{}, "1024 KB"},                              // Just below 1 MB
	}

	for _, test := range tests {
		result := FileSize(test.bytes, test.precision...)
		if result != test.expected {
			t.Errorf("FileSize(%f, %v) = %q, expected %q", test.bytes, test.precision, result, test.expected)
		}
	}
}

func TestForHumans(t *testing.T) {
	tests := []struct {
		number    float64
		precision []int
		expected  string
	}{
		// Examples from the issue description
		{1000, []int{}, "1 thousand"},
		{489939, []int{}, "490 thousand"},
		{1230000, []int{2}, "1.23 million"},

		// Additional test cases
		{0, []int{}, "0"},
		{999, []int{}, "999"},
		{1000000000, []int{}, "1 billion"},
		{1500000000000, []int{1}, "1.5 trillion"},
		{-1234567, []int{1}, "-1.2 million"},
		{1500, []int{2}, "1.50 thousand"},
		{1500000, []int{3}, "1.500 million"},
		{1500000000, []int{0}, "2 billion"},
		{1234567890123, []int{2}, "1.23 trillion"},
		{1000000000000000, []int{}, "1 quadrillion"},
		{1000000000000000000, []int{}, "1 quintillion"},
	}

	for _, test := range tests {
		result := ForHumans(test.number, test.precision...)
		if result != test.expected {
			t.Errorf("ForHumans(%f, %v) = %q, expected %q", test.number, test.precision, result, test.expected)
		}
	}
}

func TestCurrency(t *testing.T) {
	tests := []struct {
		number   float64
		options  map[string]interface{}
		expected string
	}{
		// Examples from the issue description
		{1000, nil, "$1,000.00"},
		{1000, map[string]interface{}{"in": "EUR"}, "€1,000.00"},
		{1000, map[string]interface{}{"in": "EUR", "locale": "de"}, "1.000,00 €"},
		{1000, map[string]interface{}{"in": "EUR", "locale": "de", "precision": 0}, "1.000 €"},

		// Additional test cases
		{1234.56, nil, "$1,234.56"},
		{1234.56, map[string]interface{}{"in": "USD", "locale": "en"}, "$1,234.56"},
		{1234.56, map[string]interface{}{"in": "EUR", "locale": "fr"}, "1 234,56 €"},
		{1234.56, map[string]interface{}{"in": "GBP"}, "£1,234.56"},
		{1234.56, map[string]interface{}{"in": "JPY", "precision": 0}, "¥1,235"},
		{0, nil, "$0.00"},
		{-1234.56, nil, "-$1,234.56"},
		{-1234.56, map[string]interface{}{"in": "EUR", "locale": "de"}, "-1.234,56 €"},
		{1000000, nil, "$1,000,000.00"},
		{1000000, map[string]interface{}{"in": "EUR", "locale": "fr"}, "1 000 000,00 €"},
		{0.5, map[string]interface{}{"precision": 3}, "$0.500"},

		// Edge cases
		{1234.56, map[string]interface{}{"invalid": "option"}, "$1,234.56"},    // Invalid option should be ignored
		{1234.56, map[string]interface{}{"in": ""}, "$1,234.56"},               // Empty currency code should default to USD
		{1234.56, map[string]interface{}{"locale": ""}, "$1,234.56"},           // Empty locale should default to en
		{1234.56, map[string]interface{}{"precision": "invalid"}, "$1,234.56"}, // Invalid precision should default to 2
		{1234.56, map[string]interface{}{"precision": -1}, "$1,234.56"},        // Negative precision should default to 2
	}

	for _, test := range tests {
		var result string
		if test.options == nil {
			result = Currency(test.number)
		} else {
			result = Currency(test.number, test.options)
		}
		if result != test.expected {
			t.Errorf("Currency(%f, %v) = %q, expected %q", test.number, test.options, result, test.expected)
		}
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		// Examples from the issue description
		{1, "1st"},
		{2, "2nd"},
		{21, "21st"},

		// Additional test cases
		{3, "3rd"},
		{4, "4th"},
		{11, "11th"},
		{12, "12th"},
		{13, "13th"},
		{22, "22nd"},
		{23, "23rd"},
		{24, "24th"},
		{101, "101st"},
		{102, "102nd"},
		{103, "103rd"},
		{111, "111th"},
		{112, "112th"},
		{113, "113th"},
		{0, "0th"},
		{-1, "1st"}, // Negative numbers are handled by taking the absolute value
		{-2, "2nd"},
		{-3, "3rd"},
		{-11, "11th"},
		{-21, "21st"},
	}

	for _, test := range tests {
		result := Ordinal(test.number)
		if result != test.expected {
			t.Errorf("Ordinal(%d) = %q, expected %q", test.number, result, test.expected)
		}
	}
}

func TestPairs(t *testing.T) {
	tests := []struct {
		total     int
		chunkSize int
		options   []map[string]int
		expected  [][]int
	}{
		// Examples from the issue description
		{25, 10, nil, [][]int{{0, 9}, {10, 19}, {20, 25}}},
		{25, 10, []map[string]int{{"offset": 0}}, [][]int{{0, 10}, {10, 20}, {20, 25}}},

		// Additional test cases
		{10, 5, nil, [][]int{{0, 4}, {5, 9}, {10, 10}}},
		{10, 5, []map[string]int{{"offset": 0}}, [][]int{{0, 5}, {5, 10}}},
		{0, 5, nil, nil}, // Changed from [][]int{} to nil to match the actual return value
		{5, 10, nil, [][]int{{0, 5}}},
		{15, 5, []map[string]int{{"offset": 1}}, [][]int{{0, 6}, {5, 11}, {10, 15}}},
		{100, 25, nil, [][]int{{0, 24}, {25, 49}, {50, 74}, {75, 99}, {100, 100}}},
	}

	for _, test := range tests {
		var result [][]int
		if test.options == nil {
			result = Pairs(test.total, test.chunkSize)
		} else {
			result = Pairs(test.total, test.chunkSize, test.options...)
		}

		// Special case for empty slices
		if test.total == 0 {
			if len(result) != 0 {
				t.Errorf("Pairs(%d, %d, %v) = %v, expected empty slice",
					test.total, test.chunkSize, test.options, result)
			}
			continue
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Pairs(%d, %d, %v) = %v, expected %v",
				test.total, test.chunkSize, test.options, result, test.expected)
		}
	}
}
