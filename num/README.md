# num - Number Utility Functions for Go

The `num` package provides a comprehensive set of utility functions for working with numbers in Go. It offers functions for basic math operations, formatting, statistical calculations, and more.

## Installation

```
go get github.com/gflydev/utils/num
```

## Usage

```
import "github.com/gflydev/utils/num"
```

## Functions

### Basic Math Operations

#### Clamp

Constrains a number between lower and upper bounds.

```
result := num.Clamp(5, 0, 10)
// result: 5 (within bounds)

result := num.Clamp(-5, 0, 10)
// result: 0 (below lower bound)

result := num.Clamp(15, 0, 10)
// result: 10 (above upper bound)

result := num.Clamp(5, 10, 0)
// result: 5 (lower > upper, bounds are swapped)
```

#### InRange

Checks if a number is between start and end (inclusive).

```
result := num.InRange(3, 2, 4)
// result: true (within range)

result := num.InRange(2, 2, 4)
// result: true (at lower bound)

result := num.InRange(4, 2, 4)
// result: true (at upper bound)

result := num.InRange(1, 2, 4)
// result: false (below range)

result := num.InRange(5, 2, 4)
// result: false (above range)

result := num.InRange(3, 4, 2)
// result: true (start > end, bounds are swapped)
```

#### Random

Returns a random integer between min and max (inclusive).

```
result := num.Random(1, 10)
// result: a random number between 1 and 10
```

#### Round

Rounds a number to the nearest integer or to the specified precision.

```
result := num.Round(4.7)
// result: 5 (rounded to nearest integer)

result := num.Round(4.3)
// result: 4 (rounded to nearest integer)

result := num.Round(4.5)
// result: 5 (rounded to nearest integer)

result := num.Round(-4.7)
// result: -5 (rounded to nearest integer)

result := num.Round(-4.3)
// result: -4 (rounded to nearest integer)

// With precision
result := num.Round(4.7, 1)
// result: 4.7 (rounded to 1 decimal place)

result := num.Round(4.75, 1)
// result: 4.8 (rounded to 1 decimal place)

result := num.Round(4.749, 2)
// result: 4.75 (rounded to 2 decimal places)
```

#### Floor

Rounds a number down to the nearest integer or to the specified precision.

**Parameters:**
- `n` (float64): The number to round down
- `precision` (int, optional): The number of decimal places to round to. If not provided, rounds down to the nearest integer.

**Returns:**
- (float64): The rounded down number

**Examples:**

```
result := num.Floor(4.7)
// result: 4 (rounded down to nearest integer)

result := num.Floor(4.3)
// result: 4 (rounded down to nearest integer)

result := num.Floor(4.0)
// result: 4 (rounded down to nearest integer)

result := num.Floor(-4.7)
// result: -5 (rounded down to nearest integer)

result := num.Floor(-4.3)
// result: -5 (rounded down to nearest integer)

// With precision
result := num.Floor(4.78, 1)
// result: 4.7 (rounded down to 1 decimal place)

result := num.Floor(4.753, 2)
// result: 4.75 (rounded down to 2 decimal places)

result := num.Floor(-4.78, 1)
// result: -4.8 (rounded down to 1 decimal place)
```

#### Ceil

Rounds a number up to the nearest integer or to the specified precision.

**Parameters:**
- `n` (float64): The number to round up
- `precision` (int, optional): The number of decimal places to round to. If not provided, rounds up to the nearest integer.

**Returns:**
- (float64): The rounded up number

**Examples:**

```
result := num.Ceil(4.7)
// result: 5 (rounded up to nearest integer)

result := num.Ceil(4.3)
// result: 5 (rounded up to nearest integer)

result := num.Ceil(4.0)
// result: 4 (rounded up to nearest integer)

result := num.Ceil(-4.7)
// result: -4 (rounded up to nearest integer)

result := num.Ceil(-4.3)
// result: -4 (rounded up to nearest integer)

// With precision
result := num.Ceil(4.78, 1)
// result: 4.8 (rounded up to 1 decimal place)

result := num.Ceil(4.753, 2)
// result: 4.76 (rounded up to 2 decimal places)

result := num.Ceil(-4.78, 1)
// result: -4.7 (rounded up to 1 decimal place)
```

