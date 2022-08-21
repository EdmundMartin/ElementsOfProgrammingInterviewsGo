package ch14_binary_search_trees

func FindLowestCommonAncestor(tree, first, second *BSTNode) *BSTNode {

	// Swap first and second if first is larger than second
	if first.Data > second.Data {
		first, second = second, first
	}

	for tree.Data < first.Data || tree.Data > second.Data {
		for tree.Data < first.Data {
			tree = tree.Right // LCA must be in tree's right child.
		}
		for tree.Data > second.Data {
			tree = tree.Left // LCA must be in tree's left child
		}
	}
	return tree
}
