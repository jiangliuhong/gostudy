package double_pointer

// 392. 判断子序列
func isSubsequence(s string, t string) bool {
	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if s[sp] == t[tp] {
			sp++
			tp++
		} else {
			tp++
		}
	}
	return sp == len(s)
}
