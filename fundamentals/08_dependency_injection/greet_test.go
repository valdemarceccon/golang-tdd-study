package dependency_injection_test

import (
	"bytes"
	dependency_injection "github.com/valdemarceccon/golang-tdd-study/fundamentals/08_dependency_injection"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	dependency_injection.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
