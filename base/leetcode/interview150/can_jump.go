package interview150

// 55. 跳跃游戏
//
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

func canJump(nums []int) bool {
	// 动态规划
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < i {
			return false
		}
		dp[i] = maxInt(dp[i-1], i+nums[i])
	}
	return dp[len(nums)-2] > len(nums)-1
}
func canJump_greed(nums []int) bool {
	// 贪心
	max := 0
	for i := 0; i < len(nums); i++ {
		if max >= i {
			target := i + nums[i]
			if target > max {
				max = target
			}
			if max >= len(nums)-1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// 递归实现超时
func canJumpDeep(nums []int, target int) bool {
	if target >= len(nums)-1 {
		return true
	}

	for i := nums[target]; i > 0; i-- {
		deep := canJumpDeep(nums, target+i)
		if deep {
			return true
		}
	}
	return false
}
