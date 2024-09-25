package array

//给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//示例 1:
//输入: rowIndex = 3
//输出: [1,3,3,1]
//示例 2:
//输入: rowIndex = 0
//输出: [1]
//示例 3:
//输入: rowIndex = 1
//输出: [1,1]
//提示:
//0 <= rowIndex <= 33
//进阶：
//
//你可以优化你的算法到 O(rowIndex) 空间复杂度吗？

// 1 * (6 - 1 + 1 ) / 1     6
// 6 * (6 - 2 + 1 ) / 2     15
// 15 * (6 - 3 + 1 ) / 3    20
// 20 * (6 - 4 + 1 ) / 4    15
// 15 * (6 - 5 + 1 ) / 5    6
// 6 * (6 - 6 + 1 ) /6      1

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		result[i] = result[i-1] * (rowIndex - i + 1) / i
	}
	return result
}
