# str - String Utility Functions for Go

The `str` package provides a comprehensive set of utility functions for working with strings in Go. It offers a wide range of functions to make string manipulation easier and more expressive.

## Installation

```bash
go get github.com/gflydev/utils/str
```

## Usage

```go
import "github.com/gflydev/utils/str"
```

## Functions

### ToString

Converts any value to a string representation.

```go
result := str.ToString(123)
// result: "123"

result := str.ToString(true)
// result: "true"

result := str.ToString([]int{1, 2, 3})
// result: "[1 2 3]"
```

### Length

Counts the number of Unicode characters (runes) in a string.

```go
result := str.Length("Hello, 世界")
// result: 9

result := str.Length("abc")
// result: 3
```

### Words

Splits string into an array of its words. It handles various word boundaries including camelCase, snake_case, and kebab-case.

```go
result := str.Words("hello world")
// result: ["hello", "world"]

result := str.Words("camelCase")
// result: ["camel", "case"]

result := str.Words("snake_case")
// result: ["snake", "case"]

result := str.Words("kebab-case")
// result: ["kebab", "case"]
```

### WordsPattern

Splits string into words using a custom pattern. The pattern is used as a regular expression to split the string.

```go
result := str.WordsPattern("hello-world_test", `[\-_]+`)
// result: ["hello", "world", "test"]

result := str.WordsPattern("a,b;c", `[,;]`)
// result: ["a", "b", "c"]
```

### CamelCase

Converts a string to camelCase.

```go
result := str.CamelCase("foo bar")
// result: "fooBar"

result := str.CamelCase("Foo Bar")
// result: "fooBar"

result := str.CamelCase("foo bar baz")
// result: "fooBarBaz"
```

### Capitalize

Capitalizes the first character of a string. It leaves the rest of the string unchanged. If the string is empty, it returns an empty string.

```go
result := str.Capitalize("fred")
// result: "Fred"

result := str.Capitalize("FRED")
// result: "FRED"

result := str.Capitalize("fred flintstone")
// result: "Fred flintstone"

result := str.Capitalize("")
// result: ""
```

### EndsWith

Checks if a string ends with the specified substring.

```go
result := str.EndsWith("abc", "c")
// result: true

result := str.EndsWith("abc", "bc")
// result: true

result := str.EndsWith("abc", "d")
// result: false
```

### StartsWith

Checks if a string starts with the specified substring.

```go
result := str.StartsWith("abc", "a")
// result: true

result := str.StartsWith("abc", "ab")
// result: true

result := str.StartsWith("abc", "d")
// result: false
```

### Trim

Removes leading and trailing whitespace or specified characters from a string.

```go
result := str.Trim("  abc  ")
// result: "abc"

result := str.Trim("-_-abc-_-", "-_")
// result: "abc"

result := str.Trim("abc")
// result: "abc"

result := str.Trim("")
// result: ""
```

### ToLower

Converts a string to lowercase.

```go
result := str.ToLower("FRED")
// result: "fred"

result := str.ToLower("Fred")
// result: "fred"

result := str.ToLower("fred")
// result: "fred"

result := str.ToLower("")
// result: ""
```

### ToUpper

Converts a string to uppercase. This function transforms all characters in the string to their uppercase equivalents.

```go
result := str.ToUpper("fred")
// result: "FRED"

result := str.ToUpper("Fred")
// result: "FRED"

result := str.ToUpper("FRED")
// result: "FRED"

result := str.ToUpper("")
// result: ""
```

### Split

Splits a string by a separator. This function divides a string into an array of substrings based on the specified separator.

```go
result := str.Split("a-b-c", "-")
// result: []string{"a", "b", "c"}

result := str.Split("abc", "")
// result: []string{"a", "b", "c"}

result := str.Split("a", "-")
// result: []string{"a"}

result := str.Split("", "-")
// result: []string{""}
```

### Join

Joins an array of strings with a separator. This function combines all elements of a string array into a single string, with the specified separator between each element.

```go
result := str.Join([]string{"a", "b", "c"}, "-")
// result: "a-b-c"

result := str.Join([]string{"a", "b", "c"}, "")
// result: "abc"

result := str.Join([]string{"a"}, "-")
// result: "a"

result := str.Join([]string{}, "-")
// result: ""
```

### Repeat

Repeats a string n times. This function creates a new string consisting of the original string repeated a specified number of times.

```go
result := str.Repeat("abc", 2)
// result: "abcabc"

result := str.Repeat("abc", 0)
// result: ""

result := str.Repeat("", 5)
// result: ""
```

### Replace

Replaces all occurrences of a search string with a replacement string in a subject string.

**Parameters:**
- `search`: The string to be replaced
- `replace`: The string to replace with
- `subject`: The string to perform replacements on

**Returns:**
- A new string with all occurrences of the search string replaced

