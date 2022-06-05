package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node == nil {
		return root
	}
	if rplc == nil {
		node = nil
		return root
	}
	rplc.Parent = node.Parent
	if node.Parent.Left == node {
		rplc.Parent.Left = rplc
	} else if node.Parent.Right == node {
		rplc.Parent.Right = rplc
	}
	return root
}
