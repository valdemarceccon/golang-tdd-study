package poker_test

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := pokertesting.CreateTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
