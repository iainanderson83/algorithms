package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	numbers := []int{23, 4, 42, 15, 16, 8}
	results := []int{4, 8, 15, 16, 23, 42}
	BubbleSort(numbers)

	for i := range results {
		if numbers[i] != results[i] {
			t.Fatalf("expected %d, got %d", results[i], numbers[i])
		}
	}
}
  