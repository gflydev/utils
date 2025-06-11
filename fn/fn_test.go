package fn

import (
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	counter := 0
	f := After(3, func() int {
		counter++
		return counter
	})

	// First two calls should return 0
	if result := f(); result != 0 {
		t.Errorf("After(3, func)() = %d, expected 0 on first call", result)
	}
	if result := f(); result != 0 {
		t.Errorf("After(3, func)() = %d, expected 0 on second call", result)
	}

	// Third call should return 1
	if result := f(); result != 1 {
		t.Errorf("After(3, func)() = %d, expected 1 on third call", result)
	}

	// Fourth call should return 2
	if result := f(); result != 2 {
		t.Errorf("After(3, func)() = %d, expected 2 on fourth call", result)
	}
}

func TestBefore(t *testing.T) {
	counter := 0
	f := Before(3, func() int {
		counter++
		return counter
	})

	// First two calls should return incrementing values
	if result := f(); result != 1 {
		t.Errorf("Before(3, func)() = %d, expected 1 on first call", result)
	}
	if result := f(); result != 2 {
		t.Errorf("Before(3, func)() = %d, expected 2 on second call", result)
	}

	// Third and subsequent calls should return the last result (2)
	if result := f(); result != 2 {
		t.Errorf("Before(3, func)() = %d, expected 2 on third call", result)
	}
	if result := f(); result != 2 {
		t.Errorf("Before(3, func)() = %d, expected 2 on fourth call", result)
	}
}

func TestCurry(t *testing.T) {
	add := func(a, b int) int { return a + b }
	curriedAdd := Curry(add, 2)

	add3 := curriedAdd(3)
	if result := add3(4); result != 7 {
		t.Errorf("Curry(add, 2)(3)(4) = %d, expected 7", result)
	}
}

func TestDebounce(t *testing.T) {
	counter := 0
	f := func() { counter++ }
	debounced := Debounce(f, 50*time.Millisecond)

	// Call multiple times in quick succession
	debounced()
	debounced()
	debounced()

	// Wait for the debounce period
	time.Sleep(100 * time.Millisecond)

	// Counter should have been incremented only once
	if counter != 1 {
		t.Errorf("Debounce() called the function %d times, expected 1", counter)
	}
}

func TestDelay(t *testing.T) {
	counter := 0
	f := func() { counter++ }

	start := time.Now()
	Delay(f, 50*time.Millisecond)
	elapsed := time.Since(start)

	if counter != 1 {
		t.Errorf("Delay() did not call the function, counter = %d", counter)
	}

	if elapsed < 50*time.Millisecond {
		t.Errorf("Delay() did not wait long enough, elapsed = %v", elapsed)
	}
}

func TestMemoize(t *testing.T) {
	counter := 0
	square := func(n int) int {
		counter++
		return n * n
	}

	memoizedSquare := Memoize(square)

	// First call should compute the result
	counter = 0
	if result := memoizedSquare(5); result != 25 {
		t.Errorf("Memoize(square)(5) = %d, expected 25", result)
	}
	if counter != 1 {
		t.Errorf("Memoize() called the function %d times on first call, expected 1", counter)
	}

	// Second call should use the cached result
	counter = 0
	if result := memoizedSquare(5); result != 25 {
		t.Errorf("Memoize(square)(5) = %d, expected 25 on second call", result)
	}
	if counter != 0 {
		t.Errorf("Memoize() called the function %d times on second call, expected 0", counter)
	}

	// Different argument should compute a new result
	counter = 0
	if result := memoizedSquare(6); result != 36 {
		t.Errorf("Memoize(square)(6) = %d, expected 36", result)
	}
	if counter != 1 {
		t.Errorf("Memoize() called the function %d times for a new argument, expected 1", counter)
	}
}

func TestOnce(t *testing.T) {
	counter := 0
	f := Once(func() int {
		counter++
		return counter
	})

	// First call should execute the function
	if result := f(); result != 1 {
		t.Errorf("Once(func)() = %d, expected 1 on first call", result)
	}

	// Subsequent calls should return the same result
	if result := f(); result != 1 {
		t.Errorf("Once(func)() = %d, expected 1 on second call", result)
	}

	if counter != 1 {
		t.Errorf("Once() called the function %d times, expected 1", counter)
	}
}

func TestPartial(t *testing.T) {
	add := func(a, b int) int { return a + b }
	add5 := Partial(add, 5)

	if result := add5(3); result != 8 {
		t.Errorf("Partial(add, 5)(3) = %d, expected 8", result)
	}
}

func TestRearg(t *testing.T) {
	subtract := func(a, b int) int { return a - b }
	rearged := Rearg(subtract)

	if result := rearged(5, 3); result != 5-3 {
		t.Errorf("Rearg(subtract)(5, 3) = %d, expected 2", result)
	}
}

func TestThrottle(t *testing.T) {
	counter := 0
	f := func() { counter++ }
	throttled := Throttle(f, 50*time.Millisecond)

	// Call multiple times in quick succession
	throttled()
	throttled()
	throttled()

	// Counter should have been incremented only once
	if counter != 1 {
		t.Errorf("Throttle() called the function %d times immediately, expected 1", counter)
	}

	// Wait for the throttle period
	time.Sleep(100 * time.Millisecond)

	// Call again
	throttled()

	// Counter should now be 2
	if counter != 2 {
		t.Errorf("Throttle() called the function %d times after waiting, expected 2", counter)
	}
}

func TestWrap(t *testing.T) {
	greet := func(name string) string { return "Hello, " + name }
	wrapped := Wrap(greet, func(greetFunc func(string) string, name string) string {
		return greetFunc(name) + "!"
	})

	if result := wrapped("World"); result != "Hello, World!" {
		t.Errorf("Wrap(greet, func)(\"World\") = %q, expected \"Hello, World!\"", result)
	}
}

func TestRetry(t *testing.T) {
	counter := 0
	f := func() (string, error) {
		counter++
		if counter < 3 {
			return "", &testError{"temporary error"}
		}
		return "success", nil
	}

	retried := Retry(f, 5, 10*time.Millisecond)

	result, err := retried()
	if err != nil {
		t.Errorf("Retry() returned error: %v", err)
	}
	if result != "success" {
		t.Errorf("Retry() = %q, expected \"success\"", result)
	}
	if counter != 3 {
		t.Errorf("Retry() called the function %d times, expected 3", counter)
	}
}

// Custom error type for testing Retry
type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	composed := Compose(double, addOne)

	// Compose applies functions from right to left: double(addOne(5)) = double(6) = 12
	if result := composed(5); result != 12 {
		t.Errorf("Compose(double, addOne)(5) = %d, expected 12", result)
	}
}

func TestPipe(t *testing.T) {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	piped := Pipe(double, addOne)

	// Pipe applies functions from left to right: addOne(double(5)) = addOne(10) = 11
	if result := piped(5); result != 11 {
		t.Errorf("Pipe(double, addOne)(5) = %d, expected 11", result)
	}
}

func TestNegate(t *testing.T) {
	isEven := func(x int) bool { return x%2 == 0 }
	isOdd := Negate(isEven)

	if result := isOdd(2); result != false {
		t.Errorf("Negate(isEven)(2) = %v, expected false", result)
	}

	if result := isOdd(3); result != true {
		t.Errorf("Negate(isEven)(3) = %v, expected true", result)
	}
}

func TestSpread(t *testing.T) {
	add := func(a, b int) int { return a + b }
	spreaded := Spread(add)

	if result := spreaded([]int{2, 3}); result != 5 {
		t.Errorf("Spread(add)([]int{2, 3}) = %d, expected 5", result)
	}
}
