package array

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//示例 1：
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//示例 2：
//输入：n = 1
//输出：[["Q"]]
//提示：
//1 <= n <= 9

// 解题思路 回溯
// 同一行只找一个
// 使用三个集合 记录每一列以及 两个方向的每条斜线上是否有皇后
// 从左上到右下方向 ( \ ) 同一条斜线满足行下标与列下标之差相等 比如(0,0) (3,3)
// 从右上到左下方向 ( / ) 同一条斜线满足行下标与列下标之和相等 比如(3,0) (1,2)

var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	col, right, left := map[int]bool{}, map[int]bool{}, map[int]bool{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	backtrack(queens, n, 0, col, right, left)
	return solutions
}

func backtrack(queens []int, n, row int, col, right, left map[int]bool) {
	if n == row {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}

	for i := 0; i < n; i++ {
		if col[i] {
			continue
		}
		leftIndex := row - i
		if left[leftIndex] {
			continue
		}
		rightIndex := row + i
		if right[rightIndex] {
			continue
		}
		queens[row] = i
		col[i] = true
		right[rightIndex] = true
		left[leftIndex] = true
		backtrack(queens, n, row+1, col, right, left)
		delete(right, rightIndex)
		delete(left, leftIndex)
		delete(col, i)
		queens[row] = -1
	}
	return
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
