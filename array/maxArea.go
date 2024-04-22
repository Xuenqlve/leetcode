package array

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//示例 2：
//
//输入：height = [1,1]
//输出：1
//
//
//提示：
//
//n == height.length
//2 <= n <= 105
//0 <= height[i] <= 104

// 从两头向内收，每次移动宽带-1
// 若向内 移动短板，水槽面积可能变大，若移动长板水槽面积一定变小
// 所以双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新最大面积，直到两指针相遇

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max_area, min_height, area := 0, 0, 0
	for i < j {
		if height[i] > height[j] {
			min_height = height[j]
			area = min_height * (j - i)
			for min_height >= height[j] && i < j {
				j--
			}
		} else {
			min_height = height[i]
			area = min_height * (j - i)
			for min_height >= height[i] && i < j {
				i++
			}
		}
		if area > max_area {
			max_area = area
		}
	}
	return max_area
}
