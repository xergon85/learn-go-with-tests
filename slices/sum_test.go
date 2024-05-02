package slices

import (
	"reflect"
	"testing"

	"github.com/xergon85/learn-go-with-tests/generics"
)

func TestSum(t *testing.T) {
	t.Run("an array of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d expected %d, given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("test with unknown amount of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{1, 2, 3})
		expected := []int{3, 9, 6}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	}

	t.Run("make sums of some slices' tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		generics.AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concat strings", func(t *testing.T) {
		concat := func(a, b string) string {
			return a + b
		}

		generics.AssertEqual(t, Reduce([]string{"a", "b", "c"}, concat, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		generics.AssertTrue(t, found)
		generics.AssertEqual(t, firstEvenNumber, 2)

	})
}