**Examples:**
```go
result := str.Replace("a", "b", "aaa")
// result: "bbb"

result := str.Replace("a", "b", "ccc")
// result: "ccc" (no change if search string not found)

result := str.Replace("Fred", "Barney", "Hi Fred")
// result: "Hi Barney"

result := str.Replace("d", "e", "abc")
// result: "abc" (no change if search string not found)

result := str.Replace("a", "b", "")
// result: "" (no change for empty string)
```

### Contains

Determines if a string contains a given substring.

**Parameters:**
- `s`: The string to search in
- `substr`: The substring to search for

**Returns:**
- `bool`: True if substring is found, false otherwise

**Examples:**
```go
result := str.Contains("abc", "a")
// result: true

result := str.Contains("abc", "d")
// result: false

result := str.Contains("abc", "")
// result: true (empty string is always contained)

result := str.Contains("", "")
// result: true

result := str.Contains("", "a")
// result: false (non-empty substring cannot be in empty string)
```

### Ellipsis

Trims and truncates a string to a specified length in bytes and appends an ellipsis if truncated. It ensures that UTF-8 characters are not split in the middle.

**Parameters:**
- `s`: The string to truncate
- `length`: The maximum length in bytes before truncation

**Returns:**
- `string`: The truncated string with "..." appended if truncation occurred

**Examples:**
```go
result := str.Ellipsis("This is a long text", 10)
// result: "This is..."

result := str.Ellipsis("Short", 10)
// result: "Short"

result := str.Ellipsis("Hello, 世界", 8)
// result: "Hello, ..."

result := str.Ellipsis("你好, World", 6)
// result: "你好..."

result := str.Ellipsis("Hello", 0)
// result: "..."
```

### Truncate

Truncates a string to the specified length and adds an ellipsis if truncated. It returns the original string if its length is less than or equal to maxLength, otherwise returns the truncated string with "..." appended.

**Parameters:**
- `s`: The input string to truncate
- `maxLength`: The maximum allowed length of the string

**Returns:**
- `string`: The truncated string with "..." appended if truncation occurred, otherwise original string

**Examples:**
```go
result := str.Truncate("Hello, World", 5)
// result: "Hello..."

result := str.Truncate("Hello", 10)
// result: "Hello"

result := str.Truncate("", 5)
// result: ""

result := str.Truncate("Hello", 0)
// result: ""
```

### KebabCase

Converts a string to kebab-case.

```go
result := str.KebabCase("fooBar")
// result: "foo-bar"

result := str.KebabCase("foo bar")
// result: "foo-bar"
```

### SnakeCase

Converts a string to snake_case.

```go
result := str.SnakeCase("fooBar")
// result: "foo_bar"

result := str.SnakeCase("foo bar")
// result: "foo_bar"
```

### PascalCase

Converts a string to PascalCase.

```go
result := str.PascalCase("foo bar")
// result: "FooBar"

result := str.PascalCase("foo-bar")
// result: "FooBar"
```

### Headline

Converts a string to Title Case with spaces between words. It handles various word boundaries including camelCase, snake_case, and kebab-case.

```go
result := str.Headline("foo bar")
// result: "Foo Bar"

result := str.Headline("foo-bar")
// result: "Foo Bar"

result := str.Headline("EmailNotificationSent")
// result: "Email Notification Sent"

result := str.Headline("snake_case_string")
// result: "Snake Case String"
```

### TrimStart

Removes leading whitespace or specified characters from a string.

```go
result := str.TrimStart("  abc  ")
// result: "abc  "

result := str.TrimStart("-_-abc-_-", "-_")
// result: "abc-_-"

result := str.TrimStart("abc")
// result: "abc"
```

### TrimEnd

Removes trailing whitespace or specified characters from a string.

```go
result := str.TrimEnd("  abc  ")
// result: "  abc"

result := str.TrimEnd("-_-abc-_-", "-_")
// result: "-_-abc"

result := str.TrimEnd("abc")
// result: "abc"
```

### Count

Counts the occurrences of a substring in a string.

**Parameters:**
- `s`: The string to search in
- `substr`: The substring to search for

**Returns:**
- `int`: The number of non-overlapping occurrences of the substring

**Examples:**
```go
result := str.Count("ababa", "a")
// result: 3

result := str.Count("ababa", "ab")
// result: 2

result := str.Count("abc", "d")
// result: 0

result := str.Count("", "a")
// result: 0
```

### Index

Returns the index of the first occurrence of a substring in a string. Returns -1 if the substring is not found.

**Parameters:**
- `s`: The string to search in
- `substr`: The substring to search for

**Returns:**
- `int`: The index of the first occurrence of substr in s, or -1 if not found

**Examples:**
```go
result := str.Index("abcabc", "a")
// result: 0

result := str.Index("abcabc", "bc")
// result: 1

result := str.Index("abc", "d")
// result: -1

result := str.Index("", "a")
// result: -1
```

### LastIndex

Returns the index of the last occurrence of a substring in a string. Returns -1 if the substring is not found.

**Parameters:**
- `s`: The string to search in
- `substr`: The substring to search for

**Returns:**
- `int`: The index of the last occurrence of substr in s, or -1 if not found

