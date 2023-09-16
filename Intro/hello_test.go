package main

import (
	"testing"
)

func TestHello(t *testing.T) { // t *testing.T being a function signature, t being a *testing.T pointer, which report test failure.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zak")
		want := "Hello, Zak"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
