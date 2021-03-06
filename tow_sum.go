package algos

import (
	"fmt"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]

func twoSum1(nums []int, target int) []int {
	tmp := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if key, ok := tmp[diff]; ok == true {
			return []int{key, i}
		}
		tmp[num] = i
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func RunTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res1 := twoSum1(nums, target)
	res := twoSum(nums, target)

	fmt.Printf("twoSum:result:%#v, twoSum1:result:%#v", res, res1)
}
