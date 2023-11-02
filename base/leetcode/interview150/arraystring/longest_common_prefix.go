package arraystring

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; ; i++ {
		c := ""
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return prefix
			}
			if c == "" {
				c = strs[j][i : i+1]
			} else if c != strs[j][i:i+1] {
				return prefix
			}
		}
		prefix += c
	}
}
