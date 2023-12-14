package sliding_window

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	ls, left, right, res := len(s), 0, 0, ""
	dictV, dictC := map[rune]int{}, map[rune]bool{}
	for _, item := range t {
		dictV[item]++
		dictC[item] = true
	}
	for left <= right && right <= ls {
		if isZero(dictV) {
			if res == "" {
				res = s[left:right]
			} else {
				resT := s[left:right]
				if len(resT) < len(res) {
					res = resT
				}
			}
			itemS := rune(s[left])
			if _, ok := dictC[itemS]; ok {
				dictV[itemS]++
			}
			left++
		} else if right == ls {
			break
		} else {
			itemS := rune(s[right])
			if _, ok := dictC[itemS]; ok {
				dictV[itemS]--
			}
			right++
		}
	}
	return res
}

func isZero(m map[rune]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
