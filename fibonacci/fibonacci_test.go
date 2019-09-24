package fibonacci

import "testing"

func TestRecursiveFib(t *testing.T) {
	results := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < 10; i++ {
		out := RecursiveFib(i)
		if out != results[i] {
			t.Fatalf("expected %d, got %d", results[i], out)
		}
	}
}

func TestMemoizedFib(t *testing.T) {
	results := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	out := MemoizedFib(10)
	for i := range results {
		if out[i] != results[i] {
			t.Fatalf("expected %d, got %d", results[i], out[i])
		}
	}
}

func TestJustNFib(t *testing.T) {
	result := 55
	out := JustNFib(10)
	if out != result {
		t.Fatalf("expected %d, got %d", result, out)
	}
}
