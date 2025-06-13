package fn

import (
	"sync"
)

// TransformList generic function takes a list of records and applies a transformer function to each element,
// returning a slice of transformed elements.
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

// TransformListWithError transforms a slice and collects any errors that occur during transformation
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
