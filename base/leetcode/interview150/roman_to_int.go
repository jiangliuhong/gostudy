package interview150

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
	res := 0
	for _, c := range s {
		res += dict[string(c)]
	}
	return res
}
