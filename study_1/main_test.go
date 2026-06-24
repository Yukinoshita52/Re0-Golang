package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("MuxueAvid")
	want := "Hello, MuxueAvid"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
