package main

import "fmt"

// 剑指 Offer 14- I. 剪绳子
// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
//
//示例 1：
//
//输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1
//示例2:
//
//输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 ×3 ×4 = 36
//提示：
//
//2 <= n <= 58
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/jian-sheng-zi-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n := 9
	res := cuttingRope(n)
	fmt.Println(res)
}

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}

	// 动态规划
	// 拆分成两段最大值为 f(n) = f(i) * f(n-i)  这种剪绳子可以考虑为 拆分成无数个两根绳子 取它们乘积的最大值
	// hc 用于保存每段绳子最大值
	hc := make([]int, n+1)
	// 特殊处理 1 2 3
	// 1 2 3 在不拆分的时候最大，拆分了反而变小了 所以这个地方就不取拆分后的结果 直接取它最长的情况
	hc[1] = 1
	hc[2] = 2
	hc[3] = 3

	a := 0

	for i := 4; i <= n; i++ {
		// 因为要进行拆分，所以这个地方只需要取i的一半
		for j := 1; j <= i/2; j++ {
			a = max(a, hc[j]*hc[i-j])
		}
		hc[i] = a
		a = 0
	}
	return hc[n]
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
