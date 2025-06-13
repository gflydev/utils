// Package fn provides utility functions for function manipulation.
package fn

import (
	"sync"
	"time"
)

// After creates a function that invokes func once it's called n or more times.
//
// Parameters:
//   - n: The number of calls before invoking the function
//   - fn: The function to invoke after n calls
//
// Returns:
//   - func() T: A function that will invoke fn after being called n or more times
//
// Example: fn := After(3, func() { fmt.Println("done") }); fn(); fn(); fn() // prints "done" on the third call
func After[T any](n int, fn func() T) func() T {
	var (
		count int
		mu    sync.Mutex
	)

	return func() T {
		mu.Lock()
		defer mu.Unlock()

		count++
		if count >= n {
			return fn()
		}

		var zero T
		return zero
	}
}

// Before creates a function that invokes func, with the arguments provided, at most n times.
//
// Parameters:
//   - n: The maximum number of times to invoke the function
//   - fn: The function to invoke at most n times
//
// Returns:
//   - func() T: A function that will invoke fn at most n times and return the last result afterwards
//
// Example: fn := Before(3, func() { fmt.Println("called") }); fn(); fn(); fn(); fn() // prints "called" only 3 times
func Before[T any](n int, fn func() T) func() T {
	var (
		count  int
		result T
		mu     sync.Mutex
	)

	return func() T {
		mu.Lock()
		defer mu.Unlock()

		if count < n {
			result = fn()
			count++
		}

		return result
	}
}

// Curry creates a function that accepts arguments of func and either invokes func returning its result,
// if at least arity number of arguments have been provided, or returns a function that accepts the remaining arguments.
//
// Parameters:
//   - fn: The function to curry
//   - arity: The number of arguments the function accepts
//
// Returns:
//   - func(T) func(T) R: A curried function that accepts the first argument and returns a function that accepts the second argument
//
// Example: add := func(a, b int) int { return a + b }; addCurried := Curry(add, 2); add1 := addCurried(1); add1(2) -> 3
func Curry[T, R any](fn func(T, T) R, arity int) func(T) func(T) R {
	return func(a T) func(T) R {
		return func(b T) R {
			return fn(a, b)
		}
	}
}

// Debounce creates a debounced function that delays invoking func until after wait milliseconds have elapsed
// since the last time the debounced function was invoked.
//
// Parameters:
//   - fn: The function to debounce
//   - wait: The duration to wait before invoking the function
//
// Returns:
//   - func(): A debounced function that will only execute after wait duration has passed since its last invocation
//
// Example: fn := Debounce(func() { fmt.Println("called") }, 100); fn(); fn(); fn() // prints "called" only once after 100ms
func Debounce(fn func(), wait time.Duration) func() {
	var timer *time.Timer
	var mu sync.Mutex

	return func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(wait, fn)
	}
}

// Delay invokes func after wait milliseconds.
//
// Parameters:
//   - fn: The function to delay
//   - wait: The duration to wait before invoking the function
//
// Example: Delay(func() { fmt.Println("called") }, 100) // prints "called" after 100ms
func Delay(fn func(), wait time.Duration) {
	time.AfterFunc(wait, fn)
}

// Memoize creates a function that memoizes the result of func.
//
// Parameters:
//   - fn: The function to memoize
//
// Returns:
//   - func(T) R: A memoized function that caches its results based on the input arguments
//
// Example: fibonacci := Memoize(func(n int) int { if n <= 1 { return n }; return fibonacci(n-1) + fibonacci(n-2) })
func Memoize[T comparable, R any](fn func(T) R) func(T) R {
	cache := make(map[T]R)
	var mu sync.Mutex

	return func(arg T) R {
		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[arg]; ok {
			return val
		}

		result := fn(arg)
		cache[arg] = result
		return result
	}
}

// Once creates a function that is restricted to invoking func once.
//
// Parameters:
//   - fn: The function to restrict to a single invocation
//
// Returns:
//   - func() T: A function that will only invoke fn on the first call and return the result for subsequent calls
//
// Example: initialize := Once(func() { fmt.Println("initialized") }); initialize(); initialize() // prints "initialized" only once
func Once[T any](fn func() T) func() T {
	var (
		done   bool
		result T
		mu     sync.Mutex
	)

	return func() T {
		mu.Lock()
		defer mu.Unlock()

		if !done {
			result = fn()
			done = true
		}

		return result
	}
}

// Partial creates a function that invokes func with partials prepended to the arguments it receives.
//
// Parameters:
//   - fn: The function to partially apply
//   - partial: The value to prepend to the argument list
//
// Returns:
//   - func(T) R: A function that invokes fn with partial as the first argument and the provided argument as the second
//
// Example: greet := func(greeting, name string) string { return greeting + " " + name }; sayHello := Partial(greet, "Hello"); sayHello("John") -> "Hello John"
func Partial[T, R any](fn func(T, T) R, partial T) func(T) R {
	return func(arg T) R {
		return fn(partial, arg)
	}
}

