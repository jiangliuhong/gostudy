package arraystring

import "math"

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

// 45. 跳跃游戏 II给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
func jump(nums []int) int {
	dp := make([][2]int, len(nums)) // 0 能到达的数，1 步数
	dp[0][0] = 0
	dp[0][1] = 0
	for i := 1; i < len(dp); i++ {
		dp[i][1] = math.MaxInt
	}
	for i := 0; i < len(nums); i++ {
		var target int
		if i == 0 {
			target = i + nums[i]
		} else {
			target = maxInt(dp[i-1][0], i+nums[i])
			dp[i][1] = minInt(dp[i][1], dp[i-1][1]+1)
		}
		if target >= len(nums) {
			target = len(nums) - 1
		}
		dp[i][0] = target

		dp[target][1] = minInt(dp[i][1]+1, dp[target][1])
	}
	return dp[len(nums)-1][1]
}
