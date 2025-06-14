// Package num provides utility functions for number manipulation.
package num

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"strings"
)

// Clamp constrains a number between lower and upper bounds.
//
// Parameters:
//   - n: The number to clamp
//   - lower: The lower bound
//   - upper: The upper bound
//
// Returns:
//   - float64: The clamped value
//
// Examples:
//
//	Clamp(10, 0, 5) // Returns 5 (n is above upper bound)
//	Clamp(-3, 0, 5) // Returns 0 (n is below lower bound)
//	Clamp(3, 0, 5)  // Returns 3 (n is within bounds)
func Clamp(n, lower, upper float64) float64 {
	if lower > upper {
		lower, upper = upper, lower
	}
	return math.Min(math.Max(n, lower), upper)
}

// InRange checks if a number is between start and end (inclusive).
//
// Parameters:
//   - n: The number to check
//   - start: The start of the range
//   - end: The end of the range
//
// Returns:
//   - bool: true if the number is within the range, false otherwise
//
// Examples:
//
//	InRange(3, 2, 4)  // Returns true (n is within range)
//	InRange(1, 2, 4)  // Returns false (n is below range)
//	InRange(5, 2, 4)  // Returns false (n is above range)
//	InRange(3, 4, 2)  // Returns true (start and end are automatically ordered)
func InRange(n, start, end float64) bool {
	if start > end {
		start, end = end, start
	}
	return n >= start && n <= end
}

// Random returns a random integer between min and max (inclusive).
//
// Parameters:
//   - min: The minimum value (inclusive)
//   - max: The maximum value (inclusive)
//
// Returns:
//   - int: A random integer between min and max
//
// Examples:
//
//	Random(1, 10)  // Returns a random number between 1 and 10
//	Random(5, 5)   // Always returns 5
//	Random(10, 1)  // Works the same as Random(1, 10)
func Random(min, max int) int {
	if min > max {
		min, max = max, min // Swap if min > max
	}

	if min == max {
		return min
	}

	// Use time-based seed for better randomness
	return rand.IntN(max-min+1) + min
}

