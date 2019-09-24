package graph

import (
	"github.com/iainanderson83/datastructures/queue"
)

// BFS implements a breadth first search.
// Useful if you're interested in validating a binary search tree.
func BFS(b *BinaryNode) []int {
	var (
		q        queue.Queue = queue.NewListQueue()
		currNode *BinaryNode
		results  []int
	)

	q.Enqueue(b)

	for q.Len() > 0 {
		currNode = q.Dequeue().(*BinaryNode)
		results = append(results, currNode.Value)

		if currNode.Left != nil {
			q.Enqueue(currNode.Left)
		}
		if currNode.Right != nil {
			q.Enqueue(currNode.Right)
		}
	}

	return results
}
