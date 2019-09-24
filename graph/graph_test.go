package graph

import "testing"

func TestDFS(t *testing.T) {
	root := &BinaryNode{
		Value: 0,
		Left: &BinaryNode{
			Value: 1,
			Left: &BinaryNode{
				Value: 3,
			},
			Right: &BinaryNode{
				Value: 4,
			},
		},
		Right: &BinaryNode{
			Value: 2,
			Left: &BinaryNode{
				Value: 5,
			},
			Right: &BinaryNode{
				Value: 6,
			},
		},
	}

	out := DFS(root)
	results := []int{0, 1, 3, 4, 2, 5, 6}
	for i := range out {
		if out[i] != results[i] {
			t.Fatalf("expected %d, got %d", results[i], out[i])
		}
	}
}

func TestBFS(t *testing.T) {
	root := &BinaryNode{
		Value: 0,
		Left: &BinaryNode{
			Value: 1,
			Left: &BinaryNode{
				Value: 3,
			},
			Right: &BinaryNode{
				Value: 4,
			},
		},
		Right: &BinaryNode{
			Value: 2,
			Left: &BinaryNode{
				Value: 5,
			},
			Right: &BinaryNode{
				Value: 6,
			},
		},
	}

	out := BFS(root)
	results := []int{0, 1, 2, 3, 4, 5, 6}
	for i := range out {
		if out[i] != results[i] {
			t.Fatalf("expected %d, got %d", results[i], out[i])
		}
	}
}
