package shuffle

import (
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
)

// FisherYates is an implementation of the Knuth Fisher Yates shuffle algorithm.
func FisherYates(numbers []int) {
	for i := len(numbers) - 1; i > 0; i-- {
		idx := r.Intn(i + 1)
		numbers[i], numbers[idx] = numbers[idx], numbers[i]
	}
}