**Examples:**
```go
result := str.LastIndex("abcabc", "a")
// result: 3

result := str.LastIndex("abcabc", "bc")
// result: 4

result := str.LastIndex("abc", "d")
// result: -1

result := str.LastIndex("", "a")
// result: -1
```

### Slugify

Converts a string to a URL-friendly slug.

```go
result := str.Slugify("Hello World!")
// result: "hello-world"

result := str.Slugify("This is a test")
// result: "this-is-a-test"
```

### IsEmptyOrWhitespace

Checks if a string is empty or contains only whitespace characters.

```go
result := str.IsEmptyOrWhitespace("")
// result: true

result := str.IsEmptyOrWhitespace("   ")
// result: true

result := str.IsEmptyOrWhitespace("\t\n")
// result: true

result := str.IsEmptyOrWhitespace("hello")
// result: false

result := str.IsEmptyOrWhitespace(" hello ")
// result: false
```

### ContainsAny

Checks if a string contains any of the specified substrings.

```go
result := str.ContainsAny("abc", "a", "d")
// result: true

result := str.ContainsAny("abc", "d", "e")
// result: false
```

### ToTitleCase

Converts a string to title case.

```go
result := str.ToTitleCase("hello world")
// result: "Hello World"

result := str.ToTitleCase("HELLO WORLD")
// result: "Hello World"
```

### OnlyAlphanumeric

Removes all non-alphanumeric characters from a string.

```go
result := str.OnlyAlphanumeric("Hello, World!")
// result: "HelloWorld"

result := str.OnlyAlphanumeric("123-456-789")
// result: "123456789"
```

### Mask

Masks a portion of a string, keeping a specified number of characters visible at the start and end, and replacing the rest with a mask character. If the string is too short (less than or equal to the sum of visible characters), it is returned unchanged.

```go
result := str.Mask("1234567890", 4, 2, '*')
// result: "1234****90"

result := str.Mask("abcdefghij", 0, 3, '*')
// result: "*******hij"

result := str.Mask("1234", 2, 2, '*')
// result: "1234" (no masking if string is too short)
```

### PadLeft

Pads the left side of a string with a specified character to reach the desired length. If the string is already longer than the specified length, it is returned unchanged.

```go
result := str.PadLeft("abc", ' ', 5)
// result: "  abc"

result := str.PadLeft("abc", '0', 5)
// result: "00abc"

result := str.PadLeft("hello", '*', 4)
// result: "hello" (no padding if string is already longer)
```

### PadRight

Pads the right side of a string with a specified character to reach the desired length. If the string is already longer than the specified length, it is returned unchanged.

```go
result := str.PadRight("abc", ' ', 5)
// result: "abc  "

result := str.PadRight("abc", '0', 5)
// result: "abc00"

result := str.PadRight("hello", '*', 4)
// result: "hello" (no padding if string is already longer)
```

### Reverse

Reverses the characters in a string. It properly handles UTF-8 encoded strings by working with runes.

```go
result := str.Reverse("abc")
// result: "cba"

result := str.Reverse("hello")
// result: "olleh"

result := str.Reverse("Hello, 世界")
// result: "界世 ,olleH"
```

### CountWords

Counts the number of words in a string. It uses `strings.Fields()` to split the string into words, which handles whitespace correctly.

**Parameters:**
- `s`: The string to count words in

**Returns:**
- `int`: The number of words in the string

**Examples:**
```go
result := str.CountWords("hello world")
// result: 2

result := str.CountWords("hello   world")
// result: 2

result := str.CountWords("")
// result: 0

result := str.CountWords("   ")
// result: 0

result := str.CountWords("hello")
// result: 1
```

### TruncateWords

Truncates a string to the specified number of words and adds an ellipsis if the string was truncated.

**Parameters:**
- `s`: The string to truncate
- `maxWords`: Maximum number of words to keep

**Returns:**
- `string`: The truncated string with "..." appended if truncation occurred

**Examples:**
```go
result := str.TruncateWords("hello world foo bar", 2)
// result: "hello world..."

result := str.TruncateWords("hello world", 3)
// result: "hello world"

result := str.TruncateWords("hello", 1)
// result: "hello"

result := str.TruncateWords("", 5)
// result: ""

result := str.TruncateWords("hello world", 0)
// result: ""
```

### FormatWithCommas

Formats a number as a string with commas as thousand separators.
Note: The current implementation does not actually add commas and simply returns the string representation of the number. This function may be updated in the future.

**Parameters:**
- `n`: The number to format

**Returns:**
- `string`: The formatted number string

**Examples:**
```go
result := str.FormatWithCommas(1000)
// result: "1000"

result := str.FormatWithCommas(1234567)
// result: "1234567"

result := str.FormatWithCommas(-1000)
// result: "-1000"
```

### After

Returns the portion of a string after the first occurrence of a given value.

**Parameters:**
- `s`: The string to search in
- `search`: The substring to search for

**Returns:**
- `string`: Everything after the search string, or the entire string if not found

