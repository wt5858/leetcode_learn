package main

import "fmt"

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
//
//
//
//参考以下这颗二叉搜索树：
//
//     5
//    / \
//   2   6
//  / \
// 1   3
//示例 1：
//
//输入: [1,6,3,2,5]
//输出: false
//示例 2：
//
//输入: [1,3,2,6,5]
//输出: true
//
//
//提示：
//
//数组长度 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	p := []int{1, 3, 2, 6, 5}
	r := verifyPostorder(p)
	fmt.Println(r)
}

func verifyPostorder(postorder []int) bool {
	return helper(postorder, 0, len(postorder)-1)
}

func helper(postorder []int, l int, r int) bool {
	if l > r {
		return true
	}

	root := postorder[r]
	mid := l

	// 左树的节点的值都比根节点的小  用mid可以来找到左树的分割点
	for postorder[mid] < root {
		mid++
	}

	// 右树的节点比根节点大 循环判断和根节点比较
	tmp := mid
	for tmp < r {
		tmp++
		if postorder[tmp] < root {
			return false
		}
	}

	// 递归判断
	// helper(postorder, l, mid-1) 左树
	// helper(postorder, mid, r-1) 右树
	return helper(postorder, l, mid-1) && helper(postorder, mid, r-1)
}
