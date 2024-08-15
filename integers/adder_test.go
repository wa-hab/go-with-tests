package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Add two numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d got %d", expected, sum)
		}
	})
}
