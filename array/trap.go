package array

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 示例 2：
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 提示：
//
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

// 解题思路 动态规划  规律是 这个点长度和它对应合理的最高点 的差就是这个点最大储水量
// 从左到右 再从右到左 两次循环 依次找到 这个点 对应左、右的最高点是什么，取高点中的最小值 减去这个点的高度 及为这个点的最大储水量

func trap1(height []int) int {
	n := len(height)
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], height[i])
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = Max(height[i], leftMax[i-1])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += Min(rightMax[i], leftMax[i]) - height[i]
	}
	return sum
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 解题思路2 单调栈
func trap2(height []int) int {
	stack := []int{}
	sum := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := Min(height[left], h) - height[top]
			sum += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return sum
}

// []int{4, 2, 0, 5, 2, 3}
// 解题思路 双指针
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}