// Round rounds a number to the nearest integer or to the specified precision.
//
// Parameters:
//   - n: The number to round
//   - precision: Optional. The number of decimal places to round to.
//     If not provided, rounds to the nearest integer.
//
// Returns:
//   - float64: The rounded number
//
// Examples:
//
//	Round(4.7)    // Returns 5.0 (rounded to nearest integer)
//	Round(4.7, 1) // Returns 4.7 (rounded to 1 decimal place)
//	Round(4.75, 1) // Returns 4.8 (rounded to 1 decimal place)
//	Round(-4.7)   // Returns -5.0 (rounded to nearest integer)
func Round(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Round(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Round(n*p) / p
}

// Floor rounds a number down to the nearest integer or to the specified precision.
//
// Parameters:
//   - n: The number to round down
//   - precision: Optional. The number of decimal places to round to.
//     If not provided, rounds down to the nearest integer.
//
// Returns:
//   - float64: The rounded down number
//
// Examples:
//
//	Floor(4.7)     // Returns 4.0 (rounded down to nearest integer)
//	Floor(4.78, 1) // Returns 4.7 (rounded down to 1 decimal place)
//	Floor(4.75, 1) // Returns 4.7 (rounded down to 1 decimal place)
//	Floor(-4.2)    // Returns -5.0 (rounded down to nearest integer)
func Floor(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Floor(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Floor(n*p) / p
}

// Ceil rounds a number up to the nearest integer or to the specified precision.
//
// Parameters:
//   - n: The number to round up
//   - precision: Optional. The number of decimal places to round to.
//     If not provided, rounds up to the nearest integer.
//
// Returns:
//   - float64: The rounded up number
//
// Examples:
//
//	Ceil(4.3)     // Returns 5.0 (rounded up to nearest integer)
//	Ceil(4.78, 1) // Returns 4.8 (rounded up to 1 decimal place)
//	Ceil(4.71, 1) // Returns 4.8 (rounded up to 1 decimal place)
//	Ceil(-4.7)    // Returns -4.0 (rounded up to nearest integer)
func Ceil(n float64, precision ...int) float64 {
	if len(precision) == 0 {
		return math.Ceil(n)
	}
	p := math.Pow(10, float64(precision[0]))
	return math.Ceil(n*p) / p
}

// Max returns the maximum value from a list of numbers.
//
// Parameters:
//   - numbers: A variadic list of float64 numbers
//
// Returns:
//   - float64: The maximum value from the list, or 0 if the list is empty
//
// Examples:
//
//	Max(1, 2, 3)       // Returns 3.0
//	Max(-1, -5, -3)    // Returns -1.0
//	Max(7.5, 3.2, 9.8) // Returns 9.8
//	Max()              // Returns 0.0 (empty list)
func Max(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	m := numbers[0]
	for _, n := range numbers[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

// MaxBy returns the element from a collection that produces the maximum value
// when passed through the iteratee function.
//
// Parameters:
//   - collection: A slice of any type T
//   - iteratee: A function that takes an element of type T and returns a float64
//
// Returns:
//   - T: The element that produces the maximum value, or zero value of T if collection is empty
//
// Examples:
//
//	// Find the number with the largest square
//	MaxBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * n) }) // Returns 3
//
//	// Find the person with the highest age
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}
//	MaxBy(people, func(p Person) float64 { return float64(p.Age) }) // Returns Person{"Bob", 30}
func MaxBy[T any](collection []T, iteratee func(T) float64) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}

	maxIdx := 0
	maxValue := iteratee(collection[0])

	for i := 1; i < len(collection); i++ {
		value := iteratee(collection[i])
		if value > maxValue {
			maxValue = value
			maxIdx = i
		}
	}

	return collection[maxIdx]
}

// Min returns the minimum value from a list of numbers.
//
// Parameters:
//   - numbers: A variadic list of float64 numbers
//
// Returns:
//   - float64: The minimum value from the list, or 0 if the list is empty
//
// Examples:
//
//	Min(1, 2, 3)       // Returns 1.0
//	Min(-1, -5, -3)    // Returns -5.0
//	Min(7.5, 3.2, 9.8) // Returns 3.2
//	Min()              // Returns 0.0 (empty list)
func Min(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	m := numbers[0]
	for _, n := range numbers[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

// MinBy returns the element from a collection that produces the minimum value
// when passed through the iteratee function.
//
// Parameters:
//   - collection: A slice of any type T
//   - iteratee: A function that takes an element of type T and returns a float64
//
// Returns:
//   - T: The element that produces the minimum value, or zero value of T if collection is empty
//
// Examples:
//
//	// Find the number with the smallest square
//	MinBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * n) }) // Returns 1
//
//	// Find the person with the lowest age
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}
//	MinBy(people, func(p Person) float64 { return float64(p.Age) }) // Returns Person{"Charlie", 22}
func MinBy[T any](collection []T, iteratee func(T) float64) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}

	minIdx := 0
	minValue := iteratee(collection[0])

	for i := 1; i < len(collection); i++ {
		value := iteratee(collection[i])
		if value < minValue {
			minValue = value
			minIdx = i
		}
	}

	return collection[minIdx]
}

// Sum calculates the sum of all numbers in the provided list.
//
// Parameters:
//   - numbers: A variadic list of float64 numbers
//
// Returns:
//   - float64: The sum of all numbers in the list, or 0 if the list is empty
//
// Examples:
//
//	Sum(1, 2, 3)       // Returns 6.0
//	Sum(-1, 5, 3)      // Returns 7.0
//	Sum(7.5, 3.2, 9.8) // Returns 20.5
//	Sum()              // Returns 0.0 (empty list)
func Sum(numbers ...float64) float64 {
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// SumBy calculates the sum of values in a collection after applying the iteratee function to each element.
//
// Parameters:
//   - collection: A slice of any type T
//   - iteratee: A function that takes an element of type T and returns a float64
//
// Returns:
//   - float64: The sum of all values after applying the iteratee function, or 0 if the collection is empty
//
// Examples:
//
//	// Sum of doubled values
//	SumBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * 2) }) // Returns 12.0 (2+4+6)
//
//	// Sum of ages
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}
//	SumBy(people, func(p Person) float64 { return float64(p.Age) }) // Returns 77.0 (25+30+22)
func SumBy[T any](collection []T, iteratee func(T) float64) float64 {
	var sum float64
	for _, item := range collection {
		sum += iteratee(item)
	}
	return sum
}

