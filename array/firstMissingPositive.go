package array

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//示例 1：
//输入：nums = [1,2,0]
//输出：3
//示例 2：
//输入：nums = [3,4,-1,1]
//输出：2
//示例 3：
//输入：nums = [7,8,9,11,12]
//输出：1
//提示：
//1 <= nums.length <= 5 * 105
//-231 <= nums[i] <= 231 - 1

// 解题思路 ：两次遍历 第一次 将数组内容按照规律 操作，第二次 依照规律找到符合数据
//
//
//数组中包含 x ∈ [1,N]，那么恢复后，数组的第 x−1 个元素为 x
//在恢复后，数组应当有 [1, 2, ..., N] 的形式，但其中有若干个位置上的数是错误的，每一个错误的位置就代表了一个缺失的正数
//[3, 4, -1, 1] 为例，恢复后的数组应当为 [1, -1, 3, 4]，我们就可以知道缺失的数为 2

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

//func swap(nums []int, i, j int) {
//	nums[i], nums[j] = nums[j], nums[i]
//}