**Examples:**
```go
result := str.After("hello world", "hello ")
// result: "world"

result := str.After("hello world", "not found")
// result: "hello world"

result := str.After("hello world", "")
// result: "hello world"
```

### AfterLast

Returns the portion of a string after the last occurrence of a given value.

**Parameters:**
- `s`: The string to search in
- `search`: The substring to search for

**Returns:**
- Everything after the last occurrence of search string, or entire string if not found

**Examples:**
```go
result := str.AfterLast("This is a test", "is")
// result: " a test"

result := str.AfterLast("This is a test", "not")
// result: "This is a test"

result := str.AfterLast("hello/world/test", "/")
// result: "test"

result := str.AfterLast("hello world hello", "hello ")
// result: "hello"

result := str.AfterLast("hello world", "")
// result: "hello world"
```

### Before

Returns the portion of a string before the first occurrence of a given value.

**Parameters:**
- `s`: The string to search in
- `search`: The substring to search for

**Returns:**
- Everything before the search string, or the entire string if not found

**Examples:**
```go
result := str.Before("This is a test", "is")
// result: "Th"

result := str.Before("This is a test", "not")
// result: "This is a test"

result := str.Before("hello world", " world")
// result: "hello"

result := str.Before("hello/world/test", "/")
// result: "hello"

result := str.Before("hello world", "")
// result: "hello world"
```

### BeforeLast

Returns the portion of a string before the last occurrence of a given value.

**Parameters:**
- `s`: The string to search in
- `search`: The substring to search for

**Returns:**
- Everything before the last occurrence of search string, or entire string if not found

**Examples:**
```go
result := str.BeforeLast("This is a test", "is")
// result: "This "

result := str.BeforeLast("This is a test", "not")
// result: "This is a test"

result := str.BeforeLast("hello/world/test", "/")
// result: "hello/world"

result := str.BeforeLast("hello world hello", "hello")
// result: "hello world "

result := str.BeforeLast("hello world", "")
// result: "hello world"
```

### Between

Returns the portion of a string between two values.

**Parameters:**
- `s`: The string to search in
- `start`: The starting substring
- `end`: The ending substring

**Returns:**
- The portion between start and end strings, or entire string if not found

**Examples:**
```go
result := str.Between("This is a test", "This", "test")
// result: " is a "

result := str.Between("This is a test", "not", "found")
// result: "This is a test"

result := str.Between("hello [world] test", "[", "]")
// result: "world"

result := str.Between("<div>content</div>", "<div>", "</div>")
// result: "content"

result := str.Between("hello world", "[", "]")
// result: "hello world"

result := str.Between("hello world", "", "]")
// result: "hello world"

result := str.Between("hello world", "[", "")
// result: "hello world"

result := str.Between("hello [[world]]", "[", "]")
// result: "world"
```

### BetweenFirst

Returns the portion of a string between the first occurrence of two strings.

**Parameters:**
- `s`: The string to search in
- `start`: The starting delimiter
- `end`: The ending delimiter

**Returns:**
- The portion between the first occurrence of start and end strings, or entire string if not found

**Examples:**
```go
result := str.BetweenFirst("This is a test", "This", "test")
// result: " is a "

result := str.BetweenFirst("This is a test", "not", "found")
// result: "This is a test"

result := str.BetweenFirst("[a] bc [d]", "[", "]")
// result: "a"

result := str.BetweenFirst("<div>content</div>", "<div>", "</div>")
// result: "content"

result := str.BetweenFirst("hello world", "[", "]")
// result: "hello world"

result := str.BetweenFirst("hello world", "", "]")
// result: "hello world"

result := str.BetweenFirst("hello world", "[", "")
// result: "hello world"
```

### ContainsAll

Determines if a string contains all of the given substrings.

**Parameters:**
- `s`: The string to search in
- `substrings`: Variable number of substrings to search for

**Returns:**
- `bool`: True if all substrings are found, false otherwise

**Examples:**
```go
result := str.ContainsAll("abc", "a", "b")
// result: true

result := str.ContainsAll("abc", "a", "d")
// result: false

result := str.ContainsAll("hello world", "hello", "world")
// result: true

result := str.ContainsAll("hello world", "hello", "missing")
// result: false

result := str.ContainsAll("hello world", "HELLO", "WORLD")
// result: false (case-sensitive)

result := str.ContainsAll("hello world")
// result: true (no substrings to check)
```

### Finish

Appends a single instance of the given value to a string if it does not already end with it.

**Parameters:**
- `s`: The string to append to
- `cap`: The string to append

**Returns:**
- The resulting string

**Examples:**
```go
result := str.Finish("test", "/")
// result: "test/"

result := str.Finish("test/", "/")
// result: "test/" (already ends with the cap)

result := str.Finish("hello", "!")
// result: "hello!"

result := str.Finish("hello!", "!")
// result: "hello!" (already ends with the cap)
```

### Is

Checks if a string matches a pattern. The pattern can include wildcards (*) which match any sequence of characters.

