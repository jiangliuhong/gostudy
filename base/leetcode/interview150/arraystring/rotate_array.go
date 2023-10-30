package arraystring

// 189. 轮转数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	nums2 := make([]int, len(nums))
	for index, item := range nums {
		nums2[(index+k)%len(nums)] = item
	}
	copy(nums, nums2)
}
