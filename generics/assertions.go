package generics

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, value bool) {
	t.Helper()
	if !value {
		t.Errorf("got %v, want true", value)
	}
}

func AssertFalse(t *testing.T, value bool) {
	t.Helper()
	if value {
		t.Errorf("got %v, want false", value)
	}
}
