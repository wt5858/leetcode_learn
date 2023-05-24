package main

import "fmt"

// 剑指 Offer 28. 对称的二叉树
// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
//例如，二叉树[1,2,2,3,4,4,3] 是对称的。
//
//  1
// / \
// 2  2
/// \ / \
//3 4 4 3
//但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
//
//  1
// / \
// 2  2
// \  \
// 3  3
//
//
//
//示例 1：
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//示例 2：
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//限制：
//
//0 <= 节点个数 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof
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

	res := isSymmetric(a)
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
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return check(l.Left, r.Right) && check(l.Right, r.Left)
}
