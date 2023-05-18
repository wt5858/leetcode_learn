package main

import (
	"fmt"
	"os"
)

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead操作返回 -1 )
//
//
//
//示例 1：
//
//输入：
//["CQueue","appendTail","deleteHead","deleteHead","deleteHead"]
//[[],[3],[],[],[]]
//输出：[null,null,3,-1,-1]
//示例 2：
//
//输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//提示：
//
//1 <= values <= 10000
//最多会对appendTail、deleteHead 进行10000次调用
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	cq := Constructor()
	cq.AppendTail(2)
	fmt.Fprintf(os.Stdout, "add 2 cq:%v \n", cq)
	cq.AppendTail(3)
	fmt.Fprintf(os.Stdout, "add 3 cq:%v \n", cq)
	r := cq.DeleteHead()
	fmt.Fprintf(os.Stdout, "DeleteHead cq:%v , v: %d \n", cq, r)
	r = cq.DeleteHead()
	fmt.Fprintf(os.Stdout, "DeleteHead cq:%v , v: %d \n", cq, r)
	r = cq.DeleteHead()
	fmt.Fprintf(os.Stdout, "DeleteHead cq:%v , v: %d \n", cq, r)
}

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.InStack = append(q.InStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.OutStack) == 0 {
		if len(q.InStack) == 0 {
			return -1
		}

		// 把入栈里面的都弄到出栈那个里面
		for len(q.InStack) > 0 {
			q.OutStack = append(q.OutStack, q.InStack[len(q.InStack)-1])
			q.InStack = q.InStack[:len(q.InStack)-1]
		}
	}
	v := q.OutStack[len(q.OutStack)-1]
	q.OutStack = q.OutStack[:len(q.OutStack)-1]
	return v
}
