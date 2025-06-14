# fn - Functional Programming Utility Functions for Go

The `fn` package provides a comprehensive set of utility functions for functional programming in Go. It includes functions for function composition, transformation, rate limiting, memoization, and more.

## Installation

```bash
go get github.com/gflydev/utils/fn
```

## Usage

```go
import "github.com/gflydev/utils/fn"
```

## Functions

### Function Control

#### After

Creates a function that invokes the provided function once it's called n or more times.

```go
counter := 0
f := fn.After(3, func() int {
    counter++
    return counter
})

// First two calls return 0
result := f() // result: 0
result = f()  // result: 0

// Third call executes the function and returns 1
result = f()  // result: 1

// Fourth call returns 2
result = f()  // result: 2
```

#### Before

Creates a function that invokes the provided function at most n times.

```go
counter := 0
f := fn.Before(3, func() int {
    counter++
    return counter
})

// First two calls return incrementing values
result := f() // result: 1
result = f()  // result: 2

// Third and subsequent calls return the last result
result = f()  // result: 3
result = f()  // result: 3 (function not called again)
```

#### Once

Creates a function that is restricted to invoking the provided function only once. Repeated calls to the returned function will always return the result of the first call. This is useful for initialization operations that should only be performed once.

Parameters:
- `fn`: The function to restrict to a single invocation, which returns any type

Returns:
- A function that will only invoke the original function on the first call and return the cached result for subsequent calls

Note: The returned function is thread-safe and can be safely used in concurrent environments, ensuring the original function is called exactly once even with concurrent access.

```go
counter := 0
f := fn.Once(func() int {
    counter++
    return counter
})

// First call executes the function
result := f() // result: 1

// Subsequent calls return the same result
result = f()  // result: 1
result = f()  // result: 1

// The function is only called once
// counter == 1
```

Example with initialization:

```go
var initializeApp = fn.Once(func() bool {
    // Perform expensive initialization
    fmt.Println("Initializing application...")
    // Set up resources, connections, etc.
    return true
})

// Call at application startup points
// The initialization will only happen once
initialized := initializeApp() // Prints message and returns true
initialized = initializeApp()  // Returns true without printing message
```

### Function Transformation

#### Partial

Creates a function that invokes the provided function with the first argument fixed to the specified value. This is a form of partial application, where some arguments of a function are pre-filled, resulting in a new function that takes fewer arguments.

Parameters:
- `fn`: The function to partially apply, which takes two arguments of the same type and returns any type
- `partial`: The value to prepend to the argument list (the first argument to fix)

Returns:
- A function that invokes the original function with the fixed first argument and the provided second argument

```go
add := func(a, b int) int { return a + b }
add5 := fn.Partial(add, 5)

result := add5(3) // result: 8 (equivalent to add(5, 3))

// Another example with string concatenation
greet := func(greeting, name string) string { 
    return greeting + " " + name 
}
sayHello := fn.Partial(greet, "Hello")

result = sayHello("John") // result: "Hello John"
result = sayHello("Jane") // result: "Hello Jane"
```

This is useful for creating specialized versions of more general functions by fixing some of their parameters.

#### Rearg

Creates a function that invokes the provided function with arguments arranged according to the specified indexes. This simplified version swaps the first two arguments.

Parameters:
- `fn`: The function whose arguments to rearrange, which takes two arguments of the same type and returns any type

Returns:
- A function that invokes the original function with the first two arguments swapped

```go
subtract := func(a, b int) int { return a - b }
rearged := fn.Rearg(subtract)

// Original: subtract(3, 5) = 3 - 5 = -2
// Rearged: subtract(5, 3) = 5 - 3 = 2
result := rearged(3, 5) // result: 2

// Another example with division
divide := func(a, b float64) float64 { return a / b }
divideInverted := fn.Rearg(divide)

// Original: divide(10, 2) = 10 / 2 = 5
// Rearged: divide(2, 10) = 2 / 10 = 0.2
result := divideInverted(10, 2) // result: 0.2
```

This is useful when you need to change the order of arguments without creating a new function manually.

