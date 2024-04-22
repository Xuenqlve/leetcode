package array

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//示例 1：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//输入：nums = [1]
//输出：[[1]]
//提示：
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同

func permute(nums []int) [][]int {
	result_list := [][]int{}
	tmp(nums, &result_list, []int{})
	return result_list
}

func tmp(nums []int, result_list *[][]int, result []int) {
	for i := 0; i < len(nums); i++ {
		tmp_nums := make([]int, 0, 0)
		tmp_nums = append(append(tmp_nums, nums[:i]...), nums[i+1:]...)
		result = append(result, nums[i])
		if len(tmp_nums) == 0 {
			*result_list = append(*result_list, append([]int{}, result...))
		} else {
			tmp(tmp_nums, result_list, result)
		}
		result = result[:len(result)-1]
	}
	return
}
