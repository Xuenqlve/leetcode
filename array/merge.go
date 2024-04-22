package array

import "sort"

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//示例 1：
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//提示：
//1 <= intervals.length <= 104
//intervals[i].length == 2
//0 <= starti <= endi <= 104

// 解题思路: 排序 + 双指针
//	对比 结果集中最后一位 endi 与 遍历到 starti
//	 starti 大 则 将遍历到的加入结果集中
//   starti 小 则 对比 结果集中最后一位 endi 与 遍历到的 endi 那个大，则留那个

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for _, e := range intervals[1:] {
		if result[len(result)-1][1] < e[0] {
			result = append(result, e)
		} else {
			result[len(result)-1][1] = max(result[len(result)-1][1], e[1])
		}
	}
	return result
}
