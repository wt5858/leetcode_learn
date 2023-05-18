package main

import "fmt"

// 剑指 Offer 07. 重建二叉树
// 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
//示例 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//示例 2:
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//限制：
//
//0 <= 节点个数 <= 5000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	res := buildTree(preorder, inorder)

	fmt.Println(*res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 找中序遍历的根节点的下标
	i := 0
	for j := 0; j < len(inorder); j++ {
		if inorder[j] == preorder[0] {
			i = j
			break
		}
	}

	root := &TreeNode{preorder[0], nil, nil}

	// 前序遍历的左树比中序变量的左树的key要多一个根节点，也就是多1，所以这个地方的截取是从1到i+1
	preLeft := preorder[1 : i+1]
	// 除去左树和根节点就是右树
	preRight := preorder[i+1:]

	// 中序是从左树开始，所以这个地方截图是 0到i
	inLeft := inorder[:i]
	// 除去左树和根节点就是右树
	inRight := inorder[i+1:]

	// 递归处理所有的左树
	root.Left = buildTree(preLeft, inLeft)
	// 递归处理所有的右树
	root.Right = buildTree(preRight, inRight)

	return root
}
