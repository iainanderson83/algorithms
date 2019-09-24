package graph

import (
	"github.com/iainanderson83/datastructures/stack"
)

// DFS implements a depth first search.
// Useful if you're interested in parent/child relationships.
func DFS(b *BinaryNode) []int {
	var (
		s        stack.Stack = &stack.ArrStack{}
		currNode *BinaryNode
		results  []int
	)

	s.Push(b)
	for s.Peek() != nil {
		currNode = s.Pop().(*BinaryNode)
		results = append(results, currNode.Value)

		if currNode.Right != nil {
			s.Push(currNode.Right)
		}
		if currNode.Left != nil {
			s.Push(currNode.Left)
		}
	}

	return results
}
