package maps_test

import (
	maps "github.com/valdemarceccon/golang-tdd-study/fundamentals/07_maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := maps.ErrNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := maps.Dictionary{}
		_ = dictionary.Add("test", "this is just a test")

		want := "this is just a test"

		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("add a already added word", func(t *testing.T) {
		dictionary := maps.Dictionary{}
		_ = dictionary.Add("test", "this is just a test")
		err := dictionary.Add("test", "this is just another test")

		want := "this is just a test"

		assertError(t, maps.ErrWordAlreadyDefined, err)

		assertDefinition(t, dictionary, "test", want)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := maps.Dictionary{}
	_ = dictionary.Add("test", "this is just a test")

	t.Run("update an existing word", func(t *testing.T) {
		err := dictionary.Update("test", "new definition")

		assertNotError(t, err)

		assertDefinition(t, dictionary, "test", "new definition")
	})

	t.Run("update a not existing word", func(t *testing.T) {
		got := dictionary.Update("another", "new definition")

		assertError(t, got, maps.ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	dictionary := maps.Dictionary{}
	_ = dictionary.Add("test", "this is just a test")

	t.Run("delete existing word", func(t *testing.T) {
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		assertError(t, err, maps.ErrNotFound)
	})
}

func assertDefinition(t *testing.T, dictionary maps.Dictionary, word, want string) {
	t.Helper()

	got, err := dictionary.Search(word)

	assertNotError(t, err)

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error.")
	}

	if got != want {
		t.Errorf("got %q want %q", got.Error(), want.Error())
	}
}

func assertNotError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("expected to not get an error.")
	}
}
