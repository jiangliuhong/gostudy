package arraystring

// 135. 分发糖果
//
// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
// 你需要按照以下要求，给这些孩子分发糖果：
//
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
func candy(ratings []int) int {
	candyArr := make([]int, len(ratings))
	candyArr[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyArr[i] = candyArr[i-1] + 1
		} else {
			candyArr[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candyArr[i] <= candyArr[i+1] {
			candyArr[i] = candyArr[i+1] + 1
		}
	}
	size := 0
	for _, n := range candyArr {
		size += n
	}
	return size
}