// Mean calculates the arithmetic mean (average) of a list of numbers.
//
// Parameters:
//   - numbers: A variadic list of float64 numbers
//
// Returns:
//   - float64: The arithmetic mean of the numbers, or 0 if the list is empty
//
// Examples:
//
//	Mean(1, 2, 3)       // Returns 2.0
//	Mean(2, 4, 6, 8)    // Returns 5.0
//	Mean(7.5, 3.2, 9.8) // Returns 6.833333333333333
//	Mean()              // Returns 0.0 (empty list)
func Mean(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return Sum(numbers...) / float64(len(numbers))
}

// MeanBy calculates the arithmetic mean (average) of values in a collection
// after applying the iteratee function to each element.
//
// Parameters:
//   - collection: A slice of any type T
//   - iteratee: A function that takes an element of type T and returns a float64
//
// Returns:
//   - float64: The arithmetic mean of all values after applying the iteratee function,
//     or 0 if the collection is empty
//
// Examples:
//
//	// Mean of doubled values
//	MeanBy([]int{1, 2, 3}, func(n int) float64 { return float64(n * 2) }) // Returns 4.0 (mean of 2,4,6)
//
//	// Mean of ages
//	type Person struct { Name string; Age int }
//	people := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 22}}
//	MeanBy(people, func(p Person) float64 { return float64(p.Age) }) // Returns 25.666666666666668
func MeanBy[T any](collection []T, iteratee func(T) float64) float64 {
	if len(collection) == 0 {
		return 0
	}

	return SumBy(collection, iteratee) / float64(len(collection))
}

// Abs returns the absolute value of a number.
//
// Parameters:
//   - n: The number to get the absolute value of
//
// Returns:
//   - float64: The absolute value of the number
//
// Examples:
//
//	Abs(-5)    // Returns 5.0
//	Abs(5)     // Returns 5.0
//	Abs(-3.14) // Returns 3.14
//	Abs(0)     // Returns 0.0
func Abs(n float64) float64 {
	return math.Abs(n)
}

// Pow returns the base raised to the exponent power.
//
// Parameters:
//   - base: The base number
//   - exponent: The exponent to raise the base to
//
// Returns:
//   - float64: The result of base^exponent
//
// Examples:
//
//	Pow(2, 3)    // Returns 8.0 (2^3)
//	Pow(10, 2)   // Returns 100.0 (10^2)
//	Pow(5, 0)    // Returns 1.0 (any number raised to 0 is 1)
//	Pow(2, -1)   // Returns 0.5 (2^-1 = 1/2)
//	Pow(4, 0.5)  // Returns 2.0 (square root of 4)
func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Sqrt returns the square root of a number.
//
// Parameters:
//   - n: The number to calculate the square root of
//
// Returns:
//   - float64: The square root of the number
//
// Examples:
//
//	Sqrt(9)     // Returns 3.0
//	Sqrt(2)     // Returns 1.4142135623730951
//	Sqrt(0)     // Returns 0.0
//	Sqrt(-1)    // Returns NaN (Not a Number)
//	Sqrt(25)    // Returns 5.0
func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

// Ceiling rounds a number up to the nearest integer or to the specified precision.
//
// Parameters:
//   - number: The number to round up
//   - precision: Optional. The number of decimal places to round to.
//     If not provided, rounds up to the nearest integer.
//
// Returns:
//   - float64: The rounded up number
//
// Examples:
//
//	Ceiling(4.3)      // Returns 5.0 (rounded up to nearest integer)
//	Ceiling(4.357, 2) // Returns 4.36 (rounded up to 2 decimal places)
//	Ceiling(4.352, 2) // Returns 4.36 (rounded up to 2 decimal places)
//	Ceiling(-4.7)     // Returns -4.0 (rounded up to nearest integer)
//	Ceiling(-4.7, 1)  // Returns -4.7 (rounded up to 1 decimal place)
func Ceiling(number float64, precision ...int) float64 {
	if len(precision) > 0 {
		factor := math.Pow(10, float64(precision[0]))
		return math.Ceil(number*factor) / factor
	}
	return math.Ceil(number)
}

