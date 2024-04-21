package main

import (
	"testing"
)

// TestAsterisk uses file where Resource field contains a single asterisk
func TestAsterisk(t *testing.T) {
	got := Verify("resources/asterisk.json")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// TestNoAsterisk uses file where Resource field contains words "lorem ipsum"
func TestNoAsterisk(t *testing.T) {
	got := Verify("resources/no_asterisk.json")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// TestInvalidPath uses filepath that does not exist
func TestInvalidPath(t *testing.T) {
	got := Verify("resources/test.json")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
