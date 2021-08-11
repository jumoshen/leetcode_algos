package algos

import (
	"encoding/json"
	"fmt"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	current := ret
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry

		current.Next = &ListNode{Val: sum % 10}

		current = current.Next
		carry = sum / 10
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return ret.Next
}

func RunAddTwoNumbers() {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	newList := addTwoNumbers(&l1, &l2)

	jsonList, _ := json.Marshal(newList)
	fmt.Printf("\nadd two numbers:%s", string(jsonList))
}
