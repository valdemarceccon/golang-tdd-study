package hello_test

import (
	"github.com/valdemarceccon/golang-tdd-study/fundamentals/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Hello("Valdemar")
	want := "Hello, Valdemar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
