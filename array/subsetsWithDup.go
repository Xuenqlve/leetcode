package array

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的
//子集
//（幂集）。
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//示例 1：
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//示例 2：
//输入：nums = [0]
//输出：[[],[0]]
//
//提示：
//1 <= nums.length <= 10
//-10 <= nums[i] <= 10

// 先排序排查 乱序 过滤重复元素干扰
// 使用递归 ，需要注意向结果中存储时 因为存储内容为列表，需避免指针覆盖
// *result = append(*result, append([]int{}, subset...))  用这个append([]int{},xxxx) 算是存储当前快照

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	subset := []int{}
	sort.Ints(nums)
	getSubset(nums, 0, subset, &result)
	return result
}

func getSubset(nums []int, start int, subset []int, result *[][]int) {
	if start == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		*result = append(*result, append([]int{}, subset...))
		getSubset(nums, i+1, subset, result)
		subset = subset[:len(subset)-1]
	}
}
