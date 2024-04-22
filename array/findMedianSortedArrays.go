package array

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//算法的时间复杂度应该为 O(log (m+n)) 。

//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	i, j, n1, n2 := 0, 0, len(nums1), len(nums2)
	indexMap := map[int]int{}
	sum := n1 + n2
	if (sum)%2 == 0 {
		indexMap[(sum)/2-1] = 0
		indexMap[(sum)/2] = 0
	} else {
		indexMap[sum/2] = 0
	}
	index := 0
	value := 0
	for index <= sum/2 {
		if i == n1 {
			value = nums2[j]
			j++
		} else if j == n2 {
			value = nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			value = nums2[j]
			j++
		} else {
			value = nums1[i]
			i++
		}

		if _, ok := indexMap[index]; ok {
			indexMap[index] = value
		}
		index++
	}

	if (sum)%2 == 0 {
		return (float64(indexMap[sum/2-1]) + float64(indexMap[sum/2])) / 2
	} else {
		return float64(indexMap[sum/2])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	i, j, index := len(nums1), len(nums2), len(nums1)+len(nums2)
	if (index)%2 == 1 {
		return float64(MedianSortedArrays(nums1, i, nums2, j, index/2+1))
	} else {
		return float64(MedianSortedArrays(nums1, i, nums2, j, index/2)+MedianSortedArrays(nums1, i, nums2, j, index/2+1)) / 2.0
	}
}

// 找到第k小的数
func MedianSortedArrays(nums1 []int, nums1_len int, nums2 []int, nums2_len int, k int) int {
	nums1_start, nums2_start := 0, 0
	for {

		// nums1 已经遍历完了，直接锁定 nums2 中 nums2_start + k - 1 为第k 小的数
		if nums1_start >= nums1_len {
			return nums2[nums2_start+k-1]
		}
		if nums2_start >= nums2_len {
			return nums1[nums1_start+k-1]
		}
		if k == 1 {
			return min(nums1[nums1_start], nums2[nums2_start])
		}
		// 找 nums1 的 k/2-1 的数 和 nums2 的 k/2-1 的数对比
		// 如果 nums1 的  k/2-1 的数 效应 nums2 的 k/2-1 的数, k
		nums1_index := min(nums1_start+k/2-1, nums1_len-1)
		nums2_index := min(nums2_start+k/2-1, nums2_len-1)
		if nums1[nums1_index] < nums2[nums2_index] {
			k = k - (nums1_index - nums1_start + 1)
			nums1_start = nums1_index + 1
		} else {
			k = k - (nums2_index - nums2_start + 1)
			nums2_start = nums2_index + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
