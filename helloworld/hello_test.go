package main

import (
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run(
		"saying hello to people in english", 
		func(t *testing.T) {
			got := Hello("en", "Claire")
			want := "Hello, Claire"
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"saying hello to people in spanish", 
		func(t *testing.T) {
			got := Hello("es", "Claire")
			want := "Hola, Claire"
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"saying hello to people with no language", 
		func(t *testing.T) {
			got := Hello("", "Claire")
			want := "Hello, Claire"
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"saying hello to people with no name", 
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
		})
}

// go mod init hello before running go test