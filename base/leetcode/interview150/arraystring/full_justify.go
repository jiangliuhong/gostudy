package arraystring

import (
	"strings"
)

// 68. 文本左右对齐
func fullJustify(words []string, maxWidth int) []string {
	str, strIndex := make([]string, 0), -1
	// 尽可能放单词
	for _, word := range words {
		if strIndex > -1 && len(str[strIndex])+len(word)+1 <= maxWidth {
			str[strIndex] += " " + word
		} else {
			str = append(str, "")
			strIndex++
			str[strIndex] = word
		}
	}
	// 填充空格
	res := make([]string, 0)
	for index, item := range str {
		split := strings.Split(item, " ")
		splitSize := len(split)
		blankSize := maxWidth - (len(item) - splitSize + 1)
		if index < strIndex && splitSize > 1 {
			subStr := ""
			for j := splitSize - 1; j >= 0; j-- {
				if j == 0 {
					subStr = split[j] + blank(blankSize) + subStr
				} else {
					average := blankSize / j
					subStr = blank(average) + split[j] + subStr
					blankSize = blankSize - average
				}

			}

			res = append(res, subStr)
		} else {
			// 最后一行特殊处理
			blankSize = maxWidth - len(item)
			b := blank(blankSize)
			s := item + b
			res = append(res, s)
		}

	}
	return res
}

// blank 返回长度为 n 的由空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}
