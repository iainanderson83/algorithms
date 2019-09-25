package sort

import (
	"math/rand"

	"github.com/iainanderson83/datastructures/stack"
)

// QuickSort is an implementation of the quick sort algorithm.
// O(n^2) worst case, best case O(n log n).
func QuickSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	var (
		left       []int
		right      []int
		pivot      = len(numbers) - 1
		pivotValue = numbers[pivot]
	)

	numbers = append(numbers[:pivot], numbers[pivot+1:]...)
	for _, i := range numbers {
		if i < pivotValue {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}

	return append(QuickSort(left), append([]int{pivotValue}, QuickSort(right)...)...)
}

func median(a, b, c int) int {
	return max(min(a, b), min(max(a, b), c))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FastQuickSort has some optimizations for better
// worst case:
// - insertion sort on small slices
// - no recursion
// - better pivot selection
// Code is taken from:
// https://stackoverflow.com/a/19045659
// https://yourbasic.org/golang/quicksort-optimizations/
func FastQuickSort(numbers []int) {
	if len(numbers) < 20 {
		insertionSort(numbers)
		return
	}

	var (
		s    stack.Stack = &stack.ArrStack{}
		high             = len(numbers) - 1
		low  int
	)
	s.Push(low)
	s.Push(high)

	for s.Peek() != nil {
		high = s.Pop().(int)
		low = s.Pop().(int)

		p := partition(numbers, low, high)

		if p-1 > low {
			s.Push(low)
			s.Push(p - 1)
		}
		if p+1 < high {
			s.Push(p + 1)
			s.Push(high)
		}
	}
}

func pivot(numbers []int) int {
	l := len(numbers)
	return median(rand.Intn(l), rand.Intn(l), rand.Intn(l))
}

func insertionSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = key
	}
}

func partition(numbers []int, low int, high int) int {
	p := pivot(numbers)
	i := low - 1
	for j := low; j <= high-1; i++ {
		if numbers[j] <= p {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}
