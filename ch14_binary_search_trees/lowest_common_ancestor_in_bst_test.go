package ch14_binary_search_trees

import (
	"testing"
)

func TestFindLowestCommonAncestor(t *testing.T) {

	larger := NewBSTNode(14)
	smaller := NewBSTNode(11)

	root := NewBSTNode(10)
	root.Left = NewBSTNode(7)
	root.Right = NewBSTNode(12)
	root.Right.Right = larger
	root.Right.Left = smaller

	result := FindLowestCommonAncestor(root, larger, smaller)

	if result.Data != 12 {
		t.Error("expected 12")
	}
}
