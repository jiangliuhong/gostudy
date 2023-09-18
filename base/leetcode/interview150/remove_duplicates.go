package interview150

// 26. 删除有序数组中的重复项
// 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
// 返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
func removeDuplicates26(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	left, right := 0, 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

// 80. 删除有序数组中的重复项 II
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 快慢指针写法
func removeDuplicates(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return size
	}
	slow, fast := 2, 2
	for fast < size {
		if nums[slow-2] != nums[fast] {

		}
	}
	return slow
}

// 笨拙写法
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	saveIndex := 0
	tempNum := nums[0]
	tempNumSize := 0
	for _, num := range nums {
		if num > tempNum {
			nums[saveIndex] = num
			saveIndex++
			tempNumSize = 1
			tempNum = num
		} else if tempNumSize < 2 {
			nums[saveIndex] = num
			saveIndex++
			tempNumSize++
		}
	}
	return saveIndex
}
