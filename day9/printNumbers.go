package main

import (
	"fmt"
	"strconv"
)

// 剑指 Offer 17. 打印从1到最大的n位数
// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//
//示例 1:
//
//输入: n = 1
//输出: [1,2,3,4,5,6,7,8,9]
//
//
//说明：
//
//用返回一个整数列表来代替打印
//n 为正整数
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n := 3
	res := printNumbers(n)
	fmt.Println(res)
}

// 暴力解题
//func printNumbers(n int) []int {
//	a := int(math.Pow10(n))
//	var res []int
//	for i := 1; i < a; i++ {
//		res = append(res, i)
//	}
//	return res
//}

// dfs
func printNumbers(n int) []int {
	var res []int

	// k 表示轮次，n表示数字位数，s表示对应的数字
	var dfs func(k, n int, s string)

	dfs = func(k, n int, s string) {
		if n == k {
			// 位数和轮次相同的时候，把值添加进去，也是dfs的出口
			r, _ := strconv.Atoi(s)
			res = append(res, r)
			return
		}

		for i := 0; i < 10; i++ {
			// 搜索下一个轮次
			// 这个地方的入参 为 s+strconv.Itoa(i) 而不是 strconv.Itoa(i) + s
			// 第一轮的时候是1位数 也就是1~9 第二轮的时候是2位数是 10~99 第三轮的时候是3位数 100~999
			// 所以说 这个地方应该把前面轮次的s拼接在前面
			dfs(k+1, n, s+strconv.Itoa(i))
		}
	}

	for i := 1; i <= n; i++ {
		// 有多少位数字就有多少次循环
		for j := 1; j < 10; j++ {
			dfs(1, i, strconv.Itoa(j))
		}
	}

	return res
}
