package sort

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// HeapSort is an implementation of the heap sort algorithm.
// O(n log n).
func HeapSort(numbers []int) []int {
	results := make([]int, 0, len(numbers))

	h := IntHeap(numbers)
	ptr := &h
	heap.Init(ptr)
	for h.Len() > 0 {
		results = append(results, heap.Pop(ptr).(int))
	}

	return results
}
