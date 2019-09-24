package fibonacci

// RecursiveFib is the most painful time and space
// method of solving Fibonacci for n.
func RecursiveFib(n int) int {
	if n < 2 {
		return n
	}

	return RecursiveFib(n-2) + RecursiveFib(n-1)
}

// MemoizedFib is O(n) for time and space.
func MemoizedFib(n int) []int {
	seq := []int{0, 1}
	for i := 2; i <= n; i++ {
		seq = append(seq, seq[i-2]+seq[i-1])
	}

	return seq
}

// JustNFib is a greedy algorithm and we're just returning the nth value.
func JustNFib(n int) int {
	var (
		twoAgo = 0
		oneAgo = 1
		curr   = 0
	)
	for i := 2; i <= n; i++ {
		curr = twoAgo + oneAgo
		twoAgo, oneAgo = oneAgo, curr
	}

	return curr
}
