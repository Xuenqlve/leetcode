package array

import (
	"math"
	"sort"
)

//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//返回这三个数的和。
//
//假定每组输入只存在恰好一个解。
//示例 1：
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//示例 2：
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//提示：
//
//3 <= nums.length <= 1000
//-1000 <= nums[i] <= 1000
//-104 <= target <= 104

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minnum := 9999
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-minnum)) {
				minnum = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return minnum
}
