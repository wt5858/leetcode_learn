package main

import "fmt"

// 剑指 Offer 32 - III. 从上到下打印二叉树 III
// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
//
//
//
//例如:
//给定二叉树:[3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回其层次遍历结果：
//
//[
//  [3],
//  [20,9],
//  [15,7]
//]
//
//
//提示：
//
//节点总数 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
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

	res := levelOrder2(a)
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
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queues := []*TreeNode{root}
	var res [][]int
	level := 0

	for len(queues) > 0 {
		var tmp []int
		var lists []*TreeNode
		for i := 0; i < len(queues); i++ {
			node := queues[i]
			if node.Left != nil {
				lists = append(lists, node.Left)
			}
			if node.Right != nil {
				lists = append(lists, node.Right)
			}
			tmp = append(tmp, node.Val)
		}

		if len(tmp) > 0 {
			if level%2 == 1 {
				n := len(tmp)
				for j := 0; j < n/2; j++ {
					tmp[j], tmp[n-j-1] = tmp[n-j-1], tmp[j]
				}
			}
			res = append(res, tmp)
		}
		queues = append(queues[len(queues):], lists...)
		level++
	}

	return res
}
