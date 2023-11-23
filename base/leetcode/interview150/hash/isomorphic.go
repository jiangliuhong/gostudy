package hash

// 205. 同构字符串
// 给定两个字符串 s 和 t ，判断它们是否是同构的。
func isIsomorphic(s string, t string) bool {
	dict := map[rune]rune{}
	use := map[rune]bool{}
	size := len(s)
	for i := 0; i < size; i++ {
		itemS := rune(s[i])
		itemT := rune(t[i])
		if v, ok := dict[itemS]; ok {
			if v != itemT {
				return false
			}
		} else {
			if use[itemT] {
				return false
			}
			dict[itemS] = itemT
			use[itemT] = true
		}
	}
	return true
}
