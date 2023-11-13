package sliding_window

import "math"

// 209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	left, right, tempSum, res := 0, 0, nums[0], math.MaxInt
	if tempSum == target {
		return 1
	}
	for left <= right {
		if tempSum >= target {
			resTemp := right - left + 1
			if resTemp < res {
				res = resTemp
			}
			tempSum = tempSum - nums[left]
			left++
		} else {
			right++
			if right >= len(nums) {
				break
			}
			tempSum = tempSum + nums[right]
		}
	}
	if res == math.MaxInt {
		return 0
	} else {
		return res
	}
}
