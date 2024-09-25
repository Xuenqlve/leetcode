package array

import (
	"strconv"
	"testing"
)

func TestBuildTree(t *testing.T) {
	t.Run("minimumTotal", func(t *testing.T) {
		// [[2],[3,4],[6,5,7],[4,1,8,3]]
		triangle := [][]int{
			{2},
			{3, 4},
			{6, 5, 7},
			{4, 1, 8, 3},
		}
		minV := minimumTotal(triangle)
		t.Logf("minV:%v", minV)
	})

	t.Run("getRow", func(t *testing.T) {
		num := 6
		list := getRow(num)
		t.Logf("list:%v", list)
	})

	t.Run("generate", func(t *testing.T) {
		num := 7
		list := generate(num)
		t.Logf("list:%v", list)
	})

	t.Run("sortedArrayToBST", func(t *testing.T) {
		nums := []int{-10, -3, 0, 5, 9}
		nodes := sortedArrayToBST(nums)
		nodes.RangePrint()
	})
	t.Run("allCellsDistOrder1", func(t *testing.T) {
		//输入：rows = 1, cols = 2, rCenter = 0, cCenter = 0
		//输出：[[0,0],[0,1]]
		rows := 1
		cols := 2
		rCenter := 0
		cCenter := 0
		result := allCellsDistOrder(rows, cols, rCenter, cCenter)
		t.Logf("result: %v", result)
	})
	t.Run("allCellsDistOrder2", func(t *testing.T) {
		//输入：rows = 2, cols = 2, rCenter = 0, cCenter = 1
		//输出：[[0,1],[0,0],[1,1],[1,0]]
		rows := 2
		cols := 2
		rCenter := 1
		cCenter := 1
		result := allCellsDistOrder(rows, cols, rCenter, cCenter)
		t.Logf("result: %v", result)
	})
	t.Run("allCellsDistOrder3", func(t *testing.T) {
		//输入：rows = 2, cols = 3, rCenter = 1, cCenter = 2
		//输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
		rows := 2
		cols := 3
		rCenter := 1
		cCenter := 2
		result := allCellsDistOrder(rows, cols, rCenter, cCenter)
		t.Logf("result: %v", result)
	})

	t.Run("buildTree2", func(t *testing.T) {

		inorder := []int{9, 3, 15, 20, 7}
		postorder := []int{9, 15, 7, 20, 3}
		node := buildTree2(inorder, postorder)
		node.RangePrint()
	})

	t.Run("TreeNode print", func(t *testing.T) {
		tn := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}
		tn.RangePrint()
	})
	t.Run("buildTree", func(t *testing.T) {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		node := buildTree(preorder, inorder)
		node.RangePrint()
		//输出: [3,9,20,null,null,15,7]
	})
}

func TestSubsetsWithDup(t *testing.T) {
	t.Run("subsetsWithDup", func(t *testing.T) {
		nums := []int{4, 4, 4, 1, 4}
		result := subsetsWithDup(nums)
		t.Logf("result: %v", result)
	})
}

func TestMerge2(t *testing.T) {
	t.Run("merge2", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3
		merge2(nums1, m, nums2, n)
		t.Logf("nums1:%v", nums1)
	})

	t.Run("merge2", func(t *testing.T) {
		//：nums1 = [1], m = 1, nums2 = [], n = 0
		nums1 := []int{1}
		m := 1
		nums2 := []int{}
		n := 0
		merge2(nums1, m, nums2, n)
		t.Logf("nums1:%v", nums1)
	})

	t.Run("merge2", func(t *testing.T) {
		//nums1 = [0], m = 0, nums2 = [1], n = 1
		nums1 := []int{0, 0}
		m := 0
		nums2 := []int{1, 2}
		n := 2
		merge2(nums1, m, nums2, n)
		t.Logf("nums1:%v", nums1)
	})

}

func TestMaximalRectangle(t *testing.T) {
	t.Run("maximalRectangle", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'}}
		maxA := maximalRectangle(matrix)
		t.Logf("maxA: %v", maxA)

	})

	t.Run("maximalRectangle", func(t *testing.T) {
		matrix := [][]byte{
			{'0', '1'},
			{'0', '1'}}
		maxA := maximalRectangle(matrix)
		t.Logf("maxA: %v", maxA)
	})

}

func TestLargestRectangleArea(t *testing.T) {

	t.Run("largestRectangleArea", func(t *testing.T) {
		heights := []int{1, 3, 4, 2, 1}
		area := largestRectangleArea(heights)
		t.Logf("area:%v", area)
	})

	t.Run("largestRectangleArea", func(t *testing.T) {
		heights := []int{1, 1}
		area := largestRectangleArea(heights)
		t.Logf("area:%v", area)
	})
}

