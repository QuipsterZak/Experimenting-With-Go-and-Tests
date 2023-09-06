package main

import (
	"testing"
)

func TestHello(t *testing.T) { // t *testing.T being a function signature, t being a *testing.T pointer, which report test failure.
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want) //t.Errorf report test failure by formatting the error message and marking the test as failed.
	}
}
