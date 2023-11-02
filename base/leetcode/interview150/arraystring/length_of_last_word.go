package arraystring

// 58. 最后一个单词的长度
func lengthOfLastWord(s string) int {
	size := 0
	startWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			if startWord {
				break
			}
		} else {
			startWord = true
			size++
		}
	}
	return size
}
