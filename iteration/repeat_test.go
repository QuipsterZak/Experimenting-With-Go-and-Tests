package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// ExampleRepeat demonstrates how to use the Repeat function.
func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}

func TestToUpper(t *testing.T) {
	input := "zak"
	expected := "ZAK"

	result := strings.ToUpper(input)

	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}
