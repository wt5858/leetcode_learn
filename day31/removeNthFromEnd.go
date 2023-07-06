package main

import "fmt"

// 剑指 Offer II 021. 删除链表的倒数第 n 个结点
// 给定一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。
//
//
//
//示例 1：
//
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/SLwz0R
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	n := 2

	fmt.Printf("移除后的链表:%+v", removeNthFromEnd(head, n))
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
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	var nodes []*ListNode
//
//	for node := head; node != nil; node = node.Next {
//		nodes = append(nodes, node)
//	}
//
//	x := len(nodes) - n - 1
//	if x < 0 {
//		return head.Next
//	}
//
//	pre := nodes[x]
//	pre.Next = pre.Next.Next
//
//	return head
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{0, head}
	fast := dump
	slow := dump

	// 先让快指针走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 慢指针走剩余的步数，这是慢指针对应的也就是倒数n
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dump.Next
}
