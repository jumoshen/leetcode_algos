package algos

import "sort"

func ThreeSum(nums []int) [][]int {
	var ret [][]int

	l := len(nums)

	if l < 3 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, l-1; j < k; {
			if k < l-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			sum := nums[i] + nums[j] + nums[k]

			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}
	return ret
}