#### Wrap

Creates a function that provides the original function to the wrapper function as its first argument. This allows you to execute code before or after the original function, modify arguments or return values, or add additional functionality without modifying the original function.

Parameters:
- `fn`: The function to wrap, which takes one argument and returns any type
- `wrapper`: The wrapper function that receives the original function as its first argument and the input value as its second argument

Returns:
- A function that invokes the wrapper with the original function and the provided argument

```go
greet := func(name string) string { return "Hello, " + name }
wrapped := fn.Wrap(greet, func(greetFunc func(string) string, name string) string {
    // Add an exclamation mark to the result of the original function
    return greetFunc(name) + "!"
})

result := wrapped("World") // result: "Hello, World!"

// More complex example with logging
loggedAdd := fn.Wrap(
    func(x int) int { return x + 10 },
    func(addFunc func(int) int, x int) int {
        fmt.Printf("Calling with argument: %d\n", x)
        result := addFunc(x)
        fmt.Printf("Result: %d\n", result)
        return result
    },
)

result = loggedAdd(5) // Logs "Calling with argument: 5" and "Result: 15", returns 15
```

This pattern is useful for:
- Adding logging or metrics around function calls
- Implementing retry logic
- Adding validation or transformation of inputs/outputs
- Implementing decorators or middleware patterns

#### Negate

Creates a function that negates the result of the predicate function.

```go
isEven := func(x int) bool { return x%2 == 0 }
isOdd := fn.Negate(isEven)

result := isOdd(2) // result: false
result = isOdd(3)  // result: true
```

#### Spread

Creates a function that accepts an array of arguments and applies them to the provided function.

```go
add := func(a, b int) int { return a + b }
spreaded := fn.Spread(add)

result := spreaded([]int{2, 3}) // result: 5
```

### Rate Limiting

#### Debounce

Creates a debounced function that delays invoking the provided function until after wait milliseconds have elapsed since the last time the debounced function was invoked.

```go
counter := 0
f := func() { counter++ }
debounced := fn.Debounce(f, 50*time.Millisecond)

// Call multiple times in quick succession
debounced()
debounced()
debounced()

// Wait for the debounce period
time.Sleep(100 * time.Millisecond)

// Counter is incremented only once
// counter == 1
```

#### Throttle

Creates a throttled function that only invokes the provided function at most once per every wait duration. Unlike debounce, which resets the timer on each call, throttle ensures the function is called at a regular interval regardless of how often the throttled function is invoked.

Parameters:
- `fn`: The function to throttle, which takes no arguments and returns no value
- `wait`: The minimum duration between function invocations

Returns:
- A throttled function that will only execute at most once per wait duration

Note: The throttled function is thread-safe and can be safely used in concurrent environments.

```go
counter := 0
f := func() { counter++ }
throttled := fn.Throttle(f, 50*time.Millisecond)

// Call multiple times in quick succession
throttled()
throttled()
throttled()

// Counter is incremented only once
// counter == 1

// Wait for the throttle period
time.Sleep(100 * time.Millisecond)

// Call again
throttled()

// Counter is now 2
// counter == 2
```

This is useful for rate-limiting expensive operations that might be triggered rapidly, such as:
- Handling scroll or resize events in UI applications
- Limiting API calls
- Preventing button double-clicks

### Caching and Memoization

#### Memoize

Creates a function that memoizes (caches) the result of the provided function. If the function is called with the same arguments, the cached result is returned without executing the original function again. This is useful for expensive operations that are called repeatedly with the same inputs.

Parameters:
- `fn`: The function to memoize, which takes a comparable type as input and returns any type

Returns:
- A memoized version of the function that caches results based on input arguments

Note: The memoized function is thread-safe and can be safely used in concurrent environments.

```go
counter := 0
square := func(n int) int {
    counter++
    return n * n
}

memoizedSquare := fn.Memoize(square)

// First call computes the result
counter = 0
result := memoizedSquare(5) // result: 25
// counter == 1

// Second call with same argument uses cached result
counter = 0
result = memoizedSquare(5) // result: 25
// counter == 0

// Different argument computes a new result
counter = 0
result = memoizedSquare(6) // result: 36
// counter == 1
```