// Format formats a number with grouped thousands and specified decimal places.
//
// Parameters:
//   - number: The number to format
//   - decimals: The number of decimal places to include
//   - decimalSeparator: The character to use as decimal separator
//   - thousandsSeparator: The character to use as thousands separator
//
// Returns:
//   - string: The formatted number as a string
//
// Examples:
//
//	Format(1234.5678, 2, ".", ",")    // Returns "1,234.57"
//	Format(1234567.89, 1, ",", " ")   // Returns "1 234 567,9"
//	Format(1000000, 0, ".", ",")      // Returns "1,000,000"
//	Format(-1234.56, 2, ".", ",")     // Returns "-1,234.56"
//	Format(0.5, 2, ".", ",")          // Returns "0.50"
func Format(number float64, decimals int, decimalSeparator, thousandsSeparator string) string {
	// If decimals is 0, truncate the number instead of rounding
	if decimals == 0 {
		number = math.Floor(number)
	}

	// Format the number with the specified number of decimal places
	formatStr := "%." + strconv.Itoa(decimals) + "f"
	formattedNumber := fmt.Sprintf(formatStr, number)

	// Split the number into integer and decimal parts
	parts := strings.Split(formattedNumber, ".")
	integerPart := parts[0]

	// Add thousands separator
	var result strings.Builder
	for i, char := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result.WriteString(thousandsSeparator)
		}
		result.WriteRune(char)
	}

	// Add decimal part if needed
	if decimals > 0 {
		result.WriteString(decimalSeparator)
		if len(parts) > 1 {
			result.WriteString(parts[1])
		} else {
			result.WriteString(strings.Repeat("0", decimals))
		}
	}

	return result.String()
}

// FormatPercentage formats a number as a percentage with the specified number of decimal places.
//
// Parameters:
//   - number: The number to format as a percentage (in decimal form, e.g., 0.5 for 50%)
//   - decimals: The number of decimal places to include
//
// Returns:
//   - string: The formatted percentage as a string with a % symbol
//
// Examples:
//
//	FormatPercentage(0.156, 1)  // Returns "15.6%"
//	FormatPercentage(0.5, 0)    // Returns "50%"
//	FormatPercentage(1.0, 2)    // Returns "100.00%"
//	FormatPercentage(0.0, 0)    // Returns "0%"
//	FormatPercentage(-0.25, 1)  // Returns "-25.0%"
func FormatPercentage(number float64, decimals int) string {
	return fmt.Sprintf("%.*f%%", decimals, number*100)
}

// Percent calculates what percentage one number is of another.
//
// Parameters:
//   - number: The numerator (the part)
//   - total: The denominator (the whole)
//   - decimals: Optional. The number of decimal places to round the result to.
//     If not provided, the result is not rounded.
//
// Returns:
//   - float64: The percentage value (number/total * 100), or 0 if total is 0
//
// Examples:
//
//	Percent(25, 100)    // Returns 25.0 (25% of 100)
//	Percent(1, 3, 2)    // Returns 33.33 (1/3 as a percentage, rounded to 2 decimal places)
//	Percent(1, 4)       // Returns 25.0 (1/4 as a percentage)
//	Percent(1, 0)       // Returns 0.0 (to avoid division by zero)
//	Percent(-10, 50, 1) // Returns -20.0 (negative percentage)
func Percent(number, total float64, decimals ...int) float64 {
	if total == 0 {
		return 0
	}

	percentage := (number / total) * 100

	if len(decimals) > 0 {
		factor := math.Pow(10, float64(decimals[0]))
		return math.Round(percentage*factor) / factor
	}

	return percentage
}

