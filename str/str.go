// Package str provides utility functions for string manipulation.
package str

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToString converts any value to its string representation.
// Example: ToString(123) -> "123"
func ToString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

// RuneLength returns the number of runes in a string.
// Example: RuneLength("Hello, 世界") -> 8
func RuneLength(str string) int {
	return utf8.RuneCountInString(str)
}

// Words splits string into an array of its words.
func Words(str string) []string {
	if str == "" {
		return []string{}
	}

	// Remove leading/trailing whitespace
	str = strings.TrimSpace(str)
	if str == "" {
		return []string{}
	}

	// Enhanced regular expression to handle number-letter boundaries
	// This handles:
	// - Sequences of letters followed by numbers (Int8 -> Int, 8)
	// - Numbers followed by letters (8Value -> 8, Value)
	// - CamelCase transitions
	// - Underscores, hyphens, and other separators
	wordRegex := regexp.MustCompile(`[A-Z]*[a-z]+|[A-Z]+[a-z]*|\d+|[a-z]+`)

	// Find all matches
	matches := wordRegex.FindAllString(str, -1)

	var words []string
	for _, match := range matches {
		word := strings.ToLower(strings.TrimSpace(match))
		if word != "" && isValidWord(word) {
			words = append(words, word)
		}
	}

	// Enhanced fallback: if regex didn't find anything, use custom splitting logic
	if len(words) == 0 {
		words = splitByBoundaries(str)
	}

	return words
}

// WordsPattern splits string into words using a custom pattern.
// Example: WordsPattern("hello-world_test", `[\-_]+`) -> ["hello", "world", "test"]
func WordsPattern(s, pattern string) []string {
	if s == "" {
		return []string{}
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		// Fallback to default behavior if pattern is invalid
		return Words(s)
	}

	parts := regex.Split(s, -1)
	var words []string

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			words = append(words, strings.ToLower(part))
		}
	}

	return words
}

// Camelcase converts a string to camelCase.
// Example: "foo bar" -> "fooBar"
func Camelcase(s string) string {
	words := Words(s)
	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		if word == "" {
			continue
		}
		result += Capitalize(strings.ToLower(word))
	}
	return result
}

// KebabCase converts string to kebab case.
// Example: "hello world" -> "hello-world"
// Example: "HelloWorld" -> "hello-world"
// Example: "HELLO_WORLD" -> "hello-world"
func KebabCase(s string) string {
	return changeConnector(s, "-")
}

// SnakeCase converts string to snake case.
// Example: "hello world" -> "hello_world"
// Example: "HelloWorld" -> "hello_world"
// Example: "HELLO-WORLD" -> "hello_world"
func SnakeCase(s string) string {
	return changeConnector(s, "_")
}

// PascalCase converts string to pascal case.
// Example: "hello world" -> "HelloWorld"
// Example: "hello-world" -> "HelloWorld"
// Example: "hello_world" -> "HelloWorld"
func PascalCase(s string) string {
	items := Words(s)
	for i := range items {
		items[i] = Capitalize(items[i])
	}
	return strings.Join(items, "")
}

// Capitalize capitalizes the first character of a string.
// Example: "fred" -> "Fred"
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// EndsWith checks if a string ends with the given target string.
// Example: EndsWith("abc", "c") -> true
func EndsWith(s, target string) bool {
	return strings.HasSuffix(s, target)
}

// StartsWith checks if a string starts with the given target string.
// Example: StartsWith("abc", "a") -> true
func StartsWith(s, target string) bool {
	return strings.HasPrefix(s, target)
}

// Trim removes leading and trailing whitespace or specified characters from a string.
// Example: Trim("  abc  ") -> "abc"
func Trim(s string, cutset ...string) string {
	if len(cutset) > 0 {
		return strings.Trim(s, cutset[0])
	}
	return strings.TrimSpace(s)
}

