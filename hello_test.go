package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Wahab")
	want := "Hello, Wahab"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