// Rearg creates a function that invokes func with arguments arranged according to the specified indexes.
// This is a simplified version that swaps the first two arguments.
//
// Parameters:
//   - fn: The function whose arguments to rearrange
//
// Returns:
//   - func(T, T) R: A function that invokes fn with the first two arguments swapped
//
// Example: multiply := func(a, b int) int { return a * b }; divideInstead := Rearg(multiply); divideInstead(10, 2) -> 20
func Rearg[T, R any](fn func(T, T) R) func(T, T) R {
	return func(a, b T) R {
		return fn(b, a)
	}
}

// Throttle creates a throttled function that only invokes func at most once per every wait milliseconds.
//
// Parameters:
//   - fn: The function to throttle
//   - wait: The minimum duration between function invocations
//
// Returns:
//   - func(): A throttled function that will only execute at most once per wait duration
//
// Example: fn := Throttle(func() { fmt.Println("called") }, 100*time.Millisecond); fn(); fn(); fn() // prints "called" only once per 100ms
func Throttle(fn func(), wait time.Duration) func() {
	var (
		lastInvoke time.Time
		mu         sync.Mutex
	)

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if lastInvoke.IsZero() || now.Sub(lastInvoke) >= wait {
			lastInvoke = now
			fn()
		}
	}
}

// Wrap creates a function that provides value to the wrapper function as its first argument.
//
// Parameters:
//   - fn: The function to wrap
//   - wrapper: The wrapper function that receives fn as its first argument
//
// Returns:
//   - func(T) S: A function that invokes wrapper with fn and the provided argument
//
// Example: hello := func(name string) string { return "Hello " + name }; withExclamation := Wrap(hello, func(hello func(string) string, name string) string { return hello(name) + "!" }); withExclamation("John") -> "Hello John!"
func Wrap[T, R, S any](fn func(T) R, wrapper func(func(T) R, T) S) func(T) S {
	return func(arg T) S {
		return wrapper(fn, arg)
	}
}

// Retry creates a function that retries the given function until it succeeds or reaches the maximum number of retries.
//
// Parameters:
//   - fn: The function to retry
//   - maxRetries: The maximum number of retry attempts
//   - delay: The duration to wait between retry attempts
//
// Returns:
//   - func() (T, error): A function that will retry fn up to maxRetries times with the specified delay between attempts
//
// Example: fn := Retry(func() (int, error) { return 0, errors.New("error") }, 3, 100*time.Millisecond); fn() // retries 3 times with 100ms delay
func Retry[T any](fn func() (T, error), maxRetries int, delay time.Duration) func() (T, error) {
	return func() (T, error) {
		var lastErr error
		var zero T

		for i := 0; i <= maxRetries; i++ {
			result, err := fn()
			if err == nil {
				return result, nil
			}

			lastErr = err
			if i < maxRetries {
				time.Sleep(delay)
			}
		}

		return zero, lastErr
	}
}

// Compose creates a function that is the composition of the provided functions.
// The resulting function executes from right to left (last to first).
//
// Parameters:
//   - fns: The functions to compose
//
// Returns:
//   - func(T) T: A function that is the composition of the provided functions
//
// Example: addOne := func(x int) int { return x + 1 }; double := func(x int) int { return x * 2 }; addOneThenDouble := Compose(double, addOne); addOneThenDouble(3) -> 8
func Compose[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		for i := len(fns) - 1; i >= 0; i-- {
			result = fns[i](result)
		}
		return result
	}
}

// Pipe creates a function that is the composition of the provided functions, where each function consumes the return value of the previous.
// The resulting function executes from left to right (first to last).
//
// Parameters:
//   - fns: The functions to pipe
//
// Returns:
//   - func(T) T: A function that passes its input through the provided functions in sequence
//
// Example: addOne := func(x int) int { return x + 1 }; double := func(x int) int { return x * 2 }; doubleTheAddOne := Pipe(addOne, double); doubleTheAddOne(3) -> 8
func Pipe[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		result := x
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}

// Negate creates a function that negates the result of the predicate function.
//
// Parameters:
//   - predicate: The predicate function to negate
//
// Returns:
//   - func(T) bool: A function that returns the logical NOT of the result of predicate
//
// Example: isEven := func(n int) bool { return n % 2 == 0 }; isOdd := Negate(isEven); isOdd(3) -> true
func Negate[T any](predicate func(T) bool) func(T) bool {
	return func(x T) bool {
		return !predicate(x)
	}
}

// Spread creates a function that invokes func with the array of arguments it receives.
// This is a simplified version that works with two arguments.
//
// Parameters:
//   - fn: The function to spread arguments over
//
// Returns:
//   - func([]T) R: A function that accepts an array and calls fn with elements from the array as individual arguments
//
// Example: add := func(a, b int) int { return a + b }; addArray := Spread(add); addArray([]int{1, 2}) -> 3
func Spread[T, R any](fn func(T, T) R) func([]T) R {
	return func(args []T) R {
		if len(args) < 2 {
			var zero R
			return zero
		}
		return fn(args[0], args[1])
	}
}

