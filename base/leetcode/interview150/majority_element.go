package interview150

import "sort"

// 169. 多数元素
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
func majorityElement(nums []int) int {
	sort.Ints(nums)
	// 数组排序后， n/2 位置的数一定是众数
	return nums[len(nums)/2]
}

func majorityElement2(nums []int) int {
	// 排序
	sort.Ints(nums)
	size := len(nums)
	resultSize := 1
	result := nums[0]
	for i := 1; i < size; i++ {
		num := nums[i]
		if num == result {
			resultSize++
		} else {
			resultSize = 1
			result = num
		}
		if resultSize > size/2 {
			return result
		}
	}

	return result
}
