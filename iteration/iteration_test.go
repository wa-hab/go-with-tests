package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("Testing a simple iteration", func(t *testing.T) {
		repeated := Repeat("a")

		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
