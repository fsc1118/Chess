package main

import (
	"testing"
)

func TestAbsValid(t *testing.T) {
	if abs(-1) != 1 {
		t.Errorf("Abs check was incorrect, got: %d, want: %d.", abs(-1), 1)
	}
	if abs(1) != 1 {
		t.Errorf("Abs check was incorrect, got: %d, want: %d.", abs(1), 1)
	}
	if abs(0) != 0 {
		t.Errorf("Abs check was incorrect, got: %d, want: %d.", abs(0), 0)
	}
}
