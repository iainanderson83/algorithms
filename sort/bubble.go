package sort

// BubbleSort is an implementation of the bubble sort algorithm.
// O(n^2) time + O(n) space.
func BubbleSort(numbers []int) {
	if len(numbers) <= 1 {
		return
	}

	andAgain := true
	for andAgain {
		andAgain = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i+1] < numbers[i] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				andAgain = true
			}
		}
	}
}