**Parameters:**
- `pattern`: The pattern to match against, can include * as wildcards
- `s`: The string to check

**Returns:**
- `bool`: True if the string matches the pattern, false otherwise

**Examples:**
```go
result := str.Is("foo*", "foobar")
// result: true

result := str.Is("*bar", "foobar")
// result: true

result := str.Is("foo*bar", "foobar")
// result: true

result := str.Is("foo", "foobar")
// result: false

result := str.Is("*baz", "foobar")
// result: false
```

### IsAscii

Determines if a string contains only 7-bit ASCII characters.

**Parameters:**
- `s`: The string to check

**Returns:**
- `bool`: True if string contains only ASCII characters, false otherwise

**Examples:**
```go
result := str.IsAscii("hello world")
// result: true

result := str.IsAscii("hello123!@#")
// result: true

result := str.IsAscii("こんにちは")
// result: false

result := str.IsAscii("hello世界")
// result: false
```

### Ascii

Transliterates non-ASCII characters to their ASCII equivalents. This function converts characters with diacritical marks to their basic Latin equivalents.

**Parameters:**
- `s`: The string to transliterate

**Returns:**
- `string`: The transliterated string with only ASCII characters

**Examples:**
```go
result := str.Ascii("über")
// result: "uber"

result := str.Ascii("café")
// result: "cafe"

result := str.Ascii("Crème Brûlée")
// result: "Creme Brulee"

result := str.Ascii("ñ")
// result: "n"

result := str.Ascii("Hello World")
// result: "Hello World" (already ASCII)
```

### Limit

Limits a string to a specified length and optionally appends a string to the end. If the string is already shorter than the limit, it is returned unchanged.

**Parameters:**
- `s`: The string to limit
- `limit`: The maximum length of the string in runes
- `options`: Optional parameters:
  - A string to append if the string is truncated (e.g., "...")

**Returns:**
- `string`: The truncated string, with the optional suffix appended if truncation occurred

**Examples:**
```go
result := str.Limit("This is a test", 7)
// result: "This is"

result := str.Limit("This is a test", 7, "...")
// result: "This is..."

result := str.Limit("Short", 10)
// result: "Short"

result := str.Limit("Hello world", 0)
// result: ""

result := str.Limit("", 5)
// result: ""
```

### Random

Generates a random alphanumeric string of the specified length. The string will contain a mix of uppercase letters, lowercase letters, and numbers.

**Parameters:**
- `length`: The desired length of the random string

**Returns:**
- `string`: A random alphanumeric string of the specified length

**Examples:**
```go
result := str.Random(8)
// result: "a1B2c3D4" (random 8-character string)

result := str.Random(16)
// result: "x7Y9zAbCdEfG1h2" (random 16-character string)

result := str.Random(0)
// result: "" (empty string)
```

### Password

Generates a random password with the given length. If no length is provided, the default length is 32 characters. The password will contain a mix of uppercase letters, lowercase letters, numbers, and special characters.

**Parameters:**
- `length`: The desired length of the password (optional, default: 32)

**Returns:**
- `string`: The generated random password

**Examples:**
```go
result := str.Password()
// result: "EbJo2vE-AS:U,$%_gkrV4n,q~1xy/-_4" (random 32-character password)

result := str.Password(12)
// result: "qwuar>#V|i]N" (random 12-character password)

result := str.Password(0)
// result: "" (empty string for zero length)
```

### ReplaceArray

Replaces a search string with an array of replacements sequentially. Each occurrence of the search string is replaced with the corresponding element from the replacement array. If there are more occurrences than replacements, the remaining occurrences are left unchanged.

**Parameters:**
- `search`: The string to find
- `replace`: Array of replacement strings
- `subject`: The string to perform replacements on

**Returns:**
- `string`: The resulting string after replacements

**Examples:**
```go
result := str.ReplaceArray("?", []string{"a", "b", "c"}, "? and ? and ?")
// result: "a and b and c"

result := str.ReplaceArray("?", []string{"a", "b"}, "? and ? and ?")
// result: "a and b and ?" (not enough replacements)

result := str.ReplaceArray("?", []string{"a", "b", "c", "d"}, "? and ?")
// result: "a and b" (extra replacements ignored)

result := str.ReplaceArray("not found", []string{"replacement"}, "hello world")
// result: "hello world" (search not found)
```

### ReplaceFirst

Replaces the first occurrence of a given value in a string. If the search string is not found, the original string is returned unchanged.

**Parameters:**
- `search`: The string to find
- `replace`: The string to replace with
- `subject`: The string to perform replacement on

**Returns:**
- `string`: The resulting string after replacement

**Examples:**
```go
result := str.ReplaceFirst("a", "x", "ababa")
// result: "xbaba"

result := str.ReplaceFirst("hello", "hi", "hello world hello")
// result: "hi world hello"

result := str.ReplaceFirst("not found", "replacement", "hello world")
// result: "hello world" (search not found)

result := str.ReplaceFirst("", "replacement", "hello world")
// result: "hello world" (empty search string is ignored)
```

