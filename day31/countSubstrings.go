package main

import "fmt"

// 剑指 Offer II 020. 回文子字符串的个数
// 给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
//示例 1：
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//提示：
//
//1 <= s.length <= 1000
//s 由小写英文字母组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/a7VOhD
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}

// 中心扩散法
func countSubstrings(s string) int {
	res, n := 0, len(s)

	var f func(int, int)
	f = func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			res++

			// 扩大窗口
			l--
			r++
		}
	}

	for i := 0; i < n; i++ {
		// 回文数中心为单数
		f(i, i)

		// 回文数中心为复数
		f(i, i+1)
	}

	return res
}
