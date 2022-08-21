package ch14_binary_search_trees

import "testing"

func TestIsValidBST(t *testing.T) {

	root := NewBSTNode(10)
	root.Left = NewBSTNode(7)
	root.Left.Left = NewBSTNode(5)
	root.Right = NewBSTNode(12)

	valid := IsValidBST(root)

	if valid == false {
		t.Error("Is a valid BST - wrong result")
	}

	root = NewBSTNode(10)
	root.Right = NewBSTNode(12)
	root.Right.Right = NewBSTNode(11)
	root.Left = NewBSTNode(7)
	root.Left.Left = NewBSTNode(5)

	valid = IsValidBST(root)

	if valid == true {
		t.Error("Invalid BST - wrong result")
	}
}

func TestIsValidBSTIterative(t *testing.T) {
	root := NewBSTNode(10)
	root.Left = NewBSTNode(7)
	root.Left.Left = NewBSTNode(5)
	root.Right = NewBSTNode(12)

	valid := IsValidBSTIterative(root)

	if valid == false {
		t.Error("Is a valid BST - wrong result")
	}

	root = NewBSTNode(10)
	root.Right = NewBSTNode(12)
	root.Right.Right = NewBSTNode(11)
	root.Left = NewBSTNode(7)
	root.Left.Left = NewBSTNode(5)

	valid = IsValidBSTIterative(root)

	if valid == true {
		t.Error("Invalid BST - wrong result")
	}
}
