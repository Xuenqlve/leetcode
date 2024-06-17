package array

//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
//示例 2：
//
//输入：matrix = [["0"]]
//输出：0
//示例 3：
//
//输入：matrix = [["1"]]
//输出：1
//rows == matrix.length
//cols == matrix[0].length
//1 <= row, cols <= 200
//matrix[i][j] 为 '0' 或 '1'

// 解题思路：从下往上，从右往左，先算出每个点的最大宽度，
// 取右下角节点向上遍历,
// 0 , 1 , 0 , 1
// 0 , 1 , 1 , 1  ^
// 1 , 1 , 1 , 1  |
//            <-
// matrix[2][3] 当前节点最大宽度为 4 宽度取4 高度为 1  当前最大面积 4
// matrix[1][3] 当前节点最大宽度为 3 宽度取3 高度为 2  当前最大面积 6
// matrix[0][3] 当前节点最大宽度为 1 宽度取1 高度为 3  当前最大面积 3

func maximalRectangle(matrix [][]byte) int {
	left := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		lefti := make([]int, len(matrix[i]))
		if matrix[i][0] == '1' {
			lefti[0] = 1
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				lefti[j] = lefti[j-1] + 1
			} else {
				lefti[j] = 0
			}
		}
		left[i] = lefti
	}

	maxA := 0
	for i := len(left) - 1; i >= 0; i-- {
		for j := len(left[i]) - 1; j >= 0; j-- {
			if left[i][j] == 0 {
				continue
			}
			w := left[i][j]
			h := 1
			if maxA < w*h {
				maxA = w * h
			}
			for k := i - 1; k >= 0; k-- {
				if left[k][j] == 0 {
					break
				}
				if left[k][j] < w {
					w = left[k][j]
				}
				h++
				if maxA < w*h {
					maxA = w * h
				}
			}
		}
	}
	return maxA
}
