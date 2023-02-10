package main

func maxAncestorDiff(root *TreeNode) int {
	return findAllRoot(root, []int{})
}

func findAllRoot(tree *TreeNode, root []int) int {
	if tree == nil {
		return 0
	}

	if tree.Left == nil && tree.Right == nil {
		root = append(root, tree.Val)
		maxR, minR := root[0], root[0]
		for _, v := range root {
			maxR = max(v, maxR)
			minR = min(v, minR)
		}
		return maxR - minR
	}

	root = append(root, tree.Val)
	n := len(root)
	l := findAllRoot(tree.Left, root)
	r := findAllRoot(tree.Right, root[:n])

	return max(l, r)
}