func TestRemoveDuplicates2(t *testing.T) {
	t.Run("removeDuplicates2", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 1, 1, 2, 3, 3, 3, 4, 5, 6}
		l := removeDuplicates2(nums)
		n := nums[:l]
		t.Logf("nums:%v len:%d", n, l)
	})
}

func TestExist(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "ABCCED"
		flag := exist(board, word)
		t.Logf("flag:%v", flag)
	})
}

func TestSubsets(t *testing.T) {
	t.Run("subsets", func(t *testing.T) {
		nums := []int{1, 2, 3}
		t.Logf("1<<n:%v", 1<<len(nums))
		result := subsets(nums)
		for _, info := range result {
			t.Logf("%v", info)
		}
	})
}

func TestSortColors(t *testing.T) {
	t.Run("sortColors", func(t *testing.T) {
		nums := []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)
		t.Logf("nums:%v", nums)
	})
}

func TestSearchMatrix(t *testing.T) {
	t.Run("searchMatrix", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 3
		flag := searchMatrix(matrix, target)
		t.Logf("flag:%v", flag)
	})
}

func TestFullJustify(t *testing.T) {
	t.Run("fullJustify", func(t *testing.T) {
		words := []string{"This", "is", "an", "example", "of", "text", "justification."}
		maxWidth := 16
		result := fullJustify(words, maxWidth)
		for _, str := range result {
			t.Logf(str)
		}
	})

	t.Run("spaceCompletion", func(t *testing.T) {
		str := "example of text"
		maxWidth := 16
		result := spaceCompletion(str, maxWidth, false)
		t.Logf(result)
	})

	t.Run("fullJustify", func(t *testing.T) {
		words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
		maxWidth := 16
		result := fullJustify(words, maxWidth)
		for _, str := range result {
			t.Logf(str)
		}
	})
	t.Run("fullJustify", func(t *testing.T) {
		words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
		maxWidth := 20
		result := fullJustify(words, maxWidth)
		for _, str := range result {
			t.Logf(str)
		}
	})

}

func TestPlusOne(t *testing.T) {
	t.Run("plusOne", func(t *testing.T) {
		nums := []int{1, 2, 3}
		n := plusOne(nums)
		t.Logf("n:%v", n)
	})
	t.Run("plusOne", func(t *testing.T) {
		nums := []int{4, 3, 2, 1}
		n := plusOne(nums)
		t.Logf("n:%v", n)
	})
	t.Run("plusOne", func(t *testing.T) {
		nums := []int{1, 0, 9}
		n := plusOne(nums)
		t.Logf("n:%v", n)
	})
}

func TestMinPathSum(t *testing.T) {
	t.Run("minPathSum", func(t *testing.T) {
		obstacleGrid := [][]int{
			{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
		}
		num := minPathSum(obstacleGrid)
		t.Logf("num:%v", num)
	})

	t.Run("minPathSum", func(t *testing.T) {
		obstacleGrid := [][]int{
			{1, 2, 3}, {4, 5, 6},
		}
		num := minPathSum(obstacleGrid)
		t.Logf("num:%v", num)
	})
}

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Run("uniquePathsWithObstacles", func(t *testing.T) {
		obstacleGrid := [][]int{
			{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
		}
		num := uniquePathsWithObstacles2(obstacleGrid)
		t.Logf("num:%v", num)
	})
}

func TestMerge(t *testing.T) {
	t.Run("merge", func(t *testing.T) {
		intervals := [][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}
		result := merge(intervals)
		for _, i := range result {
			t.Log(i)
		}
	})
}

func TestCanJump(t *testing.T) {
	t.Run("canJump", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}
		if canJump(nums) {
			t.Logf("ok")
		} else {
			t.Logf("fail")
		}
	})
}

func TestSpiralOrder(t *testing.T) {
	t.Run("spiralOrder", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		result := spiralOrder(matrix)
		t.Logf("result:%v", result)
	})
}

func TestMaxSubArray(t *testing.T) {
	t.Run("maxSubArray", func(t *testing.T) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		sum := maxSubArray(nums)
		t.Logf("sum:%v", sum)
	})
}

func TestSolveNQueens(t *testing.T) {
	t.Run("solveNQueens", func(t *testing.T) {
		n := 4
		result := solveNQueens(n)
		for _, info := range result {
			t.Log(info)
		}
	})
}

func TestGroupAnagrams(t *testing.T) {
	t.Run("groupAnagrams", func(t *testing.T) {
		str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		result := groupAnagrams(str)
		for _, info := range result {
			t.Log(info)
		}
	})
}

