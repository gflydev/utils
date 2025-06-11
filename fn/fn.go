// Package fn provides utility functions for function manipulation.
package fn

import (
	"sync"
	"time"
)

// After creates a function that invokes func once it's called n or more times.
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
// Example: Delay(func() { fmt.Println("called") }, 100) // prints "called" after 100ms
func Delay(fn func(), wait time.Duration) {
	time.AfterFunc(wait, fn)
}

// Memoize creates a function that memoizes the result of func.
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
// Example: greet := func(greeting, name string) string { return greeting + " " + name }; sayHello := Partial(greet, "Hello"); sayHello("John") -> "Hello John"
func Partial[T, R any](fn func(T, T) R, partial T) func(T) R {
	return func(arg T) R {
		return fn(partial, arg)
	}
}

// Rearg creates a function that invokes func with arguments arranged according to the specified indexes.
// This is a complex function that's hard to implement in Go due to type system limitations.
// For now, we'll just provide a simplified version that swaps the first two arguments.
func Rearg[T, R any](fn func(T, T) R) func(T, T) R {
	return func(a, b T) R {
		return fn(b, a)
	}
}

// Throttle creates a throttled function that only invokes func at most once per every wait milliseconds.
// Example: fn := Throttle(func() { fmt.Println("called") }, 100); fn(); fn(); fn() // prints "called" only once per 100ms
func Throttle(fn func(), wait time.Duration) func() {
	var (
		lastInvoke time.Time
		mu         sync.Mutex
	)

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if now.Sub(lastInvoke) >= wait {
			lastInvoke = now
			go fn()
		}
	}
}

// Wrap creates a function that provides value to the wrapper function as its first argument.
// Example: hello := func(name string) string { return "Hello " + name }; withExclamation := Wrap(hello, func(hello func(string) string, name string) string { return hello(name) + "!" }); withExclamation("John") -> "Hello John!"
func Wrap[T, R, S any](fn func(T) R, wrapper func(func(T) R, T) S) func(T) S {
	return func(arg T) S {
		return wrapper(fn, arg)
	}
}

// Retry creates a function that retries the given function until it succeeds or reaches the maximum number of retries.
// Example: fn := Retry(func() error { return errors.New("error") }, 3, 100); fn() // retries 3 times with 100ms delay
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
// Example: isEven := func(n int) bool { return n % 2 == 0 }; isOdd := Negate(isEven); isOdd(3) -> true
func Negate[T any](predicate func(T) bool) func(T) bool {
	return func(x T) bool {
		return !predicate(x)
	}
}

// Spread creates a function that invokes func with the array of arguments it receives.
// This is a complex function that's hard to implement in Go due to type system limitations.
// For now, we'll just provide a simplified version that works with two arguments.
func Spread[T, R any](fn func(T, T) R) func([]T) R {
	return func(args []T) R {
		if len(args) < 2 {
			var zero R
			return zero
		}
		return fn(args[0], args[1])
	}
}
