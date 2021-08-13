package algos

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	l5 := ListNode{
		Val: 5,
	}

	l3 := ListNode{
		Val:  3,
		Next: &l5,
	}

	l1 := ListNode{
		Val:  1,
		Next: &l3,
	}

	l5.Next = &l3

	res := HasCycle(&l1)

	fmt.Printf("res:%#v", res)
}
