package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	got := Repeat("b")
	fmt.Println(got)
	// Output: bbbbb
}
