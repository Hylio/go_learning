package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("hylio")
	want := "Hello, hylio"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}