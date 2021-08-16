package algos

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal1(root *TreeNode) []int {
	var values []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return values
}

func PreorderTraversal(root *TreeNode) []int {
	var values []int
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			values = append(values, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return values
}

func InorderTraversal(root *TreeNode) []int {
	var res []int

	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func PostorderTraversal(root *TreeNode) []int {
	var values []int
	var stack []*TreeNode
	var pre *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == pre {
			values = append(values, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return values
}
