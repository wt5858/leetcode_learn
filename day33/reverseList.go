package main

import "fmt"

// 剑指 Offer II 024. 反转链表
// 给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/UHnkqh
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

	fmt.Printf("新的链表:%+v", reverseList(head))
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
	cur := head
	var pre *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}