func TestRotate(t *testing.T) {
	t.Run("rotate", func(t *testing.T) {
		nums := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		rotate(nums)
		for _, list := range nums {
			t.Log(list)
		}
	})
}

func TestPermuteUnique(t *testing.T) {
	t.Run("permuteUnique", func(t *testing.T) {
		nums := []int{1, 1, 1}
		result := permuteUnique(nums)
		t.Logf("len(list):%v", len(result))
		for _, list := range result {
			t.Logf("list:%v", list)
		}
	})
}

func TestPermute(t *testing.T) {
	t.Run("permute", func(t *testing.T) {
		nums := []int{5, 4, 6, 2}
		result := permute(nums)
		t.Logf("len(list):%v", len(result))
		for _, list := range result {
			t.Logf("list:%v", list)
		}
	})

	//[[5,4,6,2],[5,4,2,6],
	//
	//
	//[5,6,4,2],[5,6,2,4],[5,2,4,6],[5,2,6,4],
	//
	//[4,5,6,2],[4,5,2,6],
	//
	//
	//[4,6,5,2],[4,6,2,5],
	//
	//
	//[4,2,5,6],[4,2,6,5],
	//
	//
	//[6,5,4,2],[6,5,2,4],[6,4,5,2],[6,4,2,5],
	//
	//
	//[6,2,5,4],[6,2,4,5],[2,5,4,6],[2,5,6,4],

	//[2,4,5,6],[2,4,6,5],[2,6,5,4],[2,6,4,5]]

}

func TestJump(t *testing.T) {
	t.Run("jump", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}
		steps := jump(nums)
		t.Logf("steps:%v", steps)
	})

	t.Run("jump", func(t *testing.T) {
		nums := []int{2, 3, 0, 1, 4}
		steps := jump(nums)
		t.Logf("steps:%v", steps)
	})

	t.Run("jump", func(t *testing.T) {
		nums := []int{1, 2, 1, 1, 1}
		steps := jump(nums)
		t.Logf("steps:%v", steps)
	})

}

func TestTrap(t *testing.T) {

	t.Run("trap", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		sum := trap2(height)
		t.Logf("area:%v", sum)
	})

	t.Run("trap", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		sum := trap2(height)
		t.Logf("area:%v", sum)
	})

}

func TestFirstMissingPositive(t *testing.T) {
	t.Run("firstMissingPositive", func(t *testing.T) {
		nums := []int{1, 2, 0}
		n := firstMissingPositive(nums)
		t.Logf("n:%v", n)
	})
	t.Run("firstMissingPositive", func(t *testing.T) {
		nums := []int{3, 4, -1, 1}
		n := firstMissingPositive(nums)
		t.Logf("n:%v", n)
	})
	t.Run("firstMissingPositive", func(t *testing.T) {
		nums := []int{7, 8, 9, 11, 12}
		n := firstMissingPositive(nums)
		t.Logf("n:%v", n)
	})
}

func TestCombinationSum2(t *testing.T) {
	t.Run("combinationSum2", func(t *testing.T) {
		candidates := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8
		result := combinationSum2(candidates, target)
		for _, info := range result {
			t.Logf("info:%v", info)
		}
	})

	t.Run("combinationSum", func(t *testing.T) {
		candidates := []int{2, 5, 2, 1, 2}
		target := 5
		result := combinationSum2(candidates, target)
		t.Logf("result len():%d", len(result))
		for _, info := range result {
			t.Logf("info:%v", info)
		}
	})
}

func TestCombinationSum(t *testing.T) {
	t.Run("combinationSum", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		result := combinationSum(candidates, target)
		for _, info := range result {
			t.Logf("info:%v", info)
		}
	})
	t.Run("combinationSum", func(t *testing.T) {
		candidates := []int{5, 6, 2, 3, 4}
		target := 11
		result := combinationSum(candidates, target)
		t.Logf("result len():%d", len(result))
		for _, info := range result {
			t.Logf("info:%v", info)
		}
	})

}

func TestSolveSudoku(t *testing.T) {
	t.Run("solveSudoku", func(t *testing.T) {
		board := [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}
		solveSudoku(board)
		for _, rows := range board {
			list := make([]int, 0, 9)
			for _, col := range rows {
				byteToInt, _ := strconv.Atoi(string(col))
				list = append(list, byteToInt)
			}
			t.Logf("%v", list)
		}

	})
}

func TestSearchInsert(t *testing.T) {
	t.Run("searchInsert", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 5
		index := searchInsert(nums, target)
		t.Logf("index:%v", index)
	})
	t.Run("searchInsert", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 2
		index := searchInsert(nums, target)
		t.Logf("index:%v", index)
	})
	t.Run("searchInsert", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		target := 7
		index := searchInsert(nums, target)
		t.Logf("index:%v", index)
	})

}

