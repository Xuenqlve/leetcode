package array

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//
//示例 1:
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//示例 2:
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]
//
//提示:
//
//1 <= preorder.length <= 3000
//inorder.length == preorder.length
//-3000 <= preorder[i], inorder[i] <= 3000
//preorder 和 inorder 均 无重复 元素
//inorder 均出现在 preorder
//preorder 保证 为二叉树的前序遍历序列
//inorder 保证 为二叉树的中序遍历序列

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) RangePrint() {
	fmt.Println("root print:", n.Val)
	if n.Left != nil {
		n.Left.RangePrint()
	}
	if n.Right != nil {
		n.Right.RangePrint()
	}
}

// 任意一颗树而言，前序遍历的形式总是
// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// 即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	fmt.Printf("index:%v inorder:%v preorder:%v\n", i, inorder, preorder)
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
