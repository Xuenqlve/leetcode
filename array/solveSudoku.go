package array

//编写一个程序，通过填充空格来解决数独问题。
//数独的解法需 遵循如下规则：
//数字 1-9 在每一行只能出现一次。
//数字 1-9 在每一列只能出现一次。
//数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
//数独部分空格内已填入了数字，空白格用 '.' 表示。
//示例 1：
//输入：board = [
//	["5","3",".",".","7",".",".",".","."],
//	["6",".",".","1","9","5",".",".","."],
//	[".","9","8",".",".",".",".","6","."],
//	["8",".",".",".","6",".",".",".","3"],
//	["4",".",".","8",".","3",".",".","1"],
//	["7",".",".",".","2",".",".",".","6"],
//	[".","6",".",".",".",".","2","8","."],
//	[".",".",".","4","1","9",".",".","5"],
//	[".",".",".",".","8",".",".","7","9"]]
//输出：[
//	["5","3","4","6","7","8","9","1","2"],
//	["6","7","2","1","9","5","3","4","8"],
//	["1","9","8","3","4","2","5","6","7"],
//	["8","5","9","7","6","1","4","2","3"],
//	["4","2","6","8","5","3","7","9","1"],
//	["7","1","3","9","2","4","8","5","6"],
//	["9","6","1","5","3","7","2","8","4"],
//	["2","8","7","4","1","9","6","3","5"],
//	["3","4","5","2","8","6","1","7","9"]
//	]
//解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
//提示：
//board.length == 9
//board[i].length == 9
//board[i][j] 是一位数字或者 '.'
//题目数据 保证 输入数独仅有一个解

// 解题思路 ：遍历 + 递归 + 回溯

func solveSudoku(board [][]byte) {
	var rows, columns [9][9]bool
	var subboxes [3][3][9]bool
	var spaces [][2]int

	for i, row := range board {
		for j, col := range row {
			if col == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				index := col - '1'
				rows[i][index] = true
				columns[j][index] = true
				subboxes[i/3][j/3][index] = true
			}
		}
	}
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for index := byte(0); index < 9; index++ {
			if !rows[i][index] && !columns[j][index] && !subboxes[i/3][j/3][index] {
				rows[i][index] = true
				columns[j][index] = true
				subboxes[i/3][j/3][index] = true
				board[i][j] = index + '1'
				if dfs(pos + 1) {
					return true
				}
				rows[i][index] = false
				columns[j][index] = false
				subboxes[i/3][j/3][index] = false
			}
		}

		return false
	}
	dfs(0)

}
