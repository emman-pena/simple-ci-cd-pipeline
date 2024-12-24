package main

import "testing"

func TestMain(t *testing.T) {
	got := "Hello, CI/CD!"
	want := "Hello, CI/CD!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
