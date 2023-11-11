package double_pointer

// 11. 盛最多水的容器
func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1
	for left < right {
		h := height[left]
		w := right - left
		if height[right] < h {
			h = height[right]
			right--
		} else {
			left++
		}
		area := h * w
		if area > max {
			max = area
		}
	}
	return max
}
