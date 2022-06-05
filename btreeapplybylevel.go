package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode(nil)
	queue = append(queue, root)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current != nil {
			f(current.Data)
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
