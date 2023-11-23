package sliding_window

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if lt > ls {
		return ""
	}
	td := map[rune]int{}
	for i := 0; i < lt; i++ {
		td[rune(t[i])]++
	}
	temp := map[rune]int{}
	for i := 0; i < ls; i++ {
		temp[rune(s[i])]++
	}
	return ""
}
