package ch14_binary_search_trees

type Interval struct {
	Left  int
	Right int
}

func RangeLookupInBST(tree *BSTNode, interval Interval) []int {
	var values []int
	rangeLookupHelper(tree, interval, &values)
	return values
}

func rangeLookupHelper(tree *BSTNode, interval Interval, values *[]int) {

	if tree == nil {
		return
	}

	if interval.Left <= tree.Data && tree.Data <= interval.Right {
		rangeLookupHelper(tree.Left, interval, values)
		*values = append(*values, tree.Data)
		rangeLookupHelper(tree.Right, interval, values)
	} else if interval.Left > tree.Data {
		rangeLookupHelper(tree.Right, interval, values)
	} else {
		rangeLookupHelper(tree.Left, interval, values)
	}
}
