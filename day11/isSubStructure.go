package main

import "fmt"

// 剑指 Offer 26. 树的子结构
// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。
//
//例如:
//给定的树 A:
//
//  3
//  / \
// 4  5
// / \
//1  2
//给定的树 B：
//
// 4
// /
//1
//返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
//
//示例 1：
//
//输入：A = [1,2,3], B = [3,1]
//输出：false
//示例 2：
//
//输入：A = [3,4,5,1,2], B = [4,1]
//输出：true
//限制：
//
//0 <= 节点个数 <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof
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

	b := &TreeNode{
		4,
		&TreeNode{
			1,
			nil,
			nil,
		},
		nil,
	}

	res := isSubStructure(a, b)
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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		// 说明B树的节点都判断完了
		return true
	}
	if A == nil {
		// A树判断完了，但是B树还有节点，说明肯定不是子树
		return false
	}

	if A.Val != B.Val {
		return false
	}

	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
