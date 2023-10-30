package arraystring

import "sort"

// 274. H 指数
// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
func hIndex(citations []int) int {
	sort.Ints(citations)
	size := len(citations)
	for i := 0; i < size; i++ {
		if size-i <= citations[i] {
			return size - i
		}
	}
	return 0
}

// 5
// 0,1,3,6,5
// 0,1,2,3,4
