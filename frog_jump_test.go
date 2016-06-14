package main

import "testing"

func TestFrogJump(t *testing.T) {
	got := Solution(10, 85, 30)
	expected := 3
	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}
}
