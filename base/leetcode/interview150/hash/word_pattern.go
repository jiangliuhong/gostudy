package hash

import "strings"

// 290. 单词规律
func wordPattern(pattern string, s string) bool {
	sa := strings.Split(s, " ")
	lp, ls := len(pattern), len(sa)
	if lp != ls {
		return false
	}
	d := map[rune]string{}
	sd := map[string]rune{}
	for i := 0; i < ls; i++ {
		pi := i % lp
		p := rune(pattern[pi])
		item := sa[i]
		if d[p] != "" && d[p] != item || sd[item] > 0 && sd[item] != p {
			return false
		} else {
			d[p] = item
			sd[item] = p
		}
	}
	return true
}
