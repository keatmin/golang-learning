package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("mate")
		want := "Hello mate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say world when empty string is provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // to notify test suite that this method is a helper. Failure will be reported in function call rather than test helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
