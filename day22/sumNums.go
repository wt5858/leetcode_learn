package main

import "fmt"

// 剑指 Offer 64. 求1+2+…+n
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
//
//
//
//示例 1：
//
//输入: n = 3
//输出:6
//示例 2：
//
//输入: n = 9
//输出:45
//
//
//限制：
//
//1 <= n<= 10000
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/qiu-12n-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n := 5
	fmt.Println(sumNums(n))
}

func sumNums(n int) int {
	_ = n > 1 && func() bool { n += sumNums(n - 1); return true }()
	return n
}
