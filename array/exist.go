package array

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
//输出：true
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
//输出：false
//m == board.length
//n = board[i].length
//1 <= m, n <= 6
//1 <= word.length <= 15
//board 和 word 仅由大小写英文字母组成
//
//
//进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

type pair2 struct {
	x int
	y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 不存在
		if board[i][j] != word[k] {
			return false
		}
		// 存在且为最后一个
		if k == len(word)-1 {
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }()
		for _, p := range directions {
			newX := i + p.x
			newY := j + p.y
			if newX < 0 || newX > h-1 || newY < 0 || newY > w-1 || vis[newX][newY] {
				continue
			}
			if check(newX, newY, k+1) {
				return true
			}
		}
		return false
	}

	for x, row := range board {
		for y := range row {
			if check(x, y, 0) {
				return true
			}
		}
	}
	return false
}