// Abbreviate formats a number to a compact form with K, M, B, T suffixes.
//
// Parameters:
//   - number: The number to format
//   - precision: Optional. The number of decimal places to include. Default is 0.
//
// Returns:
//   - string: The formatted number as a string with appropriate suffix (K for thousands,
//     M for millions, B for billions, T for trillions)
//
// Examples:
//
//	Abbreviate(1000)             // Returns "1K"
//	Abbreviate(489939)           // Returns "490K"
//	Abbreviate(1230000, 2)       // Returns "1.23M"
//	Abbreviate(1000000000)       // Returns "1B"
//	Abbreviate(1500000000000, 1) // Returns "1.5T"
//	Abbreviate(-1234567, 2)   // Returns "-1.23M"
func Abbreviate(number float64, precision ...int) string {
	absNumber := math.Abs(number)
	sign := ""
	if number < 0 {
		sign = "-"
	}

	// Default precision is 0 if not provided
	prec := 0
	if len(precision) > 0 {
		prec = precision[0]
	}

	switch {
	case absNumber >= 1_000_000_000_000:
		return sign + fmt.Sprintf("%.*f", prec, absNumber/1_000_000_000_000) + "T"
	case absNumber >= 1_000_000_000:
		return sign + fmt.Sprintf("%.*f", prec, absNumber/1_000_000_000) + "B"
	case absNumber >= 1_000_000:
		return sign + fmt.Sprintf("%.*f", prec, absNumber/1_000_000) + "M"
	case absNumber >= 1_000:
		return sign + fmt.Sprintf("%.*f", prec, absNumber/1_000) + "K"
	default:
		return sign + fmt.Sprintf("%.*f", prec, absNumber)
	}
}

// CurrencySymbol returns the symbol for the given currency code.
//
// Parameters:
//   - code: The ISO 4217 currency code (e.g., "USD", "EUR", "GBP")
//
// Returns:
//   - string: The currency symbol
func CurrencySymbol(code string) string {
	symbols := map[string]string{
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
		"CNY": "¥",
		"INR": "₹",
		"RUB": "₽",
		"BRL": "R$",
		"KRW": "₩",
		"AUD": "A$",
		"CAD": "C$",
		"CHF": "CHF",
		"HKD": "HK$",
		"SGD": "S$",
		"SEK": "kr",
		"NOK": "kr",
		"DKK": "kr",
		"PLN": "zł",
		"THB": "฿",
		"MXN": "Mex$",
		"ZAR": "R",
	}

	if symbol, ok := symbols[code]; ok {
		return symbol
	}
	return code
}

// LocaleInfo contains formatting information for a specific locale.
type LocaleInfo struct {
	DecimalSeparator   string
	ThousandsSeparator string
	SymbolPosition     string // "prefix" or "suffix"
}

// GetLocaleInfo returns formatting information for the given locale.
//
// Parameters:
//   - locale: The locale code (e.g., "en", "de", "fr")
//
// Returns:
//   - LocaleInfo: The formatting information for the locale
func GetLocaleInfo(locale string) LocaleInfo {
	locales := map[string]LocaleInfo{
		"en": {
			DecimalSeparator:   ".",
			ThousandsSeparator: ",",
			SymbolPosition:     "prefix",
		},
		"de": {
			DecimalSeparator:   ",",
			ThousandsSeparator: ".",
			SymbolPosition:     "suffix",
		},
		"fr": {
			DecimalSeparator:   ",",
			ThousandsSeparator: " ",
			SymbolPosition:     "suffix",
		},
		"es": {
			DecimalSeparator:   ",",
			ThousandsSeparator: ".",
			SymbolPosition:     "suffix",
		},
		"it": {
			DecimalSeparator:   ",",
			ThousandsSeparator: ".",
			SymbolPosition:     "suffix",
		},
		"nl": {
			DecimalSeparator:   ",",
			ThousandsSeparator: ".",
			SymbolPosition:     "prefix",
		},
		"pt": {
			DecimalSeparator:   ",",
			ThousandsSeparator: ".",
			SymbolPosition:     "prefix",
		},
		"ru": {
			DecimalSeparator:   ",",
			ThousandsSeparator: " ",
			SymbolPosition:     "suffix",
		},
		"ja": {
			DecimalSeparator:   ".",
			ThousandsSeparator: ",",
			SymbolPosition:     "prefix",
		},
		"zh": {
			DecimalSeparator:   ".",
			ThousandsSeparator: ",",
			SymbolPosition:     "prefix",
		},
	}

	if info, ok := locales[locale]; ok {
		return info
	}
	return locales["en"] // Default to English
}