#### Ceiling

Rounds a number up to the nearest integer or to the specified precision (alias for Ceil).

**Parameters:**
- `number` (float64): The number to round up
- `precision` (int, optional): The number of decimal places to round to. If not provided, rounds up to the nearest integer.

**Returns:**
- (float64): The rounded up number

**Examples:**

```
result := num.Ceiling(4.3)
// result: 5 (rounded up to nearest integer)

result := num.Ceiling(4.7)
// result: 5 (rounded up to nearest integer)

result := num.Ceiling(-4.3)
// result: -4 (rounded up to nearest integer)

result := num.Ceiling(-4.7)
// result: -4 (rounded up to nearest integer)

result := num.Ceiling(4.357, 2)
// result: 4.36 (rounded up to 2 decimal places)

result := num.Ceiling(4.351, 2)
// result: 4.36 (rounded up to 2 decimal places)

result := num.Ceiling(-4.357, 2)
// result: -4.35 (rounded up to 2 decimal places)
```

#### Max

Returns the maximum value from a list of numbers.

**Parameters:**
- `numbers` (variadic float64): A variadic list of float64 numbers

**Returns:**
- (float64): The maximum value from the list, or 0 if the list is empty

**Examples:**

```
result := num.Max(1, 2, 3)
// result: 3

result := num.Max(3, 2, 1)
// result: 3

result := num.Max(-1, -2, -3)
// result: -1

result := num.Max()
// result: 0 (empty list)
```

#### Min

Returns the minimum value from a list of numbers.

**Parameters:**
- `numbers` (variadic float64): A variadic list of float64 numbers

**Returns:**
- (float64): The minimum value from the list, or 0 if the list is empty

**Examples:**

```
result := num.Min(1, 2, 3)
// result: 1

result := num.Min(3, 2, 1)
// result: 1

result := num.Min(-1, -2, -3)
// result: -3

result := num.Min(7.5, 3.2, 9.8)
// result: 3.2

result := num.Min()
// result: 0 (empty list)
```

#### Sum

Calculates the sum of all numbers in the provided list.

**Parameters:**
- `numbers` (variadic float64): A variadic list of float64 numbers

**Returns:**
- (float64): The sum of all numbers in the list, or 0 if the list is empty

**Examples:**

```
result := num.Sum(1, 2, 3)
// result: 6

result := num.Sum(-1, 5, 3)
// result: 7

result := num.Sum(-1, -2, -3)
// result: -6

result := num.Sum(7.5, 3.2, 9.8)
// result: 20.5

result := num.Sum()
// result: 0 (empty list)
```

#### Mean

Calculates the arithmetic mean (average) of a list of numbers.

**Parameters:**
- `numbers` (variadic float64): A variadic list of float64 numbers

**Returns:**
- (float64): The arithmetic mean of the numbers, or 0 if the list is empty

**Examples:**

```
result := num.Mean(1, 2, 3)
// result: 2

result := num.Mean(1, 3, 5, 7)
// result: 4

result := num.Mean(2, 4, 6, 8, 10)
// result: 6

result := num.Mean()
// result: 0 (empty list)
```

#### Abs

Returns the absolute value of a number.

**Parameters:**
- `n` (float64): The number to get the absolute value of

**Returns:**
- (float64): The absolute value of the number

**Examples:**

```
result := num.Abs(5)
// result: 5

result := num.Abs(-5)
// result: 5

result := num.Abs(-3.14)
// result: 3.14

result := num.Abs(0)
// result: 0
```

#### Pow

Returns the base raised to the exponent power.

**Parameters:**
- `base` (float64): The base number
- `exponent` (float64): The exponent to raise the base to

**Returns:**
- (float64): The result of base^exponent

**Examples:**

