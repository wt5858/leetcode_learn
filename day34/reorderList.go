package main

import "fmt"

// 剑指 Offer II 026. 重排链表
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0→ L1→ … → Ln-1→ Ln
//请将其重新排列后变为：
//
//L0→Ln→L1→Ln-1→L2→Ln-2→ …
//
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
//示例 1:
//
//
//
//输入: head = [1,2,3,4]
//输出: [1,4,2,3]
//示例 2:
//
//
//
//输入: head = [1,2,3,4,5]
//输出: [1,5,2,4,3]
//
//
//提示：
//
//链表的长度范围为 [1, 5 * 104]
//1 <= node.val <= 1000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/LGjMqU
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

	reorderList(head)

	fmt.Printf("移除后的链表:%+v", head)
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var hc []*ListNode
	t := head

	for t != nil {
		hc = append(hc, t)
		t = t.Next
	}

	l, r := 0, len(hc)-1

	for l < r {
		hc[l].Next = hc[r]
		l++
		if l == r {
			break
		}

		hc[r].Next = hc[l]
		r--

	}
	hc[l].Next = nil
}
