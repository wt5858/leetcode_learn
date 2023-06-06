package main

import (
	"fmt"
	"strconv"
)

// 剑指 Offer 46. 把数字翻译成字符串
// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
//示例 1:
//
//输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
//
//
//提示：
//
//0 <= num < 231
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	num := 18580
	r := translateNum(num)
	fmt.Println(r)
}

func translateNum(num int) int {
	s := strconv.Itoa(num)

	r, b := 1, 1
	for i := 2; i <= len(s); i++ {
		tmp := s[i-2 : i]
		c := r
		tmp2, _ := strconv.Atoi(tmp)
		if tmp2 >= 10 && tmp2 <= 25 {
			c = r + b
		}
		b = r
		r = c
	}
	return r

}
