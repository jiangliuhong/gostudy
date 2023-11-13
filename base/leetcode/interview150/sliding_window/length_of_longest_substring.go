package sliding_window

// 3.无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	dict := map[rune]int{}
	start, size := 0, 0
	for i, r := range s {
		if v, ok := dict[r]; !ok || v < start {
			var temp int
			if i == start {
				temp = 1
			} else {
				temp = i - start + 1
			}
			if temp > size {
				size = temp
			}
		} else {
			start = dict[r] + 1
		}
		dict[r] = i
	}
	return size
}
