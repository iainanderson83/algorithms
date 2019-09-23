package sort

// MergeSort is an implementation of the merge sort algorithm.
// O(n log n) time.
func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	middle := len(numbers) / 2
	left, right := numbers[:middle], numbers[middle:]

	return mergeSortMerge(MergeSort(left), MergeSort(right))
}

func mergeSortMerge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for {
		if len(left) == 0 && len(right) == 0 {
			break
		}

		switch {
		case len(left) > 0 && len(right) > 0:
			if left[0] < right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		case len(left) > 0:
			result = append(result, left[0])
			left = left[1:]
		default:
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
