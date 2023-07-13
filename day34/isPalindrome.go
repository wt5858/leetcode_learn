package main

import "fmt"

// 剑指 Offer II 027. 回文链表
// 给定一个链表的 头节点head，请判断其是否为回文链表。
//
//如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
//
//
//
//示例 1：
//
//
//
//输入: head = [1,2,3,3,2,1]
//输出: true
//示例 2：
//
//
//
//输入: head = [1,2]
//输出: false
//
//
//提示：
//
//链表 L 的长度范围为 [1, 105]
//0<= node.val <= 9
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/aMhZSa
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					3,
					&ListNode{
						2,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var hc []*ListNode

	for t := head; t != nil; t = t.Next {
		hc = append(hc, t)
	}

	for l, r := 0, len(hc)-1; l < r; l, r = l+1, r-1 {
		if hc[l].Val != hc[r].Val {
			return false
		}
	}

	return true
}
