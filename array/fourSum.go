package array

import "sort"

//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//0 <= a, b, c, d < n
//a、b、c 和 d 互不相同
//nums[a] + nums[b] + nums[c] + nums[d] == target
//你可以按 任意顺序 返回答案 。

//示例 1：
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//示例 2：
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//提示：
//
//1 <= nums.length <= 200
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//
//
// 解题思路 排序 + 双指针 使用两重循环分别枚举前两个数，使用双指针枚举剩下的两个数
//

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	getSum(nums, target, &result, 0, []int{}, 4)
	return result
}

func getSum(nums []int, target int, result *[][]int, start int, head []int, num int) {
	if num == 2 {
		get2Sum(nums, target, result, start, head)
	} else {
		for i := start; i < len(nums)-num+1; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			getSum(nums, target-nums[i], result, i+1, append(head, nums[i]), num-1)
		}
	}
}

func get2Sum(nums []int, target int, result *[][]int, start int, head []int) {
	i, j := start, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			*result = append(*result, append(head, nums[i], nums[j]))
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			i++
			j--
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

}
