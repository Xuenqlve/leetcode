package array

//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
//示例 1:
//输入: nums = [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//示例 2:
//输入: nums = [2,3,0,1,4]
//输出: 2
//提示:
//1 <= nums.length <= 104
//0 <= nums[i] <= 1000
//题目保证可以到达 nums[n-1]
//

// 解题思路 「贪心」地进行正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数
// 维护当前能够到达的最大下标位置，记为边界,从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
// 在遍历数组时，我们不访问最后一个元素,
//			因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了
//			如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
