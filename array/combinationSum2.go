package array

import "sort"

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//candidates 中的每个数字在每个组合中只能使用 一次 。
//注意：解集不能包含重复的组合。
//示例 1:
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//示例 2:
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//提示:
//1 <= candidates.length <= 100
//1 <= candidates[i] <= 50
//1 <= target <= 30

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result_list := [][]int{}
	result := []int{}
	combination2(candidates, target, &result_list, result, 0, 0)
	return result_list
}

func combination2(candidates []int, target int, result_list *[][]int, result []int, start int, sum int) {
	for i := start; i < len(candidates); i++ {
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}
		tmp := sum + candidates[i]
		if target == tmp {
			result = append(result, candidates[i])
			*result_list = append(*result_list, result)
			return
		} else if target > tmp {
			result = append(result, candidates[i])
			combination2(candidates, target, result_list, result, i+1, tmp)
			result = result[: len(result)-1 : len(result)-1]
		}
	}
}
