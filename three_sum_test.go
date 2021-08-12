package algos

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	data := []int{-4, -3, -5, -2, 1, 4, 5, 2, 5, 6, 7, 2, 1, 10, -3, 0, -2}

	ret := ThreeSum(data)

	fmt.Println(ret)
}