This is particularly useful for recursive functions like Fibonacci:

```go
var fibonacci func(int) int
fibonacciImpl := func(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
fibonacci = fn.Memoize(fibonacciImpl)

result := fibonacci(30) // Computes efficiently using memoization
```

### Timing and Delay

#### Delay

Invokes the provided function after the specified wait duration. This function schedules the execution of the function and returns immediately.

Parameters:
- `fn`: The function to delay
- `wait`: The duration to wait before invoking the function

```go
counter := 0
f := func() { counter++ }

// Function is scheduled to run after 50ms
fn.Delay(f, 50*time.Millisecond)

// Counter is still 0 immediately after calling Delay
// counter == 0

// After waiting, the function will have executed
time.Sleep(100*time.Millisecond)
// counter == 1
```

### Error Handling

#### Retry

Creates a function that retries the given function until it succeeds or reaches the maximum number of retries. This is particularly useful for operations that might fail temporarily, such as network requests or database operations.

Parameters:
- `fn`: The function to retry, which returns a value and an error
- `maxRetries`: The maximum number of retry attempts (in addition to the initial attempt)
- `delay`: The duration to wait between retry attempts

Returns:
- A function that will retry the original function up to maxRetries times with the specified delay between attempts

Note: The function returns immediately upon success (nil error). If all attempts fail, it returns the last error encountered.

```go
counter := 0
f := func() (string, error) {
    counter++
    if counter < 3 {
        return "", errors.New("temporary error")
    }
    return "success", nil
}

retried := fn.Retry(f, 5, 10*time.Millisecond)

result, err := retried()
// result: "success", err: nil
// counter == 3 (function was called 3 times before succeeding)

// Example with API request
fetchData := fn.Retry(
    func() ([]byte, error) {
        resp, err := http.Get("https://api.example.com/data")
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()
        return io.ReadAll(resp.Body)
    },
    3,
    500*time.Millisecond,
)

data, err := fetchData()
// Retries up to 3 times with 500ms delay between attempts
```

This is useful for:
- Network requests that might fail due to temporary connectivity issues
- Database operations that might encounter transient errors
- Any operation that might succeed on a subsequent attempt after failing initially

### Function Composition

#### Compose

Creates a function that is the composition of the provided functions, where each function consumes the return value of the function that follows. The last function is invoked with the arguments of the resulting function.

```go
double := func(x int) int { return x * 2 }
addOne := func(x int) int { return x + 1 }
composed := fn.Compose(double, addOne)

// Compose applies functions from right to left: double(addOne(5)) = double(6) = 12
result := composed(5) // result: 12
```

#### Pipe

Creates a function that is the composition of the provided functions, where each function consumes the return value of the previous function. The first function is invoked with the arguments of the resulting function.

```go
double := func(x int) int { return x * 2 }
addOne := func(x int) int { return x + 1 }
piped := fn.Pipe(double, addOne)

// Pipe applies functions from left to right: addOne(double(5)) = addOne(10) = 11
result := piped(5) // result: 11
```

### List Transformation

#### TransformList

Transforms each element in a list using the provided transformer function.

Parameters:
- `records`: The slice of elements to transform
- `transformerFn`: The function to apply to each element

Returns:
- `[]R`: A new slice containing the transformed elements

```go
// Transform integers to strings
input := []int{1, 2, 3}
result := fn.TransformList(input, strconv.Itoa)
// result: []string{"1", "2", "3"}

// Transform with calculation
input = []int{1, 2, 3}
result = fn.TransformList(input, func(i int) string {
    return strconv.Itoa(i * 2)
})
// result: []string{"2", "4", "6"}

// Square numbers
numbers := []int{1, 2, 3}
squares := fn.TransformList(numbers, func(x int) int { 
    return x * x 
})
// squares = []int{1, 4, 9}
```

#### TransformMap

