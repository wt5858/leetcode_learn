package main

import "fmt"

// 剑指 Offer II 025. 链表中的两数相加
// 给定两个 非空链表 l1和 l2来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
//示例1：
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//示例2：
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//示例3：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//提示：
//
//链表的长度范围为 [1, 100]
//0 <= node.val <= 9
//输入数据保证链表代表的数字无前导 0
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/lMSNwu
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	l1 := &ListNode{
		7,
		&ListNode{
			2,
			&ListNode{
				4,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}

	l2 := &ListNode{
		5,
		&ListNode{
			6,
			&ListNode{
				4,
				nil,
			},
		},
	}

	fmt.Printf("新的链表:%+v", addTwoNumbers(l1, l2))
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 双栈+头插法
	var head *ListNode
	var n1 []*ListNode
	var n2 []*ListNode

	for l1 != nil {
		n1 = append(n1, l1)
		l1 = l1.Next
	}

	for l2 != nil {
		n2 = append(n2, l2)
		l2 = l2.Next
	}

	c := 0

	for len(n1) > 0 || len(n2) > 0 || c > 0 {
		sum := 0
		if len(n1) > 0 {
			sum += n1[len(n1)-1].Val
			n1 = n1[:len(n1)-1]
		}

		if len(n2) > 0 {
			sum += n2[len(n2)-1].Val
			n2 = n2[:len(n2)-1]
		}

		sum += c

		node := &ListNode{Val: sum % 10}
		// 头插法
		node.Next = head
		head = node

		c = sum / 10
	}

	return head
}
