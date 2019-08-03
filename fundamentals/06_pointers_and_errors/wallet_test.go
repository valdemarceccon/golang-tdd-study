package pointers_and_errors_test

import (
	pointersanderrors "github.com/valdemarceccon/golang-tdd-study/fundamentals/06_pointers_and_errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := pointersanderrors.Wallet{}

		wallet.Deposit(pointersanderrors.Bitcoin(10))

		want := pointersanderrors.Bitcoin(10)

		assertBalance(t, wallet, want)

	})
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := pointersanderrors.Wallet{}
		wallet.Deposit(pointersanderrors.Bitcoin(30))
		want := pointersanderrors.Bitcoin(20)

		err := wallet.Withdraw(pointersanderrors.Bitcoin(10))

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := pointersanderrors.Bitcoin(20)
		wallet := pointersanderrors.Wallet{}
		wallet.Deposit(startBalance)
		err := wallet.Withdraw(pointersanderrors.Bitcoin(100))

		assertBalance(t, wallet, pointersanderrors.Bitcoin(20))
		assertError(t, err, pointersanderrors.ErrInsufficientFunds)
	})

	t.Run("assert formatted stringer", func(t *testing.T) {
		wallet := pointersanderrors.Wallet{}
		wallet.Deposit(10)
		want := "10 BTC"
		got := wallet.Balance().String()

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}

	})

}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertBalance(t *testing.T, wallet pointersanderrors.Wallet, want pointersanderrors.Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
