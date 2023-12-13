package hash

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	dict := map[int]int{}
	res := make([]int, 0)
	for index, item := range nums {
		if i, ok := dict[target-item]; ok {
			res = append(res, index, i)
			return res
		}
		dict[item] = index
	}
	return res
}
