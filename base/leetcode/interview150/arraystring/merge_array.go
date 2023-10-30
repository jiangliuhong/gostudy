package arraystring

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for index, item := range nums2 {
		nums1[index+m] = item
	}
	sort.Ints(nums1)
}
