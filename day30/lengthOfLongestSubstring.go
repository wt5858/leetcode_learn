package main

import "fmt"

// 剑指 Offer II 016. 不含重复字符的最长子字符串
// 给定一个字符串 s ，请你找出其中不含有重复字符的最长连续子字符串的长度。
//
//
//
//示例1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s由英文字母、数字、符号和空格组成
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/wtcaE1
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	hc := [256]int{}
	n := len(s)
	l, r, res := 0, 0, 0
	for l < n {
		if r < n && hc[s[r]] == 0 {
			hc[s[r]]++
			r++
		} else {
			hc[s[l]]--
			l++
		}

		if res < r-l {
			res = r - l
		}
	}
	return res
}
