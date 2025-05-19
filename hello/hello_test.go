package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello using names", func(t *testing.T) {

		got := Hello("David")
		want := "Hello, David"
		assertCorrectMessage(t, got, want)
	})
	t.Run("default to 'Hello, World' when called with an empty string", func(t *testing.T) {
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
