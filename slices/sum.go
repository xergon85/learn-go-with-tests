package slices

// Sums takes a slice of numbers and returns the sum of them
func Sum(numbers []int) int {
	add := func(a, b int) int {
		return a + b
	}

	return Reduce(numbers, add, 0)
}

// SumAll takes an unknown amount of slices and returns a new slice with the sum of each slice
func SumAll(slicesToSum ...[]int) []int {

	sumAll := func(acc, x []int) []int {
		return append(acc, Sum(x))
	}

	return Reduce(slicesToSum, sumAll, []int{})
}

// SumAllTails takes an unknown amount of slices and returns a new slice with the sum of each tail
func SumAllTails(slices ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(slices, sumTail, []int{})
}

func Reduce[A, B any](collection []A, accumilator func(B, A) B, initialValue B) B {
	var result = initialValue

	for _, value := range collection {
		result = accumilator(result, value)
	}

	return result
}

func Find[T any](collection []T, predicate func(a T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var zero T
	return zero, false
}
