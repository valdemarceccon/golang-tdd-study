package arrays_and_slices_test

import (
	"github.com/valdemarceccon/golang-tdd-study/fundamentals/arrays_and_slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays_and_slices.Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7}

		got := arrays_and_slices.Sum(numbers)
		want := 22

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
