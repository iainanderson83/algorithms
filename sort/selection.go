package sort

// SelectionSort is an implementation of the selection sort algorithm.
// O(1) for pre-sorted to O(n^2) for reverse sorted.
func SelectionSort(numbers []int) {
	for i := range numbers {
		currentMinIdx := i
		for j := currentMinIdx+1; j < len(numbers); j++ {
			if numbers[j] < numbers[currentMinIdx] {
				currentMinIdx = j
			}
		}

		if currentMinIdx != i {
			numbers[i], numbers[currentMinIdx] = numbers[currentMinIdx], numbers[i]
		}
	}
}