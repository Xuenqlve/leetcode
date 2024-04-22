package array

import "sort"

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//示例 1：
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//示例 2：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//提示：
//
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10
//
// 要设定一个规则，保证在填第 idx 个数的时候重复数字只会被填入一次即可。
// 对原数组排序，保证相同的数字都相邻，然后每次填入的数一定是这个数所在重复数集合中「从左往右第一个未被填过的数字」，即如下的判断条件：

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result_list := [][]int{}
	unique(nums, &result_list, []int{})
	return result_list
}

func unique(nums []int, result_list *[][]int, result []int) {
	vas := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if vas[nums[i]] {
			continue
		}
		tmp_nums := make([]int, 0, 0)
		tmp_nums = append(append(tmp_nums, nums[:i]...), nums[i+1:]...)
		vas[nums[i]] = true
		result = append(result, nums[i])
		if len(tmp_nums) == 0 {
			*result_list = append(*result_list, append([]int{}, result...))
		} else {
			unique(tmp_nums, result_list, result)
		}
		result = result[:len(result)-1]
	}
	return
}
