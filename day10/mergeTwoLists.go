package main

import "fmt"

// 剑指 Offer 25. 合并两个排序的链表
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
//示例1：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//限制：
//
//0 <= 链表长度 <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	l1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,
				nil,
			},
		},
	}

	l2 := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}

	res := mergeTwoLists(l1, l2)
	fmt.Println(res)
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res := &ListNode{
		0,
		nil,
	}
	next := res

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			next.Next = l2
			l2 = l2.Next
		} else {
			next.Next = l1
			l1 = l1.Next
		}

		next = next.Next

		if l1 == nil {
			next.Next = l2
		}

		if l2 == nil {
			next.Next = l1
		}

	}

	return res.Next
}
