package main

import "fmt"

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
	res := reverseList(head)

	fmt.Printf("ListNode: %v", *res)
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
func reverseList(head *ListNode) *ListNode {
	// 用于保存上一个节点
	var pre *ListNode
	// 当前节点
	curr := head
	for curr != nil {
		// 下一个节点
		next := curr.Next
		// 让当前节点的下一个节点等于上一个节点（这里就是翻转）
		curr.Next = pre
		// 让上一节点等于当前节点，好进行下一次循环
		pre = curr
		// 当前节点改为下一个节点
		curr = next
	}
	return pre
}
