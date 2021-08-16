package algos

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 100,
		Left: &TreeNode{
			Val: 100,
			Left: &TreeNode{
				Val: 180,
			},
			Right: &TreeNode{
				Val: 190,
			},
		},
		Right: &TreeNode{
			Val: 200,
			Left: &TreeNode{
				Val: 210,
			},
			Right: &TreeNode{
				Val: 220,
			},
		},
	}

	fmt.Printf("先序遍历%#v\n", PreorderTraversal(root))
	fmt.Printf("中序遍历%#v\n", InorderTraversal(root))
	fmt.Printf("先序遍历递归%#v\n", PreorderTraversal1(root))
	fmt.Printf("后序遍历递归%#v\n", PostorderTraversal(root))
}
