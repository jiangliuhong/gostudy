package interview150

import "math"

// 121. 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
func maxProfit(prices []int) int {
	min, max := math.MaxInt, 0
	for _, item := range prices {
		sub := item - min
		if max < sub {
			max = sub
		}
		if min > item {
			min = item
		}
	}
	return max
}

// 122. 买卖股票的最佳时机 II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
// 返回 你能获得的 最大 利润 。
func maxProfit2(prices []int) int {
	// 动态规划三步
	// 1.状态定义 i 表示第几天，j 表示状态 买入、卖出 ，dp[i][j] 表示当前状态的现金
	// 2.转移方程 当天买入的现金 = 上一天卖出的现金 - 今天的股价 ； 当天卖出的现金 = 今天的股价 + 上一天买入的现金
	//           当天买入的现金 与上一天买入的现金比较取最大值； 当天卖出的现金 与 上一天卖出的现金比较取最大值
	// 3.确定初始值 dp[0][0] = 0 ,dp[0][1] = - prices[0]
	// 4.确定输出值
	if len(prices) == 1 {
		return 0
	}
	// 0 卖出  ， 1 买入
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxInt(dp[i-1][1]+prices[i], dp[i-1][0])
		dp[i][1] = maxInt(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func maxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
