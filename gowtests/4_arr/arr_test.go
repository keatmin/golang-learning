package arr

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("When using array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if expected != got {
			t.Errorf("Expected %d want %d, given %v", expected, got, numbers)

		}
	})

	t.Run("When using array", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if expected != got {
			t.Errorf("Expected %d want %d, given %v", expected, got, numbers)

		}
	})

}

// slices cannot be compared, use DeepEqual instead but it is not type safe
func TestSumAll(t *testing.T) {
	numbers := SumAll([]int{1, 2}, []int{3, 6})
	expected := []int{3, 9}

	if !reflect.DeepEqual(numbers, expected) {
		t.Errorf("Expected %v got %v", expected, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("with normal slice", func(t *testing.T) {
		numbers := SumAllTails([]int{1, 2}, []int{3, 6})
		expected := []int{2, 6}

		if !reflect.DeepEqual(numbers, expected) {
			t.Errorf("Expected %v got %v", expected, numbers)
		}
	})

	t.Run("with empty slice", func(t *testing.T) {
		numbers := SumAllTails([]int{}, []int{3, 3, 5})
		expected := []int{0, 8}

		if !reflect.DeepEqual(numbers, expected) {
			t.Errorf("Expected %v got %v", expected, numbers)
		}

	})
}
