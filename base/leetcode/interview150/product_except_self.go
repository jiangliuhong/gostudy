package interview150

// 238. 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	// 最简单的方法，使用除法
	// 先计算总的乘集，再分别除
	// 只有特别判断一下0，如果0出现超过两次，那返回结果所有都为0
	answer := make([]int, len(nums))
	sum, zeroNum := 1, 0
	for _, n := range nums {
		if n != 0 {
			sum = sum * n
		} else {
			zeroNum++
		}
	}
	if zeroNum < 2 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				answer[i] = sum
			} else if zeroNum != 1 {
				answer[i] = sum / nums[i]
			}
		}
	}
	return answer
}