```
result := num.Pow(2, 3)
// result: 8 (2^3)

result := num.Pow(3, 2)
// result: 9 (3^2)

result := num.Pow(2, 0)
// result: 1 (any number raised to 0 is 1)

result := num.Pow(0, 2)
// result: 0 (0 raised to any positive power is 0)

result := num.Pow(-2, 2)
// result: 4 (-2^2)

result := num.Pow(-2, 3)
// result: -8 (-2^3)

result := num.Pow(4, 0.5)
// result: 2.0 (square root of 4)
```

#### Sqrt

Returns the square root of a number.

**Parameters:**
- `n` (float64): The number to calculate the square root of

**Returns:**
- (float64): The square root of the number

**Examples:**

```
result := num.Sqrt(9)
// result: 3

result := num.Sqrt(4)
// result: 2

result := num.Sqrt(2)
// result: 1.4142135623730951

result := num.Sqrt(0)
// result: 0

result := num.Sqrt(-1)
// result: NaN (Not a Number)

result := num.Sqrt(25)
// result: 5
```

### Collection Operations

#### MaxBy

Returns the element from a collection that produces the maximum value when passed through the iteratee function.

**Parameters:**
- `collection` ([]T): A slice of any type T
- `iteratee` (func(T) float64): A function that takes an element of type T and returns a float64

**Returns:**
- (T): The element that produces the maximum value, or zero value of T if collection is empty

**Examples:**

```
// Find the number with the largest square
result := num.MaxBy([]int{1, 2, 3, 4, 5}, func(n int) float64 { return float64(n * n) })
// result: 5

// Find the person with the highest age
type person struct {
    name string
    age  int
}
people := []person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}}
result := num.MaxBy(people, func(p person) float64 { return float64(p.age) })
// result: person{"Bob", 30}

// Find the person with the longest name
result := num.MaxBy(people, func(p person) float64 { return float64(len(p.name)) })
// result: person{"Charlie", 20}

// Empty collection
result := num.MaxBy([]person{}, func(p person) float64 { return float64(p.age) })
// result: person{} (zero value)
```

#### MinBy

Returns the element from a collection that produces the minimum value when passed through the iteratee function.

**Parameters:**
- `collection` ([]T): A slice of any type T
- `iteratee` (func(T) float64): A function that takes an element of type T and returns a float64

**Returns:**
- (T): The element that produces the minimum value, or zero value of T if collection is empty

**Examples:**

```
// Find the number with the smallest square
result := num.MinBy([]int{1, 2, 3, 4, 5}, func(n int) float64 { return float64(n * n) })
// result: 1

// Find the person with the lowest age
type person struct {
    name string
    age  int
}
people := []person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}}
result := num.MinBy(people, func(p person) float64 { return float64(p.age) })
// result: person{"Charlie", 20}

// Find the person with the shortest name
result := num.MinBy(people, func(p person) float64 { return float64(len(p.name)) })
// result: person{"Bob", 30}

// Empty collection
result := num.MinBy([]person{}, func(p person) float64 { return float64(p.age) })
// result: person{} (zero value)
```

#### SumBy

Calculates the sum of values in a collection after applying the iteratee function to each element.

**Parameters:**
- `collection` ([]T): A slice of any type T
- `iteratee` (func(T) float64): A function that takes an element of type T and returns a float64

**Returns:**
- (float64): The sum of all values after applying the iteratee function, or 0 if the collection is empty

**Examples:**

```
// Sum of ages
type person struct {
    name string
    age  int
}
people := []person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}}
result := num.SumBy(people, func(p person) float64 { return float64(p.age) })
// result: 75

// Sum of name lengths
result := num.SumBy(people, func(p person) float64 { return float64(len(p.name)) })
// result: 15

// Sum of doubled numbers
result := num.SumBy([]int{1, 2, 3, 4, 5}, func(n int) float64 { return float64(n * 2) })
// result: 30

// Empty collection
result := num.SumBy([]person{}, func(p person) float64 { return float64(p.age) })
// result: 0
```

#### MeanBy

