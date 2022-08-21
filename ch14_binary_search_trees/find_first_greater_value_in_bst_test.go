package ch14_binary_search_trees

import (
	"testing"
)

func TestFindFirstValueGreaterThanK(t *testing.T) {

	root := NewBSTNode(10)
	root.Left = NewBSTNode(7)
	root.Right = NewBSTNode(12)
	root.Right.Left = NewBSTNode(11)

	result := FindFirstValueGreaterThanK(root, 10)

	if result.Data != 11 {
		t.Error("Wrong result")
	}
}
