package ch9_binary_trees

type BalanceInfo struct {
	Balanced bool
	Height   int
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkBalance(tree *BinaryTreeNode) BalanceInfo {
	if tree == nil {
		return BalanceInfo{true, -1}
	}

	leftResult := checkBalance(tree.Left)
	if !leftResult.Balanced {
		return leftResult
	}

	rightResult := checkBalance(tree.Right)
	if !rightResult.Balanced {
		return rightResult
	}

	isBalanced := abs(leftResult.Height-rightResult.Height) <= 1
	height := max(leftResult.Height, rightResult.Height) + 1
	return BalanceInfo{
		Balanced: isBalanced,
		Height:   height,
	}
}

func IsBalancedBinaryTree(tree *BinaryTreeNode) bool {
	result := checkBalance(tree)
	return result.Balanced
}
