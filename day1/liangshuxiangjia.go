package main

import "fmt"

// 两数相加
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
//
//
//示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//示例 2：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//示例 3：
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//提示：
//
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	res := addTwoNumbers(l1, l2)
	fmt.Println(fmt.Sprintf("res: %v", res))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (h *ListNode) {
	l := &ListNode{}
	c := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + c
		// 大于10的情况就向前进1
		sum, c = sum%10, sum/10

		if h != nil {
			l.Next = &ListNode{sum, nil}
			l = l.Next
		} else {
			h = &ListNode{sum, nil}
			// 这里l是指针，指向相同的内存地址，l修改对应的h也会改
			l = h
		}
	}

	// 如果最后还有加起来大于10的时候
	if c > 0 {
		l.Next = &ListNode{c, nil}
	}

	return
}
