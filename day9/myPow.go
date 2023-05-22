package main

import "fmt"

// 剑指 Offer 16. 数值的整数次方
//实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
//
//
//
//示例 1：
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//示例 2：
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//示例 3：
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//
//提示：
//
//-100.0 <x< 100.0
//-231<= n <=231-1
//-104<= xn<= 104
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	x := 2.00000
	n := -3
	r := myPow(x, n)
	fmt.Println(r)
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		// 2的-2次方 就可以换成 1/2的2次方
		return quickMul(1/x, -n)
	}
	return quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := quickMul(x, n/2)
	if n%2 == 0 {
		// 偶数的话  2的8次方就等于2的4次方 乘以 2的4次方
		return res * res
	}
	// 奇数的话  3的9次方就等于 3的4次方 乘以 3的4次方 再乘以 3
	return res * res * x
}
