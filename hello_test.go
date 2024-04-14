package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("quaqua")
	want := "Hello, quaqua"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
