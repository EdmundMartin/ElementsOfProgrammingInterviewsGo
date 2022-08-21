package ch14_binary_search_trees

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRebuildBSTFromPreOrder(t *testing.T) {

	sequence := []int{43, 23, 37, 29, 31, 41, 47, 53}

	result := RebuildBSTFromPreOrder(sequence)
	fmt.Println(result)

	var values []int
	preOrderTraverse(result, &values)

	if !reflect.DeepEqual(sequence, values) {
		t.Error("tree not re-built correctly")
	}
}

func preOrderTraverse(tree *BSTNode, values *[]int) {
	if tree == nil {
		return
	}
	*values = append(*values, tree.Data)
	preOrderTraverse(tree.Left, values)
	preOrderTraverse(tree.Right, values)
}
