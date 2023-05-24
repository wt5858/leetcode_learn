package main

import "fmt"

// 剑指 Offer 27. 二叉树的镜像
// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
//例如输入：
//
//  4
// /  \
// 2   7
/// \  / \
//1  3 6  9
//镜像输出：
//
//  4
// /  \
// 7   2
/// \  / \
//9  6 3 1
//
//
//
//示例 1：
//
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//
//限制：
//
//0 <= 节点个数 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := &TreeNode{
		3,
		&TreeNode{
			4,
			&TreeNode{
				1,
				nil,
				nil,
			},
			&TreeNode{
				2,
				nil,
				nil,
			},
		},
		&TreeNode{
			5,
			nil,
			nil,
		},
	}

	res := mirrorTree(a)
	fmt.Println(res)
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
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := mirrorTree(root.Left)
	r := mirrorTree(root.Right)
	root.Left = l
	root.Right = r

	return root
}
