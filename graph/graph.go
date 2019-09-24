package graph

// BinaryNode models a simple graph structure.
type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

// NewBinaryNode instantiates a new binary node.
func NewBinaryNode(i int) *BinaryNode {
	return &BinaryNode{Value: i}
}

// IsLeaf returns whether the node is a leaf.
func (b *BinaryNode) IsLeaf() bool {
	return b.Left == nil && b.Right == nil
}
