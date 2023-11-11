package double_pointer

// 167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	i1, i2 := 0, len(numbers)-1
	for {
		sum := numbers[i1] + numbers[i2]
		if sum > target {
			i2--
		} else if sum < target {
			i1++
		} else {
			break
		}
	}
	return []int{i1 + 1, i2 + 1}
}
func twoSum_slow(numbers []int, target int) []int {
	i1, i2 := 0, 1
	for {
		if numbers[i1]+numbers[i2] == target {
			break
		}
		if i2 == len(numbers)-1 {
			i1++
			i2 = i1 + 1
		} else {
			i2++
		}
	}
	return []int{i1 + 1, i2 + 1}
}
