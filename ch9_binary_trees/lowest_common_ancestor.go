package ch9_binary_trees

type Status struct {
	NumTargetNodes int
	Ancestor       *BinaryTreeNode
}

func LowestCommonAncestorWithOutParentPointer(tree, first, second *BinaryTreeNode) *BinaryTreeNode {
	return lcaHelper(tree, first, second).Ancestor
}

func lcaHelper(tree, first, second *BinaryTreeNode) Status {
	if tree == nil {
		return Status{NumTargetNodes: 0, Ancestor: nil}
	}

	leftResult := lcaHelper(tree.Left, first, second)
	if leftResult.NumTargetNodes == 2 {
		return leftResult
	}

	rightResult := lcaHelper(tree.Right, first, second)
	if rightResult.NumTargetNodes == 2 {
		return rightResult
	}

	count := 0
	if first == tree {
		count++
	}
	if second == tree {
		count++
	}

	numTargetNodes := leftResult.NumTargetNodes + rightResult.NumTargetNodes + count
	var ancestor *BinaryTreeNode
	if numTargetNodes == 2 {
		ancestor = tree
	}
	return Status{
		NumTargetNodes: numTargetNodes,
		Ancestor:       ancestor,
	}

}

func getDepth(node *BinaryTreeNodeWithParent) int {
	depth := 0
	for node.Parent != nil {
		depth++
		node = node.Parent
	}
	return depth
}

func LowestCommonAncestorWithParentPointer(tree, first, second *BinaryTreeNodeWithParent) *BinaryTreeNodeWithParent {

	depthFirst, depthSecond := getDepth(first), getDepth(second)

	if depthSecond > depthFirst {
		first, second = second, first
	}

	depthDiff := abs(depthFirst - depthSecond)
	for depthDiff > 0 {
		first = first.Parent
		depthDiff -= 1
	}

	for first != second {
		first, second = first.Parent, second.Parent
	}
	
	return first
}