// TrimStart removes leading whitespace or specified characters from a string.
// Example: TrimStart("  abc  ") -> "abc  "
func TrimStart(s string, cutset ...string) string {
	if len(cutset) > 0 {
		return strings.TrimLeft(s, cutset[0])
	}
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

// TrimEnd removes trailing whitespace or specified characters from a string.
// Example: TrimEnd("  abc  ") -> "  abc"
func TrimEnd(s string, cutset ...string) string {
	if len(cutset) > 0 {
		return strings.TrimRight(s, cutset[0])
	}
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// ToLower converts a string to lowercase.
// Example: ToLower("ABC") -> "abc"
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase.
// Example: ToUpper("abc") -> "ABC"
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Split splits a string by the given separator.
// Example: Split("a-b-c", "-") -> ["a", "b", "c"]
func Split(s, separator string) []string {
	return strings.Split(s, separator)
}

// Join joins an array of strings with the given separator.
// Example: Join(["a", "b", "c"], "-") -> "a-b-c"
func Join(arr []string, separator string) string {
	return strings.Join(arr, separator)
}

// Repeat repeats a string n times.
// Example: Repeat("abc", 2) -> "abcabc"
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// Replace replaces all occurrences of a substring in a string.
// Example: Replace("Hi Fred", "Fred", "Barney") -> "Hi Barney"
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Contains checks if a string contains a substring.
// Example: Contains("abc", "b") -> true
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Count counts the occurrences of a substring in a string.
// Example: Count("ababab", "ab") -> 3
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Index returns the index of the first occurrence of a substring in a string.
// Returns -1 if the substring is not found.
// Example: Index("abc", "b") -> 1
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// LastIndex returns the index of the last occurrence of a substring in a string.
// Returns -1 if the substring is not found.
// Example: LastIndex("abcabc", "b") -> 4
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// Ellipsis trims and truncates a string to a specified length in bytes and appends an ellipsis if truncated.
// It ensures that UTF-8 characters are not split in the middle.
// Example: Ellipsis("Hello, 世界", 8) -> "Hello, ..."
func Ellipsis(s string, length int) string {
	s = Trim(s)

	if len(s) <= length {
		return s
	}

	// If length is too small, just return ellipsis
	if length <= 0 {
		return "..."
	}

	// Ensure we don't break UTF-8 characters
	var result []byte
	bytesUsed := 0

	for i := 0; i < len(s); {
		_, size := utf8.DecodeRuneInString(s[i:])
		if bytesUsed+size > length {
			break
		}

		// Append the rune to the result
		result = append(result, s[i:i+size]...)
		bytesUsed += size
		i += size
	}

	return string(result) + "..."
}

// splitByBoundaries splits string by detecting word boundaries manually
func splitByBoundaries(str string) []string {
	if str == "" {
		return []string{}
	}

	var words []string
	var currentWord strings.Builder
	runes := []rune(str)

	for i, r := range runes {
		// Skip whitespace and punctuation separators
		if unicode.IsSpace(r) || isPunctuation(r) {
			if currentWord.Len() > 0 {
				word := strings.ToLower(currentWord.String())
				if isValidWord(word) {
					words = append(words, word)
				}
				currentWord.Reset()
			}
			continue
		}

		// Check for word boundaries
		if i > 0 {
			prevRune := runes[i-1]

			// Boundary conditions:
			// 1. Letter to number (Int8 -> Int|8)
			// 2. Number to letter (8Value -> 8|Value)
			// 3. Lowercase to uppercase (camelCase -> camel|Case)
			// 4. Multiple uppercase to lowercase (XMLHttp -> XML|Http)
			if shouldSplit(prevRune, r, i, runes) {
				if currentWord.Len() > 0 {
					word := strings.ToLower(currentWord.String())
					if isValidWord(word) {
						words = append(words, word)
					}
					currentWord.Reset()
				}
			}
		}

		currentWord.WriteRune(r)
	}

	// Add the last word
	if currentWord.Len() > 0 {
		word := strings.ToLower(currentWord.String())
		if isValidWord(word) {
			words = append(words, word)
		}
	}

	return words
}

// isValidWord checks if a string is a valid word (contains letters, digits, or valid symbols)
func isValidWord(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\'' && r != '\u2019' {
			return false
		}
	}

	return true
}

// shouldSplit determines if we should split at the boundary between two runes
func shouldSplit(prev, curr rune, pos int, runes []rune) bool {
	// Letter to digit boundary (Int8 -> Int|8)
	if unicode.IsLetter(prev) && unicode.IsDigit(curr) {
		return true
	}

	// Digit to letter boundary (8Value -> 8|Value)
	if unicode.IsDigit(prev) && unicode.IsLetter(curr) {
		return true
	}

	// Lowercase to uppercase boundary (camelCase -> camel|Case)
	if unicode.IsLower(prev) && unicode.IsUpper(curr) {
		return true
	}

	// Multiple uppercase followed by lowercase (XMLHttp -> XML|Http)
	if unicode.IsUpper(prev) && unicode.IsUpper(curr) && pos+1 < len(runes) {
		if unicode.IsLower(runes[pos+1]) {
			return true
		}
	}

	return false
}

// there isPunctuation checks if a rune is punctuation that should split words
func isPunctuation(r rune) bool {
	return r == '-' || r == '_' || r == '.' || r == ',' || r == ';' || r == ':' ||
		r == '!' || r == '?' || r == '(' || r == ')' || r == '[' || r == ']' ||
		r == '{' || r == '}' || r == '/' || r == '\\' || r == '|' || r == '+' ||
		r == '=' || r == '<' || r == '>' || r == '@' || r == '#' || r == '$' ||
		r == '%' || r == '^' || r == '&' || r == '*'
}

// changeConnector converts a string to a case format using the specified connector.
// It splits the input string into words, converts each word to lowercase,
// and joins them back together with the given connector string.
// Example: changeConnector("HelloWorld", "-") -> "hello-world"
func changeConnector(s, c string) string {
	words := Words(s)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, c)
}
