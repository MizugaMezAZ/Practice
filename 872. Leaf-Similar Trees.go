package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	r1 := []int{}
	r2 := []int{}
	findLeaf(root1, &r1)
	findLeaf(root2, &r2)

	if len(r1) != len(r2) {
		return false
	}

	for i := 0; i < len(r1); i++ {
		if r1[0] != r2[0] {
			return false
		}
	}

	return true
}

func findLeaf(tree *TreeNode, leaf *[]int) {
	if tree == nil {
		return
	}

	if tree.Left == nil && tree.Right == nil {
		*leaf = append(*leaf, tree.Val)
		return
	}

	findLeaf(tree.Left, leaf)
	findLeaf(tree.Right, leaf)
}
