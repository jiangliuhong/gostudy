package arraystring

// 12. 整数转罗马数字
func intToRoman(num int) string {
	dictKey := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	dictVal := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roma := ""
	for i := 0; i < len(dictKey); {
		key := dictKey[i]
		if num >= key {
			num -= key
			roma += dictVal[i]
		} else {
			i++
		}
	}
	return roma
}
