package main

import "fmt"

// 剑指 Offer 34. 二叉树中和为某一值的路径
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。
//
//
//
//示例 1：
//
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//示例 2：
//
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//示例 3：
//
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//
//提示：
//
//树中节点总数在范围 [0, 5000] 内
//-1000 <= Node.val <= 1000
//-1000 <= targetSum <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
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

	target := 8

	res := pathSum(a, target)
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
func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	var path []int

	var helper func(r *TreeNode, tar int)
	helper = func(r *TreeNode, tar int) {
		if r == nil {
			return
		}

		path = append(path, r.Val)
		// 让tar每次减少当前节点的值
		tar -= r.Val

		// 当tar=0并且没有左右节点的时候，说明路径正确
		if tar == 0 && r.Left == nil && r.Right == nil {
			// path是切片 这个地方要用深拷贝去处理
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		// 递归处理左树
		helper(r.Left, tar)
		// 递归处理右树
		helper(r.Right, tar)

		// 向上回溯的时候把路径最后的一个给剔除掉
		path = path[:len(path)-1]
	}

	helper(root, target)

	return res
}
