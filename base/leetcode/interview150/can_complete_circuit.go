package interview150

import "math"

// 134. 加油站
//
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
func canCompleteCircuit(gas []int, cost []int) int {
	size, sum, min, minIndex := len(gas), 0, math.MaxInt, -1
	for i := 0; i < size; i++ {
		sum += gas[i] - cost[i]
		if sum < min && sum < 0 {
			min = sum
			minIndex = i
		}
	}
	if sum < 0 {
		return -1
	} else {
		return minIndex + 1
	}
}

func canCompleteCircuit_bak(gas []int, cost []int) int {
	nums := make([]int, len(gas))
	sp := -1
	for i := 0; i < len(gas); i++ {
		nums[i] = gas[i] - cost[i]
		if nums[i] > 0 && sp == -1 {
			sp = i
		}
	}
	if sp == -1 {
		sp = 0
	}
	for i := sp; i < len(nums); i++ {
		start, count := i, 0
		for {
			if count < 0 {
				break
			}
			count += nums[start]
			start++
			if start >= len(nums) {
				start = 0
			}
			if start == i {
				break
			}
		}
		if start == i && count >= 0 {
			return i
		}
	}
	return -1
}
