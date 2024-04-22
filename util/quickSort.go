package util

func QuickSort(nums []int, start, end int) {
	if start > end {
		return
	}
	p := Partition(nums, start, end)
	QuickSort(nums, start, p-1)
	QuickSort(nums, p+1, end)
}

func Partition(nums []int, start, end int) int {
	i := start
	tmp := nums[end]
	for j := start; j < end; j++ {
		if nums[j] < tmp {
			if i != j {
				Swap(nums, i, j)
			}
			i++
		}

	}
	Swap(nums, end, i)
	return i
}

func Swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
