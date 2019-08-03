package integers_test

import (
	"fmt"
	"github.com/valdemarceccon/golang-tdd-study/fundamentals/02_integers"
	"testing"
)

func TestAdder(t *testing.T) {

	assertCorrectSum := func(t *testing.T, expected, got int) {
		t.Helper()

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	}

	t.Run("when adding 2 and 2 then result is 4", func(t *testing.T) {
		sum := integers.Add(2, 2)
		expected := 4

		assertCorrectSum(t, expected, sum)
	})

	t.Run("when adding 1 and 5 then result is 6", func(t *testing.T) {
		sum := integers.Add(1, 5)
		expected := 6

		assertCorrectSum(t, expected, sum)
	})

}

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
