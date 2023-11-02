package arraystring

// 42. 接雨水
//
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(height []int) int {
	if len(height) == 1 {
		return 0
	}
	size := len(height)
	leftMax, rightMax := make([]int, size), make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = maxInt(leftMax[i-1], height[i])
	}
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = maxInt(height[i], rightMax[i+1])
	}
	rain := 0
	for index, item := range height {
		rain += minInt(leftMax[index], rightMax[index]) - item
	}
	return rain
}

//func trap(height []int) int {
//	if len(height) == 1 {
//		return 0
//	}
//	leftMax, rightMax, rain, tempSum := 0, 0, 0, 0
//	for rightMax < len(height) {
//		if height[leftMax] == 0 {
//			leftMax = rightMax
//		} else if height[rightMax] >= height[leftMax] {
//			rain += height[leftMax]*(rightMax-leftMax-1) - tempSum
//			tempSum = 0
//			leftMax = rightMax
//		} else if rightMax != leftMax {
//			tempSum += height[rightMax]
//		}
//		rightMax++
//	}
//	return rain
//}
