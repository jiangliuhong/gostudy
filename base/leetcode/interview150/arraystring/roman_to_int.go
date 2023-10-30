package arraystring

// 13. 罗马数字转整数
func romanToInt(s string) int {
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res, size := 0, len(s)
	for index := 0; index < size; index++ {
		step := dict[s[index:index+1]]
		if index < size-1 {
			val2 := dict[s[index+1:index+2]]
			if val2 > step && ((val2-step)%4 == 0 || (val2-step)%9 == 0) {
				step = val2 - step
				index++
			}
		}
		res += step
	}

	return res
}
