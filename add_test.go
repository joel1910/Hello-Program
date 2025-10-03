package main

import "testing"

func TestAdd(t *testing.T) {
	// Test case
	x := 2
	y := 3
	want := 5

	got := Add(x, y)
	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", x, y, got, want)
	}
}
