package sort

import "testing"

func TestQuickSort(t *testing.T) {
	numbers := []int{23, 4, 42, 15, 16, 8}
	results := []int{4, 8, 15, 16, 23, 42}
	out := MergeSort(numbers)

	for i := range results {
		if out[i] != results[i] {
			t.Fatalf("expected %d, got %d", results[i], out[i])
		}
	}
}
