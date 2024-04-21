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

// TestIllegalFields uses file with added fields that are outside of provided description
// while astrisk is still present in Resources field
func TestIllegalFields(t *testing.T) {
	got := Verify("resources/illegal_fields.json")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// TestMultiplePolicies uses file with three policies, each with two having asterisk in Resource field,
// and one having two Resources as a list
func TestMultiplePolicies(t *testing.T) {
	got := Verify("resources/multiple_policies.json")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// TestMultiplePolicies uses file with three policies, none having asterisk in Resource field
func TestMultiplePoliciesNoAsterisk(t *testing.T) {
	got := Verify("resources/multiple_policies_no_asterisk.json")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// TestMultiplePolicies uses file with three policies, last one having asterisk in Resource field
func TestMultiplePoliciesLastAsterisk(t *testing.T) {
	got := Verify("resources/multiple_policies_last_asterisk.json")
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
