package arraystring

// 28. 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	subIndex := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i:i+1] == needle[subIndex:subIndex+1] {
			subIndex++
		} else {
			i = i - subIndex
			subIndex = 0
		}
		if subIndex >= len(needle) {
			return i - subIndex + 1
		}
	}
	return -1
}
