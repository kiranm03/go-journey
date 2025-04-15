package main

import "testing"

func TestAdd(t *testing.T) {
	// Arrange
	l, r := 1, 2
	expected := 3
	// Act
	got := Add(l, r)
	// Assert
	if got != expected {
		t.Errorf("Failed to add %d and %d. Expected %d, got %d",
			l, r, expected, got)
	}
}