Calculates the arithmetic mean (average) of values in a collection after applying the iteratee function to each element.

**Parameters:**
- `collection` ([]T): A slice of any type T
- `iteratee` (func(T) float64): A function that takes an element of type T and returns a float64

**Returns:**
- (float64): The arithmetic mean of all values after applying the iteratee function, or 0 if the collection is empty

**Examples:**

```
// Mean of ages
type person struct {
    name string
    age  int
}
people := []person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 20}}
result := num.MeanBy(people, func(p person) float64 { return float64(p.age) })
// result: 25

// Mean of name lengths
result := num.MeanBy(people, func(p person) float64 { return float64(len(p.name)) })
// result: 5

// Mean of doubled numbers
result := num.MeanBy([]int{1, 2, 3, 4, 5}, func(n int) float64 { return float64(n * 2) })
// result: 6

// Empty collection
result := num.MeanBy([]person{}, func(p person) float64 { return float64(p.age) })
// result: 0
```

### Formatting Functions

#### Format

Formats a number with grouped thousands and specified decimal places.

**Parameters:**
- `number` (float64): The number to format
- `decimals` (int): The number of decimal places to include
- `decimalSeparator` (string): The character to use as decimal separator
- `thousandsSeparator` (string): The character to use as thousands separator

**Returns:**
- (string): The formatted number as a string

**Examples:**

```
result := num.Format(1234.5678, 2, ".", ",")
// result: "1,234.57"

result := num.Format(1234567.89, 1, ",", " ")
// result: "1 234 567,9"

result := num.Format(1000000, 0, ".", ",")
// result: "1,000,000"

result := num.Format(-1234.5678, 2, ".", ",")
// result: "-1,234.57"

result := num.Format(0.5678, 3, ".", ",")
// result: "0.568"
```

#### FormatPercentage

Formats a number as a percentage with the specified number of decimal places.

**Parameters:**
- `number` (float64): The number to format as a percentage (in decimal form, e.g., 0.5 for 50%)
- `decimals` (int): The number of decimal places to include

**Returns:**
- (string): The formatted percentage as a string with a % symbol

**Examples:**

```
result := num.FormatPercentage(0.156, 1)
// result: "15.6%"

result := num.FormatPercentage(0.5, 0)
// result: "50%"

result := num.FormatPercentage(1, 2)
// result: "100.00%"

result := num.FormatPercentage(0, 1)
// result: "0.0%"

result := num.FormatPercentage(-0.25, 0)
// result: "-25%"
```

#### Percent

Calculates what percentage one number is of another.

**Parameters:**
- `number` (float64): The numerator (the part)
- `total` (float64): The denominator (the whole)
- `decimals` (int, optional): The number of decimal places to round the result to. If not provided, the result is not rounded.

**Returns:**
- (float64): The percentage value (number/total * 100), or 0 if total is 0

**Examples:**

```
result := num.Percent(25, 100)
// result: 25.0 (25% of 100)

result := num.Percent(1, 3, 2)
// result: 33.33 (1/3 as a percentage, rounded to 2 decimal places)

result := num.Percent(1, 0)
// result: 0 (to avoid division by zero)

result := num.Percent(0, 100)
// result: 0

result := num.Percent(50, 200, 1)
// result: 25.0

result := num.Percent(200, 50)
// result: 400

result := num.Percent(-25, 100)
// result: -25
```

#### Abbreviate

Formats a number to a compact form with K, M, B, T suffixes.

```
result := num.Abbreviate(1000)
// result: "1K"

result := num.Abbreviate(489939)
// result: "490K"

result := num.Abbreviate(1230000, 2)
// result: "1.23M"

result := num.Abbreviate(1000000000)
// result: "1B"

result := num.Abbreviate(1500000000000, 1)
// result: "1.5T"

result := num.Abbreviate(999)
// result: "999"

result := num.Abbreviate(-1234567, 1)
// result: "-1.2M"
```

#### ForHumans

Converts a number to a human-readable string with the appropriate unit (thousand, million, billion, trillion).