// Currency formats a number as a currency with the specified currency code and locale.
//
// Parameters:
//   - number: The number to format as currency
//   - in: Optional. The currency code (e.g., "USD", "EUR"). Default is "USD".
//   - locale: Optional. The locale code (e.g., "en", "de"). Default is "en".
//   - precision: Optional. The number of decimal places. Default is 2.
//
// Returns:
//   - string: The formatted currency string
//
// Examples:
//
//	Number::currency(1000)                                // Returns "$1,000.00"
//	Number::currency(1000, in: "EUR")                     // Returns "€1,000.00"
//	Number::currency(1000, in: "EUR", locale: "de")       // Returns "1.000,00 €"
//	Number::currency(1000, in: "EUR", locale: "de", precision: 0) // Returns "1.000 €"
func Currency(number float64, options ...map[string]interface{}) string {
	// Default values
	currencyCode := "USD"
	locale := "en"
	precision := 2

	// Parse options
	if len(options) > 0 {
		for key, value := range options[0] {
			switch key {
			case "in":
				if code, ok := value.(string); ok && code != "" {
					currencyCode = code
				}
			case "locale":
				if loc, ok := value.(string); ok && loc != "" {
					locale = loc
				}
			case "precision":
				switch v := value.(type) {
				case int:
					if v >= 0 {
						precision = v
					}
				case float64:
					if v >= 0 {
						precision = int(v)
					}
				}
			}
		}
	}

	// Get currency symbol and locale info
	symbol := CurrencySymbol(currencyCode)
	localeInfo := GetLocaleInfo(locale)

	// Handle negative numbers
	isNegative := number < 0
	absNumber := math.Abs(number)

	// Round the number for precision 0 (especially for JPY)
	if precision == 0 {
		absNumber = math.Round(absNumber)
	}

	// Format the number
	formattedNumber := Format(absNumber, precision, localeInfo.DecimalSeparator, localeInfo.ThousandsSeparator)

	// Add negative sign and currency symbol
	if isNegative {
		if localeInfo.SymbolPosition == "prefix" {
			return "-" + symbol + formattedNumber
		}
		return "-" + formattedNumber + " " + symbol
	}

	// Add currency symbol for positive numbers
	if localeInfo.SymbolPosition == "prefix" {
		return symbol + formattedNumber
	}
	return formattedNumber + " " + symbol
}

// ForHumans converts a number to a human-readable string with the appropriate unit (thousand, million, billion, trillion).
//
// Parameters:
//   - number: The number to format
//   - precision: Optional. The number of decimal places to include. Default is 0.
//
// Returns:
//   - string: The formatted number as a string with appropriate unit
//
// Examples:
//
//	ForHumans(1000)             // Returns "1 thousand"
//	ForHumans(489939)           // Returns "490 thousand"
//	ForHumans(1230000, 2)       // Returns "1.23 million"
//	ForHumans(1000000000)       // Returns "1 billion"
//	ForHumans(1500000000000, 1) // Returns "1.5 trillion"
func ForHumans(number float64, precision ...int) string {
	// Default precision is 0 if not provided
	prec := 0
	if len(precision) > 0 {
		prec = precision[0]
	}

	absNumber := math.Abs(number)
	sign := ""
	if number < 0 {
		sign = "-"
	}

	units := []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}
	unitIndex := 0

	// Find the appropriate unit
	for absNumber >= 1000 && unitIndex < len(units)-1 {
		absNumber /= 1000
		unitIndex++
	}

	// Format with the specified precision
	if unitIndex == 0 {
		// No unit for numbers less than 1000
		return sign + fmt.Sprintf("%.*f", prec, absNumber)
	}

	// Format with the specified precision and add the unit
	return sign + fmt.Sprintf("%.*f", prec, absNumber) + " " + units[unitIndex]
}

