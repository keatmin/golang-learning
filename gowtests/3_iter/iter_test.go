package iter

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	exp := "aaaaa"

	if repeated != exp {
		t.Errorf("Expected %q got %q", exp, repeated)
	}
}

// b.N determines N number to run the test to benchmark
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
