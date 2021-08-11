package algos

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	subMap := make(map[byte]int)
	for i, j := 0, 0; j < n; j++ {
		if v, ok := subMap[s[j]]; ok {
			if i < v {
				i = v
			}
		}

		if ans < (j - i + 1) {
			ans = j - i + 1
		}
		subMap[s[j]] = j + 1
	}
	return ans
}

func RunSubStringLength()  {
	fmt.Println("\nsub string length:", lengthOfLongestSubstring("abcdacdfar"))
}