### ReplaceLast

Replaces the last occurrence of a substring in a string.

```go
result := str.ReplaceLast("a", "x", "ababa")
// result: "ababx"

result := str.ReplaceLast("hello", "hi", "hello world hello")
// result: "hello world hi"

result := str.ReplaceLast("not found", "replacement", "hello world")
// result: "hello world" (search not found)

result := str.ReplaceLast("", "replacement", "hello world")
// result: "hello world" (empty search string is ignored)
```

### Start

Prepends a value to a string if it doesn't already start with it.

```go
result := str.Start("world", "hello ")
// result: "hello world"

result := str.Start("hello world", "hello ")
// result: "hello world" (already starts with prefix)

result := str.Start("hello", "")
// result: "hello" (empty prefix)

result := str.Start("", "hello")
// result: "hello" (empty string)
```

### Studly

Converts a string to StudlyCase format (no spaces or separators).

```go
result := str.Studly("hello_world")
// result: "HelloWorld"

result := str.Studly("hello-world")
// result: "HelloWorld"

result := str.Studly("hello world")
// result: "HelloWorld"

result := str.Studly("hello_WORLD")
// result: "HelloWorld"
```

### Substr

Returns a portion of a string based on start position and length.

```go
result := str.Substr("hello world", 0, 5)
// result: "hello"

result := str.Substr("hello world", 6, 5)
// result: "world"

result := str.Substr("hello world", -5, 5)
// result: "world" (negative start counts from end)

result := str.Substr("hello world", 0, -6)
// result: "hello" (negative length counts from end)

result := str.Substr("hello world", 20, 5)
// result: "" (start beyond string length)
```

### Ucfirst

Capitalizes the first character of a string.

```go
result := str.Ucfirst("hello")
// result: "Hello"

result := str.Ucfirst("hello world")
// result: "Hello world"

result := str.Ucfirst("Hello")
// result: "Hello" (already capitalized)

result := str.Ucfirst("")
// result: "" (empty string)
```

### Lcfirst

Converts the first character of a string to lowercase.

```go
result := str.Lcfirst("Hello")
// result: "hello"

result := str.Lcfirst("Hello World")
// result: "hello World"

result := str.Lcfirst("hello")
// result: "hello" (already lowercase)

result := str.Lcfirst("")
// result: "" (empty string)
```

### Ltrim

Removes specified characters from the start of a string.

```go
result := str.Ltrim("  hello", " ")
// result: "hello"

result := str.Ltrim("xxxhello", "x")
// result: "hello"

result := str.Ltrim("hello world", "hello ")
// result: "world"

result := str.Ltrim("hello", "x")
// result: "hello" (no characters to trim)

result := str.Ltrim("", "x")
// result: "" (empty string)
```

### Rtrim

Removes specified characters from the end of a string.

```go
result := str.Rtrim("hello  ", " ")
// result: "hello"

result := str.Rtrim("helloxxx", "x")
// result: "hello"

result := str.Rtrim("hello world", "world ")
// result: "hello"

result := str.Rtrim("hello", "x")
// result: "hello" (no characters to trim)

result := str.Rtrim("", "x")
// result: "" (empty string)
```

### Apa

Converts a string to title case but with the first word having only its first letter capitalized. This is similar to AP (Associated Press) style for article titles.

```go
result := str.Apa("Creating A Project")
// result: "Creating a Project"

result := str.Apa("HELLO WORLD")
// result: "Hello WORLD"

result := str.Apa("hello WORLD")
// result: "Hello WORLD"

result := str.Apa("")
// result: ""
```

### Plural

Converts a singular word to its plural form. This function handles various English pluralization rules including regular plurals, irregular plurals, and special cases.

**Parameters:**
- `s`: The singular word to pluralize

**Returns:**
- `string`: The plural form of the word

**Examples:**
```go
result := str.Plural("car")
// result: "cars"

result := str.Plural("child")
// result: "children" (irregular plural)

result := str.Plural("city")
// result: "cities" (y -> ies)

result := str.Plural("box")
// result: "boxes" (x -> xes)

result := str.Plural("day")
// result: "days" (vowel + y -> ys)

result := str.Plural("")
// result: "" (empty string)
```

### Singular

Converts a plural word to its singular form. This function handles various English singularization rules including regular plurals, irregular plurals, and special cases.

**Parameters:**
- `s`: The plural word to singularize

**Returns:**
- `string`: The singular form of the word

**Examples:**
```go
result := str.Singular("books")
// result: "book"

result := str.Singular("children")
// result: "child" (irregular plural)

result := str.Singular("cities")
// result: "city" (ies -> y)

result := str.Singular("boxes")
// result: "box" (es -> "")

result := str.Singular("days")
// result: "day" (s -> "")

result := str.Singular("")
// result: "" (empty string)
```

### Wordwrap

Wraps a string to a given number of characters. It breaks the string at word boundaries when possible, and inserts the specified break character at each wrap point.

