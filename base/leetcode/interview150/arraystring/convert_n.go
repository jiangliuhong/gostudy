package arraystring

// 6. N 字形变换
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	//res := ""
	rc := 2*numRows - 2
	ans := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j = j + rc {
			//res += string(s[j+i])
			ans = append(ans, s[j+i])
			if i > 0 && i < numRows-1 && j+rc-i < len(s) {
				//res += string(s[j+rc-i]) // 加入当前周期第二个数，例如 A 后面的P
				ans = append(ans, s[j+rc-i])

			}
		}
	}
	return string(ans)
}
