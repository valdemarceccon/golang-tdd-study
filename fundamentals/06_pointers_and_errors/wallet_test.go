package pointers_and_errors_test

import (
	pointersanderrors "github.com/valdemarceccon/golang-tdd-study/fundamentals/06_pointers_and_errors"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := pointersanderrors.Wallet{}

	wallet.Deposit(pointersanderrors.Bitcoin(10))

	got := wallet.Balance()
	want := pointersanderrors.Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
