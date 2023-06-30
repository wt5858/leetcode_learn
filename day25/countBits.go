package main

import "fmt"

// 剑指 Offer II 003. 前 n 个数字二进制中 1 的个数
// 给定一个非负整数 n，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
//
//
//
//示例 1:
//
//输入: n = 2
//输出: [0,1,1]
//解释:
//0 --> 0
//1 --> 1
//2 --> 10
//示例2:
//
//输入: n = 5
//输出: [0,1,1,2,1,2]
//解释:
//0 --> 0
//1 --> 1
//2 --> 10
//3 --> 11
//4 --> 100
//5 --> 101
//
//
//说明 :
//
//0 <= n <= 105
//
//
//进阶:
//
//给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
//要求算法的空间复杂度为O(n)。
//你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的__builtin_popcount）来执行此操作。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/w3tCBm
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	n := 5
	fmt.Println(countBits(n))
}

// 对于正整数 x，将其二进制表示右移一位，等价于将其二进制表示的最低位去掉，得到的数是 [x/2]
// 如果是偶数 bits[x] = bits[x/2]
// 奇数 bit[x] = bits[x/2] + 1
// 也就是 bit[x] = bits[x/2] + x%2
// x%2 也就等于 i&1
// [x/2] 也就可以换算为 i >> 1
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i < len(res); i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}