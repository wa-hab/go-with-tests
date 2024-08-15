package arrays

import (
	"reflect"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("It should return the sum of the array", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 1}
		got := SumArray(numbers)

		expected := 5

		checkEqualityHelper(t, got, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("It should return the sum of individual arrays", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{1, 2, 3, 4}

		sums := SumAll(numbers1, numbers2)

		expected := []int{6, 10}

		if !reflect.DeepEqual(sums, expected) {
			t.Errorf("Expected %v got %v", expected, sums)
		}
	})
}

func checkEqualityHelper(t testing.TB, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
