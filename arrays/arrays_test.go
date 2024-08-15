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

		got := SumAll(numbers1, numbers2)

		expected := []int{6, 10}

		checkDeepEqual(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Should return the sum of all elements except the first one", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkDeepEqual(t, got, expected)
	})

	t.Run("should handle empty slices safely", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{}, []int{0, 9})
		expected := []int{0, 0, 9}

		checkDeepEqual(t, got, expected)
	})
}

func checkDeepEqual(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func checkEqualityHelper(t testing.TB, got, expected int) {
	t.Helper()

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