```
result := num.ForHumans(1000)
// result: "1 thousand"

result := num.ForHumans(489939)
// result: "490 thousand"

result := num.ForHumans(1230000, 2)
// result: "1.23 million"

result := num.ForHumans(1000000000)
// result: "1 billion"

result := num.ForHumans(1500000000000, 1)
// result: "1.5 trillion"

result := num.ForHumans(-1234567, 1)
// result: "-1.2 million"
```

#### FileSize

Formats a byte size to a human-readable string with the appropriate unit (B, KB, MB, GB, TB, PB).

```
result := num.FileSize(1024)
// result: "1 KB"

result := num.FileSize(1024 * 1024)
// result: "1 MB"

result := num.FileSize(1024, 2)
// result: "1.00 KB"

result := num.FileSize(1500)
// result: "1 KB"

result := num.FileSize(1500, 2)
// result: "1.46 KB"

result := num.FileSize(1500000)
// result: "1 MB"

result := num.FileSize(1500000, 1)
// result: "1.4 MB"
```

### Currency and Locale Functions

#### CurrencySymbol

Returns the symbol for the given currency code.

```
result := num.CurrencySymbol("USD")
// result: "$"

result := num.CurrencySymbol("EUR")
// result: "€"

result := num.CurrencySymbol("GBP")
// result: "£"

result := num.CurrencySymbol("JPY")
// result: "¥"

result := num.CurrencySymbol("XYZ")
// result: "XYZ" (unknown currency code returns the code itself)
```

#### GetLocaleInfo

Returns formatting information for the given locale.

```
result := num.GetLocaleInfo("en")
// result: LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"}

result := num.GetLocaleInfo("de")
// result: LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: ".", SymbolPosition: "suffix"}

result := num.GetLocaleInfo("fr")
// result: LocaleInfo{DecimalSeparator: ",", ThousandsSeparator: " ", SymbolPosition: "suffix"}

result := num.GetLocaleInfo("xx")
// result: LocaleInfo{DecimalSeparator: ".", ThousandsSeparator: ",", SymbolPosition: "prefix"} (defaults to English)
```

#### Currency

Formats a number as a currency with the specified currency code and locale.

```
result := num.Currency(1000)
// result: "$1,000.00"

result := num.Currency(1000, map[string]interface{}{"in": "EUR"})
// result: "€1,000.00"

result := num.Currency(1000, map[string]interface{}{"in": "EUR", "locale": "de"})
// result: "1.000,00 €"

result := num.Currency(1000, map[string]interface{}{"in": "EUR", "locale": "de", "precision": 0})
// result: "1.000 €"

result := num.Currency(1234.56, map[string]interface{}{"in": "GBP"})
// result: "£1,234.56"

result := num.Currency(-1234.56, map[string]interface{}{"in": "EUR", "locale": "de"})
// result: "-1.234,56 €"
```

### Utility Functions

#### Ordinal

Converts a number to its ordinal representation.

```
result := num.Ordinal(1)
// result: "1st"

result := num.Ordinal(2)
// result: "2nd"

result := num.Ordinal(21)
// result: "21st"

result := num.Ordinal(3)
// result: "3rd"

result := num.Ordinal(4)
// result: "4th"

result := num.Ordinal(11)
// result: "11th"

result := num.Ordinal(12)
// result: "12th"

result := num.Ordinal(13)
// result: "13th"
```

#### Pairs

Splits a number into pairs of ranges based on a given chunk size.

```
result := num.Pairs(25, 10)
// result: [[0, 9], [10, 19], [20, 25]]

result := num.Pairs(25, 10, map[string]int{"offset": 0})
// result: [[0, 10], [10, 20], [20, 25]]

result := num.Pairs(10, 5)
// result: [[0, 4], [5, 9], [10, 10]]

result := num.Pairs(10, 5, map[string]int{"offset": 0})
// result: [[0, 5], [5, 10]]
```

## License

This package is licensed under the MIT License - see the LICENSE file for details.
