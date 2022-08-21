package ch14_binary_search_trees

import (
	"reflect"
	"testing"
)

func TestFindKLargestInBST(t *testing.T) {

	root := NewBSTNode(10)
	root.Right = NewBSTNode(12)
	root.Right.Right = NewBSTNode(14)
	root.Right.Right.Right = NewBSTNode(16)
	root.Right.Left = NewBSTNode(11)

	result := FindKLargestInBST(root, 3)

	if !reflect.DeepEqual(result, []int{16, 14, 12}) {
		t.Error("unexpected result")
	}
}