**Parameters:**
- `s`: The string to wrap
- `width`: The number of characters at which to wrap
- `breakChar`: The string to insert at break points

**Returns:**
- `string`: The wrapped string

**Examples:**
```go
result := str.Wordwrap("A very long sentence that needs wrapping.", 10, "\n")
// result: "A very\nlong\nsentence\nthat needs\nwrapping."

result := str.Wordwrap("Short text", 20, "\n")
// result: "Short text" (no wrapping needed)

result := str.Wordwrap("word word word", 5, "<br>")
// result: "word<br>word<br>word"

result := str.Wordwrap("", 10, "\n")
// result: "" (empty string)
```

### CharAt

Returns the character at a specified position in a string.

**Parameters:**
- `s`: The input string
- `position`: The position of the character to return (0-indexed)

**Returns:**
- The character at the specified position, or an empty string if the position is out of bounds

**Examples:**
```go
result := str.CharAt("abcdef", 2)
// result: "c"

result := str.CharAt("Hello", 0)
// result: "H"

result := str.CharAt("Hello", 4)
// result: "o"

result := str.CharAt("Hello", 5)
// result: "" (position out of bounds)

result := str.CharAt("", 0)
// result: "" (empty string)
```

### WordAt

Returns the word at a specified position in a string.

**Parameters:**
- `s`: The input string
- `position`: The position to check for a word (0-indexed)

**Returns:**
- The word at the specified position, or an empty string if the position is out of bounds

**Examples:**
```go
result := str.WordAt("This is a test", 1)
// result: "is"

result := str.WordAt("Hello world", 0)
// result: "Hello"

result := str.WordAt("Hello world", 6)
// result: "world"

result := str.WordAt("Hello world", 12)
// result: "" (position out of bounds)

result := str.WordAt("", 0)
// result: "" (empty string)
```

### ChopStart

Removes a prefix from a string if it exists. If an array of prefixes is provided, it will remove the first matching prefix.

**Parameters:**
- `s`: The string to process
- `prefixes`: The prefix or array of prefixes to remove

**Returns:**
- The string with the prefix removed

**Examples:**
```go
result := str.ChopStart("Hello world", "Hello ")
// result: "world"

result := str.ChopStart("https://laravel.com", "https://")
// result: "laravel.com"

result := str.ChopStart("http://laravel.com", []string{"https://", "http://"})
// result: "laravel.com"

result := str.ChopStart("laravel.com", "https://")
// result: "laravel.com" (no prefix to remove)

result := str.ChopStart("", "https://")
// result: "" (empty string)
```

### ChopEnd

Removes a suffix from a string if it exists. If an array of suffixes is provided, it will remove the first matching suffix.

**Parameters:**
- `s`: The string to process
- `suffixes`: The suffix or array of suffixes to remove

**Returns:**
- The string with the suffix removed

**Examples:**
```go
result := str.ChopEnd("Hello world", " world")
// result: "Hello"

result := str.ChopEnd("app/Models/Photograph.php", ".php")
// result: "app/Models/Photograph"

result := str.ChopEnd("laravel.com/index.php", []string{"/index.html", "/index.php"})
// result: "laravel.com"

result := str.ChopEnd("laravel.com", ".php")
// result: "laravel.com" (no suffix to remove)

result := str.ChopEnd("", ".php")
// result: "" (empty string)
```

### Excerpt

Extracts a portion of text around a given phrase. It returns a substring that includes the phrase and a certain number of characters around it. If the excerpt doesn't include the entire string, omission text is added at the beginning and/or end.

**Parameters:**
- `s`: The string to excerpt
- `phrase`: The phrase to search for
- `options`: Optional ExcerptOptions struct containing:
  - `Radius`: The number of characters to include around the phrase (default: 100)
  - `Omission`: The text to use for omission (default: "...")

**Returns:**
- The excerpted string with omission text if truncated

**Examples:**
```go
result := str.Excerpt("This is a test", "is")
// result: "This is a test"

result := str.Excerpt("This is my name", "my", str.ExcerptOptions{Radius: 3})
// result: "...is my na..."

result := str.Excerpt("This is my name", "my", str.ExcerptOptions{Radius: 5, Omission: "(...)"})
// result: "(...)is my name"

result := str.Excerpt("This is my name", "foo", str.ExcerptOptions{})
// result: "This is my name"

result := str.Excerpt("", "foo", str.ExcerptOptions{})
// result: ""
```

### IsJson

Determines if a string is valid JSON.

**Parameters:**
- `s`: The string to check

**Returns:**
- True if the string is valid JSON, false otherwise

**Examples:**
```go
result := str.IsJson("{\"name\":\"John\"}")
// result: true

result := str.IsJson("[1,2,3]")
// result: true

result := str.IsJson("Not JSON")
// result: false

result := str.IsJson("{first: \"John\", last: \"Doe\"}")
// result: false (invalid JSON format)

result := str.IsJson("")
// result: false
```

### Match

Returns the first match of a regular expression pattern in a string. If the pattern contains capturing groups, it returns the first captured group. Otherwise, it returns the entire match.

