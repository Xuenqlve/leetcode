package array

import "fmt"

//给你一个满足下述两条属性的 m x n 整数矩阵：
//
//每行中的整数从左到右按非严格递增顺序排列。
//每行的第一个整数大于前一行的最后一个整数。
//给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//示例 1：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//示例 2：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//提示：
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 100
//-104 <= matrix[i][j], target <= 104

func searchMatrix(matrix [][]int, target int) bool {
	num := make([]int, 0, len(matrix)*len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			num = append(num, matrix[i][j])
		}
	}
	fmt.Println(num)
	start, end := 0, len(num)-1
	for start <= end {
		h := (start + end) / 2
		if num[h] == target {
			return true
		} else if num[h] < target {
			start = h + 1
		} else {
			end = h - 1
		}
	}
	return false
}
