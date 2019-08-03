package hello_test

import (
	"github.com/valdemarceccon/golang-tdd-study/fundamentals/hello"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello.Hello("Valdemar", "")
		want := "Hello, Valdemar"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when a empty string is supplied", func(t *testing.T) {
		got := hello.Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello.Hello("Elodie", hello.Spanish)
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := hello.Hello("Elodie", hello.French)
		want := "Bonjour, Elodie"

		assertCorrectMessage(t, got, want)
	})
}