**Parameters:**
- `pattern`: The regular expression pattern to match
- `s`: The string to search in

**Returns:**
- The matched portion or first captured group, or empty string if no match

**Examples:**
```go
result := str.Match("/bar/", "foo bar")
// result: "bar"

result := str.Match("/foo (.*)/", "foo bar")
// result: "bar"

result := str.Match("foo(.*)", "foobar")
// result: "bar"

result := str.Match("/xyz/", "foo bar")
// result: ""
```

### MatchAll

Returns all matches of a regular expression pattern in a string. If the pattern contains capturing groups, it returns all captured groups. Otherwise, it returns all full matches.

**Parameters:**
- `pattern`: The regular expression pattern to match
- `s`: The string to search in

**Returns:**
- A slice containing all matches or captured groups, or an empty slice if no matches

**Examples:**
```go
result := str.MatchAll("/bar/", "bar foo bar")
// result: ["bar", "bar"]

result := str.MatchAll("/f(\\w*)/", "bar fun bar fly")
// result: ["un", "ly"]

result := str.MatchAll("(foo)(bar)", "foobar")
// result: []string{"foobar", "foo", "bar"}

result := str.MatchAll("/xyz/", "foo bar")
// result: []
```

### ReplaceMatches

Replaces all occurrences of a pattern in a string using a regular expression. The replacement can be either a string or a function that returns a string.

**Parameters:**
- `pattern`: The regular expression pattern to match (can include delimiters like `/pattern/`)
- `replace`: The replacement (string or function that takes a match array and returns a string)
- `subject`: The string to perform replacements on

**Returns:**
- A new string with all matches replaced according to the pattern and replacement

**Examples:**
```go
result := str.ReplaceMatches("(foo)(bar)", "$2$1", "foobar")
// result: "barfoo"

result := str.ReplaceMatches("/[^A-Za-z0-9]++/", "", "(+1) 501-555-1000")
// result: "15015551000"

// Using a function as replacement
result := str.ReplaceMatches("/\\d/", func(matches []string) string { 
    return "[" + matches[0] + "]" 
}, "123")
// result: "[1][2][3]"

// Empty pattern or subject returns the original string
result := str.ReplaceMatches("", "replacement", "subject")
// result: "subject"
```

### Squish

Removes all extraneous white space from a string, including extraneous white space between words.

**Parameters:**
- `s`: The string to squish

**Returns:**
- The string with all extraneous white space removed

**Examples:**
```go
result := str.Squish("  Hello   world  ")
// result: "Hello world"

result := str.Squish("    laravel    framework    ")
// result: "laravel framework"

result := str.Squish("hello      world")
// result: "hello world"

result := str.Squish("   ")
// result: ""

result := str.Squish("")
// result: ""
```

### Swap

Replaces multiple values in a string with their corresponding replacements using a map.

**Parameters:**
- `replacements`: A map of search strings to their replacements
- `subject`: The string to perform replacements on

**Returns:**
- A new string with all specified replacements applied

**Examples:**
```go
result := str.Swap(map[string]string{"Hello": "Hi", "world": "there"}, "Hello world")
// result: "Hi there"

result := str.Swap(map[string]string{"Tacos": "Burritos", "great": "fantastic"}, "Tacos are great!")
// result: "Burritos are fantastic!"

result := str.Swap(map[string]string{"a": "x", "b": "y"}, "abc")
// result: "xyc"

result := str.Swap(map[string]string{}, "hello world")
// result: "hello world" (no replacements)
```

### DoesntContain

Determines if a string does not contain a specific substring or any of the substrings in an array.

**Parameters:**
- `s`: The string to check
- `substrings`: The substring(s) to check for (can be a string or []string)

**Returns:**
- `bool`: True if the string doesn't contain the specified substring(s), false otherwise

**Examples:**
```go
result := str.DoesntContain("abc", "d")
// result: true

result := str.DoesntContain("abc", "a")
// result: false

result := str.DoesntContain("This is name", "my")
// result: true

result := str.DoesntContain("This is name", []string{"my", "foo"})
// result: true

result := str.DoesntContain("This is my name", "my")
// result: false

result := str.DoesntContain("This is my name", []string{"my", "foo"})
// result: false
```

### Remove

Removes all occurrences of a given substring or pattern from a string.

**Parameters:**
- `search`: The substring or pattern to remove
- `subject`: The string to remove occurrences from
- `options`: Optional boolean parameter to use regex (default: false)

**Returns:**
- A new string with all occurrences of the search string removed

**Examples:**
```go
result := str.Remove("a", "abcabc")
// result: "bcbc"

result := str.Remove("e", "Peter Piper picked a peck of pickled peppers.")
// result: "Ptr Pipr pickd a pck of pickld ppprs."

// Using regex
result := str.Remove("[aeiou]", "Hello World", true)
// result: "Hll Wrld"

// Empty search string returns the original string
result := str.Remove("", "hello")
// result: "hello"
```
