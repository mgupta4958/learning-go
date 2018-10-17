package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Chris")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
			got := Hello("World")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
	})
}