func TestSearchRange(t *testing.T) {
	t.Run("searchRange", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{}
		target := 0
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{2, 2}
		target := 2
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{2, 2}
		target := 3
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{0}
		target := 1
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

	t.Run("searchRange", func(t *testing.T) {
		nums := []int{1}
		target := 1
		indexs := searchRange(nums, target)
		t.Logf("index:%v", indexs)
	})

}

func TestSearch(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		target := 6
		index := search(nums, target)
		t.Logf("index:%v", index)
	})

	t.Run("search", func(t *testing.T) {
		nums := []int{1}
		target := 1
		index := search2(nums, target)
		t.Logf("index:%v", index)
	})
}

func TestNextPermutation(t *testing.T) {
	t.Run("nextPermutation", func(t *testing.T) {
		nums := []int{3, 2, 1}
		quickSort(nums, 0, 2)
		t.Logf("nums:%v", nums)
		//nextPermutation(nums)
		//t.Logf("nums:%v", nums)
	})

	t.Run("nextPermutation", func(t *testing.T) {
		nums := []int{2, 4, 3, 7, 6, 5, 1}
		nextPermutation(nums)
		t.Logf("nums:%v", nums)
	})

}

func TestRemoveElement(t *testing.T) {
	t.Run("removeElement", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		val := 3
		index := removeElement(nums, val)
		t.Logf("index:%v nums:%v", index, nums)
	})

	t.Run("removeElement", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		val := 2
		index := removeElement(nums, val)
		t.Logf("index:%v nums:%v", index, nums)
	})

}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("removeDuplicates", func(t *testing.T) {
		nums := []int{1, 1, 2}
		num := removeDuplicates(nums)
		t.Logf("nums:%v", nums)
		t.Logf("num:%v", num)
	})

	t.Run("removeDuplicates", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		num := removeDuplicates(nums)
		t.Logf("nums:%v", nums)
		t.Logf("num:%v", num)
	})

}

func TestFourSum(t *testing.T) {
	t.Run("fourSum", func(t *testing.T) {
		nums := []int{1, 0, -1, 0, -2, 2}
		target := 0
		result := fourSum(nums, target)
		for _, info := range result {
			t.Log(info)
		}
	})

	t.Run("fourSum", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2, 2}
		target := 8
		result := fourSum(nums, target)
		for _, info := range result {
			t.Log(info)
		}
	})
	//[[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	t.Run("fourSum", func(t *testing.T) {
		nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
		target := 0
		result := fourSum(nums, target)
		for _, info := range result {
			t.Log(info)
		}
	})

}

func TestThreeSumClosest(t *testing.T) {
	t.Run("threeSumClosest", func(t *testing.T) {
		nums := []int{-1, 2, 1, -4}
		target := 1
		v := threeSumClosest(nums, target)
		t.Logf("v:%v", v)
	})
}

func TestThreeSum(t *testing.T) {

	t.Run("QuickSort", func(t *testing.T) {
		nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		QuickSort(nums, 0, len(nums)-1)
		t.Logf("num:%v", nums)
	})

	t.Run("QuickSort2", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		QuickSort(nums, 0, len(nums)-1)
		t.Logf("num:%v", nums)
	})

	t.Run("threeSum", func(t *testing.T) {
		nums := []int{-1, 0, 1, 2, -1, -4}
		result := threeSum(nums)
		for _, info := range result {
			t.Log(info)
		}
	})

}

func TestMaxArea(t *testing.T) {
	t.Run("maxArea", func(t *testing.T) {
		num := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		area := maxArea(num)
		t.Logf("area:%v", area)
	})

	t.Run("maxArae", func(t *testing.T) {
		num := []int{1, 1}
		area := maxArea(num)
		t.Logf("area:%v", area)
	})
}

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		num := findMedianSortedArrays(nums1, nums2)
		t.Logf("num:%v", num)
	})

	t.Run("one", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		num := findMedianSortedArrays(nums1, nums2)
		t.Logf("num:%v", num)
	})

	t.Run("one", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{}
		num := findMedianSortedArrays(nums1, nums2)
		t.Logf("num:%v", num)
	})
}

func TestTwoSum(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		n := TwoSum(nums, target)
		t.Logf("n:%v", n)
	})

	t.Run("two", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		n := TwoSum(nums, target)
		t.Logf("n:%v", n)
	})

	t.Run("three", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		n := TwoSum(nums, target)
		t.Logf("n:%v", n)
	})

}
