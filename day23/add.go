package main

import "fmt"

// 剑指 Offer 65. 不用加减乘除做加法
// 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
//
//
//
//示例:
//
//输入: a = 1, b = 1
//输出: 2
//
//
//提示：
//
//a,b均可能是负数或 0
//结果不会溢出 32 位整数
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	a := 1
	b := -3
	fmt.Println(add(a, b))
}

// 可以用 a^b 运算表示无进位的加法
// 可以用 (a&b)<<1 表示进位

func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}
