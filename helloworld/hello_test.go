package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		name := "Wahab"
		got := Hello(name, "")
		want := englishHelloPrefix + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world when empty string is provided", func(t *testing.T) {
		got := Hello("", "")
		want := englishHelloPrefix + "World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Wahab", "Spanish")
		want := "Hola, Wahab"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Wahab", "French")
		want := "Bonjour, Wahab"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
