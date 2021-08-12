package algos

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDeleteDuplicates2(t *testing.T) {
	l := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	newList := DeleteDuplicates2(&l)
	jsonList, _ := json.Marshal(newList)
	fmt.Println(string(jsonList))
}
