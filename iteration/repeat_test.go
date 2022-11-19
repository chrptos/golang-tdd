package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// ベンチマークの取得に必要なだけ b.N回 繰り返す
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	sum := Repeat("a", 5)
	fmt.Println(sum)
	// Output: aaaaa
}

func TestReplaceAll(t *testing.T) {
	replaced := ReplaceAll("Hello, World!", "World!", "chrptos!")
	expected := "Hello, chrptos!"

	if replaced != expected {
		t.Errorf("expected %q but got %q", expected, replaced)
	}
}
