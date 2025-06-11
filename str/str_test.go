package str

import (
	"fmt"
	"testing"
)

func TestCamelcase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo bar", "fooBar"},
		{"Foo Bar", "fooBar"},
		{"foo bar baz", "fooBarBaz"},
		{"", ""},
		{"foo", "foo"},
	}

	for _, test := range tests {
		result := Camelcase(test.input)
		if result != test.expected {
			t.Errorf("Camelcase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"fred", "Fred"},
		{"FRED", "FRED"},
		{"fred flintstone", "Fred flintstone"},
		{"", ""},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		input    string
		target   string
		expected bool
	}{
		{"abc", "c", true},
		{"abc", "bc", true},
		{"abc", "abc", true},
		{"abc", "d", false},
		{"abc", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		result := EndsWith(test.input, test.target)
		if result != test.expected {
			t.Errorf("EndsWith(%q, %q) = %v, expected %v", test.input, test.target, result, test.expected)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		input    string
		target   string
		expected bool
	}{
		{"abc", "a", true},
		{"abc", "ab", true},
		{"abc", "abc", true},
		{"abc", "d", false},
		{"abc", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		result := StartsWith(test.input, test.target)
		if result != test.expected {
			t.Errorf("StartsWith(%q, %q) = %v, expected %v", test.input, test.target, result, test.expected)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"  abc  ", "", "abc"},
		{"-_-abc-_-", "-_", "abc"},
		{"abc", "", "abc"},
		{"", "", ""},
	}

	for _, test := range tests {
		var result string
		if test.cutset == "" {
			result = Trim(test.input)
		} else {
			result = Trim(test.input, test.cutset)
		}
		if result != test.expected {
			t.Errorf("Trim(%q, %q) = %q, expected %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"FRED", "fred"},
		{"Fred", "fred"},
		{"fred", "fred"},
		{"", ""},
	}

	for _, test := range tests {
		result := ToLower(test.input)
		if result != test.expected {
			t.Errorf("ToLower(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"fred", "FRED"},
		{"Fred", "FRED"},
		{"FRED", "FRED"},
		{"", ""},
	}

	for _, test := range tests {
		result := ToUpper(test.input)
		if result != test.expected {
			t.Errorf("ToUpper(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input     string
		separator string
		expected  []string
	}{
		{"a-b-c", "-", []string{"a", "b", "c"}},
		{"a", "-", []string{"a"}},
		{"", "-", []string{""}},
	}

	for _, test := range tests {
		result := Split(test.input, test.separator)
		if len(result) != len(test.expected) {
			t.Errorf("Split(%q, %q) = %v, expected %v", test.input, test.separator, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Split(%q, %q)[%d] = %q, expected %q", test.input, test.separator, i, result[i], test.expected[i])
			}
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input     []string
		separator string
		expected  string
	}{
		{[]string{"a", "b", "c"}, "-", "a-b-c"},
		{[]string{"a"}, "-", "a"},
		{[]string{}, "-", ""},
	}

	for _, test := range tests {
		result := Join(test.input, test.separator)
		if result != test.expected {
			t.Errorf("Join(%v, %q) = %q, expected %q", test.input, test.separator, result, test.expected)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected string
	}{
		{"abc", 2, "abcabc"},
		{"abc", 0, ""},
		{"", 5, ""},
	}

	for _, test := range tests {
		result := Repeat(test.input, test.n)
		if result != test.expected {
			t.Errorf("Repeat(%q, %d) = %q, expected %q", test.input, test.n, result, test.expected)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		input    string
		old      string
		new      string
		expected string
	}{
		{"Hi Fred", "Fred", "Barney", "Hi Barney"},
		{"abc", "d", "e", "abc"},
		{"", "a", "b", ""},
	}

	for _, test := range tests {
		result := Replace(test.input, test.old, test.new)
		if result != test.expected {
			t.Errorf("Replace(%q, %q, %q) = %q, expected %q", test.input, test.old, test.new, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"abc", "b", true},
		{"abc", "d", false},
		{"abc", "", true},
		{"", "", true},
	}

	for _, test := range tests {
		result := Contains(test.input, test.substr)
		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %v, expected %v", test.input, test.substr, result, test.expected)
		}
	}
}

func TestEllipsis(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		// No truncation cases
		{"Hello", 10, "Hello"},
		{"Hello", 5, "Hello"},
		{"", 5, ""},
		{"   ", 5, ""}, // Trimmed to empty string

		// ASCII truncation cases
		{"Hello, World", 5, "Hello..."},
		{"Hello, World", 7, "Hello, ..."},
		{"Hello", 2, "He..."},

		// UTF-8 character handling
		{"Hello, 世界", 8, "Hello, ..."}, // '世' is 3 bytes, so we can't include it
		{"你好, World", 6, "你好..."},      // '你' and '好' are each 3 bytes (total 6 bytes)
		{"你好, World", 3, "你..."},       // Only room for one character
		{"你好", 2, "..."},               // Not enough room for even one character

		// Edge cases
		{"Hello", 0, "..."},
		{"Hello", 1, "H..."},
	}

	for _, test := range tests {
		result := Ellipsis(test.input, test.length)
		if result != test.expected {
			t.Errorf("Ellipsis(%q, %d) = %q, expected %q", test.input, test.length, result, test.expected)
		}
	}
}

func TestIsValidWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Simple word", "hello", true},
		{"Word with digits", "hello123", true},
		{"Word with ASCII apostrophe", "don't", true},
		{"Word with Unicode apostrophe", "don\u2019t", true}, // Unicode right single quotation mark
		{"Word with invalid character", "hello!", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isValidWord(test.input)
			if result != test.expected {
				t.Errorf("isValidWord(%q) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}

func TestEllipsisOnly(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		// No truncation cases
		{"Hello", 10, "Hello"},
		{"Hello", 5, "Hello"},
		{"", 5, ""},
		{"   ", 5, ""}, // Trimmed to empty string

		// ASCII truncation cases
		{"Hello, World", 5, "Hello..."},
		{"Hello, World", 7, "Hello, ..."},
		{"Hello", 2, "He..."},

		// UTF-8 character handling
		{"Hello, 世界", 8, "Hello, ..."}, // '世' is 3 bytes, so we can't include it
		{"你好, World", 6, "你好..."},      // '你' and '好' are each 3 bytes (total 6 bytes)
		{"你好, World", 3, "你..."},       // Only room for one character
		{"你好", 2, "..."},               // Not enough room for even one character

		// Edge cases
		{"Hello", 0, "..."},
		{"Hello", 1, "H..."},
	}

	for _, test := range tests {
		result := Ellipsis(test.input, test.length)
		if result != test.expected {
			t.Errorf("Ellipsis(%q, %d) = %q, expected %q", test.input, test.length, result, test.expected)
		} else {
			fmt.Printf("PASS: Ellipsis(%q, %d) = %q\n", test.input, test.length, result)
		}
	}
}
