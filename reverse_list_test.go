package algos

import "testing"

func TestReverseList(t *testing.T) {
	l := ListNode{
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

	ReverseList(&l)
}
