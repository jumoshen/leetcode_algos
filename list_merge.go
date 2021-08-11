package algos

import (
	"encoding/json"
	"fmt"
)

// ListNode 有序链表合并
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

func MergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = new(ListNode)
	current := ret

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}
	return ret.Next
}

func RunMergeList() {
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

	newList := MergeList(&l1, &l2)

	jsonList, _ := json.Marshal(newList)
	fmt.Printf(string(jsonList))

	for newList != nil {
		fmt.Printf("\nVal:%d", newList.Val)
		newList = newList.Next
	}
}
