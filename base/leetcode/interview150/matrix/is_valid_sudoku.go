package matrix

// 36. 有效的数独数字
// 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次
func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var sub [3][3][9]int
	for index, row := range board {
		for index2, r := range row {
			if r == '.' {
				continue
			}
			index3 := r - '1'
			rows[index][index3]++
			columns[index2][index3]++
			sub[index/3][index2/3][index3]++
			if rows[index][index3] > 1 || columns[index2][index3] > 1 || sub[index/3][index2/3][index3] > 1 {
				return false
			}
		}
	}
	return true
}
