package util

import "testing"

func TestGcd(t *testing.T) {
	t.Run("gcd", func(t *testing.T) {
		a, b := 10, 6
		i := gcd(a, b)
		t.Logf("gcd:%v", i)
	})

	t.Run("lcm", func(t *testing.T) {
		a, b := 6, 10
		i := lcm(a, b)
		t.Logf("lcm:%v", i)
	})
}

func TestBinarySearch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nums := []int{1, 3, 5, 7, 10, 11, 16, 20, 23, 30, 34, 60}
		target := 3
		index := binarySearch(nums, target)
		t.Logf("index:%v", index)
	})

	t.Run("", func(t *testing.T) {
		nums := []int{0, 1, 3, 4, 5, 7, 8, 10, 23, 45, 67, 90, 122, 134, 567, 891}
		mid1 := (len(nums) - 1 + 6) / 2
		mid2 := 6 + (len(nums)-1-6)/2
		t.Logf("mid1:%v,mid2:%v", mid1, mid2)
		nums = []int{0, 1, 3, 4, 5, 7, 8, 10, 23, 45, 67, 90, 122, 134, 567}
		mid1 = (len(nums) - 1 + 6) / 2
		mid2 = 6 + (len(nums)-1-6)/2
		t.Logf("mid1:%v,mid2:%v", mid1, mid2)
	})
}

func TestQuickSort(t *testing.T) {
	t.Run("QuickSort", func(t *testing.T) {
		nums := []int{12, 32, 4, 5, 623, 1, 456, 74, 13, 7, 8, 10, 2}
		QuickSort(nums, 0, len(nums)-1)
		t.Logf("nums:%v", nums)
	})

}
