package array

//给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
//示例 1:
//输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//输出：[3,9,20,null,null,15,7]
//示例 2:
//输入：inorder = [-1], postorder = [-1]
//输出：[-1]
//提示:
//
//1 <= inorder.length <= 3000
//postorder.length == inorder.length
//-3000 <= inorder[i], postorder[i] <= 3000
//inorder 和 postorder 都由 不同 的值组成
//postorder 中每一个值都在 inorder 中
//inorder 保证是树的中序遍历
//postorder 保证是树的后序遍历

// 任意一颗树而言，后序遍历的形式总是
// [[左子树的后序遍历结果], [右子树的后序遍历结果],根节点 ]
// 即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree2(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTree2(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
	return root
}
