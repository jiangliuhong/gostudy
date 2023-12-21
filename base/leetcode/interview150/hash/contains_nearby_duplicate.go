package hash

// 219. 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}
	for index, item := range nums {
		dict[index] = item
	}
	for index, item := range nums {
		for i := index + 1; i <= k+index && i < len(nums); i++ {
			if nums[i] == item {
				return true
			}
		}
	}
	return false
}