// TransformList generic function takes a list of records and applies a transformer function to each element,
// returning a slice of transformed elements.
//
// Parameters:
//   - records: The slice of elements to transform
//   - transformerFn: The function to apply to each element
//
// Returns:
//   - []R: A new slice containing the transformed elements
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	squares := TransformList(numbers, func(x int) int { return x * x })
//	 ===> squares = []int{1, 4, 9}
func TransformList[T any, R any](records []T, transformerFn func(T) R) []R {
	if len(records) == 0 {
		return []R{}
	}

	outData := make([]R, 0, len(records))
	for i := range records {
		outData = append(outData, transformerFn(records[i]))
	}

	return outData
}

// TransformMap transforms a map of one type to a map of another type using a transformer function.
//
// Parameters:
//   - m: The map to transform
//   - transformerFn: The function to apply to each value in the map
//
// Returns:
//   - map[K]R: A new map with the same keys but transformed values
//
// Example:
//
//	ages := map[string]int{"John": 30, "Jane": 25}
//	doubled := TransformMap(ages, func(age int) int { return age * 2 })
//	 ===> doubled = map[string]int{"John": 60, "Jane": 50}
func TransformMap[K comparable, V any, R any](m map[K]V, transformerFn func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = transformerFn(v)
	}
	return result
}

// TransformListWithError transforms a slice and collects any errors that occur during transformation.
//
// Parameters:
//   - records: The slice of elements to transform
//   - transformerFn: The function to apply to each element, which may return an error
//
// Returns:
//   - []R: A new slice containing the successfully transformed elements
//   - []error: A slice containing any errors that occurred during transformation
//
// Example:
//
//	numbers := []string{"1", "2", "abc", "3"}
//	parsed, errors := TransformListWithError(numbers, func(s string) (int, error) { return strconv.Atoi(s) })
//	 ===> parsed = []int{1, 2, 3}, errors contains the error from parsing "abc"
func TransformListWithError[T any, R any](records []T, transformerFn func(T) (R, error)) ([]R, []error) {
	if len(records) == 0 {
		return []R{}, nil
	}

	outData := make([]R, 0, len(records))
	var errors []error

	for i := range records {
		result, err := transformerFn(records[i])
		if err != nil {
			errors = append(errors, err)
			continue
		}
		outData = append(outData, result)
	}

	return outData, errors
}

// TransformConcurrent transforms a slice concurrently using a specified number of workers.
//
// Parameters:
//   - records: The slice of elements to transform
//   - transformerFn: The function to apply to each element
//   - numWorkers: The number of concurrent workers to use for transformation
//
// Returns:
//   - []R: A new slice containing the transformed elements
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	squares := TransformConcurrent(numbers, func(x int) int { return x * x }, 2)
//	 ===> squares = []int{1, 4, 9, 16, 25}
func TransformConcurrent[T any, R any](records []T, transformerFn func(T) R, numWorkers int) []R {
	if len(records) == 0 {
		return []R{}
	}

	// If only a few records or numWorkers is 1, use the sequential version
	if len(records) < numWorkers || numWorkers <= 1 {
		return TransformList(records, transformerFn)
	}

	// Create result slice with capacity matching input
	result := make([]R, len(records))

	// Create a wait group to synchronize workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Calculate batch size for each worker
	batchSize := len(records) / numWorkers
	if len(records)%numWorkers != 0 {
		batchSize++
	}

	// Launch workers
	for w := 0; w < numWorkers; w++ {
		// Calculate start and end indices for this worker
		start := w * batchSize
		end := start + batchSize
		if end > len(records) {
			end = len(records)
		}

		// Skip if this worker has no work
		if start >= len(records) {
			wg.Done()
			continue
		}

		// Launch worker goroutine
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				result[i] = transformerFn(records[i])
			}
		}(start, end)
	}

	// Wait for all workers to complete
	wg.Wait()
	return result
}

// TransformBatch transforms a slice in batches and returns the combined results.
//
// Parameters:
//   - records: The slice of elements to transform
//   - transformerFn: The function to apply to each batch of elements
//   - batchSize: The size of each batch
//
// Returns:
//   - []R: A new slice containing the combined results of all batch transformations
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := TransformBatch(numbers, func(batch []int) []int {
//		var result []int
//		for _, n := range batch {
//			result = append(result, n*2)
//		}
//		return result
//	}, 2)
//	 ===> doubled = []int{2, 4, 6, 8, 10}
func TransformBatch[T any, R any](records []T, transformerFn func([]T) []R, batchSize int) []R {
	if len(records) == 0 {
		return []R{}
	}

	if batchSize <= 0 {
		batchSize = 100 // Default batch size
	}

	var result []R

	// Process in batches
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		batch := records[i:end]
		batchResult := transformerFn(batch)
		result = append(result, batchResult...)
	}

	return result
}
