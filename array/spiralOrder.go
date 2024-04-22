package array

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//示例 1：
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//示例 2：
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//提示：
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 10
//-100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) []int {
	x, y, yl, xl := 0, 0, len(matrix), len(matrix[0])
	dir := 0
	length := xl * yl
	visited := make([][]bool, yl)
	for i := 0; i < yl; i++ {
		visited[i] = make([]bool, xl)
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		//fmt.Printf("y:%d x:%d matrix[y][x]:%d dir:%d \n", y, x, matrix[y][x], dir)
		result[i] = matrix[y][x]
		visited[y][x] = true
		switch dir {
		case 0: // 向左
			if x+1 < xl && !visited[y][x+1] {
				x++
			} else {
				dir = 2
				y++
			}
		case 1: // 向左
			if x > 0 && !visited[y][x-1] {
				x--
			} else {
				dir = 3
				y--
			}
		case 2: // 向下
			if y+1 < yl && !visited[y+1][x] {
				y++
			} else {
				dir = 1
				x--
			}
		case 3: // 向上
			if y > 0 && !visited[y-1][x] {
				y--
			} else {
				dir = 0
				x++
			}
		}
	}
	return result
}