// FileSize formats a byte size to a human-readable string with the appropriate unit (B, KB, MB, GB, TB, PB).
//
// Parameters:
//   - bytes: The size in bytes
//   - precision: Optional. The number of decimal places to include. Default is 0.
//
// Returns:
//   - string: The formatted file size as a string with appropriate unit
//
// Examples:
//
//	FileSize(1024)             // Returns "1 KB"
//	FileSize(1024 * 1024)      // Returns "1 MB"
//	FileSize(1024, 2)          // Returns "1.00 KB"
//	FileSize(1500)             // Returns "1 KB"
//	FileSize(1500, 2)          // Returns "1.46 KB"
//	FileSize(1500000)          // Returns "1 MB"
//	FileSize(1500000, 1)       // Returns "1.4 MB"
func FileSize(bytes float64, precision ...int) string {
	// Default precision is 0 if not provided
	prec := 0
	if len(precision) > 0 {
		prec = precision[0]
	}

	absBytes := math.Abs(bytes)
	sign := ""
	if bytes < 0 {
		sign = "-"
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	unitIndex := 0

	// Find the appropriate unit
	for absBytes >= 1024 && unitIndex < len(units)-1 {
		absBytes /= 1024
		unitIndex++
	}

	// Format with the specified precision
	return sign + fmt.Sprintf("%.*f", prec, absBytes) + " " + units[unitIndex]
}

// Ordinal converts a number to its ordinal representation.
//
// Parameters:
//   - number: The number to convert to an ordinal
//
// Returns:
//   - string: The ordinal representation of the number
//
// Examples:
//
//	Ordinal(1)  // Returns "1st"
//	Ordinal(2)  // Returns "2nd"
//	Ordinal(3)  // Returns "3rd"
//	Ordinal(4)  // Returns "4th"
//	Ordinal(11) // Returns "11th"
//	Ordinal(21) // Returns "21st"
//	Ordinal(22) // Returns "22nd"
//	Ordinal(23) // Returns "23rd"
func Ordinal(number int) string {
	if number < 0 {
		number = -number // Handle negative numbers
	}

	// Special cases for 11, 12, 13
	if number%100 >= 11 && number%100 <= 13 {
		return fmt.Sprintf("%d%s", number, "th")
	}

	// Handle other cases based on the last digit
	switch number % 10 {
	case 1:
		return fmt.Sprintf("%d%s", number, "st")
	case 2:
		return fmt.Sprintf("%d%s", number, "nd")
	case 3:
		return fmt.Sprintf("%d%s", number, "rd")
	default:
		return fmt.Sprintf("%d%s", number, "th")
	}
}

// Pairs splits a number into pairs of ranges based on a given chunk size.
//
// Parameters:
//   - total: The total number to split into pairs
//   - chunkSize: The size of each chunk
//   - options: Optional. A map containing additional options:
//   - "offset": The offset to use for calculating the end of each range. Default is -1.
//
// Returns:
//   - [][]int: An array of pairs, where each pair is an array of two integers representing
//     the start and end of a range
//
// Examples:
//
//	Pairs(25, 10)                  // Returns [[0, 9], [10, 19], [20, 25]]
//	Pairs(25, 10, map[string]int{"offset": 0}) // Returns [[0, 10], [10, 20], [20, 25]]
func Pairs(total, chunkSize int, options ...map[string]int) [][]int {
	// Default offset is -1 (end of range is chunkSize-1)
	offset := -1

	// Parse options
	if len(options) > 0 {
		if val, ok := options[0]["offset"]; ok {
			offset = val
		}
	}

	var result [][]int
	start := 0

	// If total is 0, return an empty slice
	if total == 0 {
		return result
	}

	for start < total {
		end := start + chunkSize + offset
		if end > total {
			end = total
		}
		result = append(result, []int{start, end})
		start += chunkSize
	}

	// Add the total as a separate pair if it's exactly divisible by chunkSize
	// and the last pair doesn't already include the total
	if start == total && len(result) > 0 && result[len(result)-1][1] != total {
		result = append(result, []int{total, total})
	}

	return result
}
