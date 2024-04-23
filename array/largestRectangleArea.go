package array

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//输入： heights = [2,4]
//输出： 4
//1 <= heights.length <=105
//0 <= heights[i] <= 104

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {

			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Printf("i:%v,right[i]:%v,left[i]:%v,w:%v,h:%v\n", i, right[i], left[i], right[i]-left[i]-1, heights[i])
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

// 时间负载度为 ON 执行超时
func largestRectangleArea2(heights []int) int {
	maxArea := make([]int, len(heights))
	maxArea[0] = heights[0]
	for i := 1; i < len(heights); i++ {
		maxA := heights[i]
		minH := heights[i]
		for j := i - 1; j >= 0; j-- {
			minH = min(minH, heights[j])
			maxA = max(minH*(i-j+1), maxA)
		}
		maxArea[i] = max(maxArea[i-1], maxA)
	}
	return maxArea[len(heights)-1]
}
