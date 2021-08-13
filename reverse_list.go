package algos

import (
	"encoding/json"
	"fmt"
)

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode

	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

		currentPre, _ := json.Marshal(pre)

		fmt.Printf("%s\n", string(currentPre))
	}
	return pre
}
