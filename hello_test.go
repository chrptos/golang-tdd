package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'こんにちは、chrptos！' when a Japanese string is supplied", func(t *testing.T) {
		got := Hello("chrptos!", "Japanese")
		want := "こんにちは、chrptos!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour、chrptos！' when a French string is supplied", func(t *testing.T) {
		got := Hello("chrptos!", "French")
		want := "Bonjour, chrptos!"
		assertCorrectMessage(t, got, want)
	})

}
