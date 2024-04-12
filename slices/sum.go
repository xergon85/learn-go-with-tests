package slices

// Sums takes a slice of numbers and returns the sum of them
func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}

	return sum
}

// SumAll takes an unknown amount of slices and returns a new slice with the sum of each slice
func SumAll(slicesToSum ...[]int) []int {
	var sums []int
	for _, v := range slicesToSum {
		sums = append(sums, Sum(v))
	}

	return sums
}

// SumAllTails takes an unknown amount of slices and returns a new slice with the sum of each tail
func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, v := range slices {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
