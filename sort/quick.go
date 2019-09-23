package sort

// QuickSort is an implementation of the quick sort algorithm.
// O(n^2) worst case, best case O(n log n).
func QuickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	var left, right []int
	pivot := len(numbers) - 1
	pivotValue := numbers[pivot]
	numbers = append(numbers[:pivot], numbers[pivot+1:]...)

	for _, i := range numbers {
		if i < pivotValue {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}

	return QuickSort(append(left, append([]int{pivotValue}, right...)...))
}
