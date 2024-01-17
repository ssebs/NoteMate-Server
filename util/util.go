package util

import "testing"

// GotWantTest takes *testing.T
func GotWantTest[T comparable](got, want T, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
