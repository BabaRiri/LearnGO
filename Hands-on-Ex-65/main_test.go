package main

import "testing"

func TestMultiply(t *testing.T) {
	got := multiply(5, 10)
	want := 50.0
	if got!= want {
        t.Errorf("Got: %v, Want: %v", got, want)
    }
}