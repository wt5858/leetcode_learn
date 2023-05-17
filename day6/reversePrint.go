package main

import "fmt"

// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
//示例 1：
//
//输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//限制：
//
//0 <= 链表长度 <= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := &ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				2,
				nil,
			},
		},
	}
	res := reversePrint(a)
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
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	// 交换数组
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
