package main

import "fmt"

// 剑指 Offer 49. 丑数
// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
//示例:
//
//输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//说明:
//
//1是丑数。
//n不超过1690。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/chou-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n := 10
	r := nthUglyNumber(n)
	fmt.Println(r)
}

func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2 := dp[a] * 2
		n3 := dp[b] * 3
		n5 := dp[c] * 5
		tmp := min(min(n2, n3), n5)
		if tmp == n2 {
			a++
		}
		if tmp == n3 {
			b++
		}
		if tmp == n5 {
			c++
		}
		dp[i] = tmp
	}
	return dp[n-1]
}

func min(l, r int) int {
	if l > r {
		return r
	}
	return l
}