Transforms a map of one type to a map of another type using a transformer function.

Parameters:
- `m`: The map to transform
- `transformerFn`: The function to apply to each value in the map

Returns:
- `map[K]R`: A new map with the same keys but transformed values

```go
// Transform integers to strings
input := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
result := fn.TransformMap(input, strconv.Itoa)
// result: map[string]string{"a": "1", "b": "2", "c": "3"}

// Transform with calculation
input = map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}
result = fn.TransformMap(input, func(i int) string {
    return strconv.Itoa(i * 2)
})
// result: map[string]string{"a": "2", "b": "4", "c": "6"}

// Double ages
ages := map[string]int{"John": 30, "Jane": 25}
doubled := fn.TransformMap(ages, func(age int) int { 
    return age * 2 
})
// doubled = map[string]int{"John": 60, "Jane": 50}
```

#### TransformListWithError

Transforms a slice and collects any errors that occur during transformation.

Parameters:
- `records`: The slice of elements to transform
- `transformerFn`: The function to apply to each element, which may return an error

Returns:
- `[]R`: A new slice containing the successfully transformed elements
- `[]error`: A slice containing any errors that occurred during transformation

```go
// No errors
input := []int{1, 2, 3}
result, errs := fn.TransformListWithError(input, func(i int) (string, error) {
    return strconv.Itoa(i), nil
})
// result: []string{"1", "2", "3"}, errs: [] (empty)

// Some errors
input = []int{1, 2, 3, 4, 5}
result, errs = fn.TransformListWithError(input, func(i int) (string, error) {
    if i%2 == 0 {
        return "", errors.New("even number error")
    }
    return strconv.Itoa(i), nil
})
// result: []string{"1", "3", "5"}, len(errs): 2

// Parsing strings to integers
numbers := []string{"1", "2", "abc", "3"}
parsed, errors := fn.TransformListWithError(numbers, func(s string) (int, error) { 
    return strconv.Atoi(s) 
})
// parsed = []int{1, 2, 3}, errors contains the error from parsing "abc"
```

#### TransformConcurrent

Transforms a slice concurrently using a specified number of workers.

Parameters:
- `records`: The slice of elements to transform
- `transformerFn`: The function to apply to each element
- `numWorkers`: The number of concurrent workers to use for transformation

Returns:
- `[]R`: A new slice containing the transformed elements

```go
// Transform with multiple workers
input := []int{1, 2, 3, 4, 5, 6}
result := fn.TransformConcurrent(input, strconv.Itoa, 3)
// result: []string{"1", "2", "3", "4", "5", "6"}

// Square numbers concurrently
numbers := []int{1, 2, 3, 4, 5}
squares := fn.TransformConcurrent(numbers, func(x int) int { 
    return x * x 
}, 2)
// squares = []int{1, 4, 9, 16, 25}
```

Note: If the number of records is less than the number of workers or if `numWorkers` is 1, the function will fall back to the sequential `TransformList` implementation for better efficiency.

#### TransformBatch

Transforms a slice in batches and returns the combined results.

Parameters:
- `records`: The slice of elements to transform
- `transformerFn`: The function to apply to each batch of elements
- `batchSize`: The size of each batch

Returns:
- `[]R`: A new slice containing the combined results of all batch transformations

```go
// Transform in batches
input := []int{1, 2, 3, 4, 5}
result := fn.TransformBatch(input, func(batch []int) []string {
    var result []string
    for _, v := range batch {
        result = append(result, strconv.Itoa(v))
    }
    return result
}, 2)
// result: []string{"1", "2", "3", "4", "5"}

// Double numbers in batches
numbers := []int{1, 2, 3, 4, 5}
doubled := fn.TransformBatch(numbers, func(batch []int) []int {
    var result []int
    for _, n := range batch {
        result = append(result, n*2)
    }
    return result
}, 2)
// doubled = []int{2, 4, 6, 8, 10}
```

Note: If `batchSize` is less than or equal to 0, a default batch size of 100 will be used.

## License

This package is licensed under the MIT License - see the LICENSE file for details.
