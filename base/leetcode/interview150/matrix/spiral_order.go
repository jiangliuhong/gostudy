package matrix

// 54. 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素
func spiralOrder(matrix [][]int) []int {
	w, h := len(matrix[0]), len(matrix)
	var res []int
	x, y := 0, 0
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directIndex := 0
	visit := make([][]bool, h)
	for i := 0; i < h; i++ {
		visit[i] = make([]bool, w)
	}

	for len(res) < w*h {
		item := matrix[y][x]
		res = append(res, item)
		nx := x + direct[directIndex][0]
		ny := y + direct[directIndex][1]
		if nx >= w || nx < 0 || ny >= h || ny < 0 || visit[ny][nx] {
			directIndex = (directIndex + 1) % 4
		}
		x = x + direct[directIndex][0]
		y = y + direct[directIndex][1]
	}
	return res
}
