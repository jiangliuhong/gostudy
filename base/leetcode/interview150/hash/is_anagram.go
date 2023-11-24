package hash

// 242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	d := map[rune]int{}
	for _, item := range s {
		d[item]++
	}
	for _, item := range t {
		d[item]--
	}
	for _, v := range d {
		if v != 0 {
			return false
		}
	}
	return true
}
