package double_pointer

// 125. 验证回文串
func isPalindrome(s string) bool {
	// a - A = 32
	// a - z : 97 - 122
	// A - Z : 65 - 90
	// 0 - 9 ：48 - 57
	left, right := 0, len(s)-1
	for left <= right {
		lr := s[left]
		rr := s[right]
		// 大写 转 小写
		if lr >= 65 && lr <= 90 {
			lr += 32
		}
		if rr >= 65 && rr <= 90 {
			rr += 32
		}

		// 进行判断
		if !((lr >= 48 && lr <= 57) || (lr >= 97 && lr <= 122)) {
			left++
			continue
		}
		if !((rr >= 48 && rr <= 57) || (rr >= 97 && rr <= 122)) {
			right--
			continue
		}

		if lr != rr {
			return false
		}
		left++
		right--
	}
	return true
}
