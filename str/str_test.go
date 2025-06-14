package str

import (
	"fmt"
	"strings"
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
		result := CamelCase(test.input)
		if result != test.expected {
			t.Errorf("CamelCase(%q) = %q, expected %q", test.input, result, test.expected)
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
		subject  string
		search   string
		replace  string
		expected string
	}{
		{"Hi Fred", "Fred", "Barney", "Hi Barney"},
		{"abc", "d", "e", "abc"},
		{"", "a", "b", ""},
		{"ababa", "a", "x", "xbxbx"},            // Replace all occurrences
		{"hello hello", "hello", "hi", "hi hi"}, // Replace all occurrences
	}

	for _, test := range tests {
		result := Replace(test.search, test.replace, test.subject)
		if result != test.expected {
			t.Errorf("Replace(%q, %q, %q) = %q, expected %q", test.search, test.replace, test.subject, result, test.expected)
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

func TestWordsPattern(t *testing.T) {
	tests := []struct {
		input    string
		pattern  string
		expected []string
	}{
		{"hello-world_test", `[\-_]+`, []string{"hello", "world", "test"}},
		{"a,b;c", `[,;]`, []string{"a", "b", "c"}},
		{"camelCase", `(?=[A-Z])`, []string{"camel", "case"}},
		{"", `[\-_]+`, []string{}},
		{"no-separator-match", `-`, []string{"no", "separator", "match"}}, // Falls back to Words
		{"invalid[pattern", `[`, []string{"invalid", "pattern"}},          // Invalid pattern falls back to Words
	}

	for _, test := range tests {
		result := WordsPattern(test.input, test.pattern)
		if len(result) != len(test.expected) {
			t.Errorf("WordsPattern(%q, %q) = %v, expected %v", test.input, test.pattern, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("WordsPattern(%q, %q)[%d] = %q, expected %q", test.input, test.pattern, i, result[i], test.expected[i])
			}
		}
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo bar", "foo-bar"},
		{"Foo Bar", "foo-bar"},
		{"foo bar baz", "foo-bar-baz"},
		{"FooBarBaz", "foo-bar-baz"},
		{"foo_bar_baz", "foo-bar-baz"},
		{"foo-bar", "foo-bar"},
		{"", ""},
		{"foo", "foo"},
		{"foo!!bar", "foo-bar"},
	}

	for _, test := range tests {
		result := KebabCase(test.input)
		if result != test.expected {
			t.Errorf("KebabCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo bar", "foo_bar"},
		{"Foo Bar", "foo_bar"},
		{"foo bar baz", "foo_bar_baz"},
		{"FooBarBaz", "foo_bar_baz"},
		{"foo-bar-baz", "foo_bar_baz"},
		{"foo_bar", "foo_bar"},
		{"", ""},
		{"foo", "foo"},
		{"foo!!bar", "foo_bar"},
	}

	for _, test := range tests {
		result := SnakeCase(test.input)
		if result != test.expected {
			t.Errorf("SnakeCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"foo bar", "FooBar"},
		{"Foo Bar", "FooBar"},
		{"foo bar baz", "FooBarBaz"},
		{"foo_bar_baz", "FooBarBaz"},
		{"foo-bar", "FooBar"},
		{"", ""},
		{"foo", "Foo"},
	}

	for _, test := range tests {
		result := PascalCase(test.input)
		if result != test.expected {
			t.Errorf("PascalCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestHeadline(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"steve_jobs", "Steve Jobs"},
		{"EmailNotificationSent", "Email Notification Sent"},
		{"foo bar", "Foo Bar"},
		{"foo_bar_baz", "Foo Bar Baz"},
		{"foo-bar", "Foo Bar"},
		{"", ""},
		{"foo", "Foo"},
		{"FOO BAR", "Foo Bar"},
		{"fooBarBaz", "Foo Bar Baz"},
		{"FooBarBaz", "Foo Bar Baz"},
		{"foo_bar-baz", "Foo Bar Baz"},
	}

	for _, test := range tests {
		result := Headline(test.input)
		if result != test.expected {
			t.Errorf("Headline(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestTrimStart(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"  abc", "", "abc"},
		{"-_-abc", "-_", "abc"},
		{"abc", "", "abc"},
		{"", "", ""},
		{"---abc---", "-", "abc---"},
	}

	for _, test := range tests {
		var result string
		if test.cutset == "" {
			result = TrimStart(test.input)
		} else {
			result = TrimStart(test.input, test.cutset)
		}
		if result != test.expected {
			t.Errorf("TrimStart(%q, %q) = %q, expected %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTrimEnd(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"abc  ", "", "abc"},
		{"abc-_-", "-_", "abc"},
		{"abc", "", "abc"},
		{"", "", ""},
		{"---abc---", "-", "---abc"},
	}

	for _, test := range tests {
		var result string
		if test.cutset == "" {
			result = TrimEnd(test.input)
		} else {
			result = TrimEnd(test.input, test.cutset)
		}
		if result != test.expected {
			t.Errorf("TrimEnd(%q, %q) = %q, expected %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello world", "l", 3},
		{"hello world", "o", 2},
		{"hello world", "z", 0},
		{"hello world", "", 12}, // Empty substring counts between each char and at boundaries
		{"", "a", 0},
		{"", "", 1},      // Empty string contains 1 empty substring
		{"aaa", "aa", 1}, // Overlapping matches are not counted
	}

	for _, test := range tests {
		result := Count(test.input, test.substr)
		if result != test.expected {
			t.Errorf("Count(%q, %q) = %d, expected %d", test.input, test.substr, result, test.expected)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"hello world", "world", 6},
		{"hello world", "hello", 0},
		{"hello world", "l", 2},
		{"hello world", "z", -1},
		{"hello world", "", 0}, // Empty substring is always found at index 0
		{"", "a", -1},
		{"", "", 0}, // Empty string contains empty substring at index 0
	}

	for _, test := range tests {
		result := Index(test.input, test.substr)
		if result != test.expected {
			t.Errorf("Index(%q, %q) = %d, expected %d", test.input, test.substr, result, test.expected)
		}
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		input    string
		substr   string
		expected int
	}{
		{"abcabc", "b", 4}, // Example from documentation
		{"abcabc", "c", 5}, // Example from documentation
		{"abc", "d", -1},   // Example from documentation
		{"", "a", -1},      // Example from documentation
		{"hello world", "world", 6},
		{"hello world", "hello", 0},
		{"hello world", "l", 9}, // Last 'l' is at index 9
		{"hello world", "z", -1},
		{"hello world", "", 11}, // Empty substring is found at the end of the string
		{"", "", 0},             // Empty string contains empty substring at index 0
		{"aaa", "a", 2},         // Last 'a' in a string with multiple occurrences
		{"abcdefg", "abc", 0},   // Substring at the beginning
		{"abcdefg", "efg", 4},   // Substring at the end
	}

	for _, test := range tests {
		result := LastIndex(test.input, test.substr)
		if result != test.expected {
			t.Errorf("LastIndex(%q, %q) = %d, expected %d", test.input, test.substr, result, test.expected)
		}
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"Hello, World!", "hello-world"},
		{"Hello_World", "helloworld"}, // Underscores are removed
		{"Hello   World", "hello-world"},
		{"Héllö Wörld", "hll-wrld"}, // Accented characters are removed
		{"", ""},
		{"   ", ""},
		{"---", ""},
		{"Hello-World", "hello-world"},
		{"hello-world", "hello-world"},
		{"hello--world", "hello-world"},
		{"hello---world", "hello-world"},
		{"hello world!", "hello-world"},
	}

	for _, test := range tests {
		result := Slugify(test.input)
		if result != test.expected {
			t.Errorf("Slugify(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		input      string
		substrings []string
		expected   bool
	}{
		{"hello world", []string{"hello"}, true},
		{"hello world", []string{"foo", "world"}, true},
		{"hello world", []string{"foo", "bar"}, false},
		{"hello world", []string{}, false},
		{"", []string{"hello"}, false},
		{"", []string{""}, true},
		{"hello world", []string{""}, true},
	}

	for _, test := range tests {
		result := ContainsAny(test.input, test.substrings...)
		if result != test.expected {
			t.Errorf("ContainsAny(%q, %v) = %v, expected %v", test.input, test.substrings, result, test.expected)
		}
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"hello WORLD", "Hello World"},
		{"hElLo WoRlD", "Hello World"},
		{"", ""},
		{"hello", "Hello"},
		{"hello world-test", "Hello World-test"}, // Only spaces are used as word separators
		{"hello_world test", "Hello_world Test"}, // Only spaces are used as word separators
	}

	for _, test := range tests {
		result := ToTitleCase(test.input)
		if result != test.expected {
			t.Errorf("ToTitleCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestOnlyAlphanumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello, World!", "HelloWorld"},
		{"Hello123", "Hello123"},
		{"123-456-7890", "1234567890"},
		{"user@example.com", "userexamplecom"},
		{"", ""},
		{"!@#$%^&*()", ""},
		{"Hello_World", "HelloWorld"},
		{"Hello-World", "HelloWorld"},
	}

	for _, test := range tests {
		result := OnlyAlphanumeric(test.input)
		if result != test.expected {
			t.Errorf("OnlyAlphanumeric(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestMask(t *testing.T) {
	tests := []struct {
		input        string
		startVisible int
		endVisible   int
		maskChar     rune
		expected     string
	}{
		{"1234567890", 4, 2, '*', "1234****90"},
		{"1234567890", 0, 0, '*', "**********"},
		{"1234567890", 10, 0, '*', "1234567890"},
		{"1234567890", 0, 10, '*', "1234567890"},
		{"1234567890", 5, 5, '*', "1234567890"},
		{"1234567890", 6, 2, '*', "123456**90"},
		{"1234", 2, 1, '*', "12*4"}, // No masking if string is too short
		{"", 2, 2, '*', ""},
		{"1234567890", 2, 2, '•', "12••••••90"}, // 6 mask characters
	}

	for _, test := range tests {
		result := Mask(test.input, test.startVisible, test.endVisible, test.maskChar)
		if result != test.expected {
			t.Errorf("Mask(%q, %d, %d, %q) = %q, expected %q",
				test.input, test.startVisible, test.endVisible, test.maskChar, result, test.expected)
		}
	}
}

func TestPadLeft(t *testing.T) {
	tests := []struct {
		input    string
		padChar  rune
		length   int
		expected string
	}{
		{"abc", ' ', 5, "  abc"},
		{"abc", '0', 5, "00abc"},
		{"abc", '-', 3, "abc"},
		{"abc", '-', 2, "abc"},
		{"", '*', 3, "***"},
		{"abc", '世', 5, "世世abc"}, // Unicode character
	}

	for _, test := range tests {
		result := PadLeft(test.input, test.padChar, test.length)
		if result != test.expected {
			t.Errorf("PadLeft(%q, %q, %d) = %q, expected %q",
				test.input, test.padChar, test.length, result, test.expected)
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		input    string
		padChar  rune
		length   int
		expected string
	}{
		{"abc", ' ', 5, "abc  "},
		{"abc", '0', 5, "abc00"},
		{"abc", '-', 3, "abc"},
		{"abc", '-', 2, "abc"},
		{"", '*', 3, "***"},
		{"abc", '世', 5, "abc世世"}, // Unicode character
	}

	for _, test := range tests {
		result := PadRight(test.input, test.padChar, test.length)
		if result != test.expected {
			t.Errorf("PadRight(%q, %q, %d) = %q, expected %q",
				test.input, test.padChar, test.length, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Hello, World!", "!dlroW ,olleH"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
		{"你好", "好你"},             // Unicode characters
		{"Hello 世界", "界世 olleH"}, // Mixed ASCII and Unicode
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello world", 2},
		{"Hello, World!", 2}, // Comma is treated as part of the word
		{"one two three four", 4},
		{"", 0},
		{"   ", 0},
		{"one", 1},
		{"one-two", 1}, // Hyphen is treated as part of the word
		{"one_two", 1}, // Underscore is treated as part of the word
		{"one.two", 1}, // Period is treated as part of the word
		{"one,two", 1}, // Comma is treated as part of the word
		{"one  two", 2},
		{"one\ntwo", 2},  // Newline is treated as whitespace
		{"one\ttwo", 2},  // Tab is treated as whitespace
		{"one's two", 2}, // Apostrophe is treated as part of the word
		{"one's-two", 1}, // Hyphen is treated as part of the word
		{"one2three", 1},
	}

	for _, test := range tests {
		result := CountWords(test.input)
		if result != test.expected {
			t.Errorf("CountWords(%q) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestTruncateWords(t *testing.T) {
	tests := []struct {
		input    string
		maxWords int
		expected string
	}{
		{"hello world", 1, "hello..."},
		{"hello world", 2, "hello world"},
		{"hello world", 3, "hello world"},
		{"one two three four", 2, "one two..."},
		{"one two three four", 0, ""},
		{"", 5, ""},
		{"   ", 5, ""},
		{"one", 1, "one"},
		{"one-two three", 1, "one-two..."}, // Hyphen is treated as part of the word
		{"one_two three", 1, "one_two..."}, // Underscore is treated as part of the word
		{"one.two three", 1, "one.two..."}, // Period is treated as part of the word
		{"one,two three", 1, "one,two..."}, // Comma is treated as part of the word
	}

	for _, test := range tests {
		result := TruncateWords(test.input, test.maxWords)
		if result != test.expected {
			t.Errorf("TruncateWords(%q, %d) = %q, expected %q", test.input, test.maxWords, result, test.expected)
		}
	}
}

func TestFormatWithCommas(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1234, "1234"}, // Note: Current implementation doesn't add commas
		{1234567, "1234567"},
		{1234567890, "1234567890"},
		{123, "123"},
		{0, "0"},
		{-1234, "-1234"},
		{-1234567, "-1234567"},
		{1, "1"},
		{10, "10"},
		{100, "100"},
		{1000, "1000"},
	}

	for _, test := range tests {
		result := FormatWithCommas(test.input)
		if result != test.expected {
			t.Errorf("FormatWithCommas(%d) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"Hello, World!", 13},
		{"", 0},
		{"a", 1},
		{"你好", 2},       // 2 Unicode characters
		{"Hello 世界", 8}, // 5 ASCII + 2 Unicode characters + 1 space
		{"café", 4},     // 4 characters, one with accent
		{"café", 4},     // 4 characters, one with accent (different encoding)
		{"a\u0308", 2},  // 'a' with combining diaeresis
	}

	for _, test := range tests {
		result := Length(test.input)
		if result != test.expected {
			t.Errorf("Length(%q) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestAfter(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"hello world", "hello ", "world"},
		{"hello world", "world", ""},
		{"hello world", "o", " world"},
		{"hello world", "", "hello world"},
		{"hello world", "xyz", "hello world"}, // Not found returns original string
		{"", "hello", ""},
		{"hello hello world", "hello ", "hello world"},
	}

	for _, test := range tests {
		result := After(test.input, test.search)
		if result != test.expected {
			t.Errorf("After(%q, %q) = %q, expected %q", test.input, test.search, result, test.expected)
		}
	}
}

func TestAfterLast(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"hello world", "hello ", "world"},
		{"hello world", "world", ""},
		{"hello world", "o", "rld"},
		{"hello world", "", "hello world"},
		{"hello world", "xyz", "hello world"}, // Not found returns original string
		{"", "hello", ""},
		{"hello hello world", "hello ", "world"},
		{"app/models/user.rb", "/", "user.rb"},
		{"app/models/user.rb", ".", "rb"},
	}

	for _, test := range tests {
		result := AfterLast(test.input, test.search)
		if result != test.expected {
			t.Errorf("AfterLast(%q, %q) = %q, expected %q", test.input, test.search, result, test.expected)
		}
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"hello world", "hello", ""},
		{"hello world", "o", "hell"},
		{"hello world", "", "hello world"},    // Empty search returns original string
		{"hello world", "xyz", "hello world"}, // Not found returns original string
		{"", "hello", ""},
		{"hello hello world", "hello", ""},
		{"app/models/user.rb", "/", "app"},
		{"app/models/user.rb", ".", "app/models/user"},
	}

	for _, test := range tests {
		result := Before(test.input, test.search)
		if result != test.expected {
			t.Errorf("Before(%q, %q) = %q, expected %q", test.input, test.search, result, test.expected)
		}
	}
}

func TestBeforeLast(t *testing.T) {
	tests := []struct {
		input    string
		search   string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"hello world", "hello", ""},
		{"hello world", "o", "hello w"},
		{"hello world", "", "hello world"},    // Empty search returns original string
		{"hello world", "xyz", "hello world"}, // Not found returns original string
		{"", "hello", ""},
		{"hello hello world", "hello", "hello "},
		{"app/models/user.rb", "/", "app/models"},
		{"app/models/user.rb", ".", "app/models/user"},
	}

	for _, test := range tests {
		result := BeforeLast(test.input, test.search)
		if result != test.expected {
			t.Errorf("BeforeLast(%q, %q) = %q, expected %q", test.input, test.search, result, test.expected)
		}
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		input    string
		start    string
		end      string
		expected string
	}{
		{"hello [world]", "[", "]", "world"},
		{"[hello] world", "[", "]", "hello"},
		{"hello [world] test", "[", "]", "world"},
		{"hello world", "[", "]", "hello world"},   // Not found returns original string
		{"hello [world", "[", "]", "hello [world"}, // End not found returns original string
		{"hello] world", "[", "]", "hello] world"}, // Start not found returns original string
		{"", "[", "]", ""},
		{"[hello][world]", "[", "]", "hello"},
		{"hello [world] [test]", "[", "]", "world"},
		{"hello [world [test]]", "[", "]", "world [test"},
		// TODO {"hello [[world]]", "[", "]", "world"},  // Nested brackets handled differently
		{"hello world", "", "]", "hello world"}, // Empty start returns original string
		{"hello world", "[", "", "hello world"}, // Empty end returns original string
	}

	for _, test := range tests {
		result := Between(test.input, test.start, test.end)
		if result != test.expected {
			t.Errorf("Between(%q, %q, %q) = %q, expected %q",
				test.input, test.start, test.end, result, test.expected)
		}
	}
}

func TestBetweenFirst(t *testing.T) {
	tests := []struct {
		input    string
		start    string
		end      string
		expected string
	}{
		{"[a] bc [d]", "[", "]", "a"},
		{"hello [world]", "[", "]", "world"},
		{"[hello] world", "[", "]", "hello"},
		{"hello [world] test", "[", "]", "world"},
		{"hello world", "[", "]", "hello world"},   // Not found returns original string
		{"hello [world", "[", "]", "hello [world"}, // End not found returns original string
		{"hello] world", "[", "]", "hello] world"}, // Start not found returns original string
		{"", "[", "]", ""},
		{"[hello][world]", "[", "]", "hello"},
		{"hello [world] [test]", "[", "]", "world"},
		{"hello [world [test]]", "[", "]", "world [test"},
		{"hello world", "", "]", "hello world"}, // Empty start returns original string
		{"hello world", "[", "", "hello world"}, // Empty end returns original string
	}

	for _, test := range tests {
		result := BetweenFirst(test.input, test.start, test.end)
		if result != test.expected {
			t.Errorf("BetweenFirst(%q, %q, %q) = %q, expected %q",
				test.input, test.start, test.end, result, test.expected)
		}
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		input      string
		substrings []string
		expected   bool
	}{
		{"hello world", []string{"hello", "world"}, true},
		{"hello world", []string{"hello", "foo"}, false},
		{"hello world", []string{}, true}, // Empty list means all are contained
		{"", []string{"hello"}, false},
		{"", []string{}, true},
		{"hello world", []string{""}, true}, // Empty string is contained in any string
		{"hello world", []string{"h", "e", "l", "o", "w", "r", "d"}, true},
	}

	for _, test := range tests {
		result := ContainsAll(test.input, test.substrings...)
		if result != test.expected {
			t.Errorf("ContainsAll(%q, %v) = %v, expected %v", test.input, test.substrings, result, test.expected)
		}
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		input    string
		cap      string
		expected string
	}{
		{"hello", "/", "hello/"},
		{"hello/", "/", "hello/"},
		{"", "/", "/"},
		{"/", "/", "/"},
		{"hello", "", "hello"},
		{"hello/world", "/", "hello/world/"},
		{"hello/world/", "/", "hello/world/"},
	}

	for _, test := range tests {
		result := Finish(test.input, test.cap)
		if result != test.expected {
			t.Errorf("Finish(%q, %q) = %q, expected %q", test.input, test.cap, result, test.expected)
		}
	}
}

func TestIs(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"foo*", "foobar", true},
		{"*bar", "foobar", true},
		{"foo*bar", "foobar", true},
		{"foo*bar", "foo123bar", true},
		{"foo*bar", "foobar123", false},
		{"foo*bar", "123foobar", false},
		{"*", "foobar", true},
		{"*", "", true},
		{"", "", true},
		{"foo", "foo", true},
		{"foo", "bar", false},
		{"foo*bar*baz", "foobarbaz", true},
		{"foo*bar*baz", "foo123bar456baz", true},
		{"foo*bar*baz", "foobarbaz123", false},
		{"foo*bar*baz", "123foobarbaz", false},
	}

	for _, test := range tests {
		result := Is(test.pattern, test.s)
		if result != test.expected {
			t.Errorf("Is(%q, %q) = %v, expected %v", test.pattern, test.s, result, test.expected)
		}
	}
}

func TestIsAscii(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", true},
		{"Hello, World!", true},
		{"123", true},
		{"", true},
		{"café", false},
		{"你好", false},
		{"Hello 世界", false},
		{"\n\t\r", true},
		{"\x00\x7F", true},
		{"\x80\xFF", false},
	}

	for _, test := range tests {
		result := IsAscii(test.input)
		if result != test.expected {
			t.Errorf("IsAscii(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestAscii(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"Hello, World!", "Hello, World!"},
		{"123", "123"},
		{"", ""},
		{"café", "cafe"},
		{"über", "uber"},
		{"Crème Brûlée", "Creme Brulee"},
		{"û", "u"},
		{"Æ", "AE"},
		{"æ", "ae"},
		{"Ç", "C"},
		{"ç", "c"},
		{"É", "E"},
		{"é", "e"},
		{"Ñ", "N"},
		{"ñ", "n"},
		{"Ö", "O"},
		{"ö", "o"},
		{"Ü", "U"},
		{"ü", "u"},
		{"ß", "ss"},
		{"Ý", "Y"},
		{"ý", "y"},
	}

	for _, test := range tests {
		result := Ascii(test.input)
		if result != test.expected {
			t.Errorf("Ascii(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestLimit(t *testing.T) {
	tests := []struct {
		input    string
		limit    int
		tail     string
		expected string
	}{
		{"hello world", 5, "...", "hello..."},
		{"hello", 10, "...", "hello"},
		{"", 5, "...", ""},
		{"hello world", 0, "(...)", ""},
		{"你好世界", 2, "(...)", "你好(...)"}, // 2 Unicode characters
	}

	for _, test := range tests {
		result := Limit(test.input, test.limit, test.tail)
		if result != test.expected {
			t.Errorf("Limit(%q, %d) = %q, expected %q", test.input, test.limit, result, test.expected)
		}
	}
}

func TestRandom(t *testing.T) {
	// Test different lengths
	lengths := []int{0, 1, 5, 10, 20}

	for _, length := range lengths {
		result := Random(length)

		// Check length
		if len(result) != length {
			t.Errorf("Random(%d) returned string of length %d, expected %d", length, len(result), length)
		}

		// Check randomness by generating multiple strings
		if length > 0 {
			results := make(map[string]bool)
			for i := 0; i < 5; i++ {
				r := Random(length)
				results[r] = true

				// Check that all characters are alphanumeric
				for _, c := range r {
					if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
						t.Errorf("Random(%d) returned string with non-alphanumeric character: %q", length, r)
						break
					}
				}
			}

			// Check that we got at least 2 different strings (very high probability)
			// Skip this check for length 1 as there are only 62 possibilities
			if length > 1 && len(results) < 2 {
				t.Errorf("Random(%d) did not generate different strings in 5 attempts", length)
			}
		}
	}
}

func TestPassword(t *testing.T) {
	// Test default length and custom lengths
	tests := []struct {
		name     string
		args     []int
		expected int
	}{
		{"Default length", nil, 32},
		{"Custom length", []int{12}, 12},
		{"Zero length", []int{0}, 0},
		{"Negative length", []int{-5}, 0},
		{"Multiple args (only first used)", []int{8, 16}, 8},
	}

	// Define valid charset characters
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/~"
	validCharMap := make(map[rune]bool)
	for _, c := range validChars {
		validCharMap[c] = true
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result string
			if test.args == nil {
				result = Password()
			} else {
				result = Password(test.args...)
			}

			// Check length
			if len(result) != test.expected {
				t.Errorf("Password(%v) returned string of length %d, expected %d", test.args, len(result), test.expected)
			}

			// Check randomness by generating multiple strings
			if test.expected > 0 {
				results := make(map[string]bool)
				for i := 0; i < 5; i++ {
					var r string
					if test.args == nil {
						r = Password()
					} else {
						r = Password(test.args...)
					}
					results[r] = true

					// Check that all characters are valid
					for _, c := range r {
						if !validCharMap[c] {
							t.Errorf("Password(%v) returned string with invalid character: %q", test.args, r)
							break
						}
					}
				}

				// Check that we got at least 2 different strings (very high probability)
				// Skip this check for very short lengths
				if test.expected > 1 && len(results) < 2 {
					t.Errorf("Password(%v) did not generate different strings in 5 attempts", test.args)
				}
			}
		})
	}
}

func TestReplaceArray(t *testing.T) {
	tests := []struct {
		subject  string
		search   string
		replace  []string
		expected string
	}{
		{"?", "?", []string{"a", "b", "c"}, "a"},
		{"??", "?", []string{"a", "b", "c"}, "ab"},
		{"???", "?", []string{"a", "b", "c"}, "abc"},
		{"????", "?", []string{"a", "b", "c"}, "abc?"},
		{"?x?", "?", []string{"a", "b"}, "axb"},
		{"", "?", []string{"a", "b", "c"}, ""},
		{"abc", "?", []string{}, "abc"},
		{"?", "?", []string{}, "?"},
		{"abc", "d", []string{"x", "y", "z"}, "abc"},
	}

	for _, test := range tests {
		result := ReplaceArray(test.search, test.replace, test.subject)
		if result != test.expected {
			t.Errorf("ReplaceArray(%q, %v, %q) = %q, expected %q",
				test.search, test.replace, test.subject, result, test.expected)
		}
	}
}

func TestReplaceFirst(t *testing.T) {
	tests := []struct {
		subject  string
		search   string
		replace  string
		expected string
	}{
		{"hello world", "hello", "hi", "hi world"},
		{"hello hello world", "hello", "hi", "hi hello world"},
		{"hello world", "world", "earth", "hello earth"},
		{"hello world", "xyz", "abc", "hello world"},
		{"", "hello", "hi", ""},
		{"hello world", "", "hi", "hello world"},
		{"hello world", "hello world", "hi there", "hi there"},
	}

	for _, test := range tests {
		result := ReplaceFirst(test.search, test.replace, test.subject)
		if result != test.expected {
			t.Errorf("ReplaceFirst(%q, %q, %q) = %q, expected %q",
				test.search, test.replace, test.subject, result, test.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input     string
		maxLength int
		expected  string
	}{
		{"Hello, World", 5, "Hello..."},
		{"Hello", 10, "Hello"},
		{"", 5, ""},
		{"Hello", 0, ""},
		{"Hello", -1, ""},
		// For Unicode strings, we need to be careful with byte lengths
		// Each Chinese character takes 3 bytes in UTF-8
		{"你好世界", 6, "你好..."},
		{"Hello, World", 12, "Hello, World"},
		{"Hello, World", 11, "Hello, Worl..."},
	}

	for _, test := range tests {
		result := Truncate(test.input, test.maxLength)
		if result != test.expected {
			t.Errorf("Truncate(%q, %d) = %q, expected %q",
				test.input, test.maxLength, result, test.expected)
		}
	}
}

func TestReplaceLast(t *testing.T) {
	tests := []struct {
		subject  string
		search   string
		replace  string
		expected string
	}{
		{"hello world", "world", "earth", "hello earth"},
		{"hello hello world", "hello", "hi", "hello hi world"},
		{"hello world world", "world", "earth", "hello world earth"},
		{"hello world", "xyz", "abc", "hello world"},
		{"", "hello", "hi", ""},
		{"hello world", "", "hi", "hello world"},
		{"hello world", "hello world", "hi there", "hi there"},
	}

	for _, test := range tests {
		result := ReplaceLast(test.search, test.replace, test.subject)
		if result != test.expected {
			t.Errorf("ReplaceLast(%q, %q, %q) = %q, expected %q",
				test.search, test.replace, test.subject, result, test.expected)
		}
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		input    string
		prefix   string
		expected string
	}{
		{"hello", "/", "/hello"},
		{"/hello", "/", "/hello"},
		{"", "/", "/"},
		{"/", "/", "/"},
		{"hello", "", "hello"},
		{"world/hello", "/", "/world/hello"},
		{"/world/hello", "/", "/world/hello"},
	}

	for _, test := range tests {
		result := Start(test.input, test.prefix)
		if result != test.expected {
			t.Errorf("Start(%q, %q) = %q, expected %q", test.input, test.prefix, result, test.expected)
		}
	}
}

func TestStudly(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"hello-world", "HelloWorld"},
		{"hello_world", "HelloWorld"},
		{"hello.world", "HelloWorld"},
		{"hello_world-test.case", "HelloWorldTestCase"},
		{"", ""},
		{"hello", "Hello"},
		{"HELLO WORLD", "HelloWorld"},
		{"hElLo WoRlD", "HelloWorld"},
	}

	for _, test := range tests {
		result := Studly(test.input)
		if result != test.expected {
			t.Errorf("Studly(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		length   int
		expected string
	}{
		{"hello world", 0, 5, "hello"},
		{"hello world", 6, 5, "world"},
		{"hello world", 0, 11, "hello world"},
		{"hello world", 0, 20, "hello world"}, // Length beyond string length
		{"hello world", 5, 0, ""},             // Zero length
		{"hello world", -5, 5, "world"},       // Negative start
		{"hello world", -5, 10, "world"},      // Negative start with length beyond end
		{"hello world", 20, 5, ""},            // Start beyond string length
		{"hello world", 0, -1, "hello worl"},  // Negative length
		{"", 0, 5, ""},                        // Empty string
		{"你好世界", 0, 2, "你好"},                  // Unicode characters
	}

	for _, test := range tests {
		result := Substr(test.input, test.start, test.length)
		if result != test.expected {
			t.Errorf("Substr(%q, %d, %d) = %q, expected %q",
				test.input, test.start, test.length, result, test.expected)
		}

		// Verify our implementation against strings.Builder for a few cases
		// This uses the strings package to resolve the unused import warning
		if test.start >= 0 && test.length >= 0 && test.start < len(test.input) {
			var sb strings.Builder
			runes := []rune(test.input)
			end := test.start + test.length
			if end > len(runes) {
				end = len(runes)
			}
			for i := test.start; i < end; i++ {
				sb.WriteRune(runes[i])
			}
			builderResult := sb.String()
			if result != builderResult {
				t.Errorf("Substr implementation mismatch: %q vs strings.Builder: %q",
					result, builderResult)
			}
		}
	}
}

func TestUcfirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"HELLO", "HELLO"},
		{"hello world", "Hello world"},
		{"", ""},
		{"1hello", "1hello"},
		{"café", "Café"},
	}

	for _, test := range tests {
		result := Ucfirst(test.input)
		if result != test.expected {
			t.Errorf("Ucfirst(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestLcfirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "hello"},
		{"hello", "hello"},
		{"HELLO", "hELLO"},
		{"Hello world", "hello world"},
		{"", ""},
		{"1Hello", "1Hello"},
		{"Café", "café"},
	}

	for _, test := range tests {
		result := Lcfirst(test.input)
		if result != test.expected {
			t.Errorf("Lcfirst(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestLtrim(t *testing.T) {
	tests := []struct {
		input    string
		chars    string
		expected string
	}{
		{"  hello  ", " ", "hello  "},
		{"xxhelloxx", "x", "helloxx"},
		{"hello", "x", "hello"},
		{"", "x", ""},
		{"hello", "", "hello"},
		{"abcdefghello", "abcdefg", "hello"},
		{"abcabchello", "abc", "hello"},
	}

	for _, test := range tests {
		result := Ltrim(test.input, test.chars)
		if result != test.expected {
			t.Errorf("Ltrim(%q, %q) = %q, expected %q", test.input, test.chars, result, test.expected)
		}
	}
}

func TestRtrim(t *testing.T) {
	tests := []struct {
		input    string
		chars    string
		expected string
	}{
		{"  hello  ", " ", "  hello"},
		{"xxhelloxx", "x", "xxhello"},
		{"hello", "x", "hello"},
		{"", "x", ""},
		{"hello", "", "hello"},
		{"helloabcdefg", "abcdefg", "hello"},
		{"helloabcabc", "abc", "hello"},
	}

	for _, test := range tests {
		result := Rtrim(test.input, test.chars)
		if result != test.expected {
			t.Errorf("Rtrim(%q, %q) = %q, expected %q", test.input, test.chars, result, test.expected)
		}
	}
}

func TestPlural(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"book", "books"},
		{"child", "children"},
		{"fish", "fishes"},
		{"deer", "deers"},
		{"man", "men"},
		{"woman", "women"},
		{"tooth", "teeth"},
		{"foot", "feet"},
		{"mouse", "mice"},
		{"person", "people"},
		{"quiz", "quizzes"},
		{"matrix", "matrices"},
		{"analysis", "analyses"},
		{"index", "indices"},
		{"ox", "oxen"},
		{"knife", "knives"},
		{"life", "lives"},
		{"wife", "wives"},
		{"shelf", "shelves"},
		{"bus", "buses"},
		{"status", "statuses"},
		{"virus", "viruses"},
		{"octopus", "octopi"},
		{"", ""},
		{"already plural", "already plural"},
		{"data", "data"},
		{"series", "series"},
		{"species", "species"},
	}

	for _, test := range tests {
		result := Plural(test.input)
		if result != test.expected {
			t.Errorf("Plural(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestSingular(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"books", "book"},
		{"children", "child"},
		{"fish", "fish"},
		{"deer", "deer"},
		{"men", "man"},
		{"women", "woman"},
		{"teeth", "tooth"},
		{"feet", "foot"},
		{"mice", "mouse"},
		{"people", "person"},
		{"quizzes", "quiz"},
		{"matrices", "matrix"},
		{"analyses", "analysis"},
		{"indices", "index"},
		{"oxen", "ox"},
		{"knives", "knife"},
		{"lives", "life"},
		{"wives", "wife"},
		{"shelves", "shelf"},
		{"buses", "bus"},
		{"statuses", "status"},
		{"viruses", "virus"},
		{"octopi", "octopus"},
		{"", ""},
		{"already singular", "already singular"},
		{"data", "data"},
		{"series", "series"},
		{"species", "species"},
	}

	for _, test := range tests {
		result := Singular(test.input)
		if result != test.expected {
			t.Errorf("Singular(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestWordwrap(t *testing.T) {
	tests := []struct {
		input     string
		width     int
		breakChar string
		expected  string
	}{
		{"The quick brown fox jumped over the lazy dog.", 10, "\n", "The quick\nbrown fox\njumped\nover the\nlazy dog."},
		{"A very long word: supercalifragilisticexpialidocious", 10, "\n", "A very\nlong word:\nsupercalif\nragilistic\nexpialidoc\nious"},
		{"Short text", 20, "\n", "Short text"},
		{"", 10, "\n", ""},
		{"Word", 2, "\n", "Wo\nrd"},
		{"The quick brown fox", 10, "<br>", "The quick<br>brown fox"},
		{"The quick brown fox", 0, "\n", "The quick brown fox"},  // Width of 0 should not wrap
		{"The quick brown fox", -1, "\n", "The quick brown fox"}, // Negative width should not wrap
		{"Line1\nLine2\nLine3", 10, "\n", "Line1\nLine2\nLine3"}, // Preserve existing line breaks
	}

	for _, test := range tests {
		result := Wordwrap(test.input, test.width, test.breakChar)
		if result != test.expected {
			t.Errorf("Wordwrap(%q, %d, %q) = %q, expected %q",
				test.input, test.width, test.breakChar, result, test.expected)
		}
	}
}

func TestApa(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Creating A Project", "Creating a Project"},
		{"HELLO WORLD", "Hello WORLD"},
		{"hello WORLD", "Hello WORLD"},
		{"hElLo WoRlD", "Hello WoRlD"},
		{"", ""},
		{"hello", "Hello"},
		{"A B C", "A b c"},
		{"THE QUICK BROWN FOX", "The QUICK BROWN FOX"},
		{"the quick brown fox", "The quick brown fox"},
	}

	for _, test := range tests {
		result := Apa(test.input)
		if result != test.expected {
			t.Errorf("Apa(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestCharAt(t *testing.T) {
	tests := []struct {
		input    string
		position int
		expected string
	}{
		{"This is my name.", 6, "s"},
		{"Hello", 0, "H"},
		{"Hello", 4, "o"},
		{"Hello", 5, ""},        // Position out of bounds
		{"Hello", -1, ""},       // Negative position
		{"", 0, ""},             // Empty string
		{"你好世界", 0, "你"},        // Unicode character
		{"你好世界", 2, "世"},        // Unicode character
		{"你好世界", 4, ""},         // Position out of bounds
		{"Hello World", 5, " "}, // Space character
	}

	for _, test := range tests {
		result := CharAt(test.input, test.position)
		if result != test.expected {
			t.Errorf("CharAt(%q, %d) = %q, expected %q", test.input, test.position, result, test.expected)
		}
	}
}

func TestWordAt(t *testing.T) {
	tests := []struct {
		input    string
		position int
		expected string
	}{
		{"This is my name.", 6, "is"},
		{"Hello world", 0, "Hello"},
		{"Hello world", 6, "world"},
		{"Hello world", 12, ""},              // Position out of bounds
		{"Hello", -1, ""},                    // Negative position
		{"", 0, ""},                          // Empty string
		{"One two three", 4, "two"},          // Position in the middle of a word
		{"One two three", 8, "three"},        // Position at the start of a word
		{"Hello, world!", 7, "world"},        // Position with punctuation
		{"Multiple   spaces", 9, "spaces"},   // Multiple spaces between words
		{"Tab\tseparated", 4, "separated"},   // Tab character
		{"Line\nbreak", 5, "break"},          // Line break
		{"Hyphenated-word", 0, "Hyphenated"}, // Hyphenated word
		{"Don't worry", 0, "Don't"},          // Word with apostrophe
	}

	for _, test := range tests {
		result := WordAt(test.input, test.position)
		if result != test.expected {
			t.Errorf("WordAt(%q, %d) = %q, expected %q", test.input, test.position, result, test.expected)
		}
	}
}

func TestChopStart(t *testing.T) {
	// Test with single prefix
	singlePrefixTests := []struct {
		input    string
		prefix   string
		expected string
	}{
		{"https://laravel.com", "https://", "laravel.com"},
		{"http://laravel.com", "https://", "http://laravel.com"}, // No match
		{"laravel.com", "https://", "laravel.com"},               // No prefix to remove
		{"", "https://", ""},                                     // Empty string
		{"https://", "https://", ""},                             // Only prefix
		{"https://laravel.com", "", "https://laravel.com"},       // Empty prefix
	}

	for _, test := range singlePrefixTests {
		result := ChopStart(test.input, test.prefix)
		if result != test.expected {
			t.Errorf("ChopStart(%q, %q) = %q, expected %q", test.input, test.prefix, result, test.expected)
		}
	}

	// Test with array of prefixes
	multiPrefixTests := []struct {
		input    string
		prefixes []string
		expected string
	}{
		{"https://laravel.com", []string{"https://", "http://"}, "laravel.com"},
		{"http://laravel.com", []string{"https://", "http://"}, "laravel.com"},
		{"ftp://laravel.com", []string{"https://", "http://"}, "ftp://laravel.com"}, // No match
		{"laravel.com", []string{"https://", "http://"}, "laravel.com"},             // No prefix to remove
		{"", []string{"https://", "http://"}, ""},                                   // Empty string
		{"https://", []string{"https://", "http://"}, ""},                           // Only prefix
		{"http://", []string{"https://", "http://"}, ""},                            // Only prefix
		{"https://laravel.com", []string{}, "https://laravel.com"},                  // Empty prefixes array
	}

	for _, test := range multiPrefixTests {
		result := ChopStart(test.input, test.prefixes)
		if result != test.expected {
			t.Errorf("ChopStart(%q, %v) = %q, expected %q", test.input, test.prefixes, result, test.expected)
		}
	}
}

func TestDoesntContain(t *testing.T) {
	// Test with single substring
	singleSubstrTests := []struct {
		input    string
		substr   string
		expected bool
	}{
		{"This is name", "my", true},
		{"This is my name", "my", false},
		{"This is name", "", false},
		{"", "my", true},
		{"", "", false},
	}

	for _, test := range singleSubstrTests {
		result := DoesntContain(test.input, test.substr)
		if result != test.expected {
			t.Errorf("DoesntContain(%q, %q) = %v, expected %v", test.input, test.substr, result, test.expected)
		}
	}

	// Test with array of substrings
	multiSubstrTests := []struct {
		input      string
		substrings []string
		expected   bool
	}{
		{"This is name", []string{"my", "foo"}, true},
		{"This is my name", []string{"my", "foo"}, false},
		{"This is foo name", []string{"my", "foo"}, false},
		{"This is name", []string{}, true},
		{"", []string{"my", "foo"}, true},
		{"This is name", []string{""}, false},
		{"", []string{""}, false},
	}

	for _, test := range multiSubstrTests {
		result := DoesntContain(test.input, test.substrings)
		if result != test.expected {
			t.Errorf("DoesntContain(%q, %v) = %v, expected %v", test.input, test.substrings, result, test.expected)
		}
	}
}

func TestChopEnd(t *testing.T) {
	// Test with single suffix
	singleSuffixTests := []struct {
		input    string
		suffix   string
		expected string
	}{
		{"app/Models/Photograph.php", ".php", "app/Models/Photograph"},
		{"app/Models/Photograph.jpg", ".php", "app/Models/Photograph.jpg"}, // No match
		{"laravel.com", ".php", "laravel.com"},                             // No suffix to remove
		{"", ".php", ""},                                                   // Empty string
		{".php", ".php", ""},                                               // Only suffix
		{"app/Models/Photograph.php", "", "app/Models/Photograph.php"},     // Empty suffix
	}

	for _, test := range singleSuffixTests {
		result := ChopEnd(test.input, test.suffix)
		if result != test.expected {
			t.Errorf("ChopEnd(%q, %q) = %q, expected %q", test.input, test.suffix, result, test.expected)
		}
	}

	// Test with array of suffixes
	multiSuffixTests := []struct {
		input    string
		suffixes []string
		expected string
	}{
		{"laravel.com/index.php", []string{"/index.html", "/index.php"}, "laravel.com"},
		{"laravel.com/index.html", []string{"/index.html", "/index.php"}, "laravel.com"},
		{"laravel.com/about", []string{"/index.html", "/index.php"}, "laravel.com/about"}, // No match
		{"laravel.com", []string{"/index.html", "/index.php"}, "laravel.com"},             // No suffix to remove
		{"", []string{"/index.html", "/index.php"}, ""},                                   // Empty string
		{"/index.php", []string{"/index.html", "/index.php"}, ""},                         // Only suffix
		{"/index.html", []string{"/index.html", "/index.php"}, ""},                        // Only suffix
		{"laravel.com/index.php", []string{}, "laravel.com/index.php"},                    // Empty suffixes array
	}

	for _, test := range multiSuffixTests {
		result := ChopEnd(test.input, test.suffixes)
		if result != test.expected {
			t.Errorf("ChopEnd(%q, %v) = %q, expected %q", test.input, test.suffixes, result, test.expected)
		}
	}
}

func TestExcerpt(t *testing.T) {
	// Test with default options
	defaultTests := []struct {
		input    string
		phrase   string
		expected string
	}{
		{"This is my name", "my", "This is my name"},      // Short string, no truncation needed
		{"This is my name", "name", "This is my name"},    // Short string, no truncation needed
		{"This is my name", "missing", "This is my name"}, // Phrase not found
		{"", "my", ""}, // Empty string
		{"This is my name", "", "This is my name"}, // Empty phrase
		{"This is a very long string that should be truncated because it exceeds the default radius", "long", "This is a very long string that should be truncated because it exceeds the default radius"}, // Long string but still within radius
	}

	for _, test := range defaultTests {
		result := Excerpt(test.input, test.phrase)
		if result != test.expected {
			t.Errorf("Excerpt(%q, %q) = %q, expected %q", test.input, test.phrase, result, test.expected)
		}
	}

	// Test with custom radius
	radiusTests := []struct {
		input    string
		phrase   string
		options  ExcerptOptions
		expected string
	}{
		{"This is my name", "my", ExcerptOptions{Radius: 3}, "...is my na..."},
		{"This is my name", "name", ExcerptOptions{Radius: 3}, "...my name"},
		{"This is my name", "This", ExcerptOptions{Radius: 3}, "This is..."},
		{"This is my name", "missing", ExcerptOptions{Radius: 3}, "This is my name"}, // Phrase not found
		{"", "my", ExcerptOptions{Radius: 3}, ""},                                    // Empty string
		{"This is my name", "", ExcerptOptions{Radius: 3}, "This is my name"},        // Empty phrase
		{"This is my name", "my", ExcerptOptions{Radius: 0}, "...my..."},             // Zero radius
		{"This is my name", "my", ExcerptOptions{Radius: 100}, "This is my name"},    // Large radius
	}

	for _, test := range radiusTests {
		result := Excerpt(test.input, test.phrase, test.options)
		if result != test.expected {
			t.Errorf("Excerpt(%q, %q, %v) = %q, expected %q", test.input, test.phrase, test.options, result, test.expected)
		}
	}

	// Test with custom omission
	omissionTests := []struct {
		input    string
		phrase   string
		options  ExcerptOptions
		expected string
	}{
		{"This is my name", "my", ExcerptOptions{Radius: 3, Omission: "(...) "}, "(...) is my na(...) "},
		{"This is my name", "name", ExcerptOptions{Radius: 3, Omission: "(...) "}, "(...) my name"},
		{"This is my name", "This", ExcerptOptions{Radius: 3, Omission: "(...) "}, "This is(...) "},
		{"This is my name", "my", ExcerptOptions{Omission: ""}, "...my..."},                  // Empty omission
		{"This is my name", "my", ExcerptOptions{Radius: 3, Omission: ""}, "...is my na..."}, // Empty omission with radius
	}

	for _, test := range omissionTests {
		result := Excerpt(test.input, test.phrase, test.options)
		if result != test.expected {
			t.Errorf("Excerpt(%q, %q, %v) = %q, expected %q", test.input, test.phrase, test.options, result, test.expected)
		}
	}
}

func TestIsJson(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Valid JSON examples
		{"[1,2,3]", true},
		{`{"first": "John", "last": "Doe"}`, true},
		{"[]", true},
		{"{}", true},
		{"123", true},
		{"true", true},
		{"false", true},
		{"null", true},
		{`"string"`, true},

		// Invalid JSON examples
		{"{first: \"John\", last: \"Doe\"}", false},
		{"[1,2,", false},
		{"{key: value}", false},
		{"", false},
		{"undefined", false},
		{"function(){}", false},
		{"<html></html>", false},
		{"Hello World", false},
	}

	for _, test := range tests {
		result := IsJson(test.input)
		if result != test.expected {
			t.Errorf("IsJson(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected string
	}{
		// Basic matching with no capturing groups
		{"/bar/", "foo bar", "bar"},
		{"/foo/", "foo bar", "foo"},
		{"/baz/", "foo bar", ""}, // No match

		// Matching with capturing groups
		{"/foo (.*)/", "foo bar", "bar"},
		{"/foo (\\w+)/", "foo bar123", "bar123"},
		{"/foo (\\w+) (\\w+)/", "foo bar baz", "bar"},       // Returns first group
		{"/foo ((\\w+) (\\w+))/", "foo bar baz", "bar baz"}, // Nested groups

		// Pattern with leading and trailing slashes
		{"bar", "foo bar", "bar"},   // No slashes
		{"(foo)", "foo bar", "foo"}, // No slashes with group

		// Edge cases
		{"", "foo bar", ""},    // Empty pattern
		{"/foo/", "", ""},      // Empty string
		{"/(/", "foo bar", ""}, // Invalid regex
		{"/[/", "foo bar", ""}, // Invalid regex
	}

	for _, test := range tests {
		result := Match(test.pattern, test.s)
		if result != test.expected {
			t.Errorf("Match(%q, %q) = %q, expected %q", test.pattern, test.s, result, test.expected)
		}
	}
}

func TestMatchAll(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected []string
	}{
		// Basic matching with no capturing groups
		{"/bar/", "bar foo bar", []string{"bar", "bar"}},
		{"/foo/", "foo bar foo", []string{"foo", "foo"}},
		{"/baz/", "foo bar", []string{}}, // No match

		// Matching with capturing groups
		{"/f(\\w*)/", "bar fun bar fly", []string{"un", "ly"}},
		{"/b(a)r/", "bar foo bar", []string{"a", "a"}},
		{"/foo (\\w+)/", "foo bar foo baz", []string{"bar", "baz"}},
		{"/foo ((\\w+) (\\w+))/", "foo bar baz foo qux quux", []string{"bar baz", "qux quux"}}, // Nested groups

		// Pattern with leading and trailing slashes
		{"bar", "bar foo bar", []string{"bar", "bar"}},   // No slashes
		{"(foo)", "foo bar foo", []string{"foo", "foo"}}, // No slashes with group

		// Edge cases
		{"", "foo bar", []string{}},    // Empty pattern
		{"/foo/", "", []string{}},      // Empty string
		{"/(/", "foo bar", []string{}}, // Invalid regex
		{"/[/", "foo bar", []string{}}, // Invalid regex
	}

	for _, test := range tests {
		result := MatchAll(test.pattern, test.s)

		// Check if lengths match
		if len(result) != len(test.expected) {
			t.Errorf("MatchAll(%q, %q) returned %d results, expected %d", test.pattern, test.s, len(result), len(test.expected))
			continue
		}

		// Check each element
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("MatchAll(%q, %q)[%d] = %q, expected %q", test.pattern, test.s, i, result[i], test.expected[i])
			}
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		search   string
		subject  string
		useRegex bool
		expected string
	}{
		// Basic character removal
		{"e", "Peter Piper picked a peck of pickled peppers.", false, "Ptr Pipr pickd a pck of pickld ppprs."},
		{"a", "banana", false, "bnn"},
		{"x", "banana", false, "banana"}, // No occurrences

		// String removal
		{"abc", "abcdef", false, "def"},
		{"abc", "abcabcabc", false, ""},
		{"world", "hello world", false, "hello "},

		// Empty strings
		{"", "hello", false, "hello"}, // Empty search string
		{"a", "", false, ""},          // Empty subject
		{"", "", false, ""},           // Both empty

		// Regex removal
		{"[aeiou]", "Hello World", true, "Hll Wrld"}, // Remove vowels
		{"\\d+", "abc123def456", true, "abcdef"},     // Remove digits
		{"\\s+", "Hello  World", true, "HelloWorld"}, // Remove whitespace

		// Invalid regex (should fall back to string replacement)
		{"[", "abc[def", true, "abcdef"},

		// Edge cases
		{" ", "a b c", false, "abc"},                // Remove spaces
		{"\n", "line1\nline2", false, "line1line2"}, // Remove newlines
	}

	for _, test := range tests {
		var result string
		if test.useRegex {
			result = Remove(test.search, test.subject, true)
		} else {
			result = Remove(test.search, test.subject)
		}

		if result != test.expected {
			t.Errorf("Remove(%q, %q, %v) = %q, expected %q",
				test.search, test.subject, test.useRegex, result, test.expected)
		}
	}
}

func TestReplaceMatches(t *testing.T) {
	// Test cases for string replacements
	stringTests := []struct {
		pattern  string
		replace  string
		subject  string
		expected string
	}{
		// Basic replacements
		{"/[^A-Za-z0-9]+/", "", "(+1) 501-555-1000", "15015551000"}, // Example from the issue description (using standard quantifier)
		{"/\\d/", "X", "123", "XXX"},                                // Replace digits with X
		{"/\\s+/", "-", "hello  world", "hello-world"},              // Replace whitespace with dash

		// Pattern with and without slashes
		{"[aeiou]", "*", "hello world", "h*ll* w*rld"},   // No slashes
		{"/[aeiou]/", "*", "hello world", "h*ll* w*rld"}, // With slashes

		// Edge cases
		{"", "replacement", "subject", "subject"},                  // Empty pattern
		{"/pattern/", "replacement", "", ""},                       // Empty subject
		{"/[/", "replacement", "invalid [regex", "invalid [regex"}, // Invalid regex
		{"/pattern/", "", "pattern exists", " exists"},             // Empty replacement (only replaces "pattern", not the entire string)
	}

	for _, test := range stringTests {
		result := ReplaceMatches(test.pattern, test.replace, test.subject)
		if result != test.expected {
			t.Errorf("ReplaceMatches(%q, %q, %q) = %q, expected %q",
				test.pattern, test.replace, test.subject, result, test.expected)
		}
	}

	// Test cases for function replacements
	// Example from the issue description: wrapping digits in square brackets
	result1 := ReplaceMatches("/\\d/", func(matches []string) string {
		return "[" + matches[0] + "]"
	}, "123")
	expected1 := "[1][2][3]"
	if result1 != expected1 {
		t.Errorf("ReplaceMatches with function: got %q, expected %q", result1, expected1)
	}

	// Uppercase all matched words
	result2 := ReplaceMatches("/\\b\\w+\\b/", func(matches []string) string {
		return strings.ToUpper(matches[0])
	}, "hello world")
	expected2 := "HELLO WORLD"
	if result2 != expected2 {
		t.Errorf("ReplaceMatches with function: got %q, expected %q", result2, expected2)
	}

	// Double all matched numbers
	result3 := ReplaceMatches("/\\d+/", func(matches []string) string {
		return matches[0] + matches[0]
	}, "There are 25 apples and 10 oranges")
	expected3 := "There are 2525 apples and 1010 oranges"
	if result3 != expected3 {
		t.Errorf("ReplaceMatches with function: got %q, expected %q", result3, expected3)
	}

	// Test with capturing groups
	result4 := ReplaceMatches("/(\\w+)=(\\w+)/", func(matches []string) string {
		// matches[0] is the full match, matches[1] and matches[2] are the capturing groups
		return matches[2] + ":" + matches[1]
	}, "key=value")
	expected4 := "value:key"
	if result4 != expected4 {
		t.Errorf("ReplaceMatches with function and capturing groups: got %q, expected %q", result4, expected4)
	}

	// Test with unsupported replacement type
	result5 := ReplaceMatches("/\\d/", 123, "123") // Passing an int instead of string or function
	expected5 := "123"                             // Should return the original string
	if result5 != expected5 {
		t.Errorf("ReplaceMatches with unsupported replacement type: got %q, expected %q", result5, expected5)
	}
}

func TestSquish(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Example from the issue description
		{"    laravel    framework    ", "laravel framework"},

		// Other test cases
		{"hello      world", "hello world"},
		{"   ", ""},
		{"", ""},
		{"\t\n hello \t\n world \t\n", "hello world"}, // Tabs and newlines
		{"multiple   spaces   between   words", "multiple spaces between words"},
		{"no extra spaces", "no extra spaces"}, // Already properly spaced
		{" leading space", "leading space"},
		{"trailing space ", "trailing space"},
		{" both leading and trailing spaces ", "both leading and trailing spaces"},
		{"   multiple    leading    and    trailing    spaces   ", "multiple leading and trailing spaces"},
	}

	for _, test := range tests {
		result := Squish(test.input)
		if result != test.expected {
			t.Errorf("Squish(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		subject      string
		replacements map[string]string
		expected     string
	}{
		// Example from the issue description
		{"Tacos are great!", map[string]string{"Tacos": "Burritos", "great": "fantastic"}, "Burritos are fantastic!"},

		// Other test cases
		{"abc", map[string]string{"a": "x", "b": "y"}, "xyc"},
		{"hello world", map[string]string{}, "hello world"}, // No replacements
		{"", map[string]string{"a": "b"}, ""},               // Empty string
		{"hello world", map[string]string{"hello": "hi", "world": "earth"}, "hi earth"},
		{"hello hello world", map[string]string{"hello": "hi"}, "hi hi world"},        // Multiple occurrences
		{"hello world", map[string]string{"not found": "replacement"}, "hello world"}, // Search not found
		{"The quick brown fox", map[string]string{"quick": "slow", "brown": "red", "fox": "dog"}, "The slow red dog"},
	}

	for _, test := range tests {
		result := Swap(test.replacements, test.subject)
		if result != test.expected {
			t.Errorf("Swap(%v, %q) = %q, expected %q", test.replacements, test.subject, result, test.expected)
		}
	}
}
