package double_pointer

import (
	"sort"
)

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	size := len(nums)
	for a := 0; a < size; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		c := size - 1
		for b := a + 1; b < c; b++ {
			isappend := false
			for ; c > b; c-- {
				t := nums[a] + nums[b] + nums[c]
				if t == 0 {
					if !isappend {
						res = append(res, []int{nums[a], nums[b], nums[c]})
						isappend = true
					}
				} else if t < 0 {
					break
				}
			}

		}
	}

	return res
}
