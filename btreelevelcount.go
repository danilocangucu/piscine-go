package piscine

func BTreeLevelCount(root *TreeNode) int {
	x := 1
	y := 1
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		x += BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		y += BTreeLevelCount(root.Right)
	}
	if x >= y {
		return x
	} else {
		return y
	}
}
