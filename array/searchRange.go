package array

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//示例 1：
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//提示：
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109

func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			indexS, indexE, start, end := mid, mid, mid, mid
			for start >= i && nums[start] == target {
				indexS = start
				start--

			}
			for end <= j && nums[end] == target {
				indexE = end
				end++
			}
			return []int{indexS, indexE}
		} else if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		}
	}
	return []int{-1, -1}
